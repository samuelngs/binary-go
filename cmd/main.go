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
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	flag.Parse()

	t := binary.Scan(*dir)
	c := binary.Compose(*pkg, t)

	fmt.Println(c)

}
