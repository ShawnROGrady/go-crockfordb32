// Package crockfordb32 implements a base32 encoding using the alphabet
// designed by Douglas Crockford.
//
// With the exception of the `Encoding` function, all of the exported functions
// in this package correspond to the methods on the base32 encoding.
package crockfordb32

import "encoding/base32"

const alphabet = "0123456789abcdefghjkmnpqrstvwxyz"

// Encoding returns a base32 encoding using the Crockford alphabet.
func Encoding() *base32.Encoding {
	return base32.NewEncoding(alphabet).WithPadding(base32.NoPadding)
}

var enc = Encoding()

type CorruptInputError = base32.CorruptInputError

func AppendDecode(dst, src []byte) ([]byte, error) {
	return enc.AppendDecode(dst, src)
}

func AppendEncode(dst, src []byte) []byte {
	return enc.AppendEncode(dst, src)
}

func Decode(dst, src []byte) (n int, err error) {
	return enc.Decode(dst, src)
}

func DecodeString(s string) ([]byte, error) {
	return enc.DecodeString(s)
}

func DecodedLen(n int) int {
	return enc.DecodedLen(n)
}

func Encode(dst, src []byte) {
	enc.Encode(dst, src)
}

func EncodeToString(src []byte) string {
	return enc.EncodeToString(src)
}

func EncodedLen(n int) int {
	return enc.EncodedLen(n)
}
