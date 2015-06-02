package crypto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	token := "thisisasecuretoken"
	encrypted := Encrypt(token)
	decrypted, err := Decrypt(encrypted)

	assert.Equal(t, nil, err, fmt.Sprintf("there was an error decrypting our token: %v", err))
	assert.Equal(t, token, decrypted, fmt.Sprintf("encrypting/decrypting produced non-matching results: %v", err))
}
