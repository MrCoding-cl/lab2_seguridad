package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/kumarde/feistel"
	"log"
	"os"
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

func main() {

	//Exampling reading file
	file := readTxt("mensajedeentrada.txt") //read the text
	hash := sha1.New()                      //generate the sha1
	hash.Write([]byte(file[0]))             //using the sha1 function
	bs := hash.Sum(nil)                     //this is the hash of the original message
	println(file[0])                        //Print the original message
	println(bs)
	fmt.Printf("%x\n", bs) //print the hash

	cipher := feistel.New()
	encryptedMsg := cipher.Encrypt([]byte(file[0]))

	fmt.Printf("%x\n", encryptedMsg) //print the hash

	writeTxt(hex.EncodeToString(encryptedMsg))

	decryptedMsg := cipher.Decrypt(encryptedMsg)

	println(hex.EncodeToString(decryptedMsg))

}
