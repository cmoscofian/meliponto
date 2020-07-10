package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

var docs string = "src/docs"

func PrintUsage() {
	var bs []byte
	var err error

	if len(os.Args) > 1 {
		bs, err = ioutil.ReadFile(path.Join(docs, os.Args[1]))
		if err != nil {
			bs, _ = ioutil.ReadFile(path.Join(docs, "default"))
		}
	} else {
		bs, _ = ioutil.ReadFile(path.Join(docs, "default"))
	}

	fmt.Println(string(bs))
}
