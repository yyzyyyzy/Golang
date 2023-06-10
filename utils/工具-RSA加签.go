package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"fmt"
)

func main() {
	text := `{"appId":"05e052baf5514c95b644753886004fff","callTimestamp":"1683619324084","data":{"mobileMd5":"82327d80d3edd6cedbe0ebdbf3ebf9e1"},"nonce":"479141"}`
	privateKey := "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC86UUKAqYLcaEJOsoNK3efG/wokzQ1a+m/mWAyCgB/lgIsFZBE37RuIzW1NPH75G0E2GBwbtL6dhFALSQ/Uab0XDwjdIf4d7LeLdK14Tt9UfqyMOjZ6p2sKoajewQiSpKBl/EhYhn2TG2bWMQuO2NDZ8m2elfE4mnsJvbCrKgJXi2I0BZWz5bqj71JiKZySW4IQNrkD87CkNEQ9cXxwedPp4ZZUD4Kt+bQBc1q055bSco8Nidjm3+9CdWoCQ/52KxFmTcufPkUA85Zhc2s/IWEwZuAZC7Cj2sx4lju1x+ueyPUCE/bD5sw0KFD1vGaCtBSYsRKxDEBu4Px9wiltTxvAgMBAAECggEAHD4ryjFXLcLRNpYJeyqSecEvU8meNpZpnfGKRjDkni6JKJoYtZNUYFr3Pp4px2UBP4Yx9N08waFBNv0IvEay1Tms7AgSA/2xSjJDnFmOEVBeUKGrTlkbNZAuDObpWU8Y1DMpI7qM8xu+D0mYGulaD8vGhneIA7Ft8JZqsfQqMywVVa1yW3hcAGPb3UP4TG7UWwggXIdmFnP1/GFSAcoUtzDl3jRQ5/3YmLmzXgu0TLjcfkm3OlzocPV4j23ucHkAOqmux38gXwuiE9gbdDfUKn0R1lUybDT7zZ7mUE5+7QHd3Lgw9qsgPns03qq4IDJAFt739Aqav+hYdnLVpzVykQKBgQDv2Taa+W+LM6+7E7w44BbsKVOaWcwx6PVle7rTWIgbuY8YjOVCbJ9RXv8bVSOmZwHe9ibYjkV3h3OA03rUgq8cle9Fd3Rd8ph7+j3pk6Wnm5HPiVsSpzbaeRW5IaaiKXwsD/AfvrbbpDOogSI7tLdq6aq5SaKodSeFfEnpD67AGwKBgQDJofC/bgamecq75aSRgyukJ3HkgDzASH39KHPKpXOGQx+qLVjfmb05xTO8zcF3FLOIoyx6rHf6J/NpPRsrQXVQQbhu/jVJtj4yJOgSmRIWpvJV7253QY9u8vD+7WjyRXwY4XXxVE9e4wGHQ8myUYmGgoFqFUl8/L3/bokd3EfCPQKBgQCXOWQ1eDn9EZymHn5ZeejaBACDXETj3xcCYm3cHYDLwkZX21YdeHFHE9dS+25b16yVUKwTdDiWcZ/AxRY38SHJqztOmE+VgITl/lSU9hPHRs47hUYXz7hFLX0l2fK/Ydq5yV7EFTIj8Dbl8m2MKZhP335WDvhwsFU307KiMNIkqwKBgGV3jQOIhug32gH8anifndKZ1wK6VdgdoulG9h5AbIZgOQsWjubXIxZzWrnkgTs/u6lDFBsXt7i3ahLHoWh/JF4i9IFGg+J4R4xMbk0NosCresAsIVb6MwgOUaC4MFvbA7wxmXy8hMSnUBmVz2ZhCegavK150OWI+sQUCARy2dmxAoGBAL1mSQsitJgLGhaqudGAX165bUARdy6Kuv+qqqgp3vXT7Fs2AappcR4hf3IBysWMhVYIars6wvlqHXL0mGwgFhEPn8Vo7S4TR2hN24pvfxX5pGUqHVOy1CjVrQP6rDLrbaF3APu5Np/h6frNml8CXO4ZzaTlznuSd3Jjj62V1mWI"
	sign, _ := sign([]byte(text), privateKey)
	fmt.Println(sign)
}
func sign(data []byte, privateKey string) (string, error) {
	keyBytes, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to decode private key: %w", err)
	}

	parsedKey, err := x509.ParsePKCS8PrivateKey(keyBytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %w", err)
	}

	privKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return "", fmt.Errorf("unexpected private key type: %T", parsedKey)
	}

	hasher := sha1.New()
	_, err = hasher.Write(data)
	if err != nil {
		return "", fmt.Errorf("failed to hash data: %w", err)
	}

	sig, err := rsa.SignPKCS1v15(rand.Reader, privKey, crypto.SHA1, hasher.Sum(nil))
	if err != nil {
		return "", fmt.Errorf("failed to sign data: %w", err)
	}

	return base64.StdEncoding.EncodeToString(sig), nil
}
