package cubano

type Config struct {
  Properties map[string]interface{} `json:"props"`
  Files []string `json:"files"`
}

type function func(call otto.FunctionCall) otto.Value

type scope map[string]function

func Apply