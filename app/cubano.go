package main

import (
  "../"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "path/filepath"
)

func main() {
  defer func() {
    if r := recover(); r != nil {
      fmt.Println(r)
    }
  }()
  name := ""
  fmt.Println("blueprint: ")
  fmt.Scanf("%v", &name)
  fmt.Println(name)
  path, err := filepath.Abs("./" + name + ".json")
  fmt.Println(path)
  data, err := ioutil.ReadFile(path)
  if err != nil {
    panic(err)
  }
  var conf cubano.Config
  err = json.Unmarshal(data, &conf)
  if err != nil {
    panic(err)
  }
  err = cubano.Run(filepath.Dir(path), conf, cubano.BuildNative)
  if err != nil {
    panic(err)
  }
}
