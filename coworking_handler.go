package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CoworkingSpace struct {
	Name     string `json:"name"`
	Speed    string `json:"speed"`
	Charging string `json:"charging"`
}

var coworkingSpaces []CoworkingSpace

func getCoworkingSpacesHandler(w http.ResponseWriter, r *http.Request) {
	coworkingSpacesBytes, err := json.Marshal(coworkingSpaces)

	if err != nil {
		fmt.Println(fmt.Errorf("Error %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(coworkingSpacesBytes)
}

func createCoworkingSpaceHandler(w http.ResponseWriter, r *http.Request) {
	coworkingSpace := CoworkingSpace{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	coworkingSpace.Name = r.Form.Get("name")
	coworkingSpace.Speed = r.Form.Get("speed")
	coworkingSpace.Charging = r.Form.Get("charging")

	coworkingSpaces = append(coworkingSpaces, coworkingSpace)

	http.Redirect(w, r, "/assets", http.StatusFound)
}
