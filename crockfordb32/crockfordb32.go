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

// AppendDecode appends the Crockford base32 decoded form of src to dst and
// returns the extended buffer. If the input is malformed, it returns an error.
func AppendDecode(dst, src []byte) ([]byte, error) {
	return enc.AppendDecode(dst, src)
}

// AppendEncode appends the Crockford base32 encoded form of src to dst and
// returns the extended buffer.
func AppendEncode(dst, src []byte) []byte {
	return enc.AppendEncode(dst, src)
}

// Decode decodes src using Crockford base32 into dst. It writes at most
// DecodedLen(len(src)) bytes to dst and returns the number of bytes written.
// If src contains invalid characters, it returns a CorruptInputError.
func Decode(dst, src []byte) (n int, err error) {
	return enc.Decode(dst, src)
}

// DecodeString decodes a Crockford base32 encoded string and returns the
// decoded bytes. If s contains invalid characters, it returns a
// CorruptInputError.
func DecodeString(s string) ([]byte, error) {
	return enc.DecodeString(s)
}

// DecodedLen returns the maximum length in bytes of the decoded data
// corresponding to n bytes of Crockford base32 encoded data.
func DecodedLen(n int) int {
	return enc.DecodedLen(n)
}

// Encode encodes src using Crockford base32, writing EncodedLen(len(src))
// bytes to dst. The caller must ensure dst is large enough to hold the output.
func Encode(dst, src []byte) {
	enc.Encode(dst, src)
}

// EncodeToString returns the Crockford base32 encoding of src as a string.
func EncodeToString(src []byte) string {
	return enc.EncodeToString(src)
}

// EncodedLen returns the length in bytes of the Crockford base32 encoding of
// an input buffer of length n.
func EncodedLen(n int) int {
	return enc.EncodedLen(n)
}
