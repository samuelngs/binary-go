package binary

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
)

// Scan returns tree node
func Scan(path string, limits ...int) (*Tree, error) {
	var limit = int(50 * MEGABYTE)
	for o := range limits {
		limit = o
		break
	}
	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(abs)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tree := new(Tree)
	tree.name = filepath.Dir(abs)
	tree.path = abs
	tree.dirs = make([]*Tree, 0)
	tree.assets = make([]*Asset, 0)
	read(tree, abs, limit)
	return tree, nil
}

func read(dir *Tree, path string, limit int) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	s, err := f.Stat()
	if err != nil {
		return err
	}
	switch mode := s.Mode(); {
	case mode.IsDir():
		di, err := f.Readdir(-1)
		if err != nil {
			return err
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
					return err
				}
				d, err := readfile(f)
				if err != nil {
					return err
				}
				asset.Load(d, limit)
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
		d, err := readfile(f)
		if err != nil {
			return err
		}
		dir.assets = append(dir.assets, NewAsset(dir.Pwd(), filepath.Base(f.Name()), d, limit))
	}
	return nil
}

func readfile(f *os.File) ([]byte, error) {
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}
	buf := make([]byte, fi.Size())
	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if n == 0 {
			break
		}
	}
	return buf, nil
}
