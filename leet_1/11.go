package leet_1

var mapping = []struct {
	num int
	val string
}{
	{1, "I"},
	{4, "IV"},
	{5, "V"},
	{9, "IX"},
	{10, "X"},
	{40, "XL"},
	{50, "L"},
	{90, "XC"},
	{100, "C"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
}

func intToRoman(num int) string {
	mappingIndex := len(mapping) - 1
	result := ""
	for num > 0 {
		if num/mapping[mappingIndex].num > 0 {
			result += mapping[mappingIndex].val
			num -= mapping[mappingIndex].num
		} else {
			mappingIndex--
		}
	}
	return result
}
