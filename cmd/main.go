package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/samuelngs/binary-go"
)

var (
	dir = flag.String("dir", "./", "the file or directory path")
	pkg = flag.String("pkg", "binary", "the package name")
	max = flag.Int("max", int(50*binary.MEGABYTE), "the maximum size of each embedding binary")
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	flag.Parse()

	t := binary.Scan(*dir)
	c := binary.Compose(*pkg, *max, t)

	// log.Print(c)
	fmt.Println(c)

}
