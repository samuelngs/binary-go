# binary-go
Embedding binary data in a Go program

## Features

* Splitting files (into smaller pieces)
* Compression

## Install

Install the latest release on Mac or Linux:
```
go get -u github.com/samuelngs/binary-go/...
```

## Usage

Convert binary data to `.go` files
```sh
$ binary -dir ./test -out ./output -pkg test -max 300
┌───────┬────────────────┬──────────────────────────────────────────┬──────────────────────────┐
│  PART │           SIZE │ HASH                                     │ FILE                     │
├───────┼────────────────┼──────────────────────────────────────────┼──────────────────────────┤
│    01 │         276 KB │ 9e0d4c6ecf7afd6152bd0950ea9a22fbaeb2f58e │ README.md-1.go           │
│    01 │         304 KB │ 9b4e8073b480e2118b3ac6a73eae376a128359e4 │ LICENSE-1.go             │
│    02 │         304 KB │ 578fa140625eef79cb04f12c992d36579da0ba5e │ LICENSE-2.go             │
│    03 │         304 KB │ b3f1b26db7b0d82ad437b324128c3aaf832ea96c │ LICENSE-3.go             │
│    04 │         304 KB │ 54ca42b211e9e3120ba4a1d040d6bce7d18bcb6e │ LICENSE-4.go             │
│    05 │         304 KB │ d0a4f769fe4a235c50c48cebe76d77c1c04e0a58 │ LICENSE-5.go             │
│    06 │         304 KB │ f43cca28e65ab4e88f478784f0c84b78ec8db356 │ LICENSE-6.go             │
│    07 │         304 KB │ 7346e2c8f45188d9075cef877526f0d200917508 │ LICENSE-7.go             │
│    08 │         304 KB │ 85366250d7bd91843ff134ee945eae4405b887b0 │ LICENSE-8.go             │
│    09 │         192 KB │ c452da088bc03763ebc190a377766c1575fda3e8 │ LICENSE-9.go             │
└───────┴────────────────┴──────────────────────────────────────────┴──────────────────────────┘
```

Output
```
package test

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"log"
)

var data = map[string][]string{
	"test/README.md": []string{
		d9e0d4c6ecf7afd6152bd0950ea9a22fbaeb2f58e,
	},
	"test/LICENSE": []string{
		d9b4e8073b480e2118b3ac6a73eae376a128359e4, d578fa140625eef79cb04f12c992d36579da0ba5e, db3f1b26db7b0d82ad437b324128c3aaf832ea96c, d54ca42b211e9e3120ba4a1d040d6bce7d18bcb6e, dd0a4f769fe4a235c50c48cebe76d77c1c04e0a58, df43cca28e65ab4e88f478784f0c84b78ec8db356, d7346e2c8f45188d9075cef877526f0d200917508, d85366250d7bd91843ff134ee945eae4405b887b0, dc452da088bc03763ebc190a377766c1575fda3e8,
	},
}

var d9b4e8073b480e2118b3ac6a73eae376a128359e4 = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x52\xcf\x8e\xda\x3c\x10\xbf\xfb\x29\x46\x9c\x76\xa5\x68\xbf\xaf\x3d\xf4\xd0\x9b\x49\xcc\x62\x35\xc4\x91\x13\x96\x72\x34\x89\x21\xae\x42\x8c\x62\xa7\x68\xdf\xbe\x33\x81\xdd\xed\x56\x42\x42\x33\x9e\xdf\xbf\x99\xd4\x9d\x85\x8d\xac\x21\x77\x8d\x1d\x82\x85\x07\x2c"
var d578fa140625eef79cb04f12c992d36579da0ba5e = "\x1e\x19\x4b\xfd\xe5\x75\x74\xa7\x2e\xc2\x43\xf3\x08\x5f\xff\xff\xf2\x0d\x2a\x73\x9e\x6c\xcf\x58\x69\xc7\xb3\x0b\xc1\xf9\x01\x5c\x80\xce\x8e\xf6\xf0\x0a\xa7\xd1\x0c\xd1\xb6\x09\x1c\x47\x6b\xc1\x1f\xa1\xe9\xcc\x78\xb2\x09\x44\x0f\x66\x78\x85\x8b\x1d\x03\x02\xfc\x21\x1a\x37\xb8\xe1\x04\x06\x1a\x94\x60\x38"

...

```

## Documentation

`go doc` format documentation for this project can be viewed online without installing the package by using the GoDoc page at: https://godoc.org/github.com/samuelngs/binary-go

## Contributing

Everyone is encouraged to help improve this project. Here are a few ways you can help:

- [Report bugs](https://github.com/samuelngs/binary-go/issues)
- Fix bugs and [submit pull requests](https://github.com/samuelngs/binary-go/pulls)
- Write, clarify, or fix documentation
- Suggest or add new features

## License ##

This project is distributed under the MIT license found in the [LICENSE](./LICENSE)
file.

```
The MIT License (MIT)

Copyright (c) 2016 Samuel

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
