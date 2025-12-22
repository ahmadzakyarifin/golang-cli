# 💻 Golang CLI: Input Handling & Private Struct

Project ini mendemonstrasikan pembangunan aplikasi Command Line Interface (CLI) di Go untuk interaksi terminal, dengan fokus pada pengorganisasian data menggunakan prinsip keamanan enkapsulasi.

---

## 🏗️ Fitur dan Konsep Utama

### 1. Input Handling dengan `bufio`
Aplikasi ini menggunakan `bufio.NewReader(os.Stdin)` karena keunggulannya dibandingkan `fmt.Scan` standar:
* **Mendukung Spasi**: `bufio` membaca seluruh baris hingga tombol Enter ditekan, sehingga sangat ideal untuk menangani input seperti **Nama Lengkap**.
* **Pembersihan Data**: Menggunakan `strings.TrimSpace` untuk memastikan karakter *newline* (`\n`) tidak ikut tersimpan, sehingga data tetap bersih dan valid.

### 2. Enkapsulasi: Private (Unexported) Struct
Dalam Go, enkapsulasi dilakukan melalui aturan kapitalisasi huruf (bukan kata kunci `private`/`public`). Struct dalam project ini menggunakan huruf kecil agar bersifat **Private (Unexported)**:

```go
type mhs struct {
    nama string // Hanya bisa diakses dalam package yang sama
    nim  string // Menjaga data tidak dimodifikasi sembarangan dari luar
}