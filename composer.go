package binary

import (
	"bytes"
	"html/template"
	"log"
	"strconv"
)

// Blocks type
type Blocks []*Block

// Size returns blocks size
func (v Blocks) Size() int {
	var s int
	for _, b := range v {
		s += b.Size
	}
	return s
}

// Block struct
type Block struct {
	Package string
	Hash    string
	Data    string
	Size    int
}

// Blob data
type Blob struct {
	Package string
	Blocks  Blocks
}

var blocks, reader *template.Template

func init() {
	b, err := template.New("blocks").Delims("<<", ">>").Parse(DataTmpl)
	if err != nil {
		log.Panic(err)
	}
	blocks = b
	r, err := template.New("reader").Delims("<<", ">>").Parse(ReaderTmpl)
	if err != nil {
		log.Panic(err)
	}
	reader = r
}

// Compose to compose template
func Compose(pkg string, max int, tree *Tree) string {
	var parts []*Part
	for _, asset := range tree.Iter() {
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
			Package: pkg,
			Hash:    part.Hash(),
			Size:    part.Size(),
			Data:    w.String(),
		}
		var write bool
		for i, item := range list {
			if write || item.Size()+part.Size() > max {
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
	for _, item := range list {
		blocks.Execute(&r, &Blob{
			Package: pkg,
			Blocks:  item,
		})
	}
	return r.String()
}
