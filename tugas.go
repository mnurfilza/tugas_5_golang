package tugas5

import (
	"fmt"
	"math/rand"
)

type Pelajar struct {
	nama string
	umur int
}

func Tugas5() {
	var n1, n2, n3 Pelajar
	n1.nama = "Ahmad"
	n2.nama = "Abi"
	n3.nama = "Zaka"

	n1.umur = rand.Intn(100)
	n2.umur = rand.Intn(100)
	n3.umur = rand.Intn(100)

	fmt.Println(n1.nama, n1.umur)
	fmt.Println(n2.nama, n2.umur)
	fmt.Println(n3.nama, n3.umur)
}
