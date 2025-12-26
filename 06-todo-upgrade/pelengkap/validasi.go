package pelengkap

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"golang.org/x/text/language"
	"golang.org/x/text/message"

)

func ValidasiNama(prompt string, reader *bufio.Reader) string {
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Terjadi kesalahan!")
			continue
		}
		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("Inputan tidak boleh kosong!")
			continue
		}
		return input
	}
}

func ValidasiHarga(prompt string, reader *bufio.Reader) string {
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Terjadi kesalahan!")
			continue
		}
		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("Inputan tidak boleh kosong")
			continue
		}
		ipt, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Inputan harus berupa angka")
			continue
		}
		if ipt <= 0 {
			fmt.Println("Inputan harus lebih dari 0")
			continue
		}

		konversi := message.NewPrinter(language.Indonesian)
		formatted := konversi.Sprintf("%d",ipt)

		return  formatted
		
	}

}
func ValidasiStok(prompt string, reader *bufio.Reader) int {
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Terjadi kesalahan!")
			continue
		}
		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("Inputan tidak boleh kosong")
			continue
		}
		ipt, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Inputan harus berupa angka")
			continue
		}
		if ipt <= 0 {
			fmt.Println("Inputan harus lebih dari 0")
			continue
		}
		return ipt
		
	}

}

