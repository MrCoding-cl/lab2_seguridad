package main

import (
	"bufio"
	"log"
	"os"
)

func readTxt(name string) []string{
	file, err := os.Open(name)
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()

	return txtlines
}

func main(){
	file:=readTxt("mensajedeentrada.txt")
	for _, line := range file {
		print(line)
	}
}