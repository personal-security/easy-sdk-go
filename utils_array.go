package easysdk

func ReverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(ReverseInts(input[1:]), input[0])
}

func ReverseStrings(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(ReverseStrings(input[1:]), input[0])
}

func ReverseFloat64(input []float64) []float64 {
	if len(input) == 0 {
		return input
	}
	return append(ReverseFloat64(input[1:]), input[0])
}

func SliceHaveString(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
