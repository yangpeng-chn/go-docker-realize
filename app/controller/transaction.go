package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yangpeng-chn/go-docker-realize/app/model"
)

// TransactionController controller
type TransactionController struct {
}

func (t *TransactionController) List(w http.ResponseWriter, r *http.Request) {
	SetResponseHeaders(w, r)
	var err error
	var transactions []model.Transaction

	if transactions, err = model.ListTransactions(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if err = WriteResponse(w, http.StatusOK, transactions); err != nil {
		log.Println(err)
		return
	}
	log.Println("List transactions sucessfully")
}

func (t *TransactionController) Show(w http.ResponseWriter, r *http.Request) {
	SetResponseHeaders(w, r)
	var err error
	var transaction model.Transaction

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	if transaction, err = model.ShowTransaction(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		log.Println(err)
		return
	}

	if err = WriteResponse(w, http.StatusOK, transaction); err != nil {
		log.Println(err)
		return
	}
	log.Println("Show transaction sucessfully")
}

func (t *TransactionController) Create(w http.ResponseWriter, r *http.Request) {
	var err error
	var transaction model.Transaction
	// body, err := ioutil.ReadAll(io.LimitReader(r.Body, 10485760)) //10MB
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	if err = json.Unmarshal(body, &transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	if err := model.CreateTransaction(transaction); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write([]byte(`"msg":"created sucessfully"`)); err != nil {
		log.Println(err)
		return
	}

	log.Println("Create transaction sucessfully")
}

func (t *TransactionController) Delete(w http.ResponseWriter, r *http.Request) {
	SetResponseHeaders(w, r)
	var err error
	if err = model.DeleteTransactions(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write([]byte(`"msg":"deleted sucessfully"`)); err != nil {
		log.Println(err)
		return
	}

	log.Println("Deleted transactions sucessfully")
}
