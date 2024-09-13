package main

import "fmt"

type Mahasiswa struct {
	ID   int64
	Nama string
	NIM  int64
	IPK  float64
}

func main() {
	mhs := Mahasiswa{
		ID:   1,
		Nama: "Yanto",
		NIM:  123456,
		IPK:  3.5,
	}
	fmt.Println(mhs.ID, mhs.Nama)

	// Literal struct
	person := struct {
		Nama string
		Usia int64
	}{
		Nama: "Abdul",
		Usia: 20,
	}
	fmt.Println(person.Nama, person.Usia)

	// Slice of struct
	listMhs := []Mahasiswa{
		{
			ID:   2,
			Nama: "Bambang",
		},
		{
			ID:   3,
			Nama: "Eric",
		},
	}
	listMhs = append(listMhs, Mahasiswa{
		ID:   4,
		Nama: "Joko",
	})
	fmt.Println(listMhs[0].ID, listMhs[0].Nama)
	fmt.Println(len(listMhs))
}
