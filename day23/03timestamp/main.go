package main

// current	现在的
import (
	"fmt"
	"time"
)

func timestmpDemo1() {
	now := time.Now()            //获取当前时间
	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("现在的current timestmp1:%v\n", timestamp1)
	fmt.Printf("现在的current timestmp2:%v\n", timestamp2)
}

func timestampDemo2(timestamp int64) {
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //时
	minute := timeObj.Minute() //分
	second := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
func main() {
	timestmpDemo1()
	fmt.Println("\n====================================")
	timestampDemo2(1000000000)
	fmt.Println("\n====================================")
	//时间间隔
	fmt.Println(time.Second)
	/*
			const (
			Nanosecond  Duration = 1
			Microsecond          = 1000 * Nanosecond
			Millisecond          = 1000 * Microsecond
			Second               = 1000 * Millisecond
			Minute               = 60 * Second
			Hour                 = 60 * Minute
		)
	*/

	//now+24小时
	fmt.Println(time.Now())
	fmt.Println(time.Now().Add(24 * time.Hour))
	fmt.Println("\n====================================")
	//定时器
	timer := time.Tick(time.Second * 3)
	for t := range timer {
		fmt.Println(t) //三秒执行一次
		break
	}
	fmt.Println("\n====================================")
	//格式化时间：把语言中的时间对象，转换成字符串类型的时间
	fmt.Println(time.Now().Format("2006、01、02"))
	fmt.Println(time.Now().Format("2006/01/02 15/4/5 pm"))
	fmt.Println(time.Now().Format("2006/01/02 15/4/5.000"))
	fmt.Println("\n====================================")
	// parse	作语法分析
	//按照对应的格式解析字符串类型的时间
	timeObj, err := time.Parse("2006-01-02", "2022-10-12")
	if err != nil {
		fmt.Printf("parse time failed,err:%v\n", err)
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())
	fmt.Println("\n====================================")
	/*
		Sub:求两个时间之间的差值：
		func (t Time) Sub(u Time) Duration
		返回一个时间段t-u。如果结果超出了Duration可以表示的最大值/最小值，
		将返回最大值/最小值。要获取时间点t-d（d为Duration），可以使用t.Add(-d)。
	*/
	meetDay, err := time.Parse("2006-01-02", "2022-12-26")
	if err != nil {
		fmt.Printf("parse time failed,err:%v\n", err)
	}
	time := time.Now()
	d := meetDay.Sub(time)
	fmt.Println(d)
}
