package cubano

type Config struct {
  Properties map[string]interface{} `json:"props"`
  Files []string `json:"files"`
}
