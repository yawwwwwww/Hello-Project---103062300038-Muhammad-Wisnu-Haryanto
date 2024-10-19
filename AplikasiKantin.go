//Nama : Muhammad Wisnu Haryanto
//NIM : 103062300038
//Kelas : S1-IT-KJ-23-001

package main

import (
	"fmt"
	"os"
	"sort"
)

type Tenant struct {
	Nama            string
	TotalTransaksi  float64
	JumlahTransaksi int
}

type Transaksi struct {
	NamaTenant string
	Jumlah     float64
}

var tenants []Tenant
var transaksi []Transaksi

func tambahTenant(nama string) {
	tenant := Tenant{Nama: nama}
	tenants = append(tenants, tenant)
	sortTenantsByName()
}

func ubahTenant(namaLama string, namaBaru string) {
	index := binarySearchTenant(namaLama)
	if index != -1 {
		tenants[index].Nama = namaBaru
		sortTenantsByName()
	}
}

func hapusTenant(nama string) {
	index := binarySearchTenant(nama)
	if index != -1 {
		tenants = append(tenants[:index], tenants[index+1:]...)
	}
}

func tambahTransaksi(namaTenant string, jumlah float64) {
	index := binarySearchTenant(namaTenant)
	if index != -1 {
		transaksiBaru := Transaksi{NamaTenant: namaTenant, Jumlah: jumlah}
		transaksi = append(transaksi, transaksiBaru)
		tenants[index].TotalTransaksi += jumlah
		tenants[index].JumlahTransaksi++
		fmt.Println("Transaksi berhasil dicatat.")
	} else {
		fmt.Println("Tenant tidak ditemukan.")
	}
}

func hitungPendapatan() ([]float64, float64) {
	pendapatanTenant := make([]float64, len(tenants))
	var pendapatanAdmin float64
	for _, t := range transaksi {
		bagianTenant := t.Jumlah * 0.75
		bagianAdmin := t.Jumlah * 0.25
		index := binarySearchTenant(t.NamaTenant)
		if index != -1 {
			pendapatanTenant[index] += bagianTenant
		}
		pendapatanAdmin += bagianAdmin
	}
	return pendapatanTenant, pendapatanAdmin
}

func selectionSortTenants() {
	n := len(tenants)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if tenants[j].JumlahTransaksi > tenants[maxIdx].JumlahTransaksi {
				maxIdx = j
			}
		}
		tenants[i], tenants[maxIdx] = tenants[maxIdx], tenants[i]
	}
}

func sortTenantsByName() {
	sort.Slice(tenants, func(i, j int) bool {
		return tenants[i].Nama < tenants[j].Nama
	})
}

func binarySearchTenant(nama string) int {
	sortTenantsByName()
	left, right := 0, len(tenants)-1
	for left <= right {
		mid := left + (right-left)/2
		if tenants[mid].Nama == nama {
			return mid
		}
		if tenants[mid].Nama < nama {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func daftarTenantBerdasarkanTransaksi() {
	selectionSortTenants()
	file, _ := os.Create("daftar_tenant.txt")
	defer file.Close()

	fmt.Println("Daftar Tenant berdasarkan banyak transaksi:")
	for _, tenant := range tenants {
		output := fmt.Sprintf("Nama: %s, Jumlah Transaksi: %d, Total Uang: %.2f\n", tenant.Nama, tenant.JumlahTransaksi, tenant.TotalTransaksi)
		file.WriteString(output)
		fmt.Print(output)
	}
	fmt.Println("Daftar tenant berhasil ditulis ke daftar_tenant.txt")
}

func tampilkanPendapatanKeFile() {
	pendapatanTenant, pendapatanAdmin := hitungPendapatan()
	file, _ := os.Create("pendapatan.txt")
	defer file.Close()

	fmt.Println("Pendapatan Tenant:")
	for i, tenant := range tenants {
		output := fmt.Sprintf("Tenant Nama: %s, Pendapatan: %.2f\n", tenant.Nama, pendapatanTenant[i])
		file.WriteString(output)
		fmt.Print(output)
	}
	outputAdmin := fmt.Sprintf("Pendapatan Admin: %.2f\n", pendapatanAdmin)
	file.WriteString(outputAdmin)
	fmt.Print(outputAdmin)

	fmt.Println("Pendapatan berhasil ditulis ke pendapatan.txt")
}

func main() {
	var pilihan int
	for {
		fmt.Println(`
===================================
|         Menu                    |
| 1. Tambah Tenant                |
| 2. Ubah Tenant                  |
| 3. Hapus Tenant                 |
| 4. Tambah Transaksi             |
| 5. Tampilkan Pendapatan         |
| 6. Tampilkan Daftar Tenant      |
|    Berdasarkan Banyak Transaksi |
| 7. Keluar                       |
===================================
		`)
		fmt.Print("Pilih opsi: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			var nama string
			fmt.Print("Masukkan Nama Tenant: ")
			fmt.Scan(&nama)
			tambahTenant(nama)
			fmt.Println("Tenant berhasil ditambahkan.")
		} else if pilihan == 2 {
			var namaLama, namaBaru string
			fmt.Print("Masukkan Nama Tenant yang ingin diubah: ")
			fmt.Scan(&namaLama)
			fmt.Print("Masukkan Nama Baru: ")
			fmt.Scan(&namaBaru)
			ubahTenant(namaLama, namaBaru)
			fmt.Println("Tenant berhasil diubah.")
		} else if pilihan == 3 {
			var nama string
			fmt.Print("Masukkan Nama Tenant yang ingin dihapus: ")
			fmt.Scan(&nama)
			hapusTenant(nama)
			fmt.Println("Tenant berhasil dihapus.")
		} else if pilihan == 4 {
			var namaTenant string
			var jumlah float64
			fmt.Print("Masukkan Nama Tenant: ")
			fmt.Scan(&namaTenant)
			fmt.Print("Masukkan Jumlah Transaksi: ")
			fmt.Scan(&jumlah)
			tambahTransaksi(namaTenant, jumlah)
		} else if pilihan == 5 {
			tampilkanPendapatanKeFile()
		} else if pilihan == 6 {
			daftarTenantBerdasarkanTransaksi()
		} else if pilihan == 7 {
			fmt.Println("Terimakasih Telah Menggunakan Aplikasi Kami")
			break
		}
	}
}
