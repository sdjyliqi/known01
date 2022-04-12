package crypt

import (
	"crypto/sha1"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAESGCM(t *testing.T) {
	ostr := fmt.Sprintf("%v.%v", time.Now().Format(time.RFC1123), os.Getpid())
	hash := sha1.Sum([]byte(ostr))
	okey := hash[3:19]

	t.Logf("generate string: %v", ostr)

	encrypted, err := AESGCMEncryptString(ostr, okey)
	if err != nil {
		t.Fatalf("encrypt error: %v", err)
	}

	t.Logf("encrypted value: %v", encrypted)

	decrypted, err := AESGCMDecryptString(encrypted, okey)
	if err != nil {
		t.Fatalf("decrypt error: %v", err)
	}

	t.Logf("decrypted value: %v", decrypted)

	assert.Equal(t, ostr, decrypted)
}
