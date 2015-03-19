package cubano

import (
	"github.com/robertkrimen/otto"
)

type Config struct {
	Properties map[string]interface{} `json:"props"`
	Files      []string               `json:"files"`
}

type function func(call otto.FunctionCall) otto.Value

type Scope map[string]interface{}

func (s Scope) apply(other Scope) {
	for key, obj := range other {
		s[key] = obj
	}
}

func (s Scope) applyTo(vm *otto.Otto) {
	for key, value := range s {
		vm.Set(key, value)
	}
}
