package imagex

import (
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

func UploadImg(file multipart.File) (string, error) {
	path := "./assets/img/" // 指定保存文件的目录
	Unix := time.Now().Unix()
	UnixString := strconv.FormatInt(Unix, 10)
	avatarPath := path + UnixString + ".jpg"

	// 创建文件
	out, err := os.Create(avatarPath)
	if err != nil {
		return "", err // 文件创建失败，返回错误
	}
	defer out.Close() // 确保在函数退出前关闭文件

	// 将上传的文件内容复制到新文件
	_, err = io.Copy(out, file)
	if err != nil {
		return "", err // 文件复制失败，返回错误
	}

	return UnixString + ".jpg", nil // 返回保存的文件路径
}

func GetImg(name string) (*os.File, error) {
	path := "./assets/img/" + name
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}
