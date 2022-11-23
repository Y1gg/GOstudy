package splitstring

import (
	"reflect"
	"testing"
)

func Test0Split(t *testing.T) {
	//定义testCase结构体
	type testCase struct {
		input string
		sep   string
		want  []string
	}

	//测试用例使用map存储
	testGroup := map[string]testCase{
		"simple":   {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"more sep": {input: "afzxxcsdfxxcvrsdfxxcv", sep: "xx", want: []string{"afz", "csdf", "cvrsdf", "cv"}},
		"chinese":  {input: "张旭有张又有旭", sep: "有", want: []string{"张旭", "张又", "旭"}},
	}
	for name, tc := range testGroup {
		//使用t.Run()执行子测试
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("excepted:%#v,got:%#v", tc.want, got)
			}
		})

	}
}

//基准测试

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d:e", ":")
	}
}
