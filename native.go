package cubano

import (
  "github.com/robertkrimen/otto"
  "io/ioutil"
  "encoding/json"
  "os"
)
var CWD = ""
var Native = Scope(map[string]interface{}{
	"file":map[string]interface{}{
	    "readJson":func(call otto.FunctionCall) otto.Value {
		  filename, err := getFileName(call)
		  if err != nil {
			return makeError(err, call.Otto)
		  }
		  data, err := ioutil.ReadFile(filename)
		  if err != nil {
			return makeError(err, call.Otto)
		  }
		  var obj interface{}
		  err = json.Unmarshal(data, &obj)
		  if err != nil {
			return makeError(err, call.Otto)
		  }
		  out, err := otto.ToValue(obj)
		  if err != nil {
			return makeError(err, call.Otto)
		  }
		  return out
		},
		"read":func(call otto.FunctionCall) otto.Value {
		  fmt.Println("reading file")
		  filename, err := getFileName(call)
		  if err != nil {
			return makeError(err, call.Otto)
		  }
		  fmt.Println("reading data from file")
		  data, err := ioutil.ReadFile(filename)
		  if err != nil {
			return makeError(err, call.Otto)
		  }
		  fmt.Println("prepping data for return")
		  out, err := otto.ToValue(string(data))
		  if err != nil {
			return makeError(err, call.Otto)
		  }
		  return out
		},
		"write":func(call otto.FunctionCall) otto.Value {
		  filename, err := getFileName(call)
		  if err != nil {
			return makeError(err, call.Otto)
		  }
		  data, err := call.Argument(1).ToString()
		  if err != nil {
			return makeError(err, call.Otto)
		  }
		  err = ioutil.WriteFile(filename, []byte(data), os.ModePerm)
		  if err != nil {
			return makeError(err, call.Otto)
		  }
		  return otto.Value{}
		},
		"cwd":func(call otto.FunctionCall) otto.Value {
			return otto.ToValue(CWD)
		},
	},
})