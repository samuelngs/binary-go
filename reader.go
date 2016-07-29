package binary

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"log"
)

var data map[string][]string

// buffer data
type buffer struct {
	err  error
	data []byte
}

// Bytes to retrieve file data
func Bytes(filename string) ([]byte, error) {
	var r bytes.Buffer
	defer r.Truncate(0)
	parts, ok := data[filename]
	if !ok {
		return nil, errors.New("file does not exist")
	}
	ch := make(chan *buffer, 1)
	go func() {
		var zbuf bytes.Buffer
		defer zbuf.Truncate(0)
		for _, s := range parts {
			zbuf.Write([]byte(s))
		}
		gz, err := gzip.NewReader(&zbuf)
		if err != nil {
			ch <- &buffer{
				err: err,
			}
			return
		}
		var buf bytes.Buffer
		defer buf.Truncate(0)
		if _, err = io.Copy(&buf, gz); err != nil {
			ch <- &buffer{
				err: err,
			}
			return
		}
		if err := gz.Close(); err != nil {
			ch <- &buffer{
				err: err,
			}
			return
		}
		ch <- &buffer{
			data: buf.Bytes(),
		}
	}()
	res := <-ch
	if err := res.err; err != nil {
		return nil, res.err
	}
	return res.data, nil
}

// MustBytes to read bytes data from file
func MustBytes(filename string) []byte {
	b, err := Bytes(filename)
	if err != nil {
		log.Panic(err)
	}
	return b
}
