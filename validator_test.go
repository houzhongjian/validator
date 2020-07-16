package validator

import (
	"log"
	"testing"
)

type User struct {
	Name  string `validator:"type:string;name:昵称;required:true;length:[2-4]"`
	Phone string `validator:"type:string;name:手机号;required:true;length:[11-11]"`
	ID    int    `validator:"type:int;name:ID;min:1"`
	List  string `validator:"type:regexp;name:列表;expression:https://(\\w+).(\\w+).(\\w+)"`
	Deep  string `validator:"type:regexp;name:深度;required:false;expression:([\\d]+,){2}[\\d]+$"`
}

func TestCheck(t *testing.T) {
	user := User{
		Name:  "张san",
		Phone: "15173083374",
		ID:  1,
		List:  "https://www.baidu.com",
		Deep:  "10,20,3",
	}

	log.Println(Check(user))
}
