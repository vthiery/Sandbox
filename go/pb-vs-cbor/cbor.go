package bench

import (
	"encoding/hex"

	"github.com/ugorji/go/codec"
)

// encodeToHexCBOR encodes the input into CBOR with an hexadecimal representation
func encodeToHexCBOR(input interface{}) ([]byte, error) {
	b := make([]byte, 0, 1024)
	encoder := codec.NewEncoderBytes(&b, &codec.CborHandle{})
	if err := encoder.Encode(input); err != nil {
		return nil, err
	}

	res := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(res, b)
	return res, nil
}
