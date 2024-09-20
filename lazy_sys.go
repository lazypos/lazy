package lazy

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"time"
)

func ExitMonitor(cancel context.CancelCauseFunc, extFile string, sigs ...os.Signal) {
	os.Remove(extFile)

	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, sigs...)

	go func() {
		for {
			select {
			case sig := <-chSignal:
				cancel(fmt.Errorf("收到退出信号:%v", sig))
				return
			default:
				if reason, err := os.ReadFile(extFile); err == nil {
					cancel(fmt.Errorf("发现退出文件:%v", string(reason)))
					return
				}
			}
			time.Sleep(time.Second)
		}
	}()
}

func RunSelf(dmpFile string) bool {
	if len(os.Args) > 1 && os.Args[1] == "main" {
		return false
	}

	buf := bytes.NewBuffer(nil)
	cmdInfo := exec.Command(os.Args[0], "main")
	cmdInfo.Stdout = os.Stdout
	cmdInfo.Stderr = buf
	cmdInfo.Run()

	if buf.Len() > 0 {
		os.WriteFile(dmpFile, buf.Bytes(), 0o644)
	}
	return true
}

func GetProcessDir() string {
	exePath, err := os.Executable()
	if err != nil {
		return ""
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

func GetFileSha1(fpath string) string {
	fd, err := os.Open(fpath)
	if err != nil {
		return ""
	}
	defer fd.Close()

	hash := sha1.New()
	if _, err = io.Copy(hash, fd); err != nil {
		return ""
	}
	return hex.EncodeToString(hash.Sum(nil))
}

func WriteFileIO(fname string, ireader io.Reader) error {
	os.MkdirAll(filepath.Dir(fname), 0755)
	fh, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer fh.Close()

	if _, err = io.Copy(fh, ireader); err != nil {
		return err
	}
	return nil
}

func ReadFileIo(write io.Writer, fname string) error {
	hd, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer hd.Close()

	_, err = io.Copy(write, hd)
	if err != nil {
		return err
	}
	return nil
}

func MoveFile2(spath, dir, dpath string, cover bool) error {

	newpath := ""
	if dpath != "" {
		newpath = dpath
	} else {
		newpath = filepath.Join(dir, filepath.Base(spath))
	}

	if !cover {
		if _, err := os.Stat(newpath); err == nil {
			return fmt.Errorf("文件已经存在:%s", newpath)
		}
	}

	hd, err := os.Open(spath)
	if err != nil {
		return err
	}
	if err = WriteFileIO(newpath, hd); err != nil {
		hd.Close()
		return err
	}
	hd.Close()

	return os.Remove(spath)
}
