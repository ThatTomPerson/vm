package main

import (
	"io/ioutil"

	"github.com/davecgh/go-spew/spew"
	"github.com/ThatTomPerson/vm/pkg/compose/loader"
)

func main() {
	b, _ := ioutil.ReadFile("docker-compose.yml")

	c, _ := loader.ParseYAML(b)

	spew.Dump(c)
}
