// Program Inventori Barang
// Fitur: Tambah, Tampilkan, Hapus, Urutkan (stok & tanggal), Cari (nama & kategori)

package main

import "fmt"

// konstanta untuk batas maksimum produk
const NMAX int = 10000

// struktur tanggal
type tanggal struct {
	hari  int
	bulan int
	tahun int
}

// struktur produk
type Produk struct {
	namaBarang string
	stokBarang int
	kategori   string
	tanggal    tanggal
}

// array produk
type barang [NMAX]Produk

// fungsi utama
func main() {
	var n int
	var A barang
	fmt.Println("--------------------------------------------------")
	fmt.Println("          P R O G R A M   I N V E N T O R I       ")
	fmt.Println("--------------------------------------------------")
	fmt.Println("               MASUKKAN DATA BARANG               ")
	fmt.Println("--------------------------------------------------")
	fmt.Println("Jumlah barang: ")
	fmt.Scan(&n)

	bacaData(&A, &n)
	menu(&A, &n)
}

// fungsi untuk menampilkan menu utama
func menu(A *barang, n *int) {
	var pilihan int
	for {
		fmt.Println("--------------------------------------------------")
		fmt.Println("                   M  E  N  U                  ")
		fmt.Println("1. Tampilkan Data Barang")
		fmt.Println("2. Urutkan Barang Menaik Berdasarkan Stok")
		fmt.Println("3. Urutkan Barang Menurun Berdasarkan Stok")
		fmt.Println("4. Urutkan Barang Menaik Berdasarkan Tanggal")
		fmt.Println("5. Urutkan Barang Menurun Berdasarkan Tanggal")
		fmt.Println("6. Hapus Barang")
		fmt.Println("7. Tambah Barang")
		fmt.Println("8. Cari Barang")
		fmt.Println("9. Cari Kategori")
		fmt.Println("10. Exit")
		fmt.Println("--------------------------------------------------")
		fmt.Print("Pilih (1/2/3/4/5/6/7/8/9/10): ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			cetakData(*A, *n)
		case 2:
			// urutkan stok menaik
			urutMenaikStok(A, *n)
			cetakData(*A, *n)
		case 3:
			// urutkan stok menurun
			urutMenurunStok(A, *n)
			cetakData(*A, *n)
		case 4:
			// urutkan tanggal menaik
			urutMenaikTanggal(A, *n)
			cetakData(*A, *n)
		case 5:
			// urutkan tanggal menurun
			urutMenurunTanggal(A, *n)
			cetakData(*A, *n)
		case 6:
			// hapus barang
			hapusBarang(A, n)
		case 7:
			// cari barang berdasarkan nama
			tambahBarang(A, n)
		case 8:
			// cari barang berdasarkan kategori
			cariBarang(*A, *n)
		case 9:
			//menambah barang
			cariKategori(*A, *n)
		case 10:
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// fungsi untuk membaca data dari pengguna
func bacaData(A *barang, n *int) {
	if *n > NMAX {
		*n = NMAX
	}
	for i := 0; i < *n; i++ {
		fmt.Print("Nama barang: ")
		fmt.Scan(&A[i].namaBarang)
		fmt.Print("Kategori barang: ")
		fmt.Scan(&A[i].kategori)
		fmt.Print("Stok barang:")
		fmt.Scan(&A[i].stokBarang)
		fmt.Print("Tanggal (format dd mm yyyy): ")
		fmt.Scan(&A[i].tanggal.hari, &A[i].tanggal.bulan, &A[i].tanggal.tahun)
	}
}

// fungsi untuk mencetak semua data barang
func cetakData(A barang, n int) {
	fmt.Println("Data Barang:")
	for i := 0; i < n; i++ {
		fmt.Printf("%d. Nama: %s, Kategori: %s, Stok: %d, Tanggal: %02d-%02d-%d\n", i+1, A[i].namaBarang, A[i].kategori, A[i].stokBarang, A[i].tanggal.hari, A[i].tanggal.bulan, A[i].tanggal.tahun)
	}
}

// fungsi untuk mengurutkan stok secara menaik
func urutMenaikStok(A *barang, n int) {
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if A[i].stokBarang > A[j].stokBarang {
				A[i], A[j] = A[j], A[i]
			}
		}
	}
}

// fungsi untuk mengurutkan stok secara menurun
func urutMenurunStok(A *barang, n int) {
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if A[i].stokBarang < A[j].stokBarang {
				A[i], A[j] = A[j], A[i]
			}
		}
	}
}

// fungsi pembanding tanggal lebih lama
func lebihLama(t1, t2 tanggal) bool {
	if t1.tahun != t2.tahun {
		return t1.tahun < t2.tahun
	}
	if t1.bulan != t2.bulan {
		return t1.bulan < t2.bulan
	}
	return t1.hari < t2.hari
}

// fungsi pembanding tanggal lebih baru
func lebihBaru(t1, t2 tanggal) bool {
	if t1.tahun != t2.tahun {
		return t1.tahun > t2.tahun
	}
	if t1.bulan != t2.bulan {
		return t1.bulan > t2.bulan
	}
	return t1.hari > t2.hari
}

// fungsi untuk mengurutkan tanggal secara menaik
func urutMenaikTanggal(A *barang, n int) {
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if lebihLama(A[j].tanggal, A[i].tanggal) {
				A[i], A[j] = A[j], A[i]
			}
		}
	}
}

// fungsi untuk mengurutkan tanggal secara menurun
func urutMenurunTanggal(A *barang, n int) {
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if lebihBaru(A[j].tanggal, A[i].tanggal) {
				A[i], A[j] = A[j], A[i]
			}
		}
	}
}

// fungsi untuk menghapus barang berdasarkan nama
func hapusBarang(A *barang, n *int) {
	var nama string
	var i, idx int
	ditemukan := false

	fmt.Print("Masukkan nama barang yang ingin dihapus: ")
	fmt.Scan(&nama)

	for i = 0; i < *n; i++ {
		if A[i].namaBarang == nama {
			idx = i
			ditemukan = true
			break
		}
	}

	if ditemukan {
		for i = idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n--
		fmt.Println("Barang berhasil dihapus.")
	} else {
		fmt.Println("Barang tidak ditemukan.")
	}
}

// fungsi untuk mencari barang berdasarkan nama
func cariBarang(A barang, n int) {
	var keyword string
	fmt.Print("Masukkan nama barang yang dicari: ")
	fmt.Scan(&keyword)

	ditemukan := false
	fmt.Println("\nHasil pencarian:")
	for i := 0; i < n; i++ {
		if A[i].namaBarang == keyword {
			fmt.Printf("\nNama: %s, Kategori: %s, Stok: %d, Tanggal: %02d-%02d-%d\n", A[i].namaBarang, A[i].kategori, A[i].stokBarang, A[i].tanggal.hari, A[i].tanggal.bulan, A[i].tanggal.tahun)
			ditemukan = true
			break
		}
	}
	if !ditemukan {
		fmt.Println("Barang tidak ditemukan.")
	}
}

func tambahBarang(A *barang, n *int) {
	var nama, kategori string
	var stok int
	var t tanggal

	if *n >= NMAX {
		fmt.Println("Tidak dapat menambah barang, kapasitas maksimum tercapai.")
		return
	}

	fmt.Print("Masukkan nama barang: ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan kategori barang: ")
	fmt.Scan(&kategori)
	fmt.Print("Masukkan stok barang: ")
	fmt.Scan(&stok)
	fmt.Print("Masukkan tanggal (format dd mm yyyy): ")
	fmt.Scan(&t.hari, &t.bulan, &t.tahun)

	A[*n] = Produk{namaBarang: nama, stokBarang: stok, kategori: kategori, tanggal: t}
	*n++
	fmt.Println("Barang berhasil ditambahkan.")
}

// fungsi untuk mencari barang berdasarkan kategori
func cariKategori(A barang, n int) {
	var keyword string
	fmt.Print("Masukkan nama kategori yang dicari: ")
	fmt.Scan(&keyword)

	jumlah := 0
	fmt.Println("\nHasil pencarian:")
	for i := 0; i < n; i++ {
		if A[i].kategori == keyword {
			fmt.Printf("\nNama: %s, Kategori: %s, Stok: %d, Tanggal: %02d-%02d-%d\n", A[i].namaBarang, A[i].kategori, A[i].stokBarang, A[i].tanggal.hari, A[i].tanggal.bulan, A[i].tanggal.tahun)
			jumlah++
		}
	}
	if jumlah == 0 {
		fmt.Println("Barang tidak ditemukan.")
	}
}
