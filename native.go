package cubano

type function func(call otto.FunctionCall) otto.Value

type scope map[string]interface{}

vm.Set("sayHello", func(call otto.FunctionCall) otto.Value { 
fmt.Printf("Hello, %s.\n", call.Argument(0).String()) 
return otto.Value{} 
})

func applyTo(vm *otto.Otto) {
  s := &scope{}
  
  s.apply("json.readFile", func(call otto.FunctionCall) otto.Value {
    
  })
  
  s.apply("file.read", func(call otto.FunctionCall) otto.Value {
    
  })
  
  s.apply("file.write", func(call otto.FunctionCall) otto.Value {
    
  })
  
  s.applyTo(vm)
}



