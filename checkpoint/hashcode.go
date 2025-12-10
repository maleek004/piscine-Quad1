package checkpoint

import "strings"

// you can use bytes with int in mathematical operations without any fear

func HashCode(dec string) string {

	strSize := len(dec)
	var hashedString strings.Builder

	for _, r := range dec{
		hashValue := (int(r) + strSize) % 127
		if hashValue < 32 || hashValue == 127 { 
            hashValue += 33
        }
		hashedString.WriteByte(byte(hashValue))
	}
	return hashedString.String()
}
