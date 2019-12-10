package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for {
		nowTime := time.Now().Format("2006-01-02 15:04:05")
		fmt.Fprintf(os.Stdout, "%s\r", nowTime)
		time.Sleep(time.Second * 1)
	}

}
