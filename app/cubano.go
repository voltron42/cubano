package main

import (
  "../"
  "io/ioutil"
  "encoding/json"
  "fmt"
)

func main() {
  name := ""
  fmt.Println("blueprint: ")
  fmt.Scanf("%v", &name)
  data, err := ioutil.ReadFile(name)
  if err != nil {
    panic(err)
  }
  var conf cubano.Config
  err = json.Unmarshal(data, &conf)
  if err != nil {
    panic(err)
  }
  cubano.Run(conf, cubano.BuildNative(vm))
}



