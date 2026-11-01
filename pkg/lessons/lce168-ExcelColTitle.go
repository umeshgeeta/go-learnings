package lessons

func ConvertToTitle(columnNumber int) string {
	var result string
	crtSlice := make([]int, 0)
	crtSlice = baseTo26(crtSlice, columnNumber)
	for crtSlice[0] > 26 {
		crtSlice = baseTo26(crtSlice[1:], crtSlice[0])
	}
	for _, val := range crtSlice {
		r := 'A' + rune(val-1)
		result += string(r)
	}
	return result
}

func baseTo26(crtSlice []int, val int) []int {
	var quotient, reminder int
	quotient = val / 26
	reminder = val % 26
	if reminder > 0 {
		crtSlice = append([]int{reminder}, crtSlice...)
	} else {
		crtSlice = append([]int{26}, crtSlice...)
		quotient--
	}
	if quotient > 0 {
		crtSlice = append([]int{quotient}, crtSlice...)
	}
	return crtSlice
}
