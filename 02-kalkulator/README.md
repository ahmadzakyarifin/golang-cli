# 🧮 Kalkulator Mini (Latihan 02)

Project ini merupakan implementasi kalkulator interaktif yang menggunakan konsep **Higher-Order Functions** dan **Functional Programming** di bahasa pemrograman Go.

---

## 🏗️ Fitur dan Konsep Utama

### 1. Functional Programming (First-Class Functions)
Aplikasi ini memperlakukan fungsi sebagai objek yang fleksibel:
* **Map of Functions**: Semua operasi matematika (`Tambah`, `Kurang`, `Kali`, `Bagi`) disimpan di dalam sebuah `map` global. Hal ini memungkinkan pemanggilan fungsi secara dinamis tanpa menggunakan banyak struktur `if-else`.
* **Higher-Order Function**: Fungsi `Proses` menerima fungsi lain sebagai argumen. Ini memisahkan logika eksekusi dengan logika perhitungan matematika.

### 2. Perhitungan Presisi (`float64`)
Menggunakan tipe data `float64` untuk memastikan hasil perhitungan angka desimal tetap akurat, terutama pada operasi pembagian.

### 3. Interactive Looping & Flow Control
Program menggunakan perulangan tak terbatas (`for { ... }`) yang memberikan pengalaman pengguna lebih baik:
* **Dynamic Operator Display**: Menampilkan daftar operator yang tersedia langsung dari iterasi `map`.
* **Break Control**: Memberikan pilihan kepada pengguna untuk mengulang perhitungan atau keluar dari aplikasi dengan validasi input `y/n`.

### 4. Zero Division Safety
Dilengkapi dengan validasi pada fungsi `Bagi` untuk mencegah error sistem jika pengguna memasukkan angka nol sebagai pembagi.

---

## 🚀 Cara Menjalankan

1. Masuk ke direktori project:
   ```bash
   cd 02-kalkulator-func