package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"time"
)

func main() {
	getUserDetailByUserName(123)
}

// SmsDetail ...
type UserInfo struct {
	userId   int `gorm:"primary_key;AUTO_INCREMENT;column:id"`
	userName time.Time
	age      int
	gender   string
}

func getUserDetailByUserName(userId int) (*UserInfo, error) {
	var userInfo UserInfo
	db, err := sql.Open("mysql", "root:qishiwang5@tcp(127.0.0.1:3306)/slo_test?charset=utf8&parseTime=true")
	if err != nil {
		return &userInfo, errors.Wrap(err, "connet mysql err")
	}
	defer db.Close()
	sqlStr := fmt.Sprintf("select * from sms_record_202111 where userId = %d", userId)
	rows, err := db.Query(sqlStr)
	fmt.Println(rows)
	if err != nil && err != sql.ErrNoRows {
		return &userInfo, errors.Wrap(err, fmt.Sprintf("get Datas err sql is %s", sqlStr))
	}
	return &userInfo, nil
}
