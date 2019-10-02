package main

type Frequency struct {
	Rune rune
	Freq int
}

type Frequencies []Frequency

func (s Frequencies) Len() int      { return len(s) }
func (s Frequencies) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type ByRune struct{ Frequencies }

func (s ByRune) Less(i, j int) bool { return s.Frequencies[i].Rune < s.Frequencies[j].Rune }
