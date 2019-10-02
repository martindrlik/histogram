package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

var (
	rd   = bufio.NewReader(os.Stdin)
	freq = map[rune]int{}
)

func main() {
	for {
		r, _, err := rd.ReadRune()
		if err != nil && err != io.EOF {
			fmt.Println(err)
			os.Exit(1)
		}
		if r != 0 {
			freq[r]++
		}
		if err != nil {
			break
		}
	}
	s := make([]Frequency, 0, len(freq))
	for r, f := range freq {
		s = append(s, Frequency{r, f})
	}
	sort.Sort(ByRune{s})
	for _, f := range s {
		fmt.Println(f.Rune, f.Freq)
	}
}
