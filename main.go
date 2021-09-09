package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/cyrildever/feistel"
	"log"
	"os"
	"strings"
	"time"
)

func readTxt(name string) []string {
	//Function read a txt file and return a string list
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

func writeTxt(content string) {
	//This function make a txt with a string

	f, err := os.Create("mensajeseguro.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f.WriteString(content)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}

//We going to cipher the message with Feistel cipher permutation

func AcumulateString(file []string, backslash bool) string {
	ret := ""
	for _, s := range file {
		if backslash && len(file) > 1 {
			ret += s + "\n"
		} else {
			ret += s
		}
	}
	return ret
}
func getHash(text string) string {
	h := sha1.New()
	//encoded, _ := hex.DecodeString(text)
	h.Write([]byte(text))
	hash := hex.EncodeToString(h.Sum(nil))
	return hash
}

func goEncrypt(text string) string {
	//we set the paramet
	obfuscated, _ := cipher.Encrypt(text)
	return hex.EncodeToString(obfuscated)
}

func goDecode(text string) string {
	s, _ := hex.DecodeString(text)
	deciphered, _ := cipher.Decrypt(s)
	return deciphered
}

func cipherFile() {
	//Exampling reading file
	file := readTxt("mensajedeentrada.txt") //read the text 	//generate the sha1
	text := AcumulateString(file, true)
	hashtexto := getHash(text) //this is the hash of the original message
	BDD = append(BDD, hashtexto)
	text += "\n" + hashtexto
	// Encrypt
	originalEncrypted := goEncrypt(text)

	writeTxt(originalEncrypted) //We make the "Secure" file

}

func decryptFile() {
	secure := AcumulateString(readTxt("mensajeseguro.txt"), false)
	deciphered := goDecode(secure)

	deciphByLines := strings.Split(deciphered, "\n")
	acum := ""
	hash := ""
	for i, line := range deciphByLines {
		if len(deciphByLines)-1 != i {
			if len(deciphByLines)-2 == i {
				acum += line
			} else {
				acum += line + "\n"
			}

		} else {
			hash = line
		}
	}
	println("TEXTO DECIFRADO:\n" + acum + "\nHASH DETECTADO\n" + hash)
	hash2 := getHash(acum)
	if hash2 == hash {
		println("Si coincide el Hash generado")
		for _, s := range BDD {
			if s == hash {
				println("Y existe dentro de los hash generados dentro del programa")
				return
			}
		}
		println("WARNING!!\nNo existe dentro de la base de datos de los hash generados por el programa")
	} else {
		println("Algo ocurrio, el hash no coincide")
	}
}

func modifyFile() {
	text := AcumulateString(readTxt("mensajeseguro.txt"), false)
	writeTxt("sadasda" + text)
}
func modifyFileNotOriginal() {
	text := "6719435d0200175e02550c091e13580901560d5241440315526702090757010908020009525e0f0f0b020f5c515206575b5557505c550752540705040d0005535353"
	writeTxt(text)
}

func main() {
	cipherFile()
	decryptFile()
	time.Sleep(1 * time.Second)
	println("ERROR TEST INTEGRITY")
	modifyFile()
	decryptFile()
	time.Sleep(1 * time.Second)
	println("ERROR TEST ORIGINALITY")
	modifyFileNotOriginal()
	decryptFile()
}
