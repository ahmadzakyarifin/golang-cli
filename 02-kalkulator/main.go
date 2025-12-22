package main

import "fmt"

func Tambah(a, b float64) float64 {
	return a + b 
}
func Kurang(a, b float64) float64 {
	return a - b 
}
func Kali(a, b float64) float64 {
	return a * b
}
func Bagi(a, b float64) float64 {
	if b == 0 {
		fmt.Println("Error: Tidak dapat membagi dengan 0")
		return 0
	}
	return a / b
}

func Proses(a, b float64, Operasi func(float64, float64) float64) float64 {
	return Operasi(a, b)
}

var daftar = map[string]func(float64, float64) float64{
	"+": Tambah,
	"-": Kurang,
	"*": Kali,
	"/": Bagi,
}

func main() {
	var (
		a       float64
		b       float64
		operasi string
		lagi    string
	)

	fmt.Println("=== Kalkulator Mini ===")

	for {
		fmt.Println("\nDaftar Operator Tersedia :")
		for k := range daftar {
			fmt.Printf("[%s] ", k)
		}
		fmt.Println()

		fmt.Print("Silahkan masukkan angka a : ")
		fmt.Scan(&a)
		fmt.Print("Silahkan masukkan operator: ")
		fmt.Scan(&operasi)
		fmt.Print("Silahkan masukkan angka b : ")
		fmt.Scan(&b)

		pilih, ok := daftar[operasi]
		if ok {
			hasil := Proses(a, b, pilih)
			fmt.Printf("Hasil: %.2f\n", hasil)
		} else {
			fmt.Println("Operator tidak valid!")
		}

		fmt.Print("\nApakah Anda ingin melanjutkan? (y/n) : ")
		fmt.Scan(&lagi)
		if lagi != "y" && lagi != "Y" {
			fmt.Println("Terima kasih telah menggunakan kalkulator mini kami")
			break
		}
	}
}