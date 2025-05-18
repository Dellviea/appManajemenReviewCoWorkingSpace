package main

import "fmt"

const MAXCoWorkSpace = 100
const MAXULASAN = 100
const JFASILITAS = 3

type CoWorkSpace struct {
	ID           int
	Nama         string
	Lokasi       string
	Fasilitas    [JFASILITAS]bool
	Harga        int
	Rating       float64
	JumlahReview int
	Ulasan       [MAXULASAN]string
	RatingUser   [MAXULASAN]float64
	JumlahUlasan int
}

var dataCoWorkSpace [MAXCoWorkSpace]CoWorkSpace
var jumlahData int
var nextID int = 1

func addData(d CoWorkSpace) {
	if jumlahData < MAXCoWorkSpace {
		d.ID = nextID
		nextID++
		dataCoWorkSpace[jumlahData] = d
		jumlahData++
		fmt.Println("CoWorkSpace ditambahkan dengan ID", d.ID)
	} else {
		fmt.Println("Kapasitas penuh!")
	}
}

func changeData(id int, dataBaru CoWorkSpace) {
	idx := cariByID(id)
	if idx != -1 {
		dataBaru.ID = id
		dataBaru.JumlahReview = dataCoWorkSpace[idx].JumlahReview
		dataBaru.Rating = dataCoWorkSpace[idx].Rating
		dataBaru.JumlahUlasan = dataCoWorkSpace[idx].JumlahUlasan
		dataBaru.Ulasan = dataCoWorkSpace[idx].Ulasan
		dataBaru.RatingUser = dataCoWorkSpace[idx].RatingUser
		dataCoWorkSpace[idx] = dataBaru
		fmt.Println("CoWorkSpace berhasil diubah.")
	} else {
		fmt.Println("CoWorkSpace tidak ditemukan.")
	}
}

func deleteData(id int) {
	idx := cariByID(id)
	if idx != -1 {
		for i := idx; i < jumlahData-1; i++ {
			dataCoWorkSpace[i] = dataCoWorkSpace[i+1]
		}
		jumlahData--
		fmt.Println("CoWorkSpace berhasil dihapus.")
	} else {
		fmt.Println("CoWorkSpace tidak ditemukan.")
	}
}

func addReview(id int, newRating float64) {
	idx := cariByID(id)
	if idx != -1 {
		d := &dataCoWorkSpace[idx]
		totalRating := d.Rating * float64(d.JumlahReview)
		totalRating += newRating
		d.JumlahReview++
		d.Rating = totalRating / float64(d.JumlahReview)
		fmt.Println("Ulasan berhasil ditambahkan.")
	} else {
		fmt.Println("CoWorkSpace tidak ditemukan.")
	}
}

func cariByID(id int) int {
	for i := 0; i < jumlahData; i++ {
		if dataCoWorkSpace[i].ID == id {
			return i
		}
	}
	return -1
}

// Pencarian Sequential by Nama
func sequentialSearchByNama(nama string) int {
	for k := 0; k < jumlahData; k++ {
		if dataCoWorkSpace[k].Nama == nama {
			return k
		}
	}
	return -1
}

// Sorting Lokasi Descending (buat ngebantu Binary Search Lokasi) BAGIAN C
func sortByLokasiDescending() {
	var idx int
	var temp CoWorkSpace

	for pass := 1; pass <= jumlahData-1; pass++ {
		idx = pass - 1
		for i := pass; i < jumlahData; i++ {
			if dataCoWorkSpace[i].Lokasi > dataCoWorkSpace[idx].Lokasi {
				idx = i
			}
		}
		temp = dataCoWorkSpace[pass-1]
		dataCoWorkSpace[pass-1] = dataCoWorkSpace[idx]
		dataCoWorkSpace[idx] = temp
	}
}

// Binary Search by Lokasi Descending
func binarySearchByLokasiDescending(lokasi string) int {
	sortByLokasiDescending()
	left, right := 0, jumlahData-1
	for left <= right {
		mid := (left + right) / 2
		if lokasi > dataCoWorkSpace[mid].Lokasi {
			right = mid - 1
		} else if lokasi < dataCoWorkSpace[mid].Lokasi {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// Pengurutan Selection Sort by Harga Descending BAGIAN D
func selectionSortHargaDescending() {
	var idx int
	var temp CoWorkSpace

	for pass := 1; pass <= jumlahData-1; pass++ {
		idx = pass - 1
		for i := pass; i < jumlahData; i++ {
			if dataCoWorkSpace[i].Harga > dataCoWorkSpace[idx].Harga {
				idx = i
			}
		}
		temp = dataCoWorkSpace[pass-1]
		dataCoWorkSpace[pass-1] = dataCoWorkSpace[idx]
		dataCoWorkSpace[idx] = temp
	}
}

// Pengurutan Insertion Sort by Rating Descending
func insertionSortRatingDescending() {
	var temp CoWorkSpace

	for pass := 1; pass <= jumlahData-1; pass++ {
		i := pass
		temp = dataCoWorkSpace[pass]
		for i > 0 && temp.Rating > dataCoWorkSpace[i-1].Rating {
			dataCoWorkSpace[i] = dataCoWorkSpace[i-1]
			i--
		}
		dataCoWorkSpace[i] = temp
	}
}
