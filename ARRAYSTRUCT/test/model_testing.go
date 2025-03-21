package main

import (
	"ARRAYSTRUCT/model"
	"ARRAYSTRUCT/node"
	"fmt"
)

func main() {
	// Create Pegawai instances
	pegawais := []node.Pegawai{
		{
			ID:     1,
			Nama:   "John Doe",
			Alamat: node.Address{"Jalan Sudirman", "Jakarta", 123},
			NoTelp: "08123456789",
			Email:  "gmail.com",
		},
		{
			ID:     2,
			Nama:   "Jane Doe",
			Alamat: node.Address{"Jalan Sudirman", "Jakarta", 123},
			NoTelp: "08123456789",
			Email:  "yahoo.com",
		},
	}

	// Create Pegawai instances in the model
	for _, pegawai := range pegawais {
		model.CreatePegawai(pegawai)
	}

	// Read and print Pegawai instances
	fmt.Println("Before Update:")
	for _, emp := range model.ReadPegawai() {
		fmt.Println(emp)
	}

	// Update a Pegawai instance
	pegawai3 := node.Pegawai{
		ID:     2,
		Nama:   "John Doe",
		Alamat: node.Address{"Jalan Ahmad Jais", "Surabaya", 123},
		NoTelp: "08123456789",
		Email:  "gmail.com",
	}
	model.UpdatePegawai(pegawai3, 2)

	// Read and print Pegawai instances after update
	fmt.Println("After Update:")
	for _, emp := range model.ReadPegawai() {
		fmt.Println(emp)
	}

	// Delete a Pegawai instance
	// model.DeletePegawai(2)

	// Read and print Pegawai instances after delete
	// fmt.Println("After Delete:")
	// for _, emp := range model.ReadPegawai() {
	// 	fmt.Println(emp)
	// }
}

func main() {
	// Create a new Pegawai ke -1  instance
	pegawai1 := node.Pegawai{
		ID:     1,
		Nama:   "John Doe",
		Alamat: node.Address{"Jalan Sudirman", "Jakarta", 123},
		NoTelp: "08123456789",
		Email:  "gmail.com",
	}

	// Create a new Pegawai ke -2  instance
	pegawai2 := node.Pegawai{
		ID:     2,
		Nama:   "Jane Doe",
		Alamat: node.Address{"Jalan Sudirman", "Jakarta", 123},
		NoTelp: "08123456789",
		Email:  "yahoo.com",
	}

	model.CreatePegawai(pegawai1)
	model.CreatePegawai(pegawai2)
	// listpegawai := model.ReadPegawai()
	// fmt.Println(listpegawai)

	for _, emp := range model.ReadPegawai() {
		fmt.Println(emp)
	}

	// update testing
	pegawai3 := node.Pegawai{
		ID:     2,
		Nama:   "John Doe",
		Alamat: node.Address{"Jalan Ahmad Jais", "Surabaya", 123},
		NoTelp: "08123456789",
		Email:  "gmail.com",
	}
	model.UpdatePegawai(pegawai3, 2)
	fmt.Println("After Update")
	for _, emp := range model.ReadPegawai() {
		fmt.Println(emp)
	}

	// // delete testing
	// model.DeletePegawai(2)
	// fmt.Println("After Delete")
	// for _, emp := range model.ReadPegawai() {
	// 	fmt.Println(emp)
	// }

}
