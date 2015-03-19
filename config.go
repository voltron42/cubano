package cubano

import (
	"github.com/robertkrimen/otto"
	"strings"
)

type Config struct {
	Properties map[string]interface{} `json:"props"`
	Files      []string               `json:"files"`
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
			temp, ok1 := branch[step]
			obj, ok2 := temp.(map[string]interface{})
			if !ok1 || !ok2 {
				obj = map[string]interface{}{}
				branch[step] = obj
			}
			branch = obj
		}
		branch[path[0]] = fn
	}
	for key, value := range out {
		vm.Set(key, value)
	}
}
