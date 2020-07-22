/*
 * Product aggregation composite interface
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func DeleteSubscription(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(http.StatusOK)
}

func readOffers(path string) ([]CategoryOffers, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	categories := []CategoryOffers{}
	for _, f := range files {
		fs, err := os.Open(fmt.Sprintf("./examples/json/offers/%s", f.Name()))
		if err != nil {
			return nil, err
		}
		var filecategories []CategoryOffers
		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.

		err = json.NewDecoder(fs).Decode(&filecategories)
		//var co CategoryOffersOffers
		//err1 := json.NewDecoder(fs).Decode(&co)
		if err != nil {
			return nil, err
		}
		categories = append(categories, filecategories...)
	}
	return categories, nil
}
func GetOffers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	catName := r.URL.Query().Get("category")
	msisdn := r.URL.Query().Get("msisdn")
	categories, err := readOffers("./examples/json/offers/")
	if err != nil {
		panic(err)
	}
	if msisdn != "" {
		msisdnCategories, err := readOffers(fmt.Sprintf("./examples/json/offers/%s/", msisdn))
		if err == nil {
			categories = append(categories, msisdnCategories...)
		}
	}
	if catName != "" {
		fiiteredCategories := categories[:0]

		for _, x := range categories {
			if x.Categoryname == catName {
				fiiteredCategories = append(fiiteredCategories, x)
			}
		}
		categories = fiiteredCategories
	}
	response := OffersListResponse{Status: "OK", Category: categories}

	err2 := json.NewEncoder(w).Encode(&response)
	if err2 != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}

func GetSubscriptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func PostSubscription(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
