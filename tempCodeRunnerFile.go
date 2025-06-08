package main

import (
	"github.com/xuri/excelize/v2"
	"fmt"
)

type Golongan struct {
	Kode string
	Nama string
}

type TanggalKadaluarsa struct {
	Hari  int
	Bulan int
	Tahun int
}

type Obat struct {
	Nama         string
	Jenis        string
	Kategori     string
	Harga        int
	Stok         int
	GolonganObat Golongan
	Kadaluarsa   TanggalKadaluarsa
}

const maxObat = 100

var DaftarObat [maxObat]Obat
var JumlahObat int

var (
	GolonganBebas    = Golongan{"B", "Bebas"}
	GolonganTerbatas = Golongan{"BT", "Bebas Terbatas"}
	GolonganKeras    = Golongan{"K", "Keras"}
)

// Method untuk menambahkan obat
func tambahObat(obats *[maxObat]Obat, jumlah *int) {
	if *jumlah >= maxObat {
		fmt.Println("Data obat penuh. Tidak dapat menambah obat baru!")
		return
	}
	var nama, jenis, kodeGol, kategori string
	var harga, stok int
	var tgl TanggalKadaluarsa
	fmt.Print("Masukkan nama obat: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan jenis obat: ")
	fmt.Scanln(&jenis)
	fmt.Print("Masukkan kategori obat: ")
	fmt.Scanln(&kategori)
	fmt.Print("Masukkan harga obat: ")
	fmt.Scanln(&harga)
	fmt.Print("Masukkan stok obat: ")
	fmt.Scanln(&stok)
	fmt.Print("Masukkan golongan obat (B/BT/K): ")
	fmt.Scanln(&kodeGol)
	fmt.Print("Tanggal Kadaluarsa (dd mm yy): ")
	fmt.Scanln(&tgl.Hari, &tgl.Bulan, &tgl.Tahun)
	golongan := Golongan{}
	valid := true
	if kodeGol == "B" {
		golongan = GolonganBebas
	} else if kodeGol == "BT" {
		golongan = GolonganTerbatas
	} else if kodeGol == "K" {
		golongan = GolonganKeras
	} else {
		valid = false
	}
	if !valid {
		fmt.Println("Golongan tidak dikenali.")
		return
	}
	obats[*jumlah] = Obat{nama, jenis, kategori, harga, stok, golongan, tgl}
	*jumlah = *jumlah + 1
	fmt.Println("Obat ditambahkan.")
}

func tampilkanSemuaObat(obats [maxObat]Obat, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Belum ada data obat yang tersimpan.")
		return
	}
	fmt.Println("\n=== Daftar Obat ===")
	for i := 0; i < jumlah; i++ {
		o := obats[i]
		fmt.Printf("%s - %s - %s - %d - %d - %s (%s) - %02d/%02d/%04d\n",
			o.Nama, o.Jenis, o.Kategori, o.Harga, o.Stok,
			o.GolonganObat.Nama, o.GolonganObat.Kode,
			o.Kadaluarsa.Hari, o.Kadaluarsa.Bulan, o.Kadaluarsa.Tahun)
	}
}

func editObat(obats *[maxObat]Obat, jumlah int, nama string) {
	ditemukan := false
	for i := 0; i < jumlah; i++ {
		if obats[i].Nama == nama {
			ditemukan = true
			fmt.Println("Data ditemukan. Masukkan data baru: ")

			var namaBaru, jenisBaru, kodeGolongan, kategoriBaru string
			var hargaBaru, stokBaru int
			var tglBaru TanggalKadaluarsa

			fmt.Print("Nama baru: ")
			fmt.Scanln(&namaBaru)
			fmt.Print("Jenis baru: ")
			fmt.Scanln(&jenisBaru)
			fmt.Print("Kategori baru: ")
			fmt.Scanln(&kategoriBaru)
			fmt.Print("Harga baru: ")
			fmt.Scanln(&hargaBaru)
			fmt.Print("Stok baru: ")
			fmt.Scanln(&stokBaru)
			fmt.Print("Golongan Obat baru (B/ BT/ K): ")
			fmt.Scanln(&kodeGolongan)
			fmt.Print("Tanggal Kadaluarsa baru (dd mm yy): ")
			fmt.Scanln(&tglBaru.Hari, &tglBaru.Bulan, &tglBaru.Tahun)

			var golonganBaru Golongan
			switch kodeGolongan {
			case "B":
				golonganBaru = GolonganBebas
			case "BT":
				golonganBaru = GolonganTerbatas
			case "K":
				golonganBaru = GolonganKeras
			default:
				fmt.Println("Kode golongan tidak dikenali. Edit dibatalkan.")
				return
			}
			obats[i] = Obat{
				Nama:         namaBaru,
				Jenis:        jenisBaru,
				Kategori:     kategoriBaru,
				Harga:        hargaBaru,
				Stok:         stokBaru,
				GolonganObat: golonganBaru,
				Kadaluarsa:   tglBaru,
			}
			fmt.Println("Data berhasil diubah.")
			return
		}
	}
	if !ditemukan {
		fmt.Println("Obat tidak ditemukan")
	}
}

func cariObatLinear(obats [maxObat]Obat, jumlah int, nama string) {
	ditemukan := false
	for i := 0; i < jumlah; i++ {
		if obats[i].Nama == nama {
			o := obats[i]
			fmt.Println("Data ditemukan: ")
			fmt.Printf("%s - %s - %d - %d - %s (%s) - %02d/%02d/%04d\n",
				o.Nama, o.Jenis, o.Harga, o.Stok,
				o.GolonganObat.Nama, o.GolonganObat.Kode,
				o.Kadaluarsa.Hari, o.Kadaluarsa.Bulan, o.Kadaluarsa.Tahun)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Obat tidak ditemukan.")
	}
}

func urutkanHargaObat(obats *[maxObat]Obat, jumlah int, ascending bool) {
	for i := 0; i < jumlah-1; i++ {
		min := i
		for j := i + 1; j < jumlah; j++ {
			if (ascending && obats[j].Harga < obats[min].Harga) || (!ascending && obats[j].Harga > obats[min].Harga) {
				min = j
			}
		}
		obats[i], obats[min] = obats[min], obats[i]
	}
	fmt.Println("Selesai diurutkan berdasarkan harga.")
}

func urutkanNamaObat(obats *[maxObat]Obat, jumlah int, ascending bool) {
	for i := 1; i < jumlah; i++ {
		temp := obats[i]
		j := i - 1
		for j >= 0 && ((ascending && obats[j].Nama > temp.Nama) || (!ascending && obats[j].Nama < temp.Nama)) {
			obats[j+1] = obats[j]
			j--
		}
		obats[j+1] = temp
	}
	fmt.Println("Selesai diurutkan sesuai nama obat.")
}

func urutkanKategoriObat(obats *[maxObat]Obat, jumlah int, ascending bool) {
	for i := 1; i < jumlah; i++ {
		temp := obats[i]
		j := i - 1
		for j >= 0 && ((ascending && obats[j].Kategori > temp.Kategori) || (!ascending && obats[j].Kategori < temp.Kategori)) {
			obats[j+1] = obats[j]
			j--
		}
		obats[j+1] = temp
	}
	fmt.Println("Selesai diurutkan berdasarkan kategori obat.")
}

// Fungsi untuk menyimpan data ke file Excel
func simpanKeExcel(obats [maxObat]Obat, jumlah int) {
	f := excelize.NewFile()
	sheet := "DataObat"
	index := f.NewSheet(sheet)

	// Header
	headers := []string{"Nama", "Jenis", "Kategori", "Harga", "Stok", "Golongan", "Kode Golongan", "Kadaluarsa (dd/mm/yyyy)"}
	for i, h := range headers {
		cell := fmt.Sprintf("%c1", 'A'+i)
		f.SetCellValue(sheet, cell, h)
	}

	// Data
	for i := 0; i < jumlah; i++ {
		obat := obats[i]
		row := i + 2
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), obat.Nama)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), obat.Jenis)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), obat.Kategori)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), obat.Harga)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", row), obat.Stok)
		f.SetCellValue(sheet, fmt.Sprintf("F%d", row), obat.GolonganObat.Nama)
		f.SetCellValue(sheet, fmt.Sprintf("G%d", row), obat.GolonganObat.Kode)
		tglStr := fmt.Sprintf("%02d/%02d/%04d", obat.Kadaluarsa.Hari, obat.Kadaluarsa.Bulan, obat.Kadaluarsa.Tahun)
		f.SetCellValue(sheet, fmt.Sprintf("H%d", row), tglStr)
	}

	f.DeleteSheet("Sheet1")
	f.SetActiveSheet(index)

	err := f.SaveAs("data_obat.xlsx")
	if err != nil {
		fmt.Println("Gagal menyimpan file Excel:", err)
		return
	}
	fmt.Println("Data berhasil disimpan ke data_obat.xlsx")
}

// Fungsi menu utama
func menuUtama() {
	for {
		fmt.Println("=== SISTEM PENGELOLAAN DATA OBAT APOTEK ===")
		fmt.Println("1. Tambah Data Obat")
		fmt.Println("2. Tampilkan Semua Data Obat")
		fmt.Println("3. Edit Obat")
		fmt.Println("4. Cari Obat Linear")
		fmt.Println("5. Urutkan Harga Obat")
		fmt.Println("6. Urutkan Nama Obat")
		fmt.Println("7. Urutkan Kategori Obat")
		fmt.Println("8. Simpan Data ke Excel")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			fmt.Println("\n -- Tambah Data Obat --")
			tambahObat(&DaftarObat, &JumlahObat)
		case 2:
			fmt.Println("\n-- Tampilkan Semua Data Obat --")
			tampilkanSemuaObat(DaftarObat, JumlahObat)
		case 3:
			var n string
			fmt.Print("Masukkan nama obat yang akan diedit: ")
			fmt.Scanln(&n)
			editObat(&DaftarObat, JumlahObat, n)
		case 4:
			var n string
			fmt.Print("Masukkan nama obat yang ingin dicari: ")
			fmt.Scanln(&n)
			cariObatLinear(DaftarObat, JumlahObat, n)
		case 5:
			var input string
			var urutanHarga bool
			fmt.Print("Urutkan dari murah ke mahal? (ya/ tidak): ")
			fmt.Scanln(&input)
			urutanHarga = input == "ya" || input == "YA" || input == "Ya"
			urutkanHargaObat(&DaftarObat, JumlahObat, urutanHarga)
		case 6:
			var input string
			var urutanNama bool
			fmt.Print("Urutkan nama obat dari A ke Z? (ya/ tidak): ")
			fmt.Scanln(&input)
			urutanNama = input == "ya" || input == "YA" || input == "Ya"
			urutkanNamaObat(&DaftarObat, JumlahObat, urutanNama)
		case 7:
			var input string
			var urutanKategori bool
			fmt.Print("Urutkan kategori dari generik ke dagang? (ya/ tidak): ")
			fmt.Scanln(&input)
			urutanKategori = input == "ya" || input == "YA" || input == "Ya"
			urutkanKategoriObat(&DaftarObat, JumlahObat, urutanKategori)
		case 8:
			fmt.Println("Menyimpan data ke Excel...")
			simpanKeExcel(DaftarObat, JumlahObat)
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func main() {
	menuUtama()
}
