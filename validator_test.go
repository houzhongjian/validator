package validator

import (
	"log"
	"testing"
)

type User struct {
	Name   string `validator:"type:string;name:昵称;required:true;length:[2-4]"`
	Phone  string `validator:"type:string;name:手机号;required:true;length:[11-11]"`
	ID     int    `validator:"type:int;name:ID;min:1;max:10"`
	Domain string `validator:"type:regexp;name:域名;required:true;expression:https://(\\w+).(\\w+).(\\w+)"`
	Deep   string `validator:"type:regexp;name:深度;required:true;expression:([\\d]+,){2}[\\d]+$"`
	Mobile string `validator:"type:regexp;name:电话号码;required:true;expression:^1[345789]\\d{9}$"`
}

func TestCheck(t *testing.T) {
	user := User{}
	log.Println(Check(user))

	user = User{
		Name: "张",
	}
	log.Println(Check(user))

	user = User{
		Name: "张三",
	}
	log.Println(Check(user))

	user = User{
		Name:  "张三",
		Phone: "15183983",
	}
	log.Println(Check(user))

	user = User{
		Name:  "张三",
		Phone: "151839835555",
	}
	log.Println(Check(user))

	user = User{
		Name:  "张三",
		Phone: "15183983555",
	}
	log.Println(Check(user))

	user = User{
		Name:  "张三",
		Phone: "15183983555",
		ID:    0,
	}
	log.Println(Check(user))

	user = User{
		Name:   "张三",
		Phone:  "15183983555",
		ID:     11,
		Domain: "http://",
	}
	log.Println(Check(user))

	user = User{
		Name:   "张三",
		Phone:  "15183983555",
		ID:     1,
		Domain: "http://www.baidu.com",
	}
	log.Println(Check(user))

	user = User{
		Name:   "张三",
		Phone:  "15183983555",
		ID:     1,
		Domain: "https://www.baidu.com",
	}
	log.Println(Check(user))

	user = User{
		Name:   "张三",
		Phone:  "15183983555",
		ID:     1,
		Domain: "https://www.baidu.com",
		Deep:   "1",
	}
	log.Println(Check(user))

	user = User{
		Name:   "张三",
		Phone:  "15183983555",
		ID:     1,
		Domain: "https://www.baidu.com",
		Deep:   "1,2,3,",
	}
	log.Println(Check(user))

	user = User{
		Name:   "张三",
		Phone:  "15183983555",
		ID:     1,
		Domain: "https://www.baidu.com",
		Deep:   "1,2,3",
	}
	log.Println(Check(user))

	user = User{
		Name:   "张三",
		Phone:  "15183983555",
		ID:     1,
		Domain: "https://www.baidu.com",
		Deep:   "1,2,3",
		Mobile: "15183983555",
	}
	log.Println(Check(user))
}
