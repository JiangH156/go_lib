package utils

import (
	"Go_lib/model"
	"errors"
	"fmt"
	"regexp"
	"time"
)

const (
	emailRegex = `^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`
	phoneRegex = `^1[3456789]\d{9}$`
	TimeLayout = `2006-01-02 15:04:05`
)

func EmailRegexp(email string) error {
	// 编译正则表达式
	emailRegexp := regexp.MustCompile(emailRegex)
	// 匹配字符串
	if !emailRegexp.MatchString(email) {
		return errors.New("邮箱正则表达式匹配失败")
	}
	return nil
}

func PhoneRegexp(phone string) error {
	// 编译正则表达式
	phoneRegexp := regexp.MustCompile(phoneRegex)
	// 匹配字符串
	if !phoneRegexp.MatchString(phone) {
		return errors.New("手机号正则表达式匹配失败")
	}
	return nil
}

func ParseTime(dateStr string) (model.Time, error) {
	// 定义日期时间格式和时区信息
	timeLayout := "2006-01-02 15:04:05"

	// 解析时间字符串
	addTime, err := time.ParseInLocation(timeLayout, dateStr, time.Local)
	if err != nil {
		// 处理解析错误
		fmt.Println("parse time error:", err)
		return model.Time{}, fmt.Errorf("parse time error: %s", err)
	}

	//设置保存精度为秒
	//addTime.Truncate(time.Second)
	//fmt.Println(addTime)
	return model.Time(addTime), nil
}
