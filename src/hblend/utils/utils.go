package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

func ReadFileBytes(filename string) []byte {

	bytes, err := ioutil.ReadFile(filename)

	if nil != err {
		return []byte{}
	}

	return bytes
}

func WriteFileBytes(filename string, bytes []byte) {

	ioutil.WriteFile(filename, bytes, 0777)
}

func ReadFile(filename string) string {

	bytes := ReadFileBytes(filename)
	return string(bytes)
}

func WriteFile(filename string, text string) {

	EnsureDirs(filename)

	WriteFileBytes(filename, []byte(text))
}

func CopyFile(src, dst string) {

	EnsureDirs(dst)

	sf, err := os.Open(src)
	defer sf.Close()
	if err != nil {
		fmt.Println("I can not open file for read: `"+src+"`:", err)
		return
	}

	df, err := os.Create(dst)
	defer df.Close()
	if err != nil {
		fmt.Println("I can not open file for write: `"+dst+"`:", err)
		return
	}

	_, err = io.Copy(df, sf)
	if err != nil {
		fmt.Println("I can not copy from `"+src+"` to `"+dst+"`:", err)
	}
}

func EnsureDirs(filename string) error {
	dir := path.Dir(filename)

	if err := os.MkdirAll(dir, os.ModeDir|os.ModePerm); nil != err {
		return err
	}

	return nil
}

func CopyFileRemote(src, dst string) error {

	EnsureDirs(dst)

	response, response_err := http.Get(src)
	if response_err != nil {
		return response_err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil
	}

	output, output_err := os.Create(dst)
	if output_err != nil {
		return output_err
	}
	defer output.Close()

	_, copy_err := io.Copy(output, response.Body)
	if copy_err != nil {
		return copy_err
	}

	return nil
}

func CheckFileExists(filename string) bool {

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("File `%s` does not exist\n", filename)
		return false
	}

	return true
}

func Md5File(filename string) string {

	bytes := ReadFileBytes(filename)

	md5_adder := md5.New()
	md5_adder.Write(bytes)

	return hex.EncodeToString(md5_adder.Sum(nil))
}

func Md5String(text string) string {

	bytes := []byte(text)

	md5_adder := md5.New()
	md5_adder.Write(bytes)

	return hex.EncodeToString(md5_adder.Sum(nil))
}
