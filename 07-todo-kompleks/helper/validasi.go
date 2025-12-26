package helper

import(
	"bufio"
	"fmt"
	"strings"
	"strconv"
)



func ValidasiString(prompt string, input *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	text, err := input.ReadString('\n')
	if err != nil {
		return "", err
	}
	text = strings.TrimSpace(text)
	if text == "" {
		return "", fmt.Errorf("input tidak boleh kosong")
	}
	return text, nil
}

func ValidasiInt(prompt string, input *bufio.Reader) (int, error) {
	fmt.Print(prompt)
	angka, err := input.ReadString('\n')
	if err != nil {
		return 0, err
	}
	angka = strings.TrimSpace(angka)
	if angka == "" {
		return 0, fmt.Errorf("input tidak boleh kosong")
	}
	agk, err := strconv.Atoi(angka)
	if err != nil {
		return 0, fmt.Errorf("input harus berupa angka")
	}
	return agk, nil
}

