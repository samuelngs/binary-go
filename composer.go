package binary

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// Composer struct
type Composer struct {
	dir, out, pkg string
	max           int
	tree          *Tree
}

var blocks, reader *template.Template

func init() {
	b, err := template.New("blocks").Delims("<<", ">>").Parse(DataTmpl)
	if err != nil {
		log.Fatal(err)
	}
	blocks = b
	r, err := template.New("reader").Delims("<<", ">>").Parse(ReaderTmpl)
	if err != nil {
		log.Fatal(err)
	}
	reader = r
}

// NewComposer to create composer instance
func NewComposer(dir, out, pkg string, max int) *Composer {
	return &Composer{
		dir: dir,
		out: out,
		pkg: pkg,
		max: max,
	}
}

// Scan directory
func (v *Composer) Scan() error {
	t, err := Scan(v.dir)
	if err != nil {
		return err
	}
	v.tree = t
	return nil
}

// Compose to compose template
func (v *Composer) Compose() error {
	var parts []*Part
	for _, asset := range v.tree.Iter() {
		for _, part := range asset.Iter() {
			var exists bool
			for _, o := range parts {
				if o.Hash() == part.Hash() {
					exists = true
				}
			}
			if !exists {
				parts = append(parts, part)
			}
		}
	}
	var list = []Blocks{make(Blocks, 0)}
	for _, part := range parts {
		s := part.Size()
		w := bytes.NewBuffer(make([]byte, 0))
		for i, b := range part.Bytes() {
			w.WriteString(strconv.Itoa(int(b)))
			if i < s-1 {
				w.WriteRune(',')
			}
		}
		block := &Block{
			Package: v.pkg,
			Name:    part.File(),
			Hash:    part.Hash(),
			Size:    part.Size(),
			Data:    w.String(),
		}
		var write bool
		for i, item := range list {
			if write || item.Size()+part.Size() > v.max {
				continue
			}
			list[i] = append(list[i], block)
			write = true
			break
		}
		if !write {
			list = append(list, Blocks{block})
		}
	}
	var r bytes.Buffer
	for i, item := range list {
		var w bytes.Buffer
		var p bytes.Buffer
		blocks.Execute(&w, &Blob{
			Package: v.pkg,
			Blocks:  item,
		})
		r.Write(w.Bytes())
		p.WriteString(v.out)
		p.WriteRune(os.PathSeparator)
		p.WriteString(strconv.Itoa(i))
		p.WriteString(".binary.go")
		if err := WriteFile(p.String(), w.Bytes()); err != nil {
			return err
		}
	}
	d := &Data{
		Package: v.pkg,
		Files:   make([]*File, 0),
	}
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	for _, asset := range v.tree.Iter() {
		p, err := filepath.Rel(dir, asset.Filepath())
		if err != nil {
			return err
		}
		o := &File{
			Filepath: p,
			Hashes:   make([]string, 0),
			Size:     asset.Size(),
		}
		for _, part := range asset.Iter() {
			o.Hashes = append(o.Hashes, part.Hash())
		}
		d.Files = append(d.Files, o)
	}
	var w bytes.Buffer
	var p bytes.Buffer
	p.WriteString(v.out)
	p.WriteRune(os.PathSeparator)
	p.WriteString("binary.go")
	reader.Execute(&w, d)
	if err := WriteFile(p.String(), w.Bytes()); err != nil {
	}
	return nil
}

// WriteFile to write content to file
func WriteFile(path string, b []byte) error {
	err := ioutil.WriteFile(path, b, 0644)
	if err != nil {
		return err
	}
	return nil
}
