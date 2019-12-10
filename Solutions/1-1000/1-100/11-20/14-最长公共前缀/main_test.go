package mario

import "testing"

var testCase = []struct {
	In     []string
	Except string
}{
	{[]string{"flower", "flow", "flight"}, "fl"},
	{[]string{"dog", "racecar", "car"}, ""},
	{[]string{}, ""},
	{[]string{"c", "c"}, "c"},
	{[]string{"aaa", "aa", "aaa"}, "aa"},
}

func TestLongestCommonPrefix(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if longestCommonPrefix(tcs[i].In) != tcs[i].Except {
			t.Errorf("longest common prefix test failed on case: %d", i)
		}
	}
}
