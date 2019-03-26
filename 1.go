package go_Tools

func InArray(arr []interface{}, str string) bool {
	for _, value := range arr {
		if value == str {
			return true
		}
	}
	return false
}
