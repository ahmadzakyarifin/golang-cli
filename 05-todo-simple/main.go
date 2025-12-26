package main

import "fmt"

type Database struct{
	Nama string
	Harga uint32
	Stok uint8
}

func tampilData(database []Database){
	fmt.Printf("%-5s %-15s %-15s %-15s\n","No","Nama","Harga","Stok")
		for i, item := range database {
			fmt.Printf("%-5d %-15s %-15d %-15d\n", i+1, item.Nama, item.Harga,item.Stok)
		}
}


func main() {
	var nama string
	var harga uint32
	var stok uint8
	var pilih int8
	var nomer int
	var lagi string
	var edit int

	var database []Database
	for{
		fmt.Println("========== Daftar Program ==========")
		fmt.Println("1. Tambah Menu")
		fmt.Println("2.  Edit  Menu")
		fmt.Println("3. Tampil Menu")
		fmt.Println("4.  Hapus Menu")
		fmt.Println("5.  Keluar Program")


		fmt.Print("Masukkan pilihan Anda (1/2/3/4/5) : ")
		fmt.Scan(&pilih)

		switch pilih{
		case 1 :
			fmt.Print("Masukkan Nama : ")
			fmt.Scan(&nama)

			fmt.Print("Masukkan harga : ")
			fmt.Scan(&harga)

			fmt.Print("Masukkan stok : ")
			fmt.Scan(&stok)

			database = append(database,Database{Nama: nama,Harga: harga,Stok: stok})

		case 2 :
			

			for{

				tampilData(database)
				fmt.Print("Nomer berapa yang ingin Anda edit ? : ")
				fmt.Scan(&nomer) 
				if nomer <= 0 || nomer > len(database) {
			    	fmt.Println("Nomer tidak valid")
				}else{			
					data := &database[nomer-1]	
					
					fmt.Println("1. Nama")
					fmt.Println("2. Harga")
					fmt.Println("3. Stok")
					fmt.Print("Bagian Mana yang mau Anda edit (1/2/3) : ")
					fmt.Scan(&edit)
					switch edit{
					case 1:
						fmt.Print("Masukkan Nama : ")
						fmt.Scan(&nama)
						data.Nama =nama
					case 2:	
						fmt.Print("Masukkan harga : ")
						fmt.Scan(&harga)
						data.Harga = harga
					case 3:
						fmt.Print("Masukkan stok : ")
						fmt.Scan(&stok)
						data.Stok = stok
				}
				
			}

				fmt.Print("Edit lagi ? :  ")
				fmt.Scan(&lagi)
				if lagi == "n"{
					break
				}
			}


		case 3 : 
			if len(database) == 0 {
				fmt.Println("Masih belum ada data!")
			}else{
				tampilData(database)
			}
		case 4 :
			for{	
				tampilData(database)
				fmt.Print("Nomer berapa yang ingin Anda hapus ? :")
				fmt.Scan(&nomer)

				if nomer < 0 || nomer > len(database){
					fmt.Println("Data tidak valid!")
				}else{
					nomer = nomer-1
					database = append(database[:nomer],database[nomer+1:]... )
				}
				fmt.Print("Ingin hapus lagi ? :  ")
				fmt.Scan(&lagi)
					if lagi == "n"{
						break
					}
			}
		case 5 :
			fmt.Print("Program selesai")
			return
		}
	}
}

