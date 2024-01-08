package addresscontroller

import (
	"encoding/json"
	"errors"
	"latihan_golang/config"
	"latihan_golang/helper"
	"latihan_golang/models"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Index(w http.ResponseWriter, r *http.Request){
	var address []models.Address

	if err := config.DB.Find(&address).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
	}

	helper.Response(w, 200, "List Address", address)

}

func Create(w http.ResponseWriter, r *http.Request) {
	var address models.Address

	if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if err := config.DB.Create(&address).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success create Address", address)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	ParamsId := mux.Vars(r)["id"]

	id, _ := strconv.Atoi(ParamsId)

	var address models.Address

	if err := config.DB.First(&address, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Address Not Found", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Detail Address", address)

}

func Update(w http.ResponseWriter, r *http.Request)  {
	ParamsId := mux.Vars(r)["id"]

	id, _ := strconv.Atoi(ParamsId)

	var address models.Address

	if err := config.DB.First(&address, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Address Not Found", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if err := config.DB.Where("id = ?", id).Updates(&address).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success Update Data", nil)
}

func Delete(w http.ResponseWriter, r *http.Request)  {
	ParamsId := mux.Vars(r)["id"]

	id, _ := strconv.Atoi(ParamsId)

	var address models.Address

	res := config.DB.Delete(&address, id)
	if res.Error != nil {
		helper.Response(w, 500, res.Error.Error(), nil)
		return
	}

	if res.RowsAffected == 0 {
		helper.Response(w, 404, "Address Not Found ", nil)
		return
	}

	helper.Response(w, 200, "Success Delete Address", nil)
}