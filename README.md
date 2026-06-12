# ManajemenInvestasi-Golang

A CLI-based Investment Management App built with pure Go. Features in-memory CRUD via static arrays, along with Sequential/Binary Search and Selection/Insertion Sort algorithms. No database and no dynamic built-in functions (like `append` or slices) are used, adhering strictly to traditional algorithmic concepts.

## 🚀 Fitur Utama & Metode yang Digunakan

### 1. Manajemen Data Investasi (CRUD)

Sistem dapat menambah, mengubah, dan menghapus aset investasi.

- **Metode / Logika:** Penyimpanan data dilakukan secara _in-memory_ menggunakan Array Statis (`[MAX]Investasi`) dan variabel pelacak (`jumlahData`) di Golang.
  - **Create:** Menambahkan data baru secara langsung ke indeks kosong `array[jumlahData]`, lalu melakukan inkrementasi (menambah) nilai pelacak secara manual (`jumlahData++`).
  - **Update:** Memperbarui nilai field pada _struct_ berdasarkan iterasi ID dengan meneruskan _pointer_ (_pass by reference_), serta memanggil fungsi kalkulasi ulang secara otomatis.
  - **Delete:** Menghapus data menggunakan teknik pergeseran manual (_shifting left_). Elemen di sebelah kanan indeks yang dihapus digeser ke kiri menggunakan perulangan, kemudian nilai `jumlahData--` dikurangi agar sisa data (junk data) diabaikan.

### 2. Kalkulasi Profitabilitas Otomatis

Setiap kali data dimasukkan atau harga terkini diperbarui, sistem otomatis menghitung keuntungan/kerugian.

- **Metode / Logika:**
  Kalkulasi aritmatika sederhana di dalam fungsi terpisah yang dipanggil secara otomatis.
  - `Keuntungan Rupiah = Harga Terkini - Modal Awal`
  - `Persentase Keuntungan = (Keuntungan Rupiah / Modal Awal) * 100`

### 3. Pencarian Aset Berdasarkan Nama

Pengguna dapat mencari data investasi spesifik berdasarkan nama asetnya (misalnya: "BBCA", "GOTO").

- **Metode:** **Sequential Search (Pencarian Sekuensial)**
  - **Cara Kerja:** Algoritma menelusuri array dari indeks awal (0) hingga batas akhir data aktif (`jumlahData`) secara satu per satu (_linear_). Jika nama aset cocok dengan kata kunci, sistem akan mengembalikan indeks tersebut.
  - **Alasan Penggunaan:** Nama aset bersifat dinamis, tidak selalu berurutan, dan sangat spesifik.

### 4. Pencarian Aset Berdasarkan Jenis

Pengguna dapat mencari sekumpulan aset berdasarkan jenisnya (misalnya: "Saham" atau "Obligasi").

- **Metode:** **Binary Search (Pencarian Biner)**
  - **Cara Kerja:** Algoritma membagi rentang pencarian menjadi dua bagian dan membandingkan kata kunci dengan nilai tengah. Jika tidak cocok, pencarian dilanjutkan pada separuh bagian yang relevan.
  - **Syarat Utama:** Sebelum _Binary Search_ dieksekusi, sistem di latar belakang akan menjalankan algoritma pembantu untuk mengurutkan array secara alfabetis (A-Z) berdasarkan field Jenis Aset.

### 5. Pengurutan Portofolio (Nilai Investasi)

Pengguna dapat melihat daftar aset yang diurutkan berdasarkan Nilai Investasi (dari yang terbesar atau terkecil).

- **Metode:** **Selection Sort**
  - **Cara Kerja:** Algoritma mencari elemen dengan nilai investasi terkecil (atau terbesar) dari daftar yang belum terurut, lalu menukarnya (_swap_) secara manual dengan elemen pertama dari iterasi tersebut.

### 6. Pengurutan Portofolio (Persentase Keuntungan)

Pengguna dapat melihat daftar aset yang diurutkan berdasarkan performa (persentase keuntungan).

- **Metode:** **Insertion Sort**
  - **Cara Kerja:** Algoritma mengambil elemen satu per satu, lalu menyisipkannya (_insert_) ke posisi yang tepat dengan cara membandingkan dan menggeser data ke kanan (_shifting right_) di dalam bagian array yang sudah terurut.

### 7. Laporan Statistik & Tampilan Tabel

Menyajikan data mentah ke dalam format visual yang mudah dibaca di terminal.

- **Metode / Logika:**
  - Menggunakan _looping_ (perulangan) khusus dari 0 hingga `jumlahData` untuk merekapitulasi total modal, total nilai saat ini, dan total profit/loss agar sisa kapasitas array kosong tidak terbaca.
  - Menggunakan pemformatan string (`fmt.Printf` dengan format verb seperti `%-15s` atau `%+.2f%%`) untuk menghasilkan kolom tabel yang presisi dan rapi.

---

## 🛠️ Tech Stack

- **Bahasa Pemrograman:** Go (Golang)
- **Penyimpanan Data:** In-Memory (Static Arrays `[NMAX]Struct` & Pointers)
- **Antarmuka:** CLI (Terminal)
