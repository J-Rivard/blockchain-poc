package models

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(input string) string {
	sum := sha256.Sum256([]byte(input))

	return hex.EncodeToString(sum[:])
}

func IncrementString(input string) string {
	byteArray := []byte(input)
	length := len(byteArray)

	index := length - 1

	for index >= 0 {
		if byteArray[index] < 255 {
			byteArray[index]++
			break
		} else {
			byteArray[index] = 0
			index--
		}
	}

	return string(byteArray)
}
