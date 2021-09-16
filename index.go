package piscine

import (
	"fmt"
	"strings"
)

func Index(s string, toFind string) int {
	mot := s
	toFindword := toFind

	result := strings.Index(mot, toFindword)
	fmt.Println(result)
	return result
}
