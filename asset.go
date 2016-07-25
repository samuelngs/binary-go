package binary

import (
	"bytes"
	"compress/gzip"
	"crypto/sha1"
	"encoding/hex"
	"log"
	"path/filepath"
)

// Part holds part of the compressed data
type Part struct {
	file string
	data []byte
	size int
	hash string
}

// File returns filename
func (v *Part) File() string {
	return v.file
}

// Bytes returns part of the data
func (v *Part) Bytes() []byte {
	return v.data
}

// Size returns byte size
func (v *Part) Size() int {
	return v.size
}

// Hash returns data hash
func (v *Part) Hash() string {
	return v.hash
}

// Asset holds file information
type Asset struct {
	path string
	name string
	data [][]byte
	size int
}

// NewAsset creates gzip compressed asset
func NewAsset(path, name string, content []byte, max int) *Asset {
	asset := &Asset{path, name, nil, 0}
	if content != nil {
		asset.Load(content, max)
	}
	return asset
}

// Compress data
func (v *Asset) Compress(c []byte) []byte {
	var cmps bytes.Buffer
	zip := gzip.NewWriter(&cmps)
	zip.Write(c)
	zip.Close()
	return cmps.Bytes()
}

// Load to set content data
func (v *Asset) Load(c []byte, max int) {
	v.size = len(c)
	byt := [][]byte{}
	cps := v.Compress(c)
	buf := bytes.NewBuffer(make([]byte, 0))
	for i, b := range cps {
		if i > 0 && i%max == 0 {
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

// Size returns asset size
func (v *Asset) Size() int {
	return v.size
}

// Of returns splitted content by index
func (v *Asset) Of(i int) []byte {
	return v.data[i]
}

// Iter for-loop read data
func (v *Asset) Iter() []*Part {
	var l []*Part
	for _, b := range v.data {
		h := sha1.Sum(b)
		l = append(l, &Part{
			file: v.name,
			data: b,
			size: len(b),
			hash: hex.EncodeToString(h[:]),
		})
	}
	return l
}

// String returns string data
func (v *Asset) String() string {
	b := make([]byte, v.size)
	r := bytes.NewReader(v.Bytes())
	gz, err := gzip.NewReader(r)
	if err != nil {
		log.Fatal(err)
	}
	defer gz.Close()
	if _, err := gz.Read(b); err != nil {
		log.Fatal(err)
	}
	return string(b[:])
}
