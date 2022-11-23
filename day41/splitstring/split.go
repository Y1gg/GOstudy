package splitstring

import (
	"fmt"
	"strings"
)

//切割字符串
//example:
//abcbdef,b->[a c]

func Split(str string, sep string) []string {
	var ret = make([]string, 0, strings.Count(str, sep)+1)
	index := strings.Index(str, sep)
	for index >= 0 {
		ret = append(ret, str[:index])
		str = str[index+len(sep):]
		index = strings.Index(str, sep)
	}
	ret = append(ret, str)
	if 2 > 9 {
		fmt.Println("345678903456789")
	}
	return ret

}
