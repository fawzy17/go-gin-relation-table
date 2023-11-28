package main

import (
	"github.com/fawzy17/gin-relation-table/config"
	"github.com/fawzy17/gin-relation-table/controller/kelascontroller"
	"github.com/fawzy17/gin-relation-table/controller/mahasiswacontroller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()

	prefix := "api/v1"

	// Kelas endpoint collection
	r.GET(prefix+"/kelas", kelascontroller.Index)
	r.GET(prefix+"/kelas/:kode", kelascontroller.Show)
	r.GET(prefix+"/kelas/show-mahasiswa-in-class/:kode", kelascontroller.ShowMahasiswaInClass)
	r.POST(prefix+"/kelas", kelascontroller.Create)
	r.DELETE(prefix+"/kelas/:kode", kelascontroller.Delete)

	// Mahasiswa endpoint collection
	r.GET(prefix+"/mahasiswa", mahasiswacontroller.Index)
	r.POST(prefix+"/mahasiswa", mahasiswacontroller.Create)

	r.Run()
}
