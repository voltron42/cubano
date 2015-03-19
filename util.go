package cubano

import (
  "github.com/robertkrimen/otto"
  "filepath"
  "errors"
)

func setDeep(vm *otto.Otto, key string, prop interface{}) err {
	stdErr := errors.New("Cannot append property to non-object")
    path := strings.Split(key, ".")
	if len(path) = 1 {
		return vm.set(path[0], prop)
	}	
	value, err := vm.Get(path[0])
	if err != nil {
		return err
	}
	if !value.IsObject() {
		return stdErr
	}
	obj := value.Object()
	path = path[1:]
    for len(path) > 1 {
      step := path[0]
      path = path[1:]
      value, err = obj.Get(step)
      if err != nil {
        temp := otto.Object{}
        stub, _ := temp.Value().Export()
        err = obj.Set(step, stub)
        if err != nil {
          return err
        }
        obj = temp
      } else if !value.IsObject() {
        return stdErr
      } else {
        obj = value.Object()
      }
	}
    obj[path[0]] = fn
  return nil
}

func getFileName(call otto.FunctionCall) (string, error) {
    this := call.This
    if !this.IsObject() {
        return "", errors.New("improper scope on file method")
    }
    thisObj := this.Object()
    value, err := thisObj.Get("cwd")
    if err != nil {
        return "", err
    }
    cwd, err := value.ToString()
    if err != nil {
        return "", err
    }
    filename, err := call.Argument(0).ToString()
    if err != nil {
        return "", err
    }
    return filepath.Join(cwd, filename), nil
}

func makeError(err error, vm *otto.Otto) otto.Value {
  value, _ := vm.Call("new Error", nil, err.Error());
  return value
}




