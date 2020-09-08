package utils

import (
	"strings"
)

func filterPlayer(p string, fl []string) bool {
	for _, f := range fl {
		if strings.HasPrefix(p, f) {
			return true
		}
	}

	return false
}

func FilterStrings(pl []string, fl []string) []string {
	out := []string{}

	for _, p := range pl {
		if !filterPlayer(p, fl) {
			out = append(out, p)
		}
	}

	return out
}
