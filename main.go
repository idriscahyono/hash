package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	salt := "salt"
	text := "Hello"

	res := password(text, salt)
	fmt.Println(res)
}

func password(text, salt string) string {
	return SHA256(fmt.Sprintf("%s%s", text, salt))
}

func SHA256(msg string) string {
	h := sha256.New()
	h.Write([]byte(msg))
	return hex.EncodeToString(h.Sum(nil))
}
