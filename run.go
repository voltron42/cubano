package cubano

import (
  "github.com/robertkrimen/otto"
  "io/ioutil"
)

func Run(c Config, s Scope) {
  vm := otto.New()
  vm.Set("props",c.Properties)
  s.apply(vm)
  for _, file := range c.Files {
    data, err := ioutil.ReadFile(file)
    if err != nil {
      panic(err)
    }
    _, err = vm.Run(string(data))
    if err != nil {
      panic(err)
    }
  }
}
