package main

import (
	"fmt"
	"time"
)

func timeDemo() {
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func timestampDemo() {
	now := time.Now()        //获取当前时间
	timestamp1 := now.Unix() //时间戳
	timestamp2 := now.UnixMilli()
	timestamp3 := now.UnixMicro()
	timestamp4 := now.UnixNano() //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
	fmt.Printf("current timestamp3:%v\n", timestamp3)
	fmt.Printf("current timestamp4:%v\n", timestamp4)
}

func testLocation() {
	// 构建时区
	var timeLocation *time.Location
	timeLocation, _ = time.LoadLocation("") //UTC
	fmt.Printf("current timeLocation:%v\n", timeLocation)
	timeLocation, _ = time.LoadLocation("UTC") //UTC
	fmt.Printf("current timeLocation:%v\n", timeLocation)
	timeLocation, _ = time.LoadLocation("Local") //Local
	fmt.Printf("current timeLocation:%v\n", timeLocation)
	timeLocation, _ = time.LoadLocation("Asia/Shanghai") //使用时区码
	fmt.Printf("current timeLocation:%v\n", timeLocation)
}

func main() {

	// timeDemo()
	// timestampDemo()
	testLocation()
}
