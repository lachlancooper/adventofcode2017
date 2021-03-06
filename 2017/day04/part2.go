// http://adventofcode.com/2017/day/4
// part 2
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// isvalid returns true if s is a valid passphrase
// (i.e. contains no duplicate words)
// exact anagrams now count as duplicates
func isvalid(s string) bool {
	words := strings.Fields(s)
	m := make(map[string]bool)

	for _, w := range words {
		a := strings.Split(w, "")
		sort.Strings(a)
		w = strings.Join(a, "")

		exists := m[w]
		if exists {
			return false
		}
		m[w] = true
	}

	return true
}

func main() {
	valid, invalid := 0, 0
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		pp := scanner.Text()
		fmt.Printf("%v ", pp)
		if isvalid(pp) {
			fmt.Println("is valid.")
			valid += 1
		} else {
			fmt.Println("is not valid.")
			invalid += 1
		}
	}
	fmt.Println(valid, "valid passphrases.", invalid, "invalid passphrases.")
}
