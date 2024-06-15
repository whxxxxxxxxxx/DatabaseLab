package encrypt

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"errors"
)

//AES堆成加密

var Encrypt *Encryption

type Encryption struct {
	key string
}

func init() {
	Encrypt = NewEncryption()
}

func NewEncryption() *Encryption {
	return &Encryption{}
}

func PadPwd(srcByte []byte, blockSize int) []byte {
	padNum := blockSize - len(srcByte)%blockSize
	ret := bytes.Repeat([]byte{byte(padNum)}, padNum)
	srcByte = append(srcByte, ret...)
	return srcByte
}

func (k *Encryption) AesEncoding(src string) string {
	srcByte := []byte(src)
	block, err := aes.NewCipher([]byte(k.key))
	if err != nil {
		return src
	}
	NewSrcByte := PadPwd(srcByte, block.BlockSize())
	dst := make([]byte, len(NewSrcByte))
	block.Encrypt(dst, NewSrcByte)
	pwd := base64.StdEncoding.EncodeToString(dst)
	return pwd

}

func UnPadPwd(dst []byte) ([]byte, error) {
	if len(dst) <= 0 {
		return dst, errors.New("长度有误")
	}
	unPadNum := int(dst[len(dst)-1])
	strErr := "error"
	op := []byte(strErr)
	if len(dst) < unPadNum {
		return op, nil
	}
	str := dst[:len(dst)-unPadNum]
	return str, nil
}

func (k *Encryption) AesDecoding(pwd string) string {
	pwdByte := []byte(pwd)
	var err error
	pwdByte, err = base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		return pwd
	}
	block, errBlock := aes.NewCipher([]byte(k.key))
	if errBlock != nil {
		return pwd
	}
	dst := make([]byte, len(pwdByte))
	block.Decrypt(dst, pwdByte)
	dst, err = UnPadPwd(dst)
	if err != nil {
		return "0"
	}
	return string(dst)
}

func (k *Encryption) SetKey(key string) {
	k.key = key
}
