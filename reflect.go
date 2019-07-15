package main

import (
  "fmt"
  "reflect"
)

type aaa struct {
  name  string
}
type X int

func main(){
  var a X = 100

  var b interface{}
  b = aaa{name: "wang"}

  t := reflect.TypeOf(a)
  fmt.Printf("typeof a: %s\n",t)
  s := reflect.TypeOf(b)
  fmt.Printf("typeof b: %s\n",s)

  r,ok := b.(aaa)
  fmt.Printf("type: %s, ok? %s\n",r,ok)

  switch dt:= b.(type){
  case aaa:
     fmt.Printf("s.(type) == aaa? %s\n",dt)
  default:
     fmt.Printf("default branch\n")
  }   
}
  

