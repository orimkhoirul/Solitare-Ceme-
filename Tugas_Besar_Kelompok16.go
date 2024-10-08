package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

/*
tiles={(0,0),(0,1),(0,2),(0,3),(0,4),(0,5),(0,6),(1,1),(1,2),(1,3),(1,4),(1,5),(1,6),(2,2),(2,3),(2,4),(2,5),(2,6),(3,3),(3,4),(3,5),
(3,6),(4,4),(4,5),(4,6),(5,5),(5,6),(6,6)}
*/


const NMAX int = 2000



type User struct {
	nama  string
	score int
}

type pemain struct {
	nama  string
	skor  int
	Tiles arrTiles
}

type tiles struct {
	value1 int
	value2 int
}

type arrTiles struct {
	ArrTiles [28]tiles
	JmlTiles int
}
type arr [28]int

type arrUser struct {
	arrU    [NMAX]User
	jmlUser int
}

func clearline() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func headerUtama() {

	fmt.Println("===========================================")
	fmt.Println("=====       Domino Solitaire Ceme       =====")
	fmt.Println("=============================================")

}

func headerAwal() {

	fmt.Println("============================================")
	fmt.Println("============   SELAMAT DATANG   ============")
	fmt.Println("=====                 Di                 =====")
	fmt.Println("==========  Domino Solitaire Ceme  ===========")
	fmt.Println("============================================")
	fmt.Println("->     Mochammad Khoirullutfansyah        <-")
	fmt.Println("->      Muhammad Zaidan Dhiyaulhaq        <-")
	fmt.Println("============================================")

}

func title() {

	headerAwal()
	fmt.Println("Tekan Enter untuk mulai permainan...")
	fmt.Scanln()
	fmt.Println("Loading...")
	time.Sleep(1 * time.Second)
	clearline()
}

func Menu(p *arrUser) {

	headerUtama()

	var pilihan string

	fmt.Println("Menu : ")
	fmt.Println("1. Mulai Permainan")
	fmt.Println("2. Tampilkan Pemain")
	fmt.Println("0. Keluar dari permainan")
	fmt.Println("-------------------------")
	fmt.Print("Pilih Menu : ")
	fmt.Scan(&pilihan)

	if pilihan == "1" {
		mulai_Game(p)
	} else if pilihan == "2" {
		fmt.Println("1. Tampilkan ranking dari yang terbaik")
		fmt.Println("2. Tampilkan ranking dari yang terendah")
		fmt.Println("3. Cari skor pemain dengan nama tertentu")
		fmt.Println("0. Kembali ke menu utama")
		fmt.Println("-------------------------")
		fmt.Print("Pilih Menu : ")
		fmt.Scan(&pilihan)
		for pilihan != "0" {
			if pilihan == "1" {
				rankUserDescending(p)
				CetakUser(*p)
			} else if pilihan == "2" {
				rankUserAscending(p)
				CetakUser(*p)
			} else if pilihan == "3" {
				var nama string
				fmt.Print("Masukan nama pemain yang ingin ditampilkan :")
				fmt.Scan(&nama)
				rankUserAscendingNama(p)
				TampilPemainTertentu(*p, nama)
			}

			fmt.Println("1. Tampilkan ranking dari yang terbaik")
			fmt.Println("2. Tampilkan ranking dari yang terendah")
			fmt.Println("3. Cari skor pemain dengan nama tertentu")
			fmt.Println("0. Kembali ke menu utama")
			fmt.Println("-------------------------")
			fmt.Print("Pilih Menu : ")
			fmt.Scan(&pilihan)

		}
		Menu(p)
	}
}

// tampikan ranking pemain berdasarkan skor terbesar ke terkecil

func rankUserDescending(p *arrUser) {
	var temp User
	var j int

	for i := 0; i < p.jmlUser; i++ {
		temp = p.arrU[i]
		j = i

		for j > 0 && temp.score > p.arrU[j-1].score {
			p.arrU[j] = p.arrU[j-1]
			j--
		}
		p.arrU[j] = temp
	}
}

func rankUserAscending(p *arrUser) {

	var temp User
	var j int

	for i := 0; i < p.jmlUser; i++ {
		temp = p.arrU[i]
		j = i

		for j > 0 && temp.score < p.arrU[j-1].score {
			p.arrU[j] = p.arrU[j-1]
			j--
		}
		p.arrU[j] = temp
	}

}

func rankUserAscendingNama(p *arrUser) {
	var temp User
	var j int

	for i := 0; i < p.jmlUser; i++ {
		temp = p.arrU[i]
		j = i

		for j > 0 && temp.nama < p.arrU[j-1].nama {
			p.arrU[j] = p.arrU[j-1]
			j--
		}
		p.arrU[j] = temp
	}

}

/*fmt.Println("Kembali ke menu? (y/n)")
var pilihan string
fmt.Scan(&pilihan)
if pilihan == "y" || pilihan == "Y" {
	Menu(p)
}*/

func CetakUser(t arrUser) {
	for i := 0; i < t.jmlUser; i++ {
		fmt.Println("-------------------------")
		fmt.Println("Nama:", t.arrU[i].nama, " Skor:", t.arrU[i].score)
		fmt.Println("-------------------------")
	}
}

func isiTiles(t *arrTiles) {
	for i := 0; i < 28; i++ {
		if i < 7 {
			t.ArrTiles[i].value1 = 0
			t.ArrTiles[i].value2 = i
		} else if i < 13 {
			t.ArrTiles[i].value1 = 1
			t.ArrTiles[i].value2 = i - 6
		} else if i < 18 {
			t.ArrTiles[i].value1 = 2
			t.ArrTiles[i].value2 = i - 11
		} else if i < 22 {
			t.ArrTiles[i].value1 = 3
			t.ArrTiles[i].value2 = i - 15
		} else if i < 25 {
			t.ArrTiles[i].value1 = 4
			t.ArrTiles[i].value2 = i - 18
		} else if i < 27 {
			t.ArrTiles[i].value1 = 5
			t.ArrTiles[i].value2 = i - 20
		} else {
			t.ArrTiles[i].value1 = 6
			t.ArrTiles[i].value2 = 6
		}
	}
	t.JmlTiles = 28
}

func CekAda(t arr, n int, s int) bool {
	for i := 0; i < n; i++ {
		if t[i] == s {
			return true
		}

	}
	return false
}

func AcakTiles(t *arrTiles) {
	var indeks [28]int
	var T arrTiles
	rand.Seed(time.Now().UnixNano())
	var random int = rand.Intn(28)
	for i := 0; i < t.JmlTiles; i++ {
		for CekAda(indeks, i, random) {
			random = rand.Intn(28)
		}
		indeks[i] = random
	}

	for i := 0; i < t.JmlTiles; i++ {
		T.ArrTiles[i].value1 = t.ArrTiles[indeks[i]].value1
		T.ArrTiles[i].value2 = t.ArrTiles[indeks[i]].value2
	}
	for i := 0; i < t.JmlTiles; i++ {
		t.ArrTiles[i].value1 = T.ArrTiles[i].value1
		t.ArrTiles[i].value2 = T.ArrTiles[i].value2
	}
}

// Mengambil tile dari array t secara random
func ambilTile(tUser *arrTiles, t *arrTiles) {
	rand.Seed(time.Now().UnixNano())
	var indeks int = rand.Intn(t.JmlTiles)
	if tUser.JmlTiles < 4 {
		tUser.ArrTiles[tUser.JmlTiles].value1 = t.ArrTiles[indeks].value1
		tUser.ArrTiles[tUser.JmlTiles].value2 = t.ArrTiles[indeks].value2
		tUser.JmlTiles++
		hapusTile(t.ArrTiles[indeks].value1, t.ArrTiles[indeks].value2, t)

	}

}

// menghapus tile yang dipilih value1 dan value2, arraytile akan berkurang -1 dan array tile akan bergeser ke kiri

func hapusTile(value1, value2 int, t *arrTiles) {
	var indeks int = SeqSearch(value1, value2, *t)
	if indeks >= 0 {
		for j := indeks; j < 27; j++ {
			t.ArrTiles[j].value1 = t.ArrTiles[j+1].value1
			t.ArrTiles[j].value2 = t.ArrTiles[j+1].value2
		}
		t.JmlTiles--

	}

}

func SeqSearch(v1, v2 int, t arrTiles) int {
	for i := 0; i < t.JmlTiles; i++ {
		if t.ArrTiles[i].value1 == v1 && t.ArrTiles[i].value2 == v2 {
			return i
		}

	}
	return -1
}

func CetakTiles(t arrTiles) {
	for i := 0; i < t.JmlTiles; i++ {
		fmt.Print(t.ArrTiles[i], " ")
	}

}

func SubGame(User *pemain, Lawan *pemain, pilihan *int, bykPil *int) {
	var t arrTiles
	var decision, i int
	var bykPerintah int
	var p string
	var bykTiles int
	isiTiles(&t)
	AcakTiles(&t)
	fmt.Println("Dealing . . .")
	ambilTile(&User.Tiles, &t)
	ambilTile(&User.Tiles, &t)
	bykTiles = 2
	fmt.Println("Apakah anda masih ingin mengambil tiles (Ya)?")
	fmt.Scan(&p)
	for p == "Ya" && i < 2 {
		ambilTile(&User.Tiles, &t)
		bykTiles++
		fmt.Println("Apakah anda masih ingin mengambil tiles (Ya)?")
		fmt.Scan(&p)
		i++
	}
	if i == 2 && p == "Ya" {
		fmt.Println("Maaf anda hanya bisa mengambil maksimal 4 tiles")
	}

	ambilTile(&Lawan.Tiles, &t) //Lawan Mengambil tiles
	ambilTile(&Lawan.Tiles, &t)
	ambilTile(&Lawan.Tiles, &t)
	ambilTile(&Lawan.Tiles, &t)

	fmt.Print("Your Tiles ")
	CetakTiles(User.Tiles)
	fmt.Println()
	fmt.Print("(0=done, 9=exit, k=replace k-th tile, k = [1,2,3,4]) ")
	fmt.Println()
	fmt.Print("Decision?")
	fmt.Scan(&decision)

	for decision != 0 && bykPerintah < 2 && decision != 9 {
		if decision > bykTiles {
			fmt.Println("Tidak bisa mengganti tiles yang tidak diambil")
			bykPerintah--
		} else {
			GantiTiles(decision, User, &t)
			fmt.Print("Your Tiles ")
			CetakTiles(User.Tiles)
			fmt.Println()

		}

		fmt.Print("Decision? ")
		fmt.Scan(&decision)

		*pilihan = decision
		bykPerintah++
		*bykPil = bykPerintah
	}
	*pilihan = decision
	*bykPil = bykPerintah

}

func Tukar(m, n *int) {
	var temp int
	temp = *m
	*m = *n
	*n = temp

}

func GantiTiles(pilihan int, User *pemain, t *arrTiles) { //Mengganti Tiles sesuai plihan user
	rand.Seed(time.Now().UnixNano())
	var indeks int = rand.Intn(t.JmlTiles)

	User.Tiles.ArrTiles[pilihan-1].value1 = t.ArrTiles[indeks].value1
	User.Tiles.ArrTiles[pilihan-1].value2 = t.ArrTiles[indeks].value2

	hapusTile(t.ArrTiles[indeks].value1, t.ArrTiles[indeks].value2, t)

}

func mulai_Game(A *arrUser) {

	clearline()
	headerUtama()
	var User, lawan pemain
	var JmlMenang int
	var pilihan int = 7
	var JmlGame int = 0
	var bykPilihan int = 0
	fmt.Print("Masukan nama : ")
	fmt.Scan(&User.nama)
	fmt.Println("Your Score is 0/0")

	for pilihan != 9 {
		SubGame(&User, &lawan, &pilihan, &bykPilihan)
		if pilihan == 0 || bykPilihan == 2 {
			//cek Kemenangan
			if CekKemenangan(User, lawan) {
				fmt.Println("You Win")
				JmlMenang++
			} else {
				fmt.Println("You Lost")
			}
			fmt.Println("hasil cek sama ", CekSamaTile(User.Tiles), CekSamaTile(lawan.Tiles))
			fmt.Print("Tiles Lawan ")
			CetakTiles(lawan.Tiles)
			fmt.Println()
			JmlGame++
			fmt.Println("Your Score is ", JmlMenang, "/", JmlGame)
			fmt.Println()
			fmt.Println()

		}

		User.Tiles.JmlTiles = 0
		lawan.Tiles.JmlTiles = 0

	}
	fmt.Println("Your last score is ", JmlMenang, "/", JmlGame)
	fmt.Println("Thank You for playing with Us")
	fmt.Println("-------------------------")
	fmt.Println("Your winning rate is ", (float32(JmlMenang)/float32(JmlGame))*100, "%")
	User.skor = JmlMenang
	A.arrU[A.jmlUser].nama = User.nama //Memasukann pemain ke array untuk nantinya bisa disorting
	A.arrU[A.jmlUser].score = User.skor
	A.jmlUser++

	CekNama(A, User, A.jmlUser-1)

	Menu(A)

}

func CekSamaTile(t arrTiles) bool {
	var cek int = 0
	for i := 0; i < t.JmlTiles; i++ {
		if t.ArrTiles[i].value1 == t.ArrTiles[i].value2 {
			cek++
		}
	}
	if cek == 2 {
		return true
	}
	return false
}

func CekKemenangan(User pemain, lawan pemain) bool {
	if CekSamaTile(User.Tiles) && !CekSamaTile(lawan.Tiles) {
		return true
	} else if CekSamaTile(User.Tiles) && CekSamaTile(lawan.Tiles) {
		if JumlahTile(User.Tiles) > JumlahTile(lawan.Tiles) {
			return true
		} else {
			return false
		}
	} else {
		if JumlahTile(User.Tiles) > JumlahTile(lawan.Tiles) {
			return true
		} else {
			return false
		}
	}

}

func JumlahTile(t arrTiles) int {
	var max int = 0
	for i := 0; i < t.JmlTiles; i++ {
		if max < t.ArrTiles[i].value1+t.ArrTiles[i].value2 {
			max = t.ArrTiles[i].value1 + t.ArrTiles[i].value2
		}
	}

	return max
}

func CekNama(A *arrUser, s pemain, n int) {
	var cek bool = false
	for i := 0; i < A.jmlUser; i++ {
		if A.arrU[i].nama == s.nama && i != n {
			cek = true
			if A.arrU[i].score < s.skor {
				A.arrU[i].score = s.skor

			}
		}

	}
	if cek {
		A.jmlUser--
	}
}

func CariPemain(t arrUser, s string) int {

	//CetakUser(t)
	var ind int = -1
	var kanan, kiri, mid int
	kiri = 0
	kanan = t.jmlUser - 1
	mid = -1
	for kiri <= kanan && ind == -1 {
		mid = (kanan + kiri) / 2
		if t.arrU[mid].nama < s {
			kiri = mid + 1

		} else if t.arrU[mid].nama > s {
			kanan = mid - 1
		} else {
			ind = mid
		}
	}
	return ind
}

func TampilPemainTertentu(t arrUser, s string) {
	if CariPemain(t, s) != -1 {
		fmt.Println("-------------------------")
		fmt.Println("Nama:", t.arrU[CariPemain(t, s)].nama, " Skor:", t.arrU[CariPemain(t, s)].score)
		fmt.Println("-------------------------")
	} else {
		fmt.Println("Nama yang ingin ditampilkan tidak ada")
	}

}

func main() {
	var t arrUser
	title()
	Menu(&t)

}
