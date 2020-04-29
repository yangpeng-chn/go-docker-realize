package main_test

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"testing"
	"time"
)

const (
	baseURL = "http://localhost:8888"
)

type Transaction struct {
	ID          int    `json:"id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}

func Test(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	transactions := []Transaction{
		{
			ID:          1,
			Amount:      rand.Intn(100000),
			Description: "事務用品",
		},
		{
			ID:          2,
			Amount:      rand.Intn(200000),
			Description: "机",
		},
	}
	log.Println("amount1", transactions[0].Amount)
	for _, transaction := range transactions {
		data, err := json.Marshal(transaction)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest(
			http.MethodPost,
			baseURL+"/transactions",
			bytes.NewBuffer(data),
		)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("apikey", "secure-api-key-1")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("POST /transactions status %d != 201", resp.StatusCode)
		}
		if err := resp.Body.Close(); err != nil {
			t.Fatal(err)
		}
	}

	t.Run("list", func(t *testing.T) {
		req, err := http.NewRequest(
			http.MethodGet,
			baseURL+"/transactions",
			nil,
		)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("apikey", "secure-api-key-1")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
			return
		}
		defer resp.Body.Close()

		var got []Transaction
		if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
			t.Fatal(err)
			return
		}

		for i := range got {
			if !reflect.DeepEqual(got[i], transactions[i]) {
				t.Errorf("transactions[%d] %+v != %+v", i, got[i], transactions[i])
			}
		}
	})
}
