package match

import "strings"

func ParseCmd(command string) ([]string, string) {
	split := strings.Split(command, " ")
	return split, split[0]
}
