package encrypt

import (
	"encoding/hex"
	"strings"

	"github.com/gogf/gf/crypto/gaes"
	"github.com/yunbyte/gf-contrib/v2/consts"
)

// EncryptStringAES sensitive text encryption
func MustEncryptAES(plainText, key, iv string) string {
	encrypted, err := gaes.EncryptCBC([]byte(plainText), []byte(key), []byte(iv))
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(encrypted)
}

// DecryptStringAES sensitive text decryption
func MustDecryptAES(cipherText, key, iv string) string {
	if !strings.HasPrefix(cipherText, consts.EncryptAESPrefix) {
		return cipherText
	}
	cipherText = cipherText[4:]
	encrypted, err := hex.DecodeString(cipherText)
	if err != nil {
		panic(err)
	}
	decrypted, err := gaes.DecryptCBC(encrypted, []byte(key), []byte(iv))
	if err != nil {
		panic(err)
	}
	return string(decrypted)
}
