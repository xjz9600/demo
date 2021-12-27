package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type homeworkSecond struct {
}

func (h *homeworkSecond) homeworkDoSomething(ctx context.Context) error {
	ctx2, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	srv := &http.Server{Addr: ":8080"}
	group, errCxt := errgroup.WithContext(ctx2)
	group.Go(func() error {
		http.HandleFunc("/hello", helloServer)
		return srv.ListenAndServe()
	})
	group.Go(func() error {
		<-errCxt.Done()
		fmt.Println("http stop")
		return srv.Shutdown(errCxt)
	})
	group.Go(func() error {
		select {
		case <-errCxt.Done():
			fmt.Println("stop by context")
		case <-sign():
			fmt.Println("signl stop")
			cancel()
		}
		return nil
	})
	group.Wait()
	fmt.Println("finish")
	return nil
}

func helloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello,word!")
}

func sign() chan os.Signal {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	return sigs
}

func (h *homeworkSecond) getName() string {
	return "homeworkSecond"
}
