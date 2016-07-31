package main

import (
	"flag"
	"log"

	"github.com/samuelngs/binary-go"
)

var (
	dir = flag.String("dir", "./", "the file or directory path")
	out = flag.String("out", "./", "the output directory path")
	pkg = flag.String("pkg", "binary", "the package name")
	max = flag.Int("max", int(20*binary.MEGABYTE), "the maximum size of each embedding binary (default: 20MB)")
)

func main() {

	flag.Parse()

	c := binary.NewComposer(*dir, *out, *pkg, *max)

	if err := c.Scan(); err != nil {
		log.Fatal(err)
	}

	err := make(chan error, 1)
	go c.Compile(err)

	if e := <-err; e != nil {
		log.Fatal(err)
	}

}
