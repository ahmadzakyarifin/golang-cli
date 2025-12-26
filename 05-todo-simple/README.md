# 📝 Todo List - Simple CRUD Version

Project ini adalah implementasi sistem manajemen data (CRUD) berbasis terminal yang berfokus pada efisiensi manipulasi data di dalam memori menggunakan bahasa Go.

---

## 🌟 Keunggulan Teknis

### 1. Efisiensi Memori dengan Pointer (`&`)
Pada fitur **Edit Menu**, project ini menggunakan **Pointer Reference** untuk mengakses data langsung pada alamat memorinya:
```go
data := &database[nomer-1]