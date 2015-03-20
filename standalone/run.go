package standalone

import (
	"../"
	"os"
	"io"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"github.com/robertkrimen/otto"
)

func Run(srcpath, execpath string, conf Config) error {
	out, err := Create(srcpath, conf.Destination)
	if err != nil {
		return err
	}
	fmt.Println("writing to file")
	out.Println("<!DOCTYPE html>")
	out.Println("<html>")
	out.Println("<head>")
	if len(conf.Head.Title) > 0 {
		out.Println("<title>"+conf.Head.Title+"</title>")
	}
	if len(conf.Head.Styles) > 0 {
		out.Println("<style>")
		for _, style := range conf.Head.Styles {
			out.Stream(srcpath, style)
		}
		out.Println("</style>")
	}
	if len(conf.Head.Scripts) > 0 {
		out.Println("<script>")
		for _, script := range conf.Head.Scripts {
			out.Stream(srcpath, script)
		}
		out.Println("</script>")
	}
	out.Println("</head>")
	out.Print("<body")
	for name, attr := range conf.Body.Attributes {
		out.Print(" ")
		out.Print(name)
		out.Print("=\"")
		out.Print(attr)
		out.Print("\"")
	}
	out.Println(">")
	tpl, err := readFrom(srcpath, conf.Body.Template)
	if err != nil {
		return err
	}
	data, err := readFrom(srcpath, conf.Body.Data)
	if err != nil {
		return err
	}
	body, err := buildBody(execpath, tpl, data)
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

func buildBody(relpath, tpl, data string) (string, error) {
	conf := cubano.Config{
		Files:[]string{
			"../../blender/render/Mint/mint.js",
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
		err = cubano.Run(filepath.Join(dir, relpath), conf
	}
	return retVal, err
}

type OutFile struct {
	file *os.File
}

func Create(path, filename string) (*OutFile, error) {
	fmt.Println("creating output file")
	outpath := filepath.Join(path, filename)
	fmt.Println(outpath)
	outfile, err := os.Create(outpath)
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
	fmt.Println("streaming from:")
	srcFile := filepath.Join(path,fileName)
	fmt.Println(srcFile)
	src, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	_, err = io.Copy(o.file, src)
	return err
}