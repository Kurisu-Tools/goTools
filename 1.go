package go_Tools

import "log"

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

func LogDebugInfo(info string) {
	log.Printf("DEBUG:\x1b[1;4;35m %s  \x1b[0m", info)
}
