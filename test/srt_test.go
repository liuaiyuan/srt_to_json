package test

import (
	"github.com/liuaiyuan/srt_to_json/src/subtitle"
	"testing"
)

func TestParseSrt(t *testing.T) {
	var srt = `
1
00:00:01,280 --> 00:00:02,048
hello

2
00:00:02,304 --> 00:00:03,328
world
`

	subs := subtitle.ParseSrt(srt)
	if len(subs) != 2 {
		t.Errorf("error len != 2")
	}

	for _, sub := range subs {
		if sub.Id == "" || sub.Time == "" || sub.Text == "" {
			t.Errorf("error")
		}
	}
}

func TestParseSrtFromFile(t *testing.T) {
	subs, err := subtitle.ParseSrtFromFile("./test.srt")
	if err != nil {
		t.Error(err.Error())
		return
	}

	for _, sub := range subs {
		if sub.Id == "" || sub.Time == "" || sub.Text == "" {
			t.Errorf("error")
		}
	}
}
