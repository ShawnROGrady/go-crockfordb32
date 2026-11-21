package crockfordb32_test

import (
	"fmt"

	"github.com/ShawnROGrady/go-crockfordb32/crockfordb32"
)

func Example() {
	msg := "Hello, 世界"
	encoded := crockfordb32.EncodeToString([]byte(msg))
	fmt.Println(encoded)
	decoded, err := crockfordb32.DecodeString(encoded)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	fmt.Println(string(decoded))
	// Output:
	// 91jprv3f5gge9e4pwyarr
	// Hello, 世界
}
