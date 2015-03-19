package cubano

import (
  "fmt"
  "github.com/robertkrimen/otto"
  "io/ioutil"
  "path/filepath"
)

func Run(dir string, c Config, opts ...func(s Scope, vm *otto.Otto) error) error {
  vm := otto.New()
  vm.Set("props", c.Properties)
  s := Scope{}
  for _, option := range opts {
    err := option(s, vm)
    if err != nil {
      return err
    }
  }
  s.applyTo(vm)
  for _, file := range c.Files {
    path := filepath.Join(dir, file)
    fmt.Println(path)
    data, err := ioutil.ReadFile(path)
    if err != nil {
      return err
    }
    err = setDeep(vm, "file.cwd", filepath.Dir(path))
    if err != nil {
      return err
    }
    _, err = vm.Run(string(data))
    if err != nil {
      return err
    }
  }
  return nil
}
