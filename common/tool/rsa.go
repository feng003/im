package tool

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

type RSAUtils struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

// 初始化 RSAUtils，加载公钥或私钥
func NewRSAPrivateUtils(privateKeyPath string) (*RSAUtils, error) {
	privateKey, err := loadPrivateKey(privateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load private key: %v", err)
	}
	return &RSAUtils{privateKey: privateKey}, nil
}

func NewRSAPublicUtils(publicKeyPath string) (*RSAUtils, error) {
	publicKey, err := loadPublicKey(publicKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load public key: %v", err)
	}
	return &RSAUtils{publicKey: publicKey}, nil
}

// 从文件加载私钥
func loadPrivateKey(path string) (*rsa.PrivateKey, error) {
	privateData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading private key file: %v", err)
	}
	block, _ := pem.Decode(privateData)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("invalid private key format")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

// 从文件加载公钥
func loadPublicKey(path string) (*rsa.PublicKey, error) {
	pubData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading public key file: %v", err)
	}
	block, _ := pem.Decode(pubData)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("invalid public key format")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing public key: %v", err)
	}
	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not a valid RSA public key")
	}
	return rsaPub, nil
}

// 分段加密
func (r *RSAUtils) Encrypt(data []byte) ([]byte, error) {
	if r.publicKey == nil {
		return nil, errors.New("public key not loaded")
	}
	maxSize := r.publicKey.Size() - 11 // RSA 最大加密块大小
	var buffer bytes.Buffer
	for start := 0; start < len(data); start += maxSize {
		end := start + maxSize
		if end > len(data) {
			end = len(data)
		}
		encryptedBlock, err := rsa.EncryptPKCS1v15(rand.Reader, r.publicKey, data[start:end])
		if err != nil {
			return nil, fmt.Errorf("encryption error: %v", err)
		}
		buffer.Write(encryptedBlock)
	}
	return buffer.Bytes(), nil
}

// 分段解密
func (r *RSAUtils) Decrypt(data []byte) ([]byte, error) {
	if r.privateKey == nil {
		return nil, errors.New("private key not loaded")
	}
	blockSize := r.privateKey.Size() // RSA 最大解密块大小
	var buffer bytes.Buffer
	for start := 0; start < len(data); start += blockSize {
		end := start + blockSize
		if end > len(data) {
			end = len(data)
		}
		decryptedBlock, err := rsa.DecryptPKCS1v15(rand.Reader, r.privateKey, data[start:end])
		if err != nil {
			return nil, fmt.Errorf("decryption error: %v", err)
		}
		buffer.Write(decryptedBlock)
	}
	return buffer.Bytes(), nil
}
