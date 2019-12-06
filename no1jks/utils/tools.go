package utils

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"time"
)

var (
	chars        string
	charsReverse map[rune]int
	salt         string
)

func init() {
	s := os.Getenv("salt")
	c := os.Getenv("chars")
	cr := make(map[rune]int)
	if s == "" {
		s = "L-O-V-E"
	}
	if c == "" {
		c = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"
	}
	salt = s
	chars = c
	for idx, char := range chars {
		cr[char] = idx
	}
	charsReverse = cr
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

type EncodedString struct {
	s string
	S string
}

func MD5(str string) string {
	w := md5.New()
	_, _ = io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

func EncodeSalt(s string) (string, error) {
	if salt == "" || s == "" {
		return "", errors.New("encoded string or salt can't be empty")
	}
	ret := fmt.Sprintf("%s%s", s, salt)
	return MD5(ret), nil
}

func DecodeSalt(s string) (string, bool) {
	ret := strings.Split(s, salt)
	if len(ret) < 1 {
		return "", false
	}
	return ret[0], true
}

func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func divide(a, b int)(int, int){
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// Mask the user's phone.
func EncodeIntString(i int) (s string) {
	base := len(chars)
	var ret []byte
	for true {
		quotient, remainder := divide(i, base)
		i = quotient
		ret = append(ret, chars[remainder])
		if i == 0 {
			break
		}
	}
	return reverseString(string(ret[:]))
}

func DecodeIntString(s string) (i int) {
	arr := []rune(s)
	ret := 0
	base := len(chars)
	for _, char := range arr {
		ret = charsReverse[char] + ret * base
	}
	return ret
}

func TimeAdd(t time.Time, duration string) time.Time{
	d, e := time.ParseDuration(duration)
	if e != nil {
		panic(e)
	}
	return t.Add(d)
}

func Stamp2Str(t int64) string {
	if t == 0 {
		return "时间未知"
	}

	now := time.Now()
	then := time.Unix(t, 0)
	duration := now.Unix() - t
	if then.After(now){
		return "来自未来"
	}else if then.After(TimeAdd(now, "-1m")){
		return "刚刚"
	}else if then.After(TimeAdd(now, "-1h")){
		minute := int(math.Ceil(float64(duration / 60)))
		return  fmt.Sprintf("%d分钟前", minute)
	}else if then.After(TimeAdd(now, "-8h")){
		hour := int(math.Ceil(float64(duration / 60 / 60)))
		return  fmt.Sprintf("%d小时前", hour)
	}else if then.After(TimeAdd(now, "-24h")){
		return "1天前"
	}else {
		return then.Format(DateFormat)
	}
}
