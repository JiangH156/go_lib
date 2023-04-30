package main

import (
	"Go_lib/common"
)

func main() {
	to := []string{"2360542085@qq.com"}
	// 定义邮件主题和内容
	subject := "Test"
	body := `大聪明`
	for t := 0; t < 5; t++ {
		if err := common.SendEmail(
			to,
			nil,
			nil,
			subject,
			body,
			"",
		); err != nil {
			panic(err)
		}
	}

}
