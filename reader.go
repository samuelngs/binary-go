package binary

import (
	"bytes"
	"compress/gzip"
	"errors"
	"log"
	"strconv"
	"strings"
)

var data map[string][]string
var size map[string]int

// Bytes to retrieve file data
func Bytes(filename string) ([]byte, error) {
	var r bytes.Buffer
	defer r.Truncate(0)
	part, ok := data[filename]
	if !ok {
		return nil, errors.New("file does not exist")
	}
	size, ok := size[filename]
	if !ok {
		return nil, errors.New("could not find file size information")
	}
	cont := strings.Join(part, "")
	arry := strings.Split(cont, " ")
	for _, v := range arry {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		b := byte(n)
		r.WriteByte(b)
	}
	data := make([]byte, size)
	bytr := bytes.NewReader(r.Bytes())
	gz, err := gzip.NewReader(bytr)
	if err != nil {
		return nil, err
	}
	defer gz.Close()
	if _, err := gz.Read(data); err != nil {
		return nil, err
	}
	return data, nil
}

// MustBytes to read bytes data from file
func MustBytes(filename string) []byte {
	b, err := Bytes(filename)
	if err != nil {
		log.Panic(err)
	}
	return b
}
