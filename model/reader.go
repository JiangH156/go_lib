package model

type Reader struct {
	ReaderId string `readerId` varchar(50)   NOT NULL,
	ReaderName string `readerName` varchar(10)   DEFAULT NULL,
	Password string	`password` varchar(50)   DEFAULT NULL,
	`phone` varchar(25)   DEFAULT NULL,
	`borrowTimes` bigint(0) DEFAULT NULL,
	`ovdTimes` bigint(0) DEFAULT NULL,
	`email` varchar(255)   DEFAULT NULL,
	//姓名、密码、联系电话、借阅次数、逾期次数和电子邮箱
}
