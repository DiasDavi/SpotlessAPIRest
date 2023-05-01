package user

import (
	"encoding/json"
	"net/http"
	"spotless/database"
	"spotless/helper"
	"spotless/models"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var ResponseJson = helper.ResponseJson
var ResponseError = helper.ResponseError

func Get(w http.ResponseWriter, r *http.Request){
	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil{
		ResponseError(w, http.StatusInternalServerError, err.Error())
	}

	ResponseJson(w, http.StatusOK, users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)["id"]
	id,err := strconv.ParseInt(params, 10, 64)
	if err != nil{
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil{
		switch err{
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusBadRequest, "ID NÃO ENCONTRADO")
			return
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
		}
	}

	ResponseJson(w, http.StatusOK, user)
}

func Post(w http.ResponseWriter, r *http.Request){
	var user models.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil{
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if err := database.DB.Create(&user).Error; err != nil{
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJson(w, http.StatusOK, user)
}

func Update(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)["id"]
	id,err := strconv.ParseInt(params, 10, 64)
	if err != nil{
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil{
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if database.DB.Where("id = ?", id).Updates(&user).RowsAffected == 0{		
		ResponseError(w, http.StatusBadRequest, "Erro ao atualizar registro")
		return
	}

	user.Id = id

	ResponseJson(w, http.StatusOK, user)
}

func Delete(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)["id"]
	id,err := strconv.ParseInt(params, 10, 64)
	if err != nil{
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User

	if database.DB.Delete(&user, id).RowsAffected == 0{
		ResponseError(w, http.StatusBadRequest, "ID NÃO ENCONTRADO")
		return
	}

}