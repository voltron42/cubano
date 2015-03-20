package standalone

type Config struct {
Destination string
Title string
Styles []string
Scripts []string
Body BodyConfig
}

type BodyConfig struct {
BodyTemplate string
BodyData string
BodyAttributes map[string]string
}



