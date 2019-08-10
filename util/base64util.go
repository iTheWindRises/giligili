package util

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"strings"
)

// 读取base64字符串，解码，然后生成文件
func Base64ToFile(sourcePath, targetPath string) error {
	data, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		return err
	}
	str := string(data)
	str = strings.ReplaceAll(str, `\n`, "")
	decodeData, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(targetPath, decodeData, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
	// f, err := os.OpenFile("test/test.jpg", os.O_RDWR|os.O_CREATE, os.ModePerm)
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// f.Write(decodeData)
}

// 读取base64字符串，解码，然后生成文件
func FileToBase64(sourcePath, targetPath string) error {
	data, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		return err
	}
	str := base64.StdEncoding.EncodeToString(data)

	err = ioutil.WriteFile(targetPath, []byte(str), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
