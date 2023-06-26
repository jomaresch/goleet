package leet_1

// https://leetcode.com/problems/zigzag-conversion/
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	a := make([]string, numRows)
	direction := -1
	row := 0
	for _, char := range s {
		if row == 0 || row >= numRows-1 {
			direction *= -1
		}
		a[row] += string(char)
		row += direction
	}

	result := ""
	for _, builder := range a {
		result += builder
	}
	return result
}
