package main

import "context"

type homeWorkHandlerInterface interface {
	homeworkDoSomething(ctx context.Context) error
	getName() string
}
