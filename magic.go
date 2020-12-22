package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
//max = 6
//min = 4
//min = 2
)

/*
var data = map[string][][]byte{
	"bmp": [][]byte{[]byte{0x42, 0x4D}},
	"jpg": [][]byte{[]byte{0xFF, 0xD8}},
	//"jpg": [][]byte{[]byte{0xFF, 0xD8, 0xDD, 0xE0}, []byte{0xFF, 0xD8, 0xFF, 0xFE}},
	"png": [][]byte{[]byte{0x89, 0x50, 0x4E, 0x47}},
	"gif": [][]byte{[]byte{0x47, 0x49, 0x46, 0x38, 0x39, 0x61}, []byte{0x47, 0x49, 0x46, 0x38, 0x37, 0x61}},
}
*/

var min, max int
var data map[string][][]byte

func loadMagic(filePath string) (err error) {
	var bb []byte
	bb, err = ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	err = json.Unmarshal(bb, &data)
	if err != nil {
		return
	}
	for _, pp := range data {
		for _, p := range pp {
			l := len(p)
			if l == 0 {
				err = fmt.Errorf("abnormal magic")
				return
			}
			if l > max {
				max = l
			}
			if min == 0 {
				min = l
			} else {
				if l < min {
					min = l
				}
			}
		}
	}
	return
}
