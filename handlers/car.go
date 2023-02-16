package handlers

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func GetCarList(w http.ResponseWriter, r *http.Request) {
	json := []byte(`{
		"year": 2015,
		"selling_price": [0, 100000],
		"km_driven": 10000,
		"car_type": "city"
	}`)
	postUrl := "https://europe-west1-aller-data-science.cloudfunctions.net/test_finn"
	req, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	w.Write(body)
}
