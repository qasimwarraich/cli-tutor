package input

import "strings"

func InputFilter(s string, vocabulary []string) []string {
	split := strings.Fields(s)
    if len(split) > 0 && contains(split[0], vocabulary) {
		return split
	}
	return []string{}
}

func contains(s string, arr []string) bool {
	for _, val := range arr {
		if s == val {
			return true
		}
	}
	return false
}
