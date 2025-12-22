package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type mhs struct {
	nama string
	nim  string
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Nama : ")
	nama, _ := reader.ReadString('\n')
	nm := strings.TrimSpace(nama)

	fmt.Print("Masukkan Nim : ")
	nim, _ := reader.ReadString('\n')
	ni := strings.TrimSpace(nim)


	mahasiswa := mhs{
		nama: nm,
		nim: ni,
	}

	fmt.Printf("Nama: %s\nNim: %s\n", mahasiswa.nama, mahasiswa.nim)




}