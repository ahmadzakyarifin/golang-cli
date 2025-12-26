package fitur

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/ahmadzakyarifin/07-todo-kompleks/helper"
	"github.com/ahmadzakyarifin/07-todo-kompleks/model"
)

func Tambah(daftar *[]*model.Daftar, input *bufio.Reader) {

	for {
		nama, err := helper.ValidasiString("Masukkan nama: ", input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		harga, err :=helper.ValidasiInt("Masukkan harga: ", input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		stok, err := helper.ValidasiInt("Masukkan jumlah stok: ", input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		*daftar = append(*daftar, &model.Daftar{Nama: nama, Harga: harga, Stok: stok})

		Tampil(*daftar)

		lanjut, err := lagi("Ingin tambah lagi? (y/n): ", input)
		if err != nil || lanjut != "y" {
			break
		}
	}
}

func Edit(daftar []*model.Daftar, input *bufio.Reader) {
	

	for {
		if !Tampil(daftar) {
			break
		}

		ubah, err := helper.ValidasiString("Masukkan (nama/harga/stok): ", input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		nomer, err := helper.ValidasiInt("Masukkan nomor yang ingin diubah: ", input)
		if err != nil || nomer <= 0 || nomer > len(daftar) {
			fmt.Println("Nomor tidak valid")
			continue
		}

		idx := nomer - 1

		switch strings.ToLower(ubah) {
		case "nama":
			namaBaru, err := helper.ValidasiString("Masukkan nama baru: ", input)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			daftar[idx].Nama = namaBaru
		case "harga":
			hargaBaru, err := helper.ValidasiInt("Masukkan harga baru: ", input)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			daftar[idx].Harga = hargaBaru
		case "stok":
			stokBaru, err := helper.ValidasiInt("Masukkan stok baru: ", input)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			daftar[idx].Stok = stokBaru
		default:
			fmt.Println("Kolom tidak valid. Gunakan nama, harga, atau stok")
		}

		lanjut, err := lagi("Ingin edit lagi? (y/n): ", input)
		if err != nil || lanjut != "y" {
			break
		}
	}
}

func Hapus(daftar *[]*model.Daftar, input *bufio.Reader) {
	

	for {
		if !Tampil(*daftar) {
			break
		}

		nomer, err :=helper.ValidasiInt("Masukkan nomor yang ingin dihapus: ", input)
		if err != nil || nomer <= 0 || nomer > len(*daftar) {
			fmt.Println("Nomor tidak valid")
			continue
		}

		idx := nomer - 1
		*daftar = append((*daftar)[:idx], (*daftar)[idx+1:]...)

		lanjut, err := lagi("Ingin hapus lagi? (y/n): ", input)
		if err != nil || lanjut != "y" {
			break
		}
	}
}

func Tampil(daftar []*model.Daftar) bool {
	if len(daftar) == 0 {
		fmt.Println("Data belum ada")
		return false
	} else {
		fmt.Println("=========== Data Barang ===========")
		fmt.Printf("%-3s %-15s %-10s %-10s\n", "No", "Nama", "Harga", "Stok")
		for i, d := range daftar {
			fmt.Printf("%-3d %-15s %-10d %-10d\n", i+1, d.Nama, d.Harga, d.Stok)
		}
		fmt.Println("===================================")
		return true
	}
}

func lagi(prompt string, input *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	jawab, err := input.ReadString('\n')
	if err != nil {
		return "", err
	}
	jawab = strings.TrimSpace(jawab)
	if jawab != "y" && jawab != "n" {
		return "", fmt.Errorf("inputan harus y atau n")
	}
	return jawab, nil
}
