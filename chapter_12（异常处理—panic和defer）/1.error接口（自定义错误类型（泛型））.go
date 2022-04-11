package main

import "fmt"

type ParseError struct {
	Filename string //文件名
	Line     int    //行号
}

func (p *ParseError) Error() string {
	return fmt.Sprintf("%s:%d", p.Filename, p.Line)
}

func newParseError(filename string, line int) error {
	return &ParseError{filename, line}
}

func main() {
	e := newParseError("maingo", 1)
	fmt.Println(e.Error())
}
