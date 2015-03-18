package cubano

import (
  "github.com/robertkrimen/otto"
  "io/ioutil"
)

func Run(c Config, opts... func(vm *otto.Otto) error ) {
  vm := otto.New()
  vm.Set("props",c.Properties)
  for _, option := range opts {
  err := option(vm)
    if err != nil {
      panic(err)
    }
  }
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
