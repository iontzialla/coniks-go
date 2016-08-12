package sign

import (
	"testing"
)

// copied from official crypto.ed25519 tests
func TestVerifySignature(t *testing.T) {
	key, err := GenerateKey(nil)
	if err != nil {
		t.Fatal(err)
	}

	message := []byte("test message")
	sig := key.Sign(message)

	pk, ok := key.Public()
	if !ok {
		t.Errorf("bad PK?")
	}

	if !pk.Verify(message, sig) {
		t.Errorf("valid signature rejected")
	}

	wrongMessage := []byte("wrong message")
	if pk.Verify(wrongMessage, sig) {
		t.Errorf("signature of different message accepted")
	}
}
