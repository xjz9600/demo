package main

import (
	"context"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//getHomeWorkHanderMng().getHomeworkHandler("homeworkFirst").homeworkDoSomething(context.Background())
	getHomeWorkHanderMng().getHomeworkHandler("homeworkSecond").homeworkDoSomething(context.Background())
}

func funcName(ch chan int) {
	select {
	case i := <-ch:
		println(i)

	default:
		println("default")
	}
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:qishiwang5@tcp(127.0.0.1:3306)/slo_test?charset=utf8&parseTime=true")
}
