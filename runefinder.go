package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// AnalisarLinha analinha uma linha e retorna a runa e o nome dos caracteres Unicode
func AnalisarLinha(ucdLine string) (rune, string) {
	campos := strings.Split(ucdLine, ";") 
	codigo, _ := strconv.ParseInt(campos[0], 16, 32)
	return rune(codigo), campos[1]
}

// Listar exibe na saída padrão o código, a runa e o nome dos caracteres Unicode
// cujo nome contém o texto da consulta //
func Listar(texto io.Reader, consulta string) {
	varredor := bufio.NewScanner(texto)
	for varredor.Scan() { 
		linha := varredor.Text()
		if strings.TrimSpace(linha) == "" { 
			continue
		}
		runa, nome := AnalisarLinha(linha)
		if strings.Contains(nome, consulta) {
			fmt.Printf("U+%04X\t%[1]c\t%s\n", runa, nome)
		}
	}
}

func main() {
	ucd, err := os.Open("UnicodeData.txt") 
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func() { ucd.Close() }() 
	consulta := strings.Join(os.Args[1:], " ")
	Listar(ucd, strings.ToUpper(consulta))
}