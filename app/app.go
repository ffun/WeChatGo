package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"robot"
	"runtime/debug"
)

func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				log.Println("WARN:panic in %v - %v", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}

func main() {
	//创建一个机器人
	robot := robot.NewRobot(&robot.QinyunRobot{})
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("input query string:")
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("read err")
		}
		//并发地进行服务
		go robot.Serve(os.Stdout, "GET", string(line))
		// robot.Serve(os.Stdout, "GET", string(line))
	}
}
