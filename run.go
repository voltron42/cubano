package cubano

import (
  "github.com/robertkrimen/otto"
  "io/ioutil"
  "path/filepath"
  "fmt"
)

func Run(dir string, c Config, opts... func(vm *otto.Otto) error ) {
  vm := otto.New()
  vm.Set("props",c.Properties)
  for _, option := range opts {
  err := option(vm)
    if err != nil {
      panic(err)
    }
  }
  for _, file := range c.Files {
    path := filepath.Join(dir, file);
    fmt.Println(path)
    data, err := ioutil.ReadFile(path)
    if err != nil {
      panic(err)
    }
    _, err = vm.Run(string(data))
    if err != nil {
      panic(err)
    }
  }
}
