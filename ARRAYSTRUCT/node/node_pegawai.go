package node

type Address struct {
	Jalan string
	Kota  string
	Nomer int
}

type Pegawai struct {
	ID     int
	Nama   string
	Alamat Address
	NoTelp string
	Email  string
}
