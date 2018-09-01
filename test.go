package main

import "fmt"
import "time"

func main() {
	fmt.Println("Go言語、はじめました")
  today:=time.Now()
  fmt.Println(today.Format("2006-01-02"))
}