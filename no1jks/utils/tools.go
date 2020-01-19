package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io"
	"math"
	"math/big"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"

	mathRand "math/rand"
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

func CreateRandomString(len int) string  {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0;i < len ;i++  {
		randomInt,_ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

var AllowExtMap = map[string]bool{
	".jpg":true,
	".jpeg":true,
	".png":true,
	".webp":true,
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	//if os.IsNotExist(err) {
	//	return false, nil
	//}
	return false, err
}

func HashFileName(base string, ext string) (fileName string) {
	mathRand.Seed(time.Now().UnixNano())
	randNum := fmt.Sprintf("%d", mathRand.Intn(9999) + 1000)
	hashName := md5.Sum([]byte(time.Now().Format("2006_01_02_15_04_05_") + randNum ))

	name := fmt.Sprintf("%x",hashName) + ext
	return base + name
}

func UploadTo(fileHead *multipart.FileHeader, fileType string) (string, string, *ServiceErr) {
	ext := path.Ext(fileHead.Filename)
	if _, ok := AllowExtMap[ext]; !ok{
		logs.Info("Extension not allow" )
		return "", "", Errs[""]
	}

	baseDir := "/usr/share/nginx/html/files"
	if name, _ := os.Hostname(); name == "dahuadeMacBook-Pro.local"{
		baseDir = "/tmp/"
	}
	uploadDir := path.Join(baseDir, fileType)
	exist, _ := PathExists(uploadDir)
	if !exist {
		osErr := os.MkdirAll(uploadDir , 777)
		if osErr != nil{
			logs.Info("Make dir error", osErr, uploadDir)
			return "", "", Errs[""]
		}
	}
	fileName := HashFileName("", ext)

	uploadPath :=  path.Join(uploadDir, fileName)
	visitPath := path.Join("/static/", fileType, fileName)
	return uploadPath, visitPath, nil
}
