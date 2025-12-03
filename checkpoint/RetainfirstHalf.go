package checkpoint

func RetainfirstHalf(str string) string {
	stringLength := len(str)
	if stringLength <= 0 {
		return ""
	}
	if stringLength == 1 {
		return str
	}
	halfString := str[:stringLength/2]
	return halfString
}
