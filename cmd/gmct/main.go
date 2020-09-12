package main

import (
	"fmt"
	"os"

	"github.com/snail007/gmct/template"
	"github.com/snail007/gmct/tool"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	gmctApp := kingpin.New("gmct", "toolchain for go web framework gmc, https://github.com/snail007/gmc")
	//all subtool args defined here
	templateArgs := template.NewTemplateArgs()

	//all subtool defined here
	//template subtool
	templateCmd := gmctApp.Command("tpl", "pack or clean templates go file")
	templateArgs.Dir = templateCmd.Flag("dir", "template's template directory path, gmct will convert all template files in the folder to one go file").Default(".").String()
	templateArgs.Extension = templateCmd.Flag("ext", "extension of template file").Default(".html").String()
	templateArgs.Clean = templateCmd.Flag("clean", "clean packed file, if exists").Default("false").Bool()

	if len(os.Args) == 0 {
		os.Args = []string{""}
		gmctApp.Usage(os.Args)
		return
	}

	subToolName, err := gmctApp.Parse(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}

	var gmcToolObj tool.GMCTool
	switch subToolName {
	case "tpl":
		gmcToolObj = template.NewTemplate()
	default:
		fmt.Printf("sub command '%s' not found", subToolName)
		return
	}
	err = gmcToolObj.Start(templateArgs)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
}
