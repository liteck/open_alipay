package open_alipay

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net/url"
	"sort"
	"strings"
)

func doSign(params url.Values, privateKey []byte) (sign string, err error) {
	//对key进行升序排序.
	keys := make([]string, 0)
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var str string
	//对key=value的键值对用&连接起来，略过空值
	for _, k := range keys {
		value := params.Get(k)
		if value != "" {
			str = str + k + "=" + value + "&"
		}
	}
	str = strings.TrimRight(str, "&")
	Trace.Println(str)

	//签名
	block, _ := pem.Decode(privateKey)
	if block == nil {
		err = errors.New("私钥错误")
		return
	}

	var private *rsa.PrivateKey
	if private, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
		return
	}

	_t := crypto.SHA1.New()
	_t.Write([]byte(str))
	digest := _t.Sum(nil)
	var data []byte
	if data, err = rsa.SignPKCS1v15(rand.Reader, private, crypto.SHA1, digest); err != nil {
		return "", err
	}
	sign = base64.StdEncoding.EncodeToString(data)
	return
}

func doVerifySign(s, sign, respKey string, publicKey []byte) (signData string, pass bool) {
	idx := strings.Index(s, ",\"sign\"")
	if idx == -1 {
		return
	}
	signData = s[4+len(respKey) : idx]
	Trace.Println("verify signed data====" + signData)

	block, _ := pem.Decode(publicKey)
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Printf("Failed to parse RSA public key: %s\n", err)
		return
	}
	rsaPub, _ := pub.(*rsa.PublicKey)
	t := sha1.New()
	_, _ = io.WriteString(t, signData)
	digest := t.Sum(nil)
	data, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		fmt.Println("DecodeString sig error, reason: ", err)
		return
	}
	err = rsa.VerifyPKCS1v15(rsaPub, crypto.SHA1, digest, data)
	if err != nil {
		fmt.Println("Verify sig error, reason: ", err)
		return
	}

	pass = true
	return
}
