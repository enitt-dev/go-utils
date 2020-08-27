package convert

import (
	"encoding/hex"
	"log"
	"strings"
)

func BytesToHexStrings(b []byte) []string {
	var result []string
	s := hex.EncodeToString(b)

	for i := 0; i < len(s); i = i + 2 {
		result = append(result, s[i:i+2])
	}
	return result
}

func StringsToBytes(s []string) []byte {
	var result []byte

	str := strings.Join(s, "")

	result, err := hex.DecodeString(str)

	if err != nil {
		log.Fatal(err)
	}
	return result
}
