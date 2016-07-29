package binary

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"os"
	"path/filepath"
)

// Asset holds file information
type Asset struct {
	path string
	name string
}

// NewAsset creates gzip compressed asset
func NewAsset(path, name string) *Asset {
	return &Asset{path, name}
}

// Dirpath returns the directory path of the file
func (v *Asset) Dirpath() string {
	return filepath.Dir(v.path)
}

// Filename returns the filename with extension
func (v *Asset) Filename() string {
	return v.name
}

// Relpath returns file relative path
func (v *Asset) Relpath() string {
	pwd, err := os.Getwd()
	if err != nil {
		return v.path
	}
	rel, err := filepath.Rel(pwd, v.path)
	if err != nil {
		return v.path
	}
	return rel
}

// Filepath returns the full file path
func (v *Asset) Filepath() string {
	return v.path
}

// Pipe data
func (v *Asset) Pipe() <-chan []byte {
	fe, err := os.Open(v.path)
	if err != nil {
		log.Fatal(err)
	}
	fi, err := fe.Stat()
	if err != nil {
		log.Fatal(err)
	}
	rd := bufio.NewReader(fe)
	ch := make(chan []byte, 1)
	go func() {
		defer fe.Close()
		defer close(ch)
		buf := make([]byte, fi.Size())
		_, err := rd.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		ch <- buf
	}()
	return ch
}

// Size returns asset size
func (v *Asset) Size() int64 {
	fe, err := os.Open(v.path)
	if err != nil {
		log.Fatal(err)
	}
	defer fe.Close()
	fi, err := fe.Stat()
	if err != nil {
		log.Fatal(err)
	}
	return fi.Size()
}

// Md5 returns md5 hash
func (v *Asset) Md5() (string, error) {
	fe, err := os.Open(v.path)
	if err != nil {
		return "", err
	}
	fi, err := fe.Stat()
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan string, 1)
	go func() {
		defer fe.Close()
		b := make([]byte, fi.Size())
		fe.Read(b)
		sum := md5.New()
		sum.Write(b)
		ch <- hex.EncodeToString(sum.Sum(nil))
	}()
	return <-ch, nil
}
