package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("jbpnbwwd"))
}

func formatRuneMap(m map[rune]struct{}) string {
	if len(m) == 0 {
		return "map[]"
	}

	// Собираем ключи как строки для сортировки
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, string(k))
	}

	return fmt.Sprintf("map[%s]", strings.Join(keys, " "))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}

	var (
		maxLength                   = 0
		length                      = 0
		m         map[rune]struct{} = make(map[rune]struct{})
	)

	for i := 0; i <= len(s); i++ {
		for j, r := range s[i:] {
			fmt.Printf("i=%-3d | r=%-3d | m=%-30s | m(r)=%-30s | length=%-3d | maxLength=%-3d\n",
				j, r, fmt.Sprint(m), formatRuneMap(m), length, maxLength)
			if _, exists := m[r]; exists {
				if maxLength < length {
					maxLength = length
				}

				length = 0
				m = make(map[rune]struct{})

				break
			} else {
				m[r] = struct{}{}
				length = length + 1
			}

		}
		fmt.Println()
	}

	return maxLength
}
