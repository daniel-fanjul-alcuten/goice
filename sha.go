package ice

import (
	"crypto/sha256"
	"encoding/hex"
)

const ShaSize = sha256.Size

type Sha [ShaSize]byte

func (s *Sha) Data() []byte {
	return (*s)[:]
}

func (s *Sha) String() string {
	return hex.EncodeToString(s.Data())
}
