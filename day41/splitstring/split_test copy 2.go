package splitstring

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	type testCase struct {
		input string
		sep   string
		want  []string
	}

	testGroup := []testCase{
		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		{input: "afzxxcsdfxxcvrsdfxxcv", sep: "xx", want: []string{"afz", "csdf", "cvrsdf", "cv"}},
		{input: "张旭有张又有旭", sep: "有", want: []string{"张旭", "张又", "旭"}},
	}
	for _, tc := range testGroup {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("excepted:%#v,got:%#v", tc.want, got)
		}
	}
}
