package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ahmadzakyarifin/06-todo-upgrade/pelengkap"
)

type Daftar struct {
	No   int
	Menu string
}

type Database struct {
	Nama  string
	Harga string 
	Stok  int
}

func data() []Daftar {
	return []Daftar{
		{1, "Tambah Barang"},
		{2, "Tampil Barang"},
		{3, "Edit Barang"},
		{4, "Hapus Barang"},
		{5, "Keluar"},
	}
}

func tampilMenu(menu []Daftar) {
	fmt.Println("\n========== MENU ==========")
	for _, item := range menu {
		fmt.Printf("%d. %s\n", item.No, item.Menu)
	}
	fmt.Println("==========================")
}

func tambahBarang(reader *bufio.Reader, database []Database) []Database {
	nama := pelengkap.ValidasiNama("Masukkan nama : ", reader)
	harga := pelengkap.ValidasiHarga("Masukkan harga : ", reader)
	stok := pelengkap.ValidasiStok("Masukkan stok : ", reader)

	fmt.Println("Barang berhasil ditambahkan.")
	return append(database, Database{Nama: nama, Harga: harga, Stok: stok})
}

func tampilBarang(database []Database) bool {
	if len(database) == 0 {
		fmt.Println("\n[!] Belum ada data barang.")
		return false
	}
	fmt.Println("\n========= DATA BARANG =========")
	fmt.Printf("%-3s | %-15s | %-12s | %-5s\n", "No", "Nama", "Harga", "Stok")
	fmt.Println("-----------------------------------------------")
	for i, item := range database {
		fmt.Printf("%-3d | %-15s | Rp%-10s | %-5d\n", i+1, item.Nama, item.Harga, item.Stok)
	}
	fmt.Println("================================")
	return true
}

func editBarang(reader *bufio.Reader, database []Database) {
	if !tampilBarang(database) {
		return
	}

	fmt.Print("Masukkan nomor barang yang ingin diedit: ")
	var no int
	fmt.Scan(&no)

	reader.ReadString('\n') 

	if no <= 0 || no > len(database) {
		fmt.Println("[!] Nomor tidak valid.")
		return
	}

	idx := no - 1
	fmt.Println("Apa yang ingin diubah? (nama/harga/stok)")
	fmt.Print("Pilihan: ")
	pilihan, _ := reader.ReadString('\n')
	pilihan = strings.TrimSpace(strings.ToLower(pilihan))

	switch pilihan {
	case "nama":
		database[idx].Nama = pelengkap.ValidasiNama("Masukkan nama baru: ", reader)
	case "harga":
		database[idx].Harga = pelengkap.ValidasiHarga("Masukkan harga baru: ", reader)
	case "stok":
		database[idx].Stok = pelengkap.ValidasiStok("Masukkan stok baru: ", reader)
	default:
		fmt.Println("[!] Pilihan kolom tidak valid.")
		return
	}
	fmt.Println("Data berhasil diperbarui.")
}


func hapusBarang(database []Database) []Database {
	if !tampilBarang(database) {
		return database
	}

	fmt.Print("Masukkan nomor barang yang ingin dihapus: ")
	var no int
	fmt.Scan(&no)

	if no <= 0 || no > len(database) {
		fmt.Println("[!] Nomor tidak valid.")
		return database
	}

	idx := no - 1

	database = append(database[:idx], database[idx+1:]...)
	fmt.Println("Barang berhasil dihapus.")
	return database
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	menu := data()
	var database []Database

	for {
		tampilMenu(menu)
		fmt.Print("Pilih menu (1-5): ")
		var pilihan int
		fmt.Scan(&pilihan)

		reader.ReadString('\n')

		switch pilihan {
		case 1:
			database = tambahBarang(reader, database)
		case 2:
			tampilBarang(database)
		case 3:
			editBarang(reader, database)
		case 4:
			database = hapusBarang(database)
		case 5:
			fmt.Println("Terima kasih telah menggunakan program!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}