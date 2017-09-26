package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

func AnalisarLinha(ucdLine string) (rune, string) {
	campos := strings.Split(ucdLine, ";") 
	código, _ := strconv.ParseInt(campos[0], 16, 32)
	return rune(código), campos[1]
}

func Listar(texto io.Reader, consulta string) {
	runa, nome := '?', "QUESTION MARK"
	fmt.Printf("U+%04X\t%[1]c\t%s\n", runa, nome)
}