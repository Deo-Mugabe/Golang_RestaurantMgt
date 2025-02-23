package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Deo-Mugabe/Golang_RestaurantMgt/models"
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/services"
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/utility"
)

func GetTablesHandler(w http.ResponseWriter, r *http.Request) {
	tables, err := services.GetTables()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utility.JsonResponse(w, http.StatusOK, tables)
}

func GetTableHandler(w http.ResponseWriter, r *http.Request) {

	id, err := utility.ParseID(r)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}
	table, err := services.GetTable(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utility.JsonResponse(w, http.StatusOK, table)
}

func CreateTableHandler(w http.ResponseWriter, r *http.Request) {
	var table models.Table
	if err := json.NewDecoder(r.Body).Decode(&table); err != nil {
		http.Error(w, "Invalid Inputs", http.StatusBadRequest)
		return
	}
	err := services.CreateTable(&table)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utility.JsonResponse(w, http.StatusCreated, table)
}

func UpdateTableHandler(w http.ResponseWriter, r *http.Request) {

	id, err := utility.ParseID(r)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}
	var updatedTable models.Table
	if err := json.NewDecoder(r.Body).Decode(&updatedTable); err != nil {
		http.Error(w, "Invalid Inputs", http.StatusBadRequest)
		return
	}
	err = services.UpdateTable(uint(id), &updatedTable)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utility.JsonResponse(w, http.StatusAccepted, updatedTable)
}

func DeleteTableHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utility.ParseID(r)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}
	err = services.DeleteTable(uint(id))
	if err != nil {
		if err.Error() == "user not found" {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(`{"Table deleted successfully"}`))
}
