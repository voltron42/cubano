package main

import (
  "../"
  "io/ioutil"
  "encoding/json"
  "fmt"
  "path/filepath"
)

func main() {
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
  cubano.Run(filepath.Dir(path), conf, cubano.BuildNative)
}



