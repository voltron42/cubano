package cubano

import (
  "github.com/robertkrimen/otto"
)

func Run(c Config) {
  vm := otto.New()
  vm.Set("props",c.Properties)
  applyNative(vm)
  for (file string : c.Files) {
    data, err := ioutil.ReadFile(file)
    if err != nil {
      panic(err)
    }
    _, err := vm.Run(string(data))
    if err != nil {
      panic(err)
    }
  }
}
