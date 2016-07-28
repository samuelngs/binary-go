package binary

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"log"
	"strings"
)

var data map[string][]string

// Bytes to retrieve file data
func Bytes(filename string) (b []byte, e error) {
	var r bytes.Buffer
	defer r.Truncate(0)
	part, ok := data[filename]
	if !ok {
		return nil, errors.New("file does not exist")
	}
	ch := make(chan []byte, 1)
	go func() {
		cont := strings.Join(part, "")
		data := []byte(cont)
		gz, err := gzip.NewReader(bytes.NewBuffer(data))
		if err != nil {
			e = err
			ch <- nil
			return
		}
		var buf bytes.Buffer
		if _, err = io.Copy(&buf, gz); err != nil {
			e = err
			ch <- nil
			return
		}
		if err := gz.Close(); err != nil {
			e = err
			ch <- nil
			return
		}
		ch <- buf.Bytes()
	}()
	return <-ch, nil
}

// MustBytes to read bytes data from file
func MustBytes(filename string) []byte {
	b, err := Bytes(filename)
	if err != nil {
		log.Panic(err)
	}
	return b
}
