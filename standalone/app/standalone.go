package main

import (
  "../"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "path/filepath"
)

func main() {
/*
defer func() {
    if r := recover(); r != nil {
      fmt.Println(r)
    }
  }()
  */
  name := ""
  fmt.Println("blueprint: ")
  fmt.Scanf("%v", &name)
  fmt.Println(name)
  path, err := filepath.Abs(name)
  if err != nil {
    panic(err)
  }
  fmt.Println(path)
  data, err := ioutil.ReadFile(path)
  if err != nil {
    panic(err)
  }
  fmt.Println(data)
  var conf standalone.Config
  err = json.Unmarshal(data, &conf)
  if err != nil {
    panic(err)
  }
  fmt.Println(conf)
  err = standalone.Run(filepath.Dir(path), "..", conf)
  if err != nil {
    panic(err)
  }
}
