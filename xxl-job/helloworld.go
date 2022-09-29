// Package xxl_job
// @author： Boice
// @createTime：2022/9/29 17:38
package main

import (
	"context"
	"fmt"
	xxl "github.com/xxl-job/xxl-job-executor-go"
	"log"
)

func main() {
	exec := xxl.NewExecutor(
		xxl.ServerAddr("http://127.0.0.1:8080/xxl-job-admin"),
		xxl.AccessToken("testToken"),   // token (default null)
		xxl.ExecutorIp("192.168.1.7"),  // executor ip(By default, it will be automatically obtained)
		xxl.ExecutorPort("9999"),       // default 9999
		xxl.RegistryKey("golang-jobs"), // job name
		xxl.SetLogger(&logger{}),       // custom log
	)
	exec.Init()
	// register task handler
	exec.RegTask("hello world", func(cxt context.Context, param *xxl.RunReq) (msg string) {
		fmt.Println("hello world task" + param.ExecutorHandler + " param：" + param.ExecutorParams + " log_id:" + xxl.Int64ToStr(param.LogID))
		return "done"
	})
	log.Fatal(exec.Run())

}

// xxl.Logger impl
type logger struct{}

func (l *logger) Info(format string, a ...interface{}) {
	fmt.Println(fmt.Sprintf("custom log - "+format, a...))
}

func (l *logger) Error(format string, a ...interface{}) {
	log.Println(fmt.Sprintf("custom log - "+format, a...))
}
