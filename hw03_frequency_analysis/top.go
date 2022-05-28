package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type Word struct {
	word  string
	count int
}

type words []Word

func (s words) Len() int { return len(s) }
func (s words) Less(i, j int) bool {
	if s[i].count < s[j].count {
		return true
	}
	if s[i].count == s[j].count {
		return strings.Compare(s[i].word, s[j].word) == 1
	}
	return false
}
func (s words) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func Top10(inStr string) []string {
	// Place your code here.
	if inStr == "" {
		return nil
	}
	arrSplit := strings.Fields(inStr)
	cache := make(map[string]int)
	for i := 0; i < len(arrSplit); i++ {
		key := strings.ToLower(arrSplit[i])
		re := regexp.MustCompile(`^(.*)\.$`)
		key = re.ReplaceAllString(key, `$1`)
		if key == "-" {
			continue
		}
		value, ok := cache[key]
		if !ok {
			cache[key] = 1
		} else {
			cache[key] = value + 1
		}
	}

	ws := make(words, 0)
	for k, v := range cache {
		ws = append(ws, Word{count: v, word: k})
	}
	sort.Sort(sort.Reverse(ws))
	var arrRet []string
	numTop10 := func() int {
		if len(ws) > 10 {
			return 10
		}
		return len(ws)
	}()
	for i := 0; i < numTop10; i++ {
		arrRet = append(arrRet, ws[i].word)
	}
	return arrRet
}
