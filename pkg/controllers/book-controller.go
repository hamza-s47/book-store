package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hamza-s47/book-store/pkg/models"
	"github.com/hamza-s47/book-store/pkg/utils"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, err := json.Marshal(b)
	if err != nil {
		log.Panicf("Error while marsalling %v", err)

		http.Error(w, "Failed to Marshal book data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, err := json.Marshal(newBooks)
	if err != nil {
		log.Printf("Can not marshalling Book %v", err)

		http.Error(w, "Failed to marshal books data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		log.Panicf("Error while parsing ID %v", err)

		http.Error(w, "Failed to parse ID", http.StatusInternalServerError)
		return
	}

	bookDetails, db := models.GetBookById(ID)
	if db.Error != nil {
		log.Printf("Error fetching book: %v", db.Error)
		http.Error(w, "Error fetching book", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(bookDetails)
	if err != nil {
		log.Printf("Can not marshalling Book %v", err)

		http.Error(w, "Failed to marshal books data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		log.Panicf("Error while parsing ID %v", err)

		http.Error(w, "Failed to parse ID", http.StatusInternalServerError)
		return
	}

	updatedBook := &models.Book{}
	utils.ParseBody(r, updatedBook)

	book, db := models.UpdateBook(ID, *updatedBook)
	if db.Error != nil {
		log.Printf("Error while updating book: %v", db.Error)
		http.Error(w, "Failed to update book", http.StatusInternalServerError)
		return
	}

	if book == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		log.Panicf("Error while parsing ID %v", err)

		http.Error(w, "Failed to parse ID", http.StatusInternalServerError)
		return
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
