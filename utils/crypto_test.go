package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncrypt(t *testing.T) {
	a := "user:passport@tcp(114.55.139.105:3306)/brain?charset=utf8mb4&"
	s, err := Encrypt(a)
	assert.Nil(t, err)
	t.Logf(s)

	a = "root:Bit0123456789!@tcp(10.233.146.47:16315)/brain?charset=utf8mb4"
	s, err = Encrypt(a)
	assert.Nil(t, err)
	t.Logf(s)
}

func TestDecrypt(t *testing.T) {
	a := "wdzHhOX/SSdWWziV4TDy0AYqXfr0dwPoVWNGPbgg26gLOoV0731EyR/b49lfJSSf6dnK0C9s5Il4QyRmaFsNTc6XOtu1ApToSaYGns+OVasYdbGpKsbRqyYRroZ0sirBC8VEyx8FbcWlXQ=="
	s, err := Decrypt(a)
	assert.Nil(t, err)
	t.Logf("return %#v", s)
}

func TestEmptyString(t *testing.T) {
	a := ""
	s, err := Encrypt(a)
	if err != nil {
		t.Fatalf("ecrypt with error: %v", err.Error())
	}
	aa, err := Decrypt(s)
	if err != nil {
		t.Fatalf("decrypt with error: %v", err.Error())
	}
	assert.Equal(t, aa, a)
}
