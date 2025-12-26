package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ahmadzakyarifin/07-todo-kompleks/fitur"
	"github.com/ahmadzakyarifin/07-todo-kompleks/helper"
	"github.com/ahmadzakyarifin/07-todo-kompleks/model"
)



func main() {
	input := bufio.NewReader(os.Stdin)
	var daftar []*model.Daftar

	menu := map[int]func(){
		1: func() { fitur.Tambah(&daftar, input) },
		2: func() { fitur.Edit(daftar, input) },
		3: func() { fitur.Hapus(&daftar, input) },
		4: func() { fitur.Tampil(daftar) },
		5: func() { 
            fmt.Println("Terima kasih! Program selesai")
            os.Exit(0)
        },
	}

	for {
		fmt.Println("\n===== TodoList =====")
		fmt.Print("1. Tambah\n2. Edit\n3. Hapus\n4. Tampil\n5. Keluar\n")
		pilihan, err := helper.ValidasiInt("Masukkan pilihan Anda: ", input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		if fn, ada := menu[pilihan]; ada {
			fn()
		} else {
			fmt.Println("Pilihan tidak tersedia")
		}
	}
}
