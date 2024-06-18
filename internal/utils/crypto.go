package utils

import "crypto/aes"

// PKCS7UnPadding 
//  @param encrypt 
//  @return []byte 
func PKCS7UnPadding(encrypt []byte) []byte {
	length := len(encrypt)
	unPadding := int(encrypt[length-1])
	return encrypt[:(length - unPadding)]
}

// DecryptAes128Ecb 
//  @param data 
//  @param key 
//  @return []byte 
func DecryptAes128Ecb(data, key []byte) []byte {
	cipher, _ := aes.NewCipher(key)
	decrypted := make([]byte, len(data))
	size := 16
	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		cipher.Decrypt(decrypted[bs:be], data[bs:be])
	}
	return decrypted
}
