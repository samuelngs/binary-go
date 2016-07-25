package binary

import (
	"bytes"
	"compress/gzip"
	"errors"
	"log"
)

// DataTmpl template
const DataTmpl string = `package << .Package >>
<< range .Blocks >>
var d<< .Hash >> = []byte{<< .Data >>}
<< end >>
`

// ReaderTmpl template
const ReaderTmpl string = `
package << .Package >>

import (
	"bytes"
	"compress/gzip"
	"errors"
	"log"
)

var file = map[string][][]byte{
	<< range .Files ->>"<< .Filepath >>": [][]byte{
		<< range .Hashes ->>
		d<< . >>,
		<<- end >>
	},
	<< end >>
}

var size = map[string]int{
	<< range .Files ->>
	"<< .Filepath >>": << .Size >>,
	<< end >>
}

// Get returns file in bytes format
func Get(filename string) ([]byte, error) {
	var b bytes.Buffer
	data, ok := file[filename]
	if !ok {
		return nil, errors.New("file does not exist")
	}
	length := size[filename]
	for _, part := range data {
		b.Write(part)
	}
	res := make([]byte, length)
	r := bytes.NewReader(b.Bytes())
	gz, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	defer gz.Close()
	if _, err := gz.Read(res); err != nil {
		return nil, err
	}
	return res, nil
}

// MustGet to retrieve data, panic if fail
func MustGet(filename string) []byte {
	b, err := Get(filename)
	if err != nil {
		log.Panic(err)
	}
	return b
}

`

// Data struct
type Data struct {
	Package string
	Files   []*File
}

// File struct
type File struct {
	Filepath string
	Hashes   []string
	Size     int
}

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
	Name    string
	Hash    string
	Data    string
	Size    int
}

// Blob data
type Blob struct {
	Package string
	Blocks  Blocks
}

var file map[string][][]byte
var size map[string]int

// Get returns file in bytes format
func Get(filename string) ([]byte, error) {
	var b bytes.Buffer
	data, ok := file[filename]
	if !ok {
		return nil, errors.New("file does not exist")
	}
	length := size[filename]
	for _, part := range data {
		b.Write(part)
	}
	res := make([]byte, length)
	r := bytes.NewReader(b.Bytes())
	gz, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	defer gz.Close()
	if _, err := gz.Read(res); err != nil {
		return nil, err
	}
	return res, nil
}

// MustGet to retrieve data, panic if fail
func MustGet(filename string) []byte {
	b, err := Get(filename)
	if err != nil {
		log.Panic(err)
	}
	return b
}
