package input

import (
	"os/exec"
	"strings"
)

func InputFilter(s string, vocabulary []string) []string {
	split := strings.Fields(s)
	if len(split) > 0 && contains(split[0], vocabulary) {
		return split
	}
	return []string{}
}

func RunCommand(filtered_input []string) string {
	var cmd *exec.Cmd

	if len(filtered_input) > 1 {
		args := filtered_input[1:]
		cmd = exec.Command(filtered_input[0], args...)
	} else {
		cmd = exec.Command(filtered_input[0])
	}
	output, _ := cmd.CombinedOutput()
	return string(output)
}

func contains(s string, arr []string) bool {
	for _, val := range arr {
		if s == val {
			return true
		}
	}
	return false
}
