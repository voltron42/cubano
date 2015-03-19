package cubano

import (
  "fmt"
  "github.com/robertkrimen/otto"
  "io/ioutil"
  "path/filepath"
)

func Run(dir string, c Config) error {
  vm := otto.New()
  vm.Set("props", c.Properties)
  Native.applyTo(vm)
  for _, file := range c.Files {
    path := filepath.Join(dir, file)
    fmt.Println(path)
    data, err := ioutil.ReadFile(path)
    if err != nil {
      return err
    }
	CWD = filepath.Dir(path)
    _, err = vm.Run(string(data))
    if err != nil {
      return err
    }
  }
  return nil
}
