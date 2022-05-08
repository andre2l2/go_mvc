package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/andre2l2/go_mvc/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	getAllProducts := models.GetAllProducs()
	temp.ExecuteTemplate(w, "Index", getAllProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
} 

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		total := r.FormValue("total")

		parsePriceToFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error to converted value: ", err)
		}

		parseTotalToInt, err := strconv.Atoi(total)
		if err != nil {
			log.Println("Error to convert value: ", err)
		}

		models.CreatNewProduct(name, description, parsePriceToFloat, parseTotalToInt)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idPoduct := r.URL.Query().Get("id")
	models.DeleteProduct(idPoduct)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	product := models.GetOneProduct(idProduct)
	temp.ExecuteTemplate(w, "Edit", product)	
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		total := r.FormValue("total")

		parsePriceToFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error to converted value: ", err)
		}

		parseTotalToInt, err := strconv.Atoi(total)
		if err != nil {
			log.Println("Error to convert value: ", err)
		}

		models.UpdateProduct(id, name, description, parsePriceToFloat, parseTotalToInt)
	}
	http.Redirect(w, r, "/", 301)
}