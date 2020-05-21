package main

import (
	"github.com/scott-x/gutils/fs"
	"github.com/scott-x/gutils/cmd"
	"github.com/scott-x/gutils/cl"
	// "github.com/scott-x/gutils/cmd"
)

var TEMPLATE_FOLDER = fs.GetEnv("GOPATH")+"/src/github.com/scott-x/py/pkg"

func main() {
	cmd.AddQuestion("name", "What's your pip installed name: ", "Please input correct name: ", "^[a-z]+")
	answers := cmd.Exec()
    name := answers["name"]
    fs.CreateDirIfNotExist("./scott")
    err :=fs.CopyFolder(TEMPLATE_FOLDER,"./scott")
    if err!=nil{
    	panic(err)
    }
    fs.Rename("./scott/slimz","./scott/"+name)
    err = fs.ReadAndReplace("./scott/setup.py",map[string]string{
    	"slimz": name,
    })
    if err!=nil{
    	panic(err)
    }
    cmd.Info( name +" was created successfully, please make sure setuptools、 wheel、 twine are installed. If not, please install with following command:")
    cl.BoldCyan.Println("$ python3 -m pip install --user --upgrade setuptools wheel twine")
}
