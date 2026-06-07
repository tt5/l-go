package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	logfile := path.Join(os.TempDir(), "log.log")
	fmt.Println(logfile)
	f, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	mylog := log.New(f, "mylog ", log.LstdFlags)
	mylog.Println("Hi")
}
