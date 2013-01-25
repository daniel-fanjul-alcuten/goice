package ice

import (
	"crypto/sha256"
	"testing"
)

func TestShaSize(t *testing.T) {
	if ShaSize != 32 {
		t.Error("ShaSize", ShaSize)
	}
}

func TestShaString(t *testing.T) {
	sha, h := Sha{}, sha256.New()
	h.Write([]byte("foo\n"))
	h.Sum(sha[:0])
	str := sha.String()
	if str != "b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c" {
		t.Error("Unexpected String()", str)
	}
}
