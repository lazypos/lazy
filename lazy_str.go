package lazy

import (
	"regexp"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

const ILegalString = `'"?- |()%<>$&*\\/,;:.~`
const filemaxlen = 99

func LegalString(src string) bool {
	if src == "" || len(src) > 50 {
		return false
	}
	if strings.ContainsAny(src, ILegalString) {
		return false
	}
	for _, v := range src {
		if !unicode.IsPrint(v) {
			return false
		}
	}
	return true
}

func CheckUser(user string) string {
	if len(user) > 50 || len(user) < 6 {
		return "账户长度不合规！"
	}
	if !LegalString(user) {
		return "账户格式不正确！"
	}
	return ""
}

func CheckNick(nick string) string {
	nicklen := utf8.RuneCountInString(nick)
	if nicklen > 7 || nicklen < 2 {
		return "昵称长度不正确！"
	}
	if !LegalString(nick) {
		return "昵称不合规！"
	}
	return ""
}

func CheckPass(pass string) string {
	if len(pass) > 32 || len(pass) < 6 {
		return "密码长度不正确！"
	}
	return ""
}

func VeriMail(mail string) bool {
	if len(mail) > 64 {
		return false
	}
	reg := regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)
	return reg.MatchString(mail)
}

func GetNowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func ChangeString(src string) string {
	idx := 0
	tmp := src
	for idx != -1 {
		idx = strings.IndexAny(tmp, ILegalString)
		if idx != -1 {
			tmp = tmp[:idx] + "_" + tmp[idx+1:]
		}
	}
	tt := ""
	for _, v := range tmp {
		if !unicode.IsPrint(v) {
			tt += string("_")
		} else {
			tt += string(v)
		}
	}
	return tt
}

func GetFileBase(name string) string {
	idx := strings.LastIndex(name, ".")
	if idx == -1 {
		return name
	}
	return name[:idx]
}

func GetLastIndexString(str, find string) string {
	idx := strings.LastIndex(str, find)
	if idx == -1 {
		return str
	}
	return str[idx+1:]
}

func GetStringBase(str, find string) string {
	idx := strings.LastIndex(str, find)
	if idx == -1 {
		return str
	}
	return str[:idx]
}

func SplitLast(src, sep string) (string, string) {
	idx := strings.LastIndex(src, sep)
	if idx == -1 {
		return src, ""
	}
	return src[:idx], src[idx+len(sep):]
}
