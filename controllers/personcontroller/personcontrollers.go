package personcontroller

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

func Index(w http.ResponseWriter, r *http.Request) {
	var person []models.Person

	if err := config.DB.Find(&person).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
	}

	helper.Response(w, 200, "List Person", person)

}

func Create(w http.ResponseWriter, r *http.Request) {
	var person models.Person

	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if err := config.DB.Create(&person).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success create Person", person)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	ParamsId := mux.Vars(r)["id"]

	id, _ := strconv.Atoi(ParamsId)

	var person models.Person

	if err := config.DB.First(&person, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Person Not Found", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Detail Person", person)

}

func Update(w http.ResponseWriter, r *http.Request)  {
	ParamsId := mux.Vars(r)["id"]

	id, _ := strconv.Atoi(ParamsId)

	var person models.Person

	if err := config.DB.First(&person, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Person Not Found", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if err := config.DB.Where("id = ?", id).Updates(&person).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success Update Data", nil)
}

func Delete(w http.ResponseWriter, r *http.Request)  {
	ParamsId := mux.Vars(r)["id"]

	id, _ := strconv.Atoi(ParamsId)

	var person models.Person

	res := config.DB.Delete(&person, id)
	if res.Error != nil {
		helper.Response(w, 500, res.Error.Error(), nil)
		return
	}

	if res.RowsAffected == 0 {
		helper.Response(w, 404, "Person Not Found ", nil)
		return
	}

	helper.Response(w, 200, "Success Delete Person", nil)
}