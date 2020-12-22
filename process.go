package main

import (
	"fmt"
	"io"
	"os"
)

func processFile(filePath string) (err error) {
	var in *os.File
	in, err = os.Open(filePath)
	if err != nil {
		return
	}
	defer func() {
		err2 := in.Close()
		if err != nil {
			err = err2
		}
	}()

	buf := make([]byte, max)

	var pos int64
	var n int
	for {
		//fmt.Println("pos =", pos)
		_, err = in.Seek(pos, 0)
		if err != nil {
			break
		}

		n, err = in.Read(buf)
		if err != nil {
			break
		}
		if n < min {
			break
		}
		ext, skip, ok := checkBytes(buf)
		if ok {
			fmt.Printf("%d %s\n", pos, ext)
			pos += int64(skip)
			//break
		}

		pos++

	}

	if err == io.EOF {
		err = nil
	}

	return
}

func checkBytes(bb []byte) (ext string, skip int, ok bool) {
	for e, pp := range data {
		for _, p := range pp {
			if matchBytes(bb, p) {
				ext = e
				skip = len(p)
				ok = true
				return
			}
		}
	}
	return
}

func matchBytes(aa, bb []byte) bool {
	if len(aa) < len(bb) {
		return false
	}
	for i, b := range bb {
		if aa[i] != b {
			return false
		}
	}
	return true
}
