package main

import (
    "errors"
    "fmt"
    "unicode/utf8"
)


/* rune an integer that represents a Unicode code point */

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil}

func ReverseByte(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b), nil}


func main() {

	input := "Some sentence about life, today we've got Saturday, November"
	rev, revErr := Reverse(input)
	doubleRev, doublerevErr := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
	fmt.Printf("double reversed: %q, err: %v\n", doubleRev, doublerevErr)
	
}
