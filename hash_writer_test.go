package ice

import (
	"bytes"
	"testing"
)

func TestHashWriter(t *testing.T) {
	str, buffer := "foo\n", &bytes.Buffer{}
	writer := NewHashWriter(buffer, 1)
	n, err := writer.Write([]byte(str))
	if n != len(str) {
		t.Error("Unexpected n", n)
	}
	if err != nil {
		t.Error("Unexpected err", err)
	}
	err = writer.Close()
	if err != nil {
		t.Error("Unexpected err", err)
	}
	sha := writer.Sha()
	if sha == nil {
		t.Error("Unexpected sha", sha)
	} else if sha.String() != "b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c" {
		t.Error("Unexpected sha.String()", sha.String())
	}
}
