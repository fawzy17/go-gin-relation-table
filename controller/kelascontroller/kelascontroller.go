package kelascontroller

import (
	"net/http"

	"github.com/fawzy17/gin-relation-table/config"
	"github.com/fawzy17/gin-relation-table/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(ctx *gin.Context) {
	var kelas []models.Kelas

	config.DB.Find(&kelas)

	if err := config.DB.Find(&kelas).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ctx.AbortWithStatusJSON(http.StatusNotFound, models.Response{
				StatusCode: http.StatusNotFound,
				IsSuccess:  false,
				Message:    "Record Not Found",
				Data:       nil,
			})
			return
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
				StatusCode: http.StatusInternalServerError,
				IsSuccess:  false,
				Message:    "Internal Server Error",
				Data:       nil,
			})
			return
		}
	}

	kelasResponse := make([]models.KelasResponse, 0)
	for _, k := range kelas {
		kelasResponse = append(kelasResponse, models.KelasResponse{
			KodeKelas: k.Kode,
			Nama:      k.Name,
			Jumlah:    k.Jumlah,
		})
	}

	ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		IsSuccess:  true,
		Message:    "Success",
		Data:       kelasResponse,
	})
}

func Show(ctx *gin.Context) {
	var kelas models.Kelas

	kode := ctx.Param("kode")

	if err := config.DB.Where("kode = ?", kode).First(&kelas).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ctx.AbortWithStatusJSON(http.StatusNotFound, models.Response{
				StatusCode: http.StatusNotFound,
				IsSuccess:  false,
				Message:    "Record Not Found",
				Data:       nil,
			})
			return
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
				StatusCode: http.StatusInternalServerError,
				IsSuccess:  false,
				Message:    "Internal Server Error",
				Data:       nil,
			})
			return
		}
	}

	kelasResponse := models.KelasResponse{
		KodeKelas: kelas.Kode,
		Nama:      kelas.Name,
		Jumlah:    kelas.Jumlah,
	}

	ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		IsSuccess:  true,
		Data:       kelasResponse,
		Message:    "Success",
	})

}

func Create(ctx *gin.Context) {
	var kelas models.Kelas

	if err := ctx.ShouldBindJSON(&kelas); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			IsSuccess:  false,
			Data:       nil,
			Message:    "Bad Request",
		})
		return
	}

	if err := config.DB.Create(&kelas).Error; err != nil {
		ctx.JSON(http.StatusNotAcceptable, models.Response{
			StatusCode: http.StatusNotAcceptable,
			IsSuccess:  false,
			Data:       nil,
			Message:    "Failed",
		})
		return
	}

	kelasResponse := models.KelasResponse{
		KodeKelas: kelas.Kode,
		Nama:      kelas.Name,
		Jumlah:    kelas.Jumlah,
	}

	ctx.JSON(http.StatusCreated, models.Response{
		StatusCode: http.StatusCreated,
		IsSuccess:  true,
		Data:       kelasResponse,
		Message:    "Success",
	})

}

func Delete(ctx *gin.Context) {
	kode := ctx.Param("kode")

	var kelas models.Kelas

	if err := config.DB.Where("kode = ?", kode).First(&kelas).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ctx.AbortWithStatusJSON(http.StatusNotFound, models.Response{
				StatusCode: http.StatusNotFound,
				IsSuccess:  false,
				Data:       nil,
				Message:    "Record Not Found",
			})
			return
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
				StatusCode: http.StatusInternalServerError,
				IsSuccess:  false,
				Data:       nil,
				Message:    "Internal Server Error",
			})
			return
		}
	}

	if config.DB.Delete(&kelas).RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, models.Response{
			StatusCode: http.StatusNotAcceptable,
			IsSuccess:  false,
			Data:       nil,
			Message:    "Failed",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		IsSuccess:  true,
		Data:       kelas,
		Message:    "Deleted",
	})

}

func ShowMahasiswaInClass(ctx *gin.Context) {
	var mahasiswas []models.Mahasiswa
	var kelas models.Kelas

	kode := ctx.Param("kode")

	if err := config.DB.Where("kode = ?", kode).First(&kelas).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ctx.AbortWithStatusJSON(http.StatusNotFound, models.Response{
				StatusCode: http.StatusNotFound,
				IsSuccess:  false,
				Data:       nil,
				Message:    "Failed",
			})
			return
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
				StatusCode: http.StatusInternalServerError,
				IsSuccess:  false,
				Data:       nil,
				Message:    "Failed",
			})
			return
		}
	}
	config.DB.Where("kelas_kode = ?", kode).Find(&mahasiswas)

	mahasiswaResponse := make([]map[string]interface{}, 0)
	for _, m := range mahasiswas {
		mahasiswaResponse = append(mahasiswaResponse, map[string]interface{}{
			"id":         m.ID,
			"nama":       m.Nama,
			"npm":        m.Npm,
			"kelas_kode": m.KelasKode,
		})
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status_code": http.StatusOK,
		"is_success":  true,
		"data": map[string]interface{}{
			"KodeKelas": kelas.Kode,
			"Nama":      kelas.Name,
			"Jumlah":    kelas.Jumlah,
			"mahasiswa": mahasiswaResponse,
		},
		"message": "Success",
	})

}
