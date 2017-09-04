package str

import (
	"strings"
)

func Str(s string) (string, string, []string, []string) {
	s1 := strings.TrimSpace(s)
	s2 := strings.Trim(s, "1")
	s3 := strings.Fields(s)
	s4 := strings.Split(s, "1")

	return s1, s2, s3, s4


}

