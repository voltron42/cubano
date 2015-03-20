package standalone

import (
	"../"
	"os"
	"io"
	"io/ioutil"
	"path/filepath"
	"github.com/robertkrimen/otto"
)

func Run(dir string, conf Config) error {
	out, err := Create(dir, conf.Destination)
	if err != nil {
		return err
	}
	out.Println("<!DOCTYPE html>")
	out.Println("<html>")
	out.Println("<head>")
	out.Println("<title>"+conf.Head.Title+"</title>")
	out.Println("<style>")
	for _, style := range conf.Head.Styles {
		out.Stream(dir, style)
	}
	out.Println("</style>")
	out.Println("<script>")
	for _, script := range conf.Head.Scripts {
		out.Stream(dir, script)
	}
	out.Println("</script>")
	out.Println("</head>")
	out.Println("<body")
	for name, attr := range conf.Body.Attributes {
		out.Print(" ")
		out.Print(name)
		out.Print("=\"")
		out.Print(attr)
		out.Print("\"")
	}
	out.Println(">")
	tpl, err := readFrom(dir, conf.Body.Template)
	if err != nil {
		return err
	}
	data, err := readFrom(dir, conf.Body.Data)
	if err != nil {
		return err
	}
	body, err := buildBody(tpl, data)
	if err != nil {
		return err
	}
	out.Println(body)
	out.Println("</body>")
	out.Println("</html>")
	return nil
}

func readFrom(path, filename string) (string, error) {
	out, err := ioutil.ReadFile(filepath.Join(path, filename))
	return string(out), err
}

func buildBody(tpl, data string) (string, error) {
	conf := cubano.Config{
		Files:[]string{
			
			"./standalone.js",
		},
		Properties:map[string]interface{}{
			"job":map[string]string{
				"tpl":(tpl),
				"data":(data),
			},
		},
	}
	retVal := ""
	cubano.Native["callback"] = func(call otto.FunctionCall) otto.Value {
		retVal = call.Argument(0).String()
		return otto.Value{}
	}
	dir, err := os.Getwd()
	if err == nil {
		err = cubano.Run(dir, conf)
	}
	return retVal, err
}

type OutFile struct {
	file *os.File
}

func Create(path, filename string) (*OutFile, error) {
	
	outfile, err := os.Create(filepath.Join(path, filename))
	return &OutFile{outfile}, err
}

func (o *OutFile) Print(line string) error {
	_, err := o.file.WriteString(line)
	return err
}

func (o *OutFile) Println(line string) error {
	return o.Print(line + "\n")
}

func (o *OutFile) Stream(path, fileName string) error {
	src, err := os.Open(filepath.Join(path,fileName))
	if err != nil {
		return err
	}
	_, err = io.Copy(o.file, src)
	return err
}