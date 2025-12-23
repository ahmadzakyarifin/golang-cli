package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Menu struct {
	Kode  string
	Nama  string
	Harga int
}

type Pesanan struct {
	Nama   string
	Harga  int
	Jumlah int
	Total  int
}


var p = message.NewPrinter(language.Indonesian)

func dataMenu() []Menu {
	return []Menu{
		{"A", "Pulpen", 5000},
		{"B", "Pensil", 3000},
		{"C", "Penghapus", 2000},
		{"D", "Buku Tulis", 10000},
		{"E", "Penggaris", 4000},
	}
}

func tampilMenu(menu []Menu) {
	fmt.Println("=================== Daftar Menu =====================")
	fmt.Printf("%-5s %-20s %-10s\n", "Kode", "Nama Barang", "Harga")
	fmt.Println("-----------------------------------------------------")
	for _, item := range menu {
		p.Printf("%-5s %-20s Rp%d\n", item.Kode, item.Nama, item.Harga)
	}
	fmt.Println("=====================================================")
}

func validasiPilih(input *bufio.Reader, opsi []string) string {
	for {
		fmt.Print("Masukkan pilihan (", strings.Join(opsi, "/"), ") : ")
		pilih, _ := input.ReadString('\n')
		pilih = strings.TrimSpace(strings.ToUpper(pilih))
		if pilih == "" {
			fmt.Println("Input tidak boleh kosong!")
			continue
		}
		if len(pilih) != 1 || !unicode.IsLetter(rune(pilih[0])) {
			fmt.Println("Input harus 1 huruf.")
			continue
		}
		for _, item := range opsi {
			if item == pilih {
				return pilih
			}
		}
		fmt.Println("Pilihan tidak tersedia.")
	}
}

func validasiJumlah(input *bufio.Reader) int {
	for {
		fmt.Print("Masukkan jumlah pesanan : ")
		jumlah, _ := input.ReadString('\n')
		jumlah = strings.TrimSpace(jumlah)
		jml, err := strconv.Atoi(jumlah)
		if err != nil || jml <= 0 {
			fmt.Println("Input harus berupa angka positif.")
			continue
		}
		return jml
	}
}

func pilihMenu(menu []Menu, kode string, jumlah int) (Pesanan, bool) {
	for _, item := range menu {
		if item.Kode == kode {
			total := item.Harga * jumlah
			return Pesanan{
				Nama:   item.Nama,
				Harga:  item.Harga,
				Jumlah: jumlah,
				Total:  total,
			}, true
		}
	}
	return Pesanan{}, false
}

func cetakKeranjang(keranjang []Pesanan) int {
	var total int
	fmt.Println("\n============= Struk Pembelian =============")
	fmt.Printf("%-20s %-10s %-15s %-15s\n", "Nama Barang", "Jumlah", "Harga", "Total")
	fmt.Println("------------------------------------------------------")
	for _, item := range keranjang {
		p.Printf("%-20s %-10d Rp%-13d Rp%-13d\n", item.Nama, item.Jumlah, item.Harga, item.Total)
		total += item.Total
	}
	fmt.Println("------------------------------------------------------")
	return total
}

func hitungDiskon(total int) (int, int, int) {
	var diskonPersen int
	var diskonNominal int

	switch {
	case total > 150000:
		diskonPersen = 10
	case total > 50000:
		diskonPersen = 5
	default:
		diskonPersen = 0
	}

	diskonNominal = total * diskonPersen / 100
	bayar := total - diskonNominal
	return diskonPersen, diskonNominal, bayar
}

func cetakPembayaran(total int, diskonPersen int, diskonNominal int, bayar int) {
	fmt.Println("========= Rincian Pembayaran =========")
	p.Printf("%-25s : Rp%d\n", "Total Belanja", total)
	p.Printf("%-25s : %d%%\n", "Persentase Diskon", diskonPersen)
	p.Printf("%-25s : Rp%d\n", "Potongan Harga", diskonNominal)
	fmt.Println("--------------------------------------")
	p.Printf("%-25s : Rp%d\n", "TOTAL BAYAR", bayar)
	fmt.Println("======================================")
}

func main() {
	menu := dataMenu()
	tampilMenu(menu)

	keranjang := []Pesanan{}
	input := bufio.NewReader(os.Stdin)
	opsi := []string{"A", "B", "C", "D", "E"}

	for {
		kode := validasiPilih(input, opsi)
		jumlah := validasiJumlah(input)
		pesanan, ok := pilihMenu(menu, kode, jumlah)
		if ok {
			keranjang = append(keranjang, pesanan)
		}

		var lanjut string
        for {
            fmt.Print("Tambah lagi (y/n)? : ")
            lanjut, _ = input.ReadString('\n')
            lanjut = strings.TrimSpace(strings.ToLower(lanjut))

            if lanjut == "y" || lanjut == "n" {
                break 
            }
            fmt.Println("Input salah! Harap masukkan 'y' untuk ya atau 'n' untuk tidak.")
        }

        if lanjut == "n" {
            break
        }

	}

	total := cetakKeranjang(keranjang)
	persen, potong, bayar := hitungDiskon(total)
	cetakPembayaran(total, persen, potong, bayar)
}