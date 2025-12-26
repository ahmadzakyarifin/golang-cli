# 📦 06-todo-upgrade: Complete Modular Inventory System

Project ini merupakan versi penyempurnaan menyeluruh dari sistem manajemen inventaris sebelumnya. Fokus utama pada versi ini adalah implementasi **Arsitektur Modular** dan siklus **CRUD (Create, Read, Update, Delete)** yang fungsional, aman, dan user-friendly.

---

## 🌟 Fitur Unggulan (Final Upgrade)

### 1. Full CRUD Cycle Management
Aplikasi telah mendukung siklus pengelolaan data lengkap di dalam memori:
* **Create**: Menambah barang baru dengan validasi tipe data yang ketat.
* **Read**: Menampilkan daftar barang dalam format tabel yang rapi dan sejajar.
* **Update**: Fitur **Edit** yang memungkinkan perubahan atribut spesifik (Nama, Harga, atau Stok) pada item tertentu.
* **Delete**: Menghapus data secara dinamis menggunakan teknik *re-slicing* yang efisien.

### 2. Arsitektur Modular (Custom Package)
Logika program dipisahkan secara cerdas untuk meningkatkan skalabilitas dan keterbacaan kode:
* **`main.go`**: Sebagai orkestrator yang menangani alur navigasi menu dan pengelolaan state database.
* **`pelengkap/`**: Package khusus yang menangani validasi input. Pemisahan ini memungkinkan fungsi validasi bersifat *reusable* (dapat digunakan kembali).

### 3. Localization & Smart Formatting (IDR)
Mengintegrasikan library `golang.org/x/text` untuk memberikan pengalaman pengguna yang lebih profesional:
* **Auto-Formatting**: Mengubah input angka mentah menjadi format mata uang Rupiah dengan pemisah ribuan titik (contoh: `150000` -> `150.000`).

### 4. Robust Input & Buffer Handling
Mengatasi kendala umum pada aplikasi CLI Go saat menangani input:
* **Space-Safe Input**: Menggunakan `bufio.NewReader` sehingga input nama yang mengandung spasi tetap terbaca dengan benar.
* **Buffer Cleaning**: Menangani sisa karakter *newline* (`\n`) agar perpindahan antar input menu tidak terlewati (*skipping input*).

---

## 🏗️ Struktur Proyek
```text
.
├── main.go            # Central Logic: CRUD Operations & Menu Flow
└── pelengkap/         
    └── validasi.go    # Data Integrity: Validation Logic (Nama, Harga, Stok)