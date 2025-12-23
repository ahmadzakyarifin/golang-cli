package main

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"strings"
)

func main() {
	var (
		pilih      uint8
		jumlah     int
		lagi       string
		dataNama   []string
		dataHarga  []int
		dataTotal  []int
		dataJumlah []int
	)

	p := message.NewPrinter(language.Indonesian)

	for {
		fmt.Println("============= Daftar Barang =============")
		p.Printf("%-3s %-10s Rp%10d\n", "1.", "Kemeja", 80000)
		p.Printf("%-3s %-10s Rp%10d\n", "2.", "Celana", 70000)
		p.Printf("%-3s %-10s Rp%10d\n", "3.", "Jaz", 180000)
		p.Printf("%-3s %-10s Rp%10d\n", "4.", "Kaos", 40000)
		fmt.Println("=========================================")

		fmt.Print("Masukkan pilihan barang Anda (angka 1-4): ")
		_, err := fmt.Scan(&pilih)
		if err != nil {
			fmt.Println("Input tidak valid!")
			continue
		}

		var (
			nama  string
			harga int
			total int
		)

		switch pilih {
		case 1:
			nama, harga = "Kemeja", 80000
		case 2:
			nama, harga = "Celana", 70000
		case 3:
			nama, harga = "Jaz", 180000
		case 4:
			nama, harga = "Kaos", 40000
		default:
			fmt.Println("Pilihan tidak tersedia!")
			continue
		}

		
		for {
			fmt.Print("Masukkan jumlah pesanan Anda: ")
			_, err := fmt.Scan(&jumlah)
			if err != nil || jumlah <= 0 {
				fmt.Println("Jumlah tidak valid!")
				continue
			}else{
				total = jumlah * harga
				break
			}
			
		}

		dataNama = append(dataNama, nama)
		dataHarga = append(dataHarga, harga)
		dataTotal = append(dataTotal, total)
		dataJumlah = append(dataJumlah, jumlah)

		fmt.Print("Ingin pesan lagi (y/n)? ")
		fmt.Scan(&lagi)
		lagi = strings.ToLower(strings.TrimSpace(lagi))
		if lagi != "y" {
			break
		}
	}

	fmt.Println("\n================= Rincian Pesanan =================")
	fmt.Printf("%-3s %-15s %-10s %-12s %-12s\n", "No", "Nama", "Jumlah", "Harga", "Total")
	fmt.Println("---------------------------------------------------")
	totalSebelumDiskon := 0
	for i := range dataNama {
		p.Printf("%-3d %-15s %-10d Rp%10d Rp%10d\n", i+1, dataNama[i], dataJumlah[i], dataHarga[i], dataTotal[i])
		totalSebelumDiskon += dataTotal[i]
	}

	var diskonPersen float32
	switch {
	case totalSebelumDiskon >= 150000:
		diskonPersen = 0.10
	case totalSebelumDiskon >= 50000:
		diskonPersen = 0.05
	default:
		diskonPersen = 0
	}
	diskonRupiah := float32(totalSebelumDiskon) * diskonPersen
	totalSetelahDiskon := float32(totalSebelumDiskon) - diskonRupiah

	fmt.Println("---------------------------------------------------")
	fmt.Println("=============== Rincian Pembayaran ================")
	p.Printf("Total pembayaran sebelum diskon : Rp%s\n", p.Sprintf("%d", totalSebelumDiskon))
	p.Printf("Besar Diskon (%.0f%%)              : Rp%s\n", diskonPersen*100, p.Sprintf("%.0f", diskonRupiah))
	p.Printf("Total pembayaran setelah diskon : Rp%s\n", p.Sprintf("%.0f", totalSetelahDiskon))
	fmt.Println("===================================================")
}
