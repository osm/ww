package ww

import (
	"reflect"
	"testing"
)

type wwTest struct {
	width  int
	input  string
	output []string
}

var wwTests = []wwTest{
	{
		width:  1,
		input:  "aaa bbb ccc",
		output: []string{"aaa bbb ccc"},
	},
	{
		width:  5,
		input:  "a b c d e f g h i",
		output: []string{"a b c", "d e f", "g h i"},
	},
	{
		width:  5,
		input:  "aaaaaa b c d efg hijk",
		output: []string{"aaaaaa", "b c d", "efg", "hijk"},
	},
	{
		width:  5,
		input:  "aaaaa bbbbb cccc dddd eee fff gg hh i j",
		output: []string{"aaaaa", "bbbbb", "cccc", "dddd", "eee", "fff", "gg hh", "i j"},
	},
	{
		width:  5,
		input:  "aaaaa bbbbb cccc dddd eee fff gg hh i j llllll",
		output: []string{"aaaaa", "bbbbb", "cccc", "dddd", "eee", "fff", "gg hh", "i j", "llllll"},
	},
	{
		width:  5,
		input:  "a bb c d e f gggg hh ii j",
		output: []string{"a bb", "c d e", "f", "gggg", "hh ii", "j"},
	},
	{
		width:  5,
		input:  "aa bbb c dddd eee f gggggg hh iii jjjj kkkkk llll m n o p",
		output: []string{"aa", "bbb c", "dddd", "eee f", "gggggg", "hh", "iii", "jjjj", "kkkkk", "llll", "m n o", "p"},
	},
}

func Test(t *testing.T) {
	for _, wwt := range wwTests {
		t.Run("", func(t *testing.T) {
			o := Wrap(wwt.input, wwt.width)
			if !reflect.DeepEqual(wwt.output, o) {
				t.Errorf("failed to wrap input %s", wwt.input)
				t.Logf("output  : %#v", o)
				t.Logf("expected: %#v", wwt.output)
			}
		})
	}
}
