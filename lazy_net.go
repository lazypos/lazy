package lazy

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"time"
)

func QueryHttps(useTls bool, url, meth string, body io.Reader, header map[string]string) (string, error) {
	req, err := http.NewRequest(meth, url, body)
	if err != nil {
		return "", fmt.Errorf("NewRequest失败：%v", err)
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}

	client := &http.Client{Timeout: time.Duration(60 * time.Second)}
	if useTls {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client.Transport = tr
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("http请求失败：%v", err)
	}
	defer resp.Body.Close()

	text, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("http读取body数据失败：%v", err)
	}
	return string(text), nil
}

func QueryHttpsFile(useTls bool, url, meth, fname string, header map[string]string) error {
	req, err := http.NewRequest(meth, url, nil)
	if err != nil {
		return fmt.Errorf("NewRequest失败：%v", err)
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}

	client := &http.Client{Timeout: time.Duration(60 * time.Second)}
	if useTls {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client.Transport = tr
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("http请求失败：%v", err)
	}
	defer resp.Body.Close()

	return WriteFileIO(fname, resp.Body)
}

func QueryHttpsIO(useTls bool, url, meth string, out io.Writer, header map[string]string) error {
	req, err := http.NewRequest(meth, url, nil)
	if err != nil {
		return fmt.Errorf("NewRequest失败：%v", err)
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}

	client := &http.Client{Timeout: time.Duration(60 * time.Second)}
	if useTls {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client.Transport = tr
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("http请求失败：%v", err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
