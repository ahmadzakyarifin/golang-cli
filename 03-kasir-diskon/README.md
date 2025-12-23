# 🛒 Simulasi Pesanan Barang & Diskon 

Project ini adalah simulasi sederhana untuk menghitung pesanan belanjaan. Di sini saya belajar cara menangani input banyak barang, menghitung diskon otomatis, dan merapikan tampilan harga.

---

## 🏗️ Apa Saja yang Saya Pelajari di Sini?

### 1. Menggunakan Library Luar (External Library)
Di project ini saya menggunakan bantuan kode tambahan bernama `golang.org/x/text`. 
* **Kenapa pakai ini?** Supaya angka harga (seperti `100000`) otomatis tercetak rapi menjadi `100.000` mengikuti format uang Indonesia. Jadi saya tidak perlu repot membuat fungsi tambah titik manual.

### 2. Penggunaan Go Mod (`go.mod`)
Saya menyertakan file `go.mod` dalam project ini karena dua alasan utama:
* **Catatan Library**: Karena ada kode dari luar yang digunakan, Go butuh file ini sebagai daftar belanja library apa saja yang perlu di-download.
* **Identitas Project**: `go.mod` membantu Go mengenali bahwa folder ini adalah satu kesatuan project. Ini sangat berguna kalau nanti project-nya makin besar dan dibagi-bagi ke banyak folder agar tetap saling terhubung.

### 3. Logika Diskon Sederhana
Aplikasi ini sudah bisa menghitung potongan harga secara otomatis:
* **Diskon 10%**: Jika total belanja mencapai Rp 150.000 atau lebih.
* **Diskon 5%**: Jika total belanja minimal Rp 50.000.
* **Normal**: Jika belanja di bawah Rp 50.000 (tidak dapat diskon).

### 4. Penyimpanan Data Sementara
Nama barang, harga, dan jumlah pesanan disimpan dalam daftar (**Slices**) sebelum akhirnya diringkas menjadi struk pembayaran di akhir aplikasi.

---

## 🚀 Cara Menjalankan

1. **Siapkan Library-nya** (Ketik ini di terminal folder project):
   ```bash
   go mod init 04-kasir-diskon
   go mod tidy