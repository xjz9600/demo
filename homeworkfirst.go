package main

import (
	"context"
	"github.com/pkg/errors"
	"studyProgect/demo/Dao"
)

type homeworkFirst struct {
}

func (h *homeworkFirst) homeworkDoSomething(ctx context.Context) error {
	_, err := Dao.GetUserDetailByUserName(123)
	if errors.Is(err, Dao.NotFound) {
		//to do something
		return nil
	}
	if err != nil {
		return err
	}
	return nil
}

func (h *homeworkFirst) getName() string {
	return "homeworkFirst"
}
