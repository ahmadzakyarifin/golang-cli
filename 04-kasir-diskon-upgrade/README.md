# 🛍️ Kasir Alat Tulis Modular - Upgrade Version (Project 04)

Project ini adalah versi **upgrade** dari simulasi kasir sebelumnya. Fokus utama di versi ini adalah penerapan **Struktur Data (Struct)**, **Modularitas Fungsi**, dan penggunaan **Go Modules** agar kode lebih profesional, aman, dan mudah dikelola.

---

## 🏗️ Apa yang Baru di Versi Upgrade Ini?

### 1. Penggunaan Go Modules (`go mod init`)
Saya menyertakan sistem **Go Modules** dalam project ini. Ini sangat penting karena:
* **Manajemen Library**: Menjadi wadah wajib untuk mengelola library eksternal (`golang.org/x/text`).
* **Identitas Project**: Memberikan identitas resmi pada folder project agar semua file di dalamnya saling terhubung dengan rapi.

### 2. Data Tidak Mudah Tertukar (Struct)
Data barang tidak lagi disimpan dalam daftar terpisah (*parallel slices*). Sekarang, Nama, Harga, dan Jumlah dibungkus dalam satu **Struct**:
* **Data Integrity**: Memastikan satu paket data (misal: "Pulpen" dan harganya) terkunci rapat dalam satu unit memori, sehingga tidak akan ada risiko data tertukar indeksnya.

### 3. Struktur Kode Modular
Kode tidak lagi menumpuk di fungsi `main`, melainkan dibagi menjadi fungsi-fungsi kecil yang spesifik:
* `validasiPilih` & `validasiJumlah`: Menangani keamanan input.
* `hitungDiskon`: Menangani logika matematika.
* `cetakPembayaran`: Menangani tampilan struk.
Ini membuat kode jauh lebih bersih (*Clean Code*) dan mudah untuk dikembangkan.

### 4. Validasi Input yang Kuat
Program kini lebih tahan terhadap *human error*. Menggunakan perulangan (**Nested Loop**), program akan terus meminta input sampai user memasukkan data yang valid, terutama saat pemilihan barang dan konfirmasi "Tambah lagi (y/n)".

---

## 🚀 Cara Menjalankan

1. **Inisialisasi & Ambil Library** (Jalankan di terminal folder ini):
   ```bash
   go mod init 04-kasir-diskon-upgrade
   go mod tidy