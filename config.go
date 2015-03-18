package cubano

import (
"github.com/robertkrimen/otto"
"strings"
)

type Config struct {
  Properties map[string]interface{} `json:"props"`
  Files []string `json:"files"`
}

type function func(call otto.FunctionCall) otto.Value

type Scope map[string]function

func (s Scope) apply(other Scope) {
	for key, fn := range other {
		s[key] = fn
	}
}

func (s Scope) applyTo(vm *otto.Otto) {
  out := map[string]interface{}{}
  for key, fn := range s {
    path := strings.Split(key, ".")
    branch := out
    for len(path) > 1 {
      step := path[0]
      path = path[1:]
      temp, ok := branch[step]
      if !ok {
        temp = map[string]interface{}{}
        branch[step] = temp
      }
      branch = temp
    }
    branch[path[0]] = fn
  }
  for key, value := range out {
    vm.Set(key, value)
  }
}

func SetDeep(vm *otto.Otto, key string, prop interface{}) err {
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
      temp, ok := branch[step]
      if !ok {
        temp = map[string]interface{}{}
        branch[step] = temp
      }
      branch = temp
	}
    branch := out
    obj[path[0]] = fn

}


