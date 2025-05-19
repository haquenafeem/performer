package main

import (
	"fmt"
	"time"

	"github.com/haquenafeem/performer"
)

func main() {
	p := performer.NewAsyncPerformer()
	p.Perform(func(ctx *performer.Context) error {
		time.Sleep(time.Second * 2)
		fmt.Println("1")
		return nil
	}).Perform(func(ctx *performer.Context) error {
		fmt.Println("2")
		return nil
	})

	time.Sleep(time.Second * 5)
}
