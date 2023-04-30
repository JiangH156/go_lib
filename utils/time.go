package utils

import (
	"fmt"
	"time"
)

func ParseTime(dateStr string) (time.Time, error) {

	// 解析时间字符串
	addTime, err := time.ParseInLocation(time.DateTime, dateStr, time.Local)
	if err != nil {
		// 处理解析错误
		fmt.Println("parse time error:", err)
		return time.Time{}, fmt.Errorf("parse time error: %s", err)
	}

	//设置保存精度为秒
	//addTime.Truncate(time.Second)
	//fmt.Println(addTime)
	return addTime, nil
}

func NowTime() time.Time {
	t := time.Now()
	timeStr := t.Format(time.DateTime)
	date, _ := time.ParseInLocation(time.DateTime, timeStr, time.Local)
	return date
}
