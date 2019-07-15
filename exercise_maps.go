package main

import (
	"golang.org/x/tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
    c := strings.Fields(s)
    for _,v := range c{
      m[v]+=1
    }
    return m
}

func main() {
    var s = "i am learning go"
	wc.Test(WordCount(s))
}
