**# Aplikasi Manajemen Portofolio Investasi

Aplikasi berbasis terminal (CLI) untuk mengelola portofolio investasi dengan fitur CRUD, kalkulasi otomatis keuntungan/kerugian, pencarian (sequential & binary search), pengurutan (selection sort & insertion sort), serta laporan portofolio dan statistik.

## Daftar Isi

- [Daftar Isi](#daftar-isi)
- [Fitur Utama](#fitur-utama)
  - [Fitur A: Manajemen Data (CRUD)](#fitur-a-manajemen-data-crud)
  - [Fitur B: Kalkulasi Keuntungan/Kerugian Otomatis](#fitur-b-kalkulasi-keuntungankerugian-otomatis)
  - [Fitur C: Pencarian Aset (Searching)](#fitur-c-pencarian-aset-searching)
  - [Fitur D: Pengurutan Aset (Sorting)](#fitur-d-pengurutan-aset-sorting)
  - [Fitur E: Laporan Portofolio \& Statistik](#fitur-e-laporan-portofolio--statistik)
- [Struktur File \& Direktori](#struktur-file--direktori)
- [└── 📄 ui.go         # Tampilan tabel, statistik, dan cetak menu](#--uigo----------tampilan-tabel-statistik-dan-cetak-menu)
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

  Sistem memasukkan data baru ke dalam indeks array yang masih kosong. Indeks ini dilacak menggunakan variabel pelacak (contoh: `jumlahData`). Setelah data masuk, nilai `jumlahData` ditambah 1 secara manual. Sistem juga otomatis membuat ID unik berdasarkan nilai pelacak tersebut.

- **Ubah Data (Update)** Pengguna dapat memperbarui **Jenis Aset** atau **Harga Terkini** berdasarkan pencarian ID atau Nama tertentu melalui iterasi dari indeks awal hingga batas data yang aktif.

- **Hapus Data (Delete)** Menghapus data investasi tertentu dari daftar **array statis**. Karena tidak boleh menggunakan `append` atau memotong array, penghapusan dilakukan dengan cara **menggeser data (shifting)** secara manual. Elemen-elemen di sebelah kanan indeks yang dihapus akan digeser satu per satu ke kiri menggunakan perulangan, kemudian variabel `jumlahData` dikurangi 1.

---

### Fitur B: Kalkulasi Keuntungan/Kerugian Otomatis

- Berjalan **otomatis** setiap kali ada manipulasi data (terutama saat proses Tambah atau Ubah Harga Terkini).

- **Rumus Nominal Keuntungan/Kerugian:** `Keuntungan Rupiah = Harga Terkini - Modal Awal`

- **Rumus Persentase Keuntungan:** `Persentase Keuntungan = (Keuntungan Rupiah / Modal Awal) * 100`

---

### Fitur C: Pencarian Aset (Searching)

- **Sequential Search** Digunakan untuk mencari elemen berdasarkan **Nama Aset**.  
  Sistem memeriksa satu per satu elemen array dari indeks 0 hingga batas `jumlahData`.

- **Binary Search** Digunakan untuk mencari elemen berdasarkan **Jenis Aset**.  
  Sistem membagi array aktif menjadi dua bagian secara berulang.  
  **Syarat mutlak:** Array harus diurutkan terlebih dahulu berdasarkan Jenis Aset secara alfabetis (A–Z) sebelum algoritma ini dijalankan.

---

### Fitur D: Pengurutan Aset (Sorting)

- **Selection Sort** Digunakan untuk mengurutkan aset berdasarkan **Nilai Investasi** (secara _ascending_ atau _descending_).  
  Sistem mencari nilai ekstrem (terkecil/terbesar) dari rentang indeks 0 hingga `jumlahData` lalu menukarnya ke posisi yang benar.

- **Insertion Sort** Digunakan untuk mengurutkan aset berdasarkan **Persentase Keuntungan** (secara _ascending_ atau _descending_).  
  Sistem menyisipkan setiap elemen ke posisi yang tepat pada bagian array yang sudah terurut.

---

### Fitur E: Laporan Portofolio & Statistik

- **Tampilan Tabel** Menyajikan portofolio dalam bentuk kolom teratur di terminal menggunakan _string formatting_. Hanya mencetak data dari indeks 0 hingga batas `jumlahData`.

- **Ringkasan Statistik** Menampilkan total akumulasi dari data yang aktif berupa:
  - Total Modal Keseluruhan
  - Total Nilai Terkini
  - Total Keuntungan/Kerugian

---

## Struktur File & Direktori

Untuk menjaga kode tetap rapi dan mempermudah pembagian kerja kelompok, aplikasi ini **menggunakan** pendekatan **Single Package, Multiple Files**. Semua file berada di dalam satu folder yang sama dan menggunakan deklarasi `package main`.

```text
📁 investasi-app/
├── 📄 main.go       # Titik masuk aplikasi & alur menu utama
├── 📄 models.go     # Definisi struct Investasi, array statis global (misal [100]Investasi), dan variabel pelacak (jumlahData)
├── 📄 crud.go       # Logika Tambah, Ubah, Hapus aset secara manual
├── 📄 kalkulasi.go  # Logika perhitungan profit/loss
├── 📄 searching.go  # Algoritma Sequential Search & Binary Search
├── 📄 sorting.go    # Algoritma Selection Sort & Insertion Sort
└── 📄 ui.go         # Tampilan tabel, statistik, dan cetak menu
---

## Daftar Fungsi / Prosedur

Berikut adalah daftar fungsi terpisah (**modular**) yang wajib diimplementasikan di dalam kode program. *(Catatan: Penulisan `[MAX]` di bawah ini merepresentasikan batasan ukuran array statis).*

### Fungsi Utama & Alur Menu

| Fungsi                 | Deskripsi                                                                                                               |
| ---------------------- | ----------------------------------------------------------------------------------------------------------------------- |
| `func main()`          | Titik masuk utama program. Menampung variabel portofolio dan menangani perulangan menu utama (`for` dan `switch-case`). |
| `func tampilkanMenu()` | Prosedur untuk mencetak teks pilihan menu ke layar terminal.                                                            |

### Fungsi CRUD (Manajemen Data)

| Fungsi                                                                | Deskripsi                                                                                                                          |
| --------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `func tambahAset(portofolio *[MAX]Investasi, jumlahData *int)`        | Meminta input pengguna dan memasukkan data baru ke indeks yang kosong, lalu menambah nilai `jumlahData`.                           |
| `func ubahAset(portofolio *[MAX]Investasi, jumlahData int, id int)`   | Mengubah _field_ data pada indeks tertentu berdasarkan ID dan otomatis memanggil fungsi kalkulasi ulang.                             |
| `func hapusAset(portofolio *[MAX]Investasi, jumlahData *int, id int)` | Menghapus elemen array pada indeks tertentu dengan menggeser data ke kiri secara manual, lalu mengurangi nilai `jumlahData`.       |

### Fungsi Perhitungan

| Fungsi                                   | Deskripsi                                                                                               |
| ---------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `func hitungKeuntungan(aset *Investasi)` | Menghitung nominal _profit/loss_ serta persentasenya, lalu memperbarui nilai di dalam _struct_ terkait. |

### Fungsi Algoritma Searching

| Fungsi                                                                                   | Deskripsi                                                                                                             |
| ---------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- |
| `func sequentialSearchNama(portofolio [MAX]Investasi, jumlahData int, keyword string) int` | Mencari kecocokan Nama Aset secara linier. Mengembalikan indeks array jika ditemukan, atau `-1` jika tidak ditemukan. |
| `func binarySearchJenis(portofolio [MAX]Investasi, jumlahData int, keyword string) int`    | Mencari Jenis Aset dengan membelah array secara biner setelah array dipastikan terurut.                               |
| `func sortUntukBinarySearch(portofolio *[MAX]Investasi, jumlahData int)`                   | Fungsi pembantu untuk mengurutkan Jenis Aset (A–Z) secara manual agar _Binary Search_ dapat bekerja dengan valid.     |

### Fungsi Algoritma Sorting

| Fungsi                                                                                            | Deskripsi                                                                                           |
| ------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `func selectionSortNilaiInvestasi(portofolio *[MAX]Investasi, jumlahData int, isDescending bool)` | Mengurutkan elemen berdasarkan _field_ Nilai Investasi menggunakan metode **Selection Sort**.       |
| `func insertionSortPersentase(portofolio *[MAX]Investasi, jumlahData int, isDescending bool)`     | Mengurutkan elemen berdasarkan _field_ Persentase Keuntungan menggunakan metode **Insertion Sort**. |

### Fungsi Representasi Visual

| Fungsi                                                               | Deskripsi                                                                                                |
| -------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `func tampilkanTabel(portofolio [MAX]Investasi, jumlahData int)`     | Mencetak isi portofolio yang aktif (dari indeks 0 hingga batas `jumlahData`) dalam format tabel.         |
| `func tampilkanStatistik(portofolio [MAX]Investasi, jumlahData int)` | Menghitung agregat total modal, total nilai saat ini, dan total keuntungan portofolio lalu mencetaknya.  |

---

**Tags:** golang, cli, investment, portfolio, crud, array, sequential-search, binary-search, selection-sort, insertion-sort**