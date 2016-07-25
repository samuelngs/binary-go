package binary

import (
	"bytes"
	"io"
	"log"
	"os"
	"path/filepath"
)

// Scan returns tree node
func Scan(path string, limits ...int) *Tree {
	var limit = int(50 * MEGABYTE)
	for o := range limits {
		limit = o
		break
	}
	abs, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Open(abs)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	tree := new(Tree)
	tree.name = filepath.Dir(abs)
	tree.path = abs
	tree.dirs = make([]*Tree, 0)
	tree.assets = make([]*Asset, 0)
	read(tree, abs, limit)
	return tree
}

func read(dir *Tree, path string, limit int) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	switch mode := s.Mode(); {
	case mode.IsDir():
		di, err := f.Readdir(-1)
		if err != nil {
			log.Fatal(err)
		}
		for _, fi := range di {
			b := bytes.NewBuffer(make([]byte, 0))
			b.WriteString(dir.Pwd())
			b.WriteRune(os.PathSeparator)
			b.WriteString(fi.Name())
			abs := b.String()
			switch mode := fi.Mode(); {
			case mode.IsRegular():
				asset := NewAsset(abs, fi.Name(), nil, limit)
				f, err := os.Open(abs)
				if err != nil {
					log.Fatal(err)
				}
				asset.Load(readfile(f), limit)
				dir.assets = append(dir.assets, asset)
			case mode.IsDir():
				tree := new(Tree)
				tree.name = filepath.Dir(abs)
				tree.path = abs
				tree.dirs = make([]*Tree, 0)
				tree.assets = make([]*Asset, 0)
				tree.parent = dir
				dir.dirs = append(dir.dirs, tree)
				read(tree, abs, limit)
			}
		}
	case mode.IsRegular():
		dir.assets = append(dir.assets, NewAsset(dir.Pwd(), filepath.Base(f.Name()), readfile(f), limit))
	}
}

func readfile(f *os.File) []byte {
	fi, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, fi.Size())
	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}
	}
	return buf
}
