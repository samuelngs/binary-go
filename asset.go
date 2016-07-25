package binary

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"path/filepath"
)

// Asset holds file information
type Asset struct {
	path string
	name string
	data [][]byte
}

// NewAsset creates gzip compressed asset
func NewAsset(path, name string, content []byte, size int) *Asset {
	asset := &Asset{path, name, nil}
	if content != nil {
		asset.Load(content, size)
	}
	return asset
}

// Compress data
func (v *Asset) Compress(c []byte) []byte {
	var cmps bytes.Buffer
	zip := gzip.NewWriter(&cmps)
	zip.Write(c)
	zip.Close()
	fmt.Printf(" %-100.100v %10d KB => %10d KB\n", v.path, len(c), len(cmps.Bytes()))
	return cmps.Bytes()
}

// Load to set content data
func (v *Asset) Load(c []byte, size int) {
	byt := [][]byte{}
	cps := v.Compress(c)
	buf := bytes.NewBuffer(make([]byte, 0))
	for i, b := range cps {
		if i > 0 && i%size == 0 {
			byt = append(byt, buf.Bytes())
			buf = bytes.NewBuffer(make([]byte, 0))
		}
		buf.WriteByte(b)
	}
	byt = append(byt, buf.Bytes())
	v.data = byt
}

// Dirpath returns the directory path of the file
func (v *Asset) Dirpath() string {
	return filepath.Dir(v.path)
}

// Filename returns the filename with extension
func (v *Asset) Filename() string {
	return v.name
}

// Filepath returns the full file path
func (v *Asset) Filepath() string {
	return v.path
}

// Parts returns the total num of parts of the content
func (v *Asset) Parts() int {
	return len(v.data)
}

// Bytes returns asset content in bytes
func (v *Asset) Bytes() []byte {
	return bytes.Join(v.data, make([]byte, 0))
}

// Of returns splitted content by index
func (v *Asset) Of(i int) []byte {
	return v.data[i]
}
