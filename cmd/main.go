package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/samuelngs/binary-go"
)

var (
	dir = flag.String("dir", "./", "the file or directory path")
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	flag.Parse()

	t := binary.Scan(*dir)
	c := binary.Compose(t)

	fmt.Println("result: ", c)

}
