package main

import (
	"fmt"
	"go-project/subject"
	"go-project/user"
)

func main() {
	mtk := subject.NewMatematika()

	mhs := user.NewMahasiswa(mtk)
	mhs.SetNama("Yanto")
	fmt.Println(mhs.GetProfile())
}
