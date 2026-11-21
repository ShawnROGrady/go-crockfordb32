package crockfordb32_test

import (
	"slices"
	"testing"

	"github.com/ShawnROGrady/go-crockfordb32/crockfordb32"
)

// addRoundTripSeeds seeds f with the known-answer inputs plus a few edge cases.
func addRoundTripSeeds(f *testing.F) {
	f.Add([]byte{})
	f.Add([]byte("abc123"))
	f.Add([]byte{0x00, 0x00, 0x00, 0x00, 0x00})
	f.Add([]byte{0xff, 0xff, 0xff, 0xff, 0xff})
	for _, mapping := range basicMappings {
		f.Add([]byte(mapping.input))
	}
}

func FuzzEncodeToStringDecodeStringRoundTrip(f *testing.F) {
	addRoundTripSeeds(f)
	f.Fuzz(func(t *testing.T, in []byte) {
		enc := crockfordb32.EncodeToString(in)
		if len(enc) != crockfordb32.EncodedLen(len(in)) {
			t.Errorf("len(crockfordb32.EncodeToString(%q)) = %d; want %d", in, len(enc), crockfordb32.EncodedLen(len(in)))
		}

		dec, err := crockfordb32.DecodeString(enc)
		if err != nil {
			t.Fatalf("Failed to crockfordb32.DecodeString(%q): %v\n\tin = %q", enc, err, in)
		}

		if !slices.Equal(dec, in) {
			t.Errorf("crockfordb32.DecodeString(%q) = %q; want %q", enc, dec, in)
		}
	})
}

func FuzzEncodeDecodeRoundTrip(f *testing.F) {
	addRoundTripSeeds(f)
	f.Fuzz(func(t *testing.T, in []byte) {
		enc := make([]byte, crockfordb32.EncodedLen(len(in)))
		crockfordb32.Encode(enc, in)

		dec := make([]byte, crockfordb32.DecodedLen(len(enc)))
		n, err := crockfordb32.Decode(dec, enc)
		if err != nil {
			t.Fatalf("Failed to crockfordb32.Decode(dec, %q): %v\n\tin = %q", enc, err, in)
		}

		if !slices.Equal(dec[:n], in) {
			t.Errorf("crockfordb32.Decode(dec, %q) = %q; want %q", enc, dec[:n], in)
		}
	})
}

func FuzzAppendEncodeAppendDecodeRoundTrip(f *testing.F) {
	addRoundTripSeeds(f)
	f.Fuzz(func(t *testing.T, in []byte) {
		const prefix = "prefix:"

		enc := crockfordb32.AppendEncode([]byte(prefix), in)
		if string(enc[:len(prefix)]) != prefix {
			t.Fatalf("crockfordb32.AppendEncode(%q, %q) = %q; want prefix %q preserved", prefix, in, enc, prefix)
		}

		dec, err := crockfordb32.AppendDecode([]byte(prefix), enc[len(prefix):])
		if err != nil {
			t.Fatalf("Failed to crockfordb32.AppendDecode(%q, %q): %v\n\tin = %q", prefix, enc[len(prefix):], err, in)
		}

		if want := append([]byte(prefix), in...); !slices.Equal(dec, want) {
			t.Errorf("crockfordb32.AppendDecode(%q, %q) = %q; want %q", prefix, enc[len(prefix):], dec, want)
		}
	})
}
