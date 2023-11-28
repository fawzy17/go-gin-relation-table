package mahasiswacontroller

import (
	"net/http"

	"github.com/fawzy17/gin-relation-table/config"
	"github.com/fawzy17/gin-relation-table/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(ctx *gin.Context) {
	var mahasiswas []models.Mahasiswa

	if err := config.DB.Preload("Kelas").Find(&mahasiswas).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ctx.AbortWithStatusJSON(http.StatusNotFound, models.Response{
				StatusCode: http.StatusNotFound,
				IsSuccess:  false,
				Message:    "Failed",
				Data:       nil,
			})
			return
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
				StatusCode: http.StatusInternalServerError,
				IsSuccess:  false,
				Message:    "Failed",
				Data:       nil,
			})
			return
		}

	}

	mahasiswaResponse := make([]models.MahasiswaResponse, 0)

	for _, m := range mahasiswas {
		mahasiswaResponse = append(mahasiswaResponse, models.MahasiswaResponse{
			ID:        m.ID,
			Nama:      m.Nama,
			Npm:       m.Npm,
			KelasKode: m.KelasKode,
			Kelas:     m.Kelas,
		})
	}

	ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		IsSuccess:  true,
		Data:       mahasiswaResponse,
		Message:    "Success",
	})

}

func Create(ctx *gin.Context) {
	var mahasiswa models.Mahasiswa
	var kelas models.Kelas

	if err := ctx.ShouldBindJSON(&mahasiswa); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			IsSuccess:  false,
			Message:    http.StatusText(http.StatusBadRequest),
			Data:       nil,
		})
		return
	}

	if err := config.DB.Where("kode = ?", mahasiswa.KelasKode).Find(&kelas).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			IsSuccess:  false,
			Message:    "Kelas tidak ditemukan",
			Data:       nil,
		})
		return
	}

	if err := config.DB.Create(&mahasiswa).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, models.Response{
			StatusCode: http.StatusNotAcceptable,
			IsSuccess:  false,
			Message:    "Gagal membuat data",
			Data:       nil,
		})
		return
	}

	newKelas := models.Kelas{
		Kode:   kelas.Kode,
		Name:   kelas.Name,
		Jumlah: kelas.Jumlah + 1,
	}
	config.DB.Model(&kelas).Where("kode = ?", kelas.Kode).Updates(&newKelas)

	mahasiswaResponse := models.MahasiswaResponse{
		ID:        mahasiswa.ID,
		Nama:      mahasiswa.Nama,
		Npm:       mahasiswa.Npm,
		KelasKode: mahasiswa.KelasKode,
		Kelas:     newKelas,
	}

	ctx.JSON(http.StatusCreated, models.Response{
		StatusCode: http.StatusCreated,
		IsSuccess:  true,
		Message:    "Data berhasil dibuat",
		Data:       mahasiswaResponse,
	})
}
