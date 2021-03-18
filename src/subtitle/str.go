package subtitle

import (
	"io/ioutil"
	"regexp"
	"strings"
)

func lastNonEmptyLine(linesArray []string) int {
	var idx = len(linesArray) - 1

	for idx >= 0 && linesArray[idx] == "" {
		idx--
	}

	return idx
}

func nextNonEmptyLine(linesArray []string, position int) int {
	var idx = position

	for linesArray[idx] == "" {
		idx++
	}

	return idx
}

func ParseSrt(srt string) []Sub {
	lines := regexp.MustCompile(`(?:\r\n|\r|\n)`).Split(srt, -1)

	var endIdx = lastNonEmptyLine(lines)

	var subs []Sub

	var i int
	for i = 0; i < endIdx; i++ {
		var sub Sub

		i = nextNonEmptyLine(lines, i)

		sub.Id = lines[i]

		i++

		sub.Time = lines[i]

		var text []string
		for i < endIdx && lines[i] != "" {
			i++
			text = append(text, lines[i])
		}

		sub.Text = strings.Join(text, " ")

		subs = append(subs, sub)
	}

	return subs
}

func ParseSrtFromFile(filename string) ([]Sub, error) {
	if srt, err := ioutil.ReadFile(filename); err != nil {
		return nil, err
	} else {
		return ParseSrt(string(srt)), nil
	}
}
