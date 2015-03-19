package cubano

import (
  "github.com/robertkrimen/otto"
  "io/ioutil"
  "encoding/json"
  "os"
)

func BuildNative(s Scope, vm *otto.Otto) error {
  s.apply(Scope(map[string]function{
    "file.readJson":function(func(call otto.FunctionCall) otto.Value {
      filename, err := getFileName(call)
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
      filename, err := getFileName(call)
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
      filename, err := getFileName(call)
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
  }))
  return nil
}

