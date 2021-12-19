package main

import "fmt"

var mngInstance *homeworkHandlerMng

type homeworkHandlerMng struct {
	handler map[string]homeWorkHandlerInterface
}

func getHomeWorkHanderMng() *homeworkHandlerMng {
	return mngInstance
}

func (mng *homeworkHandlerMng) registHandler(handerInterface homeWorkHandlerInterface) {
	mng.handler[handerInterface.getName()] = handerInterface
}

func init() {
	mngInstance = &homeworkHandlerMng{handler: map[string]homeWorkHandlerInterface{}}
	mngInstance.registHandler(&homeworkFirst{})
	mngInstance.registHandler(&homeworkSecond{})
}

func (mng *homeworkHandlerMng) getHomeworkHandler(homeworkName string) homeWorkHandlerInterface {
	handler, isok := mngInstance.handler[homeworkName]
	if !isok {
		panic(fmt.Sprintf("次作业还没做呢"))
	}
	return handler
}
