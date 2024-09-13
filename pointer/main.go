package main

import "fmt"

func main() {
	// OS menyediakan tempat pada memori
	i := 10

	fmt.Println(i)
	fmt.Println(&i)

	// var "x" menampung alamat memori dari variable "i"
	x := &i

	fmt.Println(x) // menampilkan alamat memori "0xc00009e018"

	// ubah var "i"
	i = 20

	fmt.Println(*x) // untuk mendapat nilai asli dari alamat memori

	var n int64
	getNilai(&n)
	fmt.Println(n) // 10

	setNilai(nil)

	angka := 1
	foo(&angka)
	foo(&angka)
	foo(&angka)
	fmt.Println("angka", angka)
}

func foo(i *int) {
	*i++
}

// mendapatkan nilai dari fungsi melalui argument variable
func getNilai(n *int64) {
	result := int64(10)
	*n = result
}

// membuat argument fungsi bersifat opsional
func setNilai(n *int64) {

}
