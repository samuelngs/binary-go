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

var file = map[string][]string{
	<< range .Assets >>
	"<< .Filepath >>": []string{ << .Hashes >> },
	<< end >>
}

var data = map[string][]byte{
	<< range .Data >>"<< .Hash >>": << .Ref >>,<< end >>
}

var size = map[string]int{
	<< range .Sizes >>"<< .Hash >>": << .Size >>,<< end >>
}

// Get returns file in bytes format
func Get(filename string) ([]byte, error) {
	var b bytes.Buffer
	hashes, ok := file[filename]
	if !ok {
		return nil, errors.New("file does not exist")
	}
	length := size[filename]
	for _, hash := range hashes {
		b.Write(data[hash])
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

var file map[string][]string
var data map[string][]byte
var size map[string]int

// Get returns file in bytes format
func Get(filename string) ([]byte, error) {
	var b bytes.Buffer
	hashes, ok := file[filename]
	if !ok {
		return nil, errors.New("file does not exist")
	}
	length := size[filename]
	for _, hash := range hashes {
		b.Write(data[hash])
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
