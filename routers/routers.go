package routers

import (
	"github.com/gorilla/mux"
	"orders_by/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	// Membuat Data RestApi
	router.HandleFunc("/orders", controllers.CreateOrder).Methods("POST")
	// Membaca Semua Data RestApi
	router.HandleFunc("/orders", controllers.GetAllOrder).Methods("GET")
	// Membaca Data RestApi Berdasarkan ID
	router.HandleFunc("/orders/{orderId}", controllers.GetOrderById).Methods("GET")
	// Mengupdate Data RestApi Berdasarkan ID
	router.HandleFunc("/orders/{orderId}", controllers.UpdateOrder).Methods("PUT")
	// Menghapus Data RestApi Berdasarkan ID
	router.HandleFunc("/orders/{orderId}", controllers.DeleteOrder).Methods("DELETE")

	controllers.Connection()

	return router
}
