package points

// countAlpha helper function for counting alphanumeric characters
func countAlpha(str string) int64 {
	bytes := []byte(str)

	total := int64(0)
	for _, by := range bytes {
		if ('a' <= by && by <= 'z') || ('A' <= by && by <= 'Z') || ('0' <= by && by <= '9') {
			total++
		}
	}
	return total
}
