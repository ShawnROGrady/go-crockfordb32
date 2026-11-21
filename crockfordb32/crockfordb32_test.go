package crockfordb32_test

import (
	"errors"
	"slices"
	"testing"

	"github.com/ShawnROGrady/go-crockfordb32/crockfordb32"
)

var basicMappings = []struct {
	input string
	enc   string
}{
	{input: "f", enc: "cr"},
	{input: "fo", enc: "csqg"},
	{input: "foo", enc: "csqpy"},
	{input: "foob", enc: "csqpyrg"},
	{input: "fooba", enc: "csqpyrk1"},
	{input: "foobar", enc: "csqpyrk1e8"},
	{input: "test", enc: "ehjq6x0"},
}

func TestEncodeToString(t *testing.T) {
	for _, mapping := range basicMappings {
		var (
			in   = []byte(mapping.input)
			want = mapping.enc
		)

		got := crockfordb32.EncodeToString(in)
		if got != want {
			t.Errorf("crockfordb32.EncodeToString(%q) = %q; want %q", in, got, want)
		}
	}
}

func TestEncode(t *testing.T) {
	for _, mapping := range basicMappings {
		var (
			in   = []byte(mapping.input)
			want = []byte(mapping.enc)

			got = make([]byte, crockfordb32.EncodedLen(len(in)))
		)

		crockfordb32.Encode(got, in)
		if !slices.Equal(got, want) {
			t.Errorf("crockfordb32.Encode(got, %q) = %q; want %q", in, got, want)
		}
	}
}

func TestDecodeString(t *testing.T) {
	for _, mapping := range basicMappings {
		var (
			in   = mapping.enc
			want = []byte(mapping.input)
		)

		got, err := crockfordb32.DecodeString(in)
		if err != nil {
			t.Errorf("Failed to crockfordb32.DecodeString(%q): %v", in, err)
			continue
		}

		if !slices.Equal(got, want) {
			t.Errorf("crockfordb32.DecodeString(%q) = %q; want %q", in, got, want)
		}
	}
}

func TestDecode(t *testing.T) {
	for _, mapping := range basicMappings {
		var (
			in   = []byte(mapping.enc)
			want = []byte(mapping.input)

			got = make([]byte, crockfordb32.DecodedLen(len(in)))
		)

		n, err := crockfordb32.Decode(got, in)
		if err != nil {
			t.Errorf("Failed to crockfordb32.Decode(got, %q): %v", in, err)
			continue
		}

		got = slices.Delete(got, n, len(got))
		if !slices.Equal(got, want) {
			t.Errorf("crockfordb32.Decode(got, %q) = %q; want %q", in, got, want)
		}
	}
}

func TestAppendEncode(t *testing.T) {
	for _, prefix := range []string{"", "prefix:"} {
		for _, mapping := range basicMappings {
			var (
				dst  = []byte(prefix)
				in   = []byte(mapping.input)
				want = []byte(prefix + mapping.enc)
			)

			got := crockfordb32.AppendEncode(dst, in)
			if !slices.Equal(got, want) {
				t.Errorf("crockfordb32.AppendEncode(%q, %q) = %q; want %q", prefix, in, got, want)
			}
		}
	}
}

func TestAppendDecode(t *testing.T) {
	for _, prefix := range []string{"", "prefix:"} {
		for _, mapping := range basicMappings {
			var (
				dst  = []byte(prefix)
				in   = []byte(mapping.enc)
				want = []byte(prefix + mapping.input)
			)

			got, err := crockfordb32.AppendDecode(dst, in)
			if err != nil {
				t.Errorf("Failed to crockfordb32.AppendDecode(%q, %q): %v", prefix, in, err)
				continue
			}

			if !slices.Equal(got, want) {
				t.Errorf("crockfordb32.AppendDecode(%q, %q) = %q; want %q", prefix, in, got, want)
			}
		}
	}
}

func TestDecodeStringCorruptInput(t *testing.T) {
	in := "foo bar"
	dec, err := crockfordb32.DecodeString(in)
	if err == nil {
		t.Fatalf("Unexpectedly no error returned by crockfordb32.DecodeString(in); dec = %q", dec)
	}

	var errCorruptInput crockfordb32.CorruptInputError
	if !errors.As(err, &errCorruptInput) {
		t.Errorf("errors.As(err, &errCorruptInput) = false; want true\n\terr [%T] = %v", err, err)
	}
}

func TestDecodeCorruptInput(t *testing.T) {
	in := "foo bar"

	dec := make([]byte, crockfordb32.DecodedLen(len(in)))
	n, err := crockfordb32.Decode(dec, []byte(in))
	if err == nil {
		t.Fatalf("Unexpectedly no error returned by crockfordb32.Decode(dec, []byte(in)); dec = %q; n = %d", dec, n)
	}

	var errCorruptInput crockfordb32.CorruptInputError
	if !errors.As(err, &errCorruptInput) {
		t.Errorf("errors.As(err, &errCorruptInput) = false; want true\n\terr [%T] = %v", err, err)
	}
}

func TestAppendDecodeCorruptInput(t *testing.T) {
	var (
		prefix = "prefix:"
		in     = []byte("foo bar")
	)

	got, err := crockfordb32.AppendDecode([]byte(prefix), in)
	if err == nil {
		t.Fatalf("Unexpectedly no error returned by crockfordb32.AppendDecode(%q, %q); got = %q", prefix, in, got)
	}

	var errCorruptInput crockfordb32.CorruptInputError
	if !errors.As(err, &errCorruptInput) {
		t.Errorf("errors.As(err, &errCorruptInput) = false; want true\n\terr [%T] = %v", err, err)
	}
}
