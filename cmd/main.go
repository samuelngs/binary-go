package main

import (
	"flag"
	"log"
	"runtime"

	"github.com/samuelngs/binary-go"
)

var (
	dir = flag.String("dir", "./", "the file or directory path")
	out = flag.String("out", "./", "the output directory path")
	pkg = flag.String("pkg", "binary", "the package name")
	max = flag.Int("max", int(20*binary.MEGABYTE), "the maximum size of each embedding binary")
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	flag.Parse()

	c := binary.NewComposer(*dir, *out, *pkg, *max)

	if err := c.Scan(); err != nil {
		log.Fatal(err)
	}

	if err := c.Compose(); err != nil {
		log.Fatal(err)
	}

}
