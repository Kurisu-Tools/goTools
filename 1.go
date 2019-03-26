package go_Tools

func InStringArray(arr *[]string, str string) bool {
	for _, value := range *arr {
		if value == str {
			return true
		}
	}
	return false
}

func InIntArray(arr *[]int, str int) bool {
	for _, value := range *arr {
		if value == str {
			return true
		}
	}
	return false
}
