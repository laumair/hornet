package hornet

import (
	"encoding/hex"
	"fmt"

	iotago "github.com/iotaledger/iota.go"
)

var (
	// NullMessageID is the ID of the genesis message.
	NullMessageID = Hash(make([]byte, iotago.MessageHashLength))
)

// Hash is the binary representation of a Hash.
type Hash []byte

// Hex converts the binary Hash to its hex string representation.
func (h Hash) Hex() string {
	if len(h) == iotago.MessageHashLength {
		return hex.EncodeToString(h)
	}

	panic(fmt.Sprintf("Unknown hash length (%d)", len(h)))
}

// ID converts the binary Hash to an array representation.
func (h Hash) ID() (id [iotago.MessageHashLength]byte) {
	if len(h) == iotago.MessageHashLength {
		copy(id[:], h[:iotago.MessageHashLength])
		return
	}

	panic(fmt.Sprintf("Unknown hash length (%d)", len(h)))
}

// Hashes is a slice of Hash.
type Hashes []Hash

// Hex converts the binary Hashes to their hex string representation.
func (h Hashes) Hex() []string {
	var results []string
	for _, hash := range h {
		results = append(results, hash.Hex())
	}
	return results
}
