package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"orders_by/models"
	"strconv"
)

var db *gorm.DB

func Connection() {
	var err error
	dataSourceName := "root:@tcp(localhost:3306)/?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	db.Exec("CREATE DATABASE orders_by")
	db.Exec("USE orders_by")

	// Migrasi untuk membuat tabel untuk skema Pesanan dan Item
	db.AutoMigrate(&models.Order{}, models.Item{})
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	json.NewDecoder(r.Body).Decode(&order)
	// Membuat pesanan baru dengan menyisipkan catatan di tabel `pesanan` dan `item`
	db.Create(&order)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func GetAllOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var orders []models.Order
	db.Preload("Items").Find(&orders)
	json.NewEncoder(w).Encode(orders)
}

func GetOrderById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]

	var order models.Order
	db.Preload("Items").First(&order, inputOrderID)
	json.NewEncoder(w).Encode(order)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	var updatedOrder models.Order
	json.NewDecoder(r.Body).Decode(&updatedOrder)
	db.Save(&updatedOrder)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedOrder)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	// Ubah parameter string `orderId` ke uint64
	id64, _ := strconv.ParseUint(inputOrderID, 10, 64)
	// Konversikan uint64 ke uint
	idToDelete := uint(id64)

	db.Where("order_id = ?", idToDelete).Delete(&models.Item{})
	db.Where("order_id = ?", idToDelete).Delete(&models.Order{})
	w.WriteHeader(http.StatusNoContent)
}
