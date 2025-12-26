# 🚀 06-todo-kompleks: Ultimate Pointer-to-Slice Architecture

Proyek ini merupakan **High-Level Upgrade** dari versi `06-todo-upgrade`. Di sini, aplikasi telah bertransformasi total dari struktur modular sederhana menjadi arsitektur berbasis **Pointer Management** dan **Function Mapping** yang umum digunakan dalam skala *enterprise*.

---

## 📈 Evolusi: Apa yang Meningkat dari `06-todo-upgrade`?

| Fitur | 06-todo-upgrade | 06-todo-kompleks (Final) |
| :--- | :--- | :--- |
| **Manajemen Memori** | Value-based (Data disalin) | **Pointer-to-Slice of Pointers** (Referensi memori) |
| **Struktur Folder** | 1 Sub-package (`pelengkap`) | **Full Clean Architecture** (`fitur`, `helper`, `model`) |
| **Pola Kontrol** | `switch-case` (Statis) | **Function Mapping via Map** (Dinamis/Scalable) |
| **Pass-by-Value** | Data dikirim balik lewat *return* | **Pass-by-Reference** (Manipulasi data langsung) |

---

## 💾 Spesifikasi Penyimpanan: In-Memory Pointer Engine

Sistem penyimpanan data di versi ini dirancang untuk efisiensi RAM maksimal:

1. **Double Pointer Mechanism**: Menggunakan struktur data `*[]*model.Daftar`. 
   - **Slice of Pointers**: List tidak menyimpan data mentah, melainkan alamat memori objek. Ini mempercepat proses manipulasi data karena sistem hanya memindahkan alamat (8 byte) bukan seluruh data struct.
   - **Pointer to Slice**: Memungkinkan fungsi di package `fitur` untuk menambah atau menghapus elemen list asli di `main.go` tanpa perlu melakukan *return value*.
2. **In-Memory Storage**: Karena data disimpan di RAM, operasi CRUD memiliki latensi nol dengan kompleksitas akses index $O(1)$.
3. **Smart Re-slicing**: Penghapusan data menggunakan logika `append((*daftar)[:idx], (*daftar)[idx+1:]...)` yang memastikan kapasitas slice dikelola secara optimal oleh *Garbage Collector* Go.


## 📝 Catatan Teknis: Kenapa Pointer-to-Slice?

Sedikit catatan teknis mengenai alasan saya menggunakan struktur data `*[]*model.Daftar` (Pointer ke Slice berisi Pointer ke Struct) dalam proyek ini. Keputusan ini diambil untuk efisiensi memori dan memastikan data yang diubah benar-benar tersimpan.

### 1. Kenapa Pointer ke Struct (`*model.Daftar`)?
Secara default, struct di Go bersifat *Pass-by-Value*. Artinya, jika saya mengirim data barang ke sebuah fungsi tanpa pointer, Go akan membuat **salinan (copy)** dari data tersebut.
- **Masalah:** Jika tidak pakai pointer, memori jadi boros karena terjadi duplikasi data. Selain itu, kalau saya mengedit data di dalam fungsi, yang berubah hanya salinannya, bukan data aslinya.
- **Solusi:** Saya pakai pointer agar fungsi langsung mengarah ke alamat memori data asli. Jadi, tidak ada salinan yang tidak perlu.

### 2. Kenapa Pointer ke Slice (`*[]...`) di Fitur Tambah & Hapus?
Khusus untuk fitur `Tambah` dan `Hapus`, saya menggunakan **Pointer ke Slice**.
- **Alasan:** Operasi seperti `append` (tambah) atau slicing ulang (hapus) akan mengubah struktur header slice, yaitu panjang (`len`) dan kapasitas (`cap`).
- Jika saya hanya melempar slice biasa, perubahan `len` dan `cap` hanya terjadi di dalam fungsi tersebut (lokal) dan tidak akan terbaca di `main`.
- Dengan pointer ke slice, saya memastikan bahwa perubahan jumlah data (bertambah/berkurang) ikut berubah di variabel utama yang ada di `main`.

### 3. Kenapa Cukup Slice Biasa di Fitur Edit & Tampil?
Untuk `Edit` dan `Tampil`, saya tidak perlu pointer ke slice, cukup slice yang berisi pointer struct-nya saja.
- **Alasan:** Di fitur ini, saya **tidak mengotak-atik slice** (tidak mengubah `len` atau `cap`). Saya hanya memanipulasi isinya saja (mengganti nama, harga, atau stok).
- Karena elemen di dalamnya sudah berupa pointer struct, perubahan data tetap tersimpan tanpa perlu membawa pointer slice-nya.
---

## 🏗️ Struktur Proyek (Modular Layering)

Aplikasi dipisahkan menjadi beberapa layer untuk mematuhi prinsip **Separation of Concerns (SoC)**:

```text
.
├── main.go             # Entry Point: Menu Mapping & Loop Utama
├── go.mod              # Module: [github.com/ahmadzakyarifin/06-todo-kompleks](https://github.com/ahmadzakyarifin/06-todo-kompleks)
├── model/              # LAYER 1: Blueprint Data (Data Entity)
│   └── daftar.go       # Struct tunggal sebagai Single Source of Truth
├── helper/             # LAYER 2: Global Validators (Error Handling)
│   └── validasi.go     # Validasi String & Int dengan return 'error' idiomatik
└── fitur/              # LAYER 3: Business Logic (CRUD Operations)
    └── operasional.go  # Logika utama Tambah, Edit, Hapus, & Tampil