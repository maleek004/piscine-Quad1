package checkpoint

import (
	"fmt"
	//"strconv"
	"strings"
)

func FromTo(from int, to int) string {
	if from < 0 || from > 99 {
		return "Invalid\n"
	}
	if to < 0 || to > 99 {
		return "Invalid\n"
	}
	if from == to {
		return fmt.Sprintf("%02d", from) + "\n"
	}
	var res strings.Builder
	if from > to {
		for from >= to {
			if from == to {
				res.WriteString(fmt.Sprintf("%02d", from) + "\n")
			} else {
				res.WriteString(fmt.Sprintf("%02d", from))
				res.WriteString(", ")
			}
			from--
		}
	} else {
		for from <= to {
			if from == to {
				res.WriteString(fmt.Sprintf("%02d", from) + "\n")
			} else {
				res.WriteString(fmt.Sprintf("%02d", from))
				res.WriteString(", ")
			}
			from++
		}
	}
	return res.String()
}
