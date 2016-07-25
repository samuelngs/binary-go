package binary

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"strconv"
)

var blocks, reader *template.Template

func init() {
	b, err := template.New("blocks").Parse(DataTmpl)
	if err != nil {
		log.Panic(err)
	}
	blocks = b
	r, err := template.New("reader").Parse(ReaderTmpl)
	if err != nil {
		log.Panic(err)
	}
	reader = r
}

// Compose to compose template
func Compose(pkg string, tree *Tree) string {
	for _, asset := range tree.Iter() {
		for _, part := range asset.Iter() {
			w := bytes.NewBuffer(make([]byte, 0))
			l := len(part.Bytes())
			for i, b := range part.Bytes() {
				w.WriteString(strconv.Itoa(int(b)))
				if i < l-1 {
					w.WriteRune(',')
					w.WriteRune(' ')
				}
			}
			blocks.Execute(os.Stdout, struct {
				Package string
				Hash    string
				Data    string
			}{
				Package: pkg,
				Hash:    part.Hash(),
				Data:    w.String(),
			})
			// fmt.Println(part.Hash())
		}
	}
	return ""
}
