package splitstring

import (
	"reflect"
	"testing"
)

func Test1Split(t *testing.T) {
	got := Split("bcdbefgbhijk", "b")
	want := []string{"", "cd", "efg", "hijk"}
	if !reflect.DeepEqual(got, want) {
		//测试用例失败了
		t.Errorf("want : %v,but got : %v\n", want, got)
	}
}

func Test2Split(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		//测试用例失败了
		t.Errorf("want : %v,but got : %v\n", want, got)
	}
}

func Test3Split(t *testing.T) {
	got := Split("abcde", "cd")
	want := []string{"ab", "e"}
	if !reflect.DeepEqual(got, want) {
		//测试用例失败了
		t.Fatalf("want : %v,but got : %v\n", want, got)
	}
}
