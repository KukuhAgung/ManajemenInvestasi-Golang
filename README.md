# ManajemenInvestasi-Golang
A CLI-based Investment Management App built with pure Go. Features in-memory CRUD via slices, along with Sequential/Binary Search and Selection/Insertion Sort algorithms. No database required.

## 🚀 Fitur Utama & Metode yang Digunakan
### 1. Manajemen Data Investasi (CRUD)
Sistem dapat menambah, mengubah, dan menghapus aset investasi.
* **Metode / Logika:** Penyimpanan data dilakukan secara *in-memory* menggunakan `slice` bertipe `struct` di Golang. 
  * **Create:** Menggunakan fungsi `append()` untuk menambah data ke akhir *slice*.
  * **Update:** Memperbarui nilai field pada *struct* berdasarkan indeks array dengan meneruskan *pointer* (*pass by reference*).
  * **Delete:** Menghapus elemen pada indeks ke-`i` menggunakan teknik manipulasi *slice* Go: `append(slice[:i], slice[i+1:]...)`.

### 2. Kalkulasi Profitabilitas Otomatis
Setiap kali data dimasukkan atau harga terkini diperbarui, sistem otomatis menghitung keuntungan/kerugian.
* **Metode / Logika:**
  Kalkulasi aritmatika sederhana di dalam fungsi. 
  * `Keuntungan Rupiah = Harga Terkini - Modal Awal`
  * `Persentase Keuntungan = (Keuntungan Rupiah / Modal Awal) * 100`

### 3. Pencarian Aset Berdasarkan Nama
Pengguna dapat mencari data investasi spesifik berdasarkan nama asetnya (misalnya: "BBCA", "GOTO").
* **Metode:** **Sequential Search (Pencarian Sekuensial)**
  * **Cara Kerja:** Algoritma akan menelusuri array dari indeks awal (0) hingga akhir secara satu per satu (*linear*). Jika nama aset cocok dengan kata kunci, sistem akan mengembalikan indeks tersebut.
  * **Alasan Penggunaan:** Nama aset bersifat dinamis, tidak selalu berurutan, dan sangat spesifik.

### 4. Pencarian Aset Berdasarkan Jenis
Pengguna dapat mencari sekumpulan aset berdasarkan jenisnya (misalnya: "Saham" atau "Obligasi").
* **Metode:** **Binary Search (Pencarian Biner)**
  * **Cara Kerja:** Algoritma akan membagi array menjadi dua bagian dan membandingkan kata kunci dengan nilai tengah. Jika tidak cocok, pencarian dilanjutkan pada separuh bagian yang relevan.
  * **Syarat Utama:** Sebelum *Binary Search* dieksekusi, sistem di latar belakang akan menjalankan algoritma *sorting* (pengurutan alfabetis A-Z pada field Jenis Aset) terlebih dahulu agar kondisi array *sorted* terpenuhi.

### 5. Pengurutan Portofolio (Nilai Investasi)
Pengguna dapat melihat daftar aset yang diurutkan berdasarkan Nilai Investasi (dari yang terbesar atau terkecil).
* **Metode:** **Selection Sort**
  * **Cara Kerja:** Algoritma mencari elemen dengan nilai investasi terkecil (atau terbesar) dari daftar yang belum terurut, lalu menukarnya (*swap*) dengan elemen pertama dari daftar tersebut. Proses ini diulang hingga seluruh array terurut.

### 6. Pengurutan Portofolio (Persentase Keuntungan)
Pengguna dapat melihat daftar aset yang diurutkan berdasarkan performa (persentase keuntungan).
* **Metode:** **Insertion Sort**
  * **Cara Kerja:** Algoritma mengambil elemen satu per satu dari array yang belum terurut, lalu menyisipkannya (*insert*) ke posisi yang tepat di dalam bagian array yang sudah terurut sebelumnya. Cocok untuk data yang sebagian besar mungkin sudah berurutan.

### 7. Laporan Statistik & Tampilan Tabel
Menyajikan data mentah ke dalam format visual yang mudah dibaca di terminal.
* **Metode / Logika:**
  * Menggunakan *looping* (perulangan) untuk mengumpulkan total modal, total nilai saat ini, dan total profit/loss dari seluruh elemen array.
  * Menggunakan pemformatan string (`fmt.Printf` dengan verb seperti `%-15s` atau `%10.2f`) untuk menghasilkan kolom tabel yang presisi dan rapi.

---

## 🛠️ Tech Stack

* **Bahasa Pemrograman:** Go (Golang)
* **Penyimpanan Data:** In-Memory (Slices & Structs)
* **Antarmuka:** CLI (Terminal)
