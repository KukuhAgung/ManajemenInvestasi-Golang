# Aplikasi Manajemen Portofolio Investasi

Aplikasi berbasis terminal (CLI) untuk mengelola portofolio investasi dengan fitur CRUD, kalkulasi otomatis keuntungan/kerugian, pencarian (sequential & binary search), pengurutan (selection sort & insertion sort), serta laporan portofolio dan statistik.

## Daftar Isi

- [Aplikasi Manajemen Portofolio Investasi](#aplikasi-manajemen-portofolio-investasi)
  - [Daftar Isi](#daftar-isi)
  - [Fitur Utama](#fitur-utama)
    - [Fitur A: Manajemen Data (CRUD)](#fitur-a-manajemen-data-crud)
    - [Fitur B: Kalkulasi Keuntungan/Kerugian Otomatis](#fitur-b-kalkulasi-keuntungankerugian-otomatis)
    - [Fitur C: Pencarian Aset (Searching)](#fitur-c-pencarian-aset-searching)
    - [Fitur D: Pengurutan Aset (Sorting)](#fitur-d-pengurutan-aset-sorting)
    - [Fitur E: Laporan Portofolio \& Statistik](#fitur-e-laporan-portofolio--statistik)
  - [Struktur File \& Direktori](#struktur-file--direktori)
  - [Daftar Fungsi / Prosedur](#daftar-fungsi--prosedur)
    - [Fungsi Utama \& Alur Menu](#fungsi-utama--alur-menu)
    - [Fungsi CRUD (Manajemen Data)](#fungsi-crud-manajemen-data)
    - [Fungsi Perhitungan](#fungsi-perhitungan)
    - [Fungsi Algoritma Searching](#fungsi-algoritma-searching)
    - [Fungsi Algoritma Sorting](#fungsi-algoritma-sorting)
    - [Fungsi Representasi Visual](#fungsi-representasi-visual)

---

## Fitur Utama

### Fitur A: Manajemen Data (CRUD)

- **Tambah Data (Create)** Pengguna memasukkan:
  - Nama Aset
  - Jenis Aset
  - Modal Awal

  Sistem **secara otomatis** membuat ID unik (menggunakan auto-increment sederhana atau berbasis panjang array).

- **Ubah Data (Update)** Pengguna dapat memperbarui **Jenis Aset** atau **Harga Terkini** berdasarkan pencarian ID atau Nama tertentu.

- **Hapus Data (Delete)** Menghapus data investasi tertentu dari daftar _slice_.

---

### Fitur B: Kalkulasi Keuntungan/Kerugian Otomatis

- Berjalan **otomatis** setiap kali ada manipulasi data (terutama saat proses Tambah atau Ubah Harga Terkini).

- **Rumus Nominal Keuntungan/Kerugian:** `Keuntungan Rupiah = Harga Terkini - Modal Awal`

- **Rumus Persentase Keuntungan:** `Persentase Keuntungan = (Keuntungan Rupiah / Modal Awal) * 100`

---

### Fitur C: Pencarian Aset (Searching)

- **Sequential Search** Digunakan untuk mencari elemen berdasarkan **Nama Aset**.  
  Sistem memeriksa satu per satu elemen array dari indeks 0 hingga akhir (karena sifat nama aset yang acak/tidak berurutan).

- **Binary Search** Digunakan untuk mencari elemen berdasarkan **Jenis Aset**.  
  Sistem membagi array menjadi dua bagian secara berulang.  
  **Syarat mutlak:** Array harus diurutkan terlebih dahulu berdasarkan Jenis Aset secara alfabetis (A–Z) sebelum algoritma ini dijalankan.

---

### Fitur D: Pengurutan Aset (Sorting)

- **Selection Sort** Digunakan untuk mengurutkan aset berdasarkan **Nilai Investasi** (secara _ascending_ atau _descending_).  
  Sistem mencari nilai ekstrem (terkecil/terbesar) lalu menukarnya ke posisi yang benar.

- **Insertion Sort** Digunakan untuk mengurutkan aset berdasarkan **Persentase Keuntungan** (secara _ascending_ atau _descending_).  
  Sistem menyisipkan setiap elemen ke posisi yang tepat pada bagian array yang sudah terurut.

---

### Fitur E: Laporan Portofolio & Statistik

- **Tampilan Tabel** Menyajikan portofolio dalam bentuk kolom teratur di terminal menggunakan _string formatting_.

- **Ringkasan Statistik** Menampilkan total akumulasi dari seluruh data berupa:
  - Total Modal Keseluruhan
  - Total Nilai Terkini
  - Total Keuntungan/Kerugian

---

## Struktur File & Direktori

Untuk menjaga kode tetap rapi dan mempermudah pembagian kerja kelompok, aplikasi ini menggunakan pendekatan **Single Package, Multiple Files**. Semua file berada di dalam satu folder yang sama dan menggunakan deklarasi `package main`.

```text
📁 investasi-app/
├── 📄 main.go       # Titik masuk aplikasi & alur menu utama
├── 📄 models.go     # Definisi struct Investasi dan variabel slice global
├── 📄 crud.go       # Logika Tambah, Ubah, Hapus aset
├── 📄 kalkulasi.go  # Logika perhitungan profit/loss
├── 📄 searching.go  # Algoritma Sequential Search & Binary Search
├── 📄 sorting.go    # Algoritma Selection Sort & Insertion Sort
└── 📄 ui.go         # Tampilan tabel, statistik, dan cetak menu
```

---

## Daftar Fungsi / Prosedur

Berikut adalah daftar fungsi terpisah (**modular**) yang wajib diimplementasikan di dalam kode program.

### Fungsi Utama & Alur Menu

| Fungsi                 | Deskripsi                                                                                                               |
| ---------------------- | ----------------------------------------------------------------------------------------------------------------------- |
| `func main()`          | Titik masuk utama program. Menampung variabel portofolio dan menangani perulangan menu utama (`for` dan `switch-case`). |
| `func tampilkanMenu()` | Prosedur untuk mencetak teks pilihan menu ke layar terminal.                                                            |

### Fungsi CRUD (Manajemen Data)

| Fungsi                                            | Deskripsi                                                                                                      |
| ------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `func tambahAset(portofolio *[]Investasi)`        | Meminta input pengguna dan memasukkan data baru menggunakan fungsi `append`.                                   |
| `func ubahAset(portofolio *[]Investasi, id int)`  | Mengubah _field_ data pada indeks tertentu dan otomatis memanggil fungsi kalkulasi ulang.                      |
| `func hapusAset(portofolio *[]Investasi, id int)` | Menghapus elemen _slice_ pada indeks ke‑i menggunakan teknik penyambungan `append(slice[:i], slice[i+1:]...)`. |

### Fungsi Perhitungan

| Fungsi                                   | Deskripsi                                                                                               |
| ---------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `func hitungKeuntungan(aset *Investasi)` | Menghitung nominal _profit/loss_ serta persentasenya, lalu memperbarui nilai di dalam _struct_ terkait. |

### Fungsi Algoritma Searching

| Fungsi                                                                  | Deskripsi                                                                                                             |
| ----------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- |
| `func sequentialSearchNama(portofolio []Investasi, keyword string) int` | Mencari kecocokan Nama Aset secara linier. Mengembalikan indeks array jika ditemukan, atau `-1` jika tidak ditemukan. |
| `func binarySearchJenis(portofolio []Investasi, keyword string) []int`  | Mencari Jenis Aset dengan membelah array secara biner setelah array dipastikan terurut.                               |
| `func sortUntukBinarySearch(portofolio []Investasi)`                    | Fungsi pembantu untuk mengurutkan Jenis Aset (A–Z) agar _Binary Search_ dapat bekerja dengan valid.                   |

### Fungsi Algoritma Sorting

| Fungsi                                                                        | Deskripsi                                                                                           |
| ----------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `func selectionSortNilaiInvestasi(portofolio []Investasi, isDescending bool)` | Mengurutkan elemen berdasarkan _field_ Nilai Investasi menggunakan metode **Selection Sort**.       |
| `func insertionSortPersentase(portofolio []Investasi, isDescending bool)`     | Mengurutkan elemen berdasarkan _field_ Persentase Keuntungan menggunakan metode **Insertion Sort**. |

### Fungsi Representasi Visual

| Fungsi                                            | Deskripsi                                                                                                |
| ------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `func tampilkanTabel(portofolio []Investasi)`     | Mencetak seluruh isi portofolio dalam format tabel yang rapi.                                            |
| `func tampilkanStatistik(portofolio []Investasi)` | Menghitung agregat total modal, total nilai saat ini, dan total keuntungan portofolio, lalu mencetaknya. |

---

**Tags:** golang, cli, investment, portfolio, crud, array, sequential-search, binary-search, selection-sort, insertion-sort
