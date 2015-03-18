package cubano

import (
  "github.com/robertkrimen/otto"
  "io/ioutil"
  "encoding/json"
  "os"
)

func BuildNative(vm *otto.Otto) error {
  Scope(map[string]function{
    "json.readFile":function(func(call otto.FunctionCall) otto.Value {
      filename, err := call.Argument(0).ToString()
      if err != nil {
        return makeError(err, vm)
      }
      data, err := ioutil.ReadFile(filename)
      if err != nil {
        return makeError(err, vm)
      }
      var obj interface{}
      err = json.Unmarshal(data, &obj)
      if err != nil {
        return makeError(err, vm)
      }
      out, err := otto.ToValue(obj)
      if err != nil {
        return makeError(err, vm)
      }
      return out
    }),
    "file.read":function(func(call otto.FunctionCall) otto.Value {
      filename, err := call.Argument(0).ToString()
      if err != nil {
        return makeError(err, vm)
      }
      data, err := ioutil.ReadFile(filename)
      if err != nil {
        return makeError(err, vm)
      }
      out, err := otto.ToValue(string(data))
      if err != nil {
        return makeError(err, vm)
      }
      return out
    }),
    "file.write":function(func(call otto.FunctionCall) otto.Value {
      filename, err := call.Argument(0).ToString()
      if err != nil {
        return makeError(err, vm)
      }
      data, err := call.Argument(1).ToString()
      if err != nil {
        return makeError(err, vm)
      }
      err = ioutil.WriteFile(filename, []byte(data), os.ModePerm)
      if err != nil {
        return makeError(err, vm)
      }
      return otto.Value{}
    }),
  }).apply(vm)
  return nil
}
func makeError(err error, vm *otto.Otto) otto.Value {
  value, _ := vm.Call("new Error", nil, err.Error());
  return value
}




