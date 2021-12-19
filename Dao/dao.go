package Dao

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/pkg/errors"
	"time"
)

var NotFound = errors.New("not found")

// UserContactInfo ...
type UserContactInfo struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	UserID       string
	WeUserID     string
	UserNO       string
	UserName     string
	TelMobile    string
	DeptNames    string
	DeptIds      string
	WechatID     string
	LastSendTime time.Time
	LastCallTime time.Time
	IsWeFans     bool
	Enable       bool
	Duty         uint
}

func GetUserDetailByUserName(userId int) (*UserContactInfo, error) {
	o := orm.NewOrm()
	var userInfo UserContactInfo
	sqlStr := fmt.Sprintf("SELECT * FROM userContactInfo WHERE user_id = %d", userId)
	err := o.Raw(sqlStr).QueryRow(&userInfo)
	if err == orm.ErrNoRows {
		return &userInfo, errors.Wrap(NotFound, fmt.Sprintf("can not find userInfo userId is %d", userId))
	}
	if err != nil {
		return &userInfo, errors.Wrap(err, fmt.Sprintf("get Datas err sql is %s", sqlStr))
	}
	return &userInfo, nil
}
