package main

import (
	"flag"
	"github.com/go-generator/demo_cmd/core"
	"log"
)

func main() {
	var t = flag.String("t", "template", "template folder")
	var i = flag.String("i", "javaWithModel.json", "input json")
	var o = flag.String("o", "javaOutput", "output folder")
	var p = flag.String("p", "projectName", "project name")
	var l = flag.String("l", "java", "coding language")

	flag.Parse()
	result := ""
	output, err := core.GenerateFromFile(*t, *p, *i, &result, *l)
	if err != "" {
		log.Panicln(err)
	}
	err = core.OutputStructToFiles(*o, output)
	if err != "" {
		log.Panicln(err)
	}
	log.Print("Successful")

}
