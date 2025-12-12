package checkpoint

import "strings"

// you can use bytes with int in mathematical operations without any fear

// func HashCode(dec string) string {

// 	strSize := len(dec)
// 	var hashedString strings.Builder

// 	for _, r := range dec{
// 		hashValue := (int(r) + strSize) % 127
// 		if hashValue < 32 || hashValue == 127 {
//             hashValue += 33
//         }
// 		hashedString.WriteByte(byte(hashValue))
// 	}
// 	return hashedString.String()
// }

func HashCode(dec string) string {
	var res strings.Builder
	strSize := len(dec)
	for _, r := range dec {
		ascii := int(r)
		hash := (ascii + strSize) % 127
		if hash < 32 {
			hash += 33
		}
		res.WriteByte(byte(hash))
	}
	return res.String()
}
