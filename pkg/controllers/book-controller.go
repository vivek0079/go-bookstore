package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-bookstore/pkg/configs"
	"go-bookstore/pkg/models"
	"go-bookstore/pkg/utils"
	"net/http"
	"strconv"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome"))
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	fmt.Println(books)
	res, _ := json.Marshal(books)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idString := params["id"]

	id, err := strconv.ParseInt(idString, 0, 0);
	if err != nil {
		fmt.Println("invalid id. not a valid integer")
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid id. not a valid integer"))
		return
	}
	book := models.GetBookById(id)
	res, _ := json.Marshal(book)
	fmt.Println(res)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	b := book.CreateBook()
	res, _ := json.Marshal(b)
	fmt.Println(res)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idString := params["id"]

	id, err := strconv.ParseInt(idString, 0, 0)
	if err != nil {
		fmt.Println("invalid id. not a valid integer")
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid id. not a valid integer"))
		return
	}
	book := models.DeleteBook(id)
	res, _ := json.Marshal(book)
	fmt.Println(res)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)

	params := mux.Vars(r)
	idString := params["id"]

	id, err := strconv.ParseInt(idString, 0, 0)
	if err != nil {
		fmt.Println("invalid id. not a valid integer")
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid id. not a valid integer"))
		return
	}

	db := configs.GetDb()
	db.Model(&book).Where("id = ?", id).Updates(&book)

	res, _ := json.Marshal(book)
	fmt.Println(res)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}