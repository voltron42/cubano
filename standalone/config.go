package standalone

type Config struct {
Destination string `json:"dest"`
Head HeadConfig `json:"head"`
Body BodyConfig `json:"body"`
}

type HeadConfig struct {
Title string `json:"title"`
Styles []string `json:"styles"`
Scripts []string `json:"scripts"`
}
type BodyConfig struct {
Attributes map[string]string `json:"attrs"`
Template string `json:"tpl"`
Data string `json:"data"`
}



