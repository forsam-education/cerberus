package controllers

import (
	"encoding/json"
	"github.com/forsam-education/cerberus/models"
	"github.com/forsam-education/cerberus/proxy"
	"github.com/forsam-education/cerberus/state"
	"github.com/forsam-education/cerberus/utils"
	"github.com/volatiletech/sqlboiler/boil"
	"net/http"
)

// ListServices returns a JSON list of services.
func ListServices(w http.ResponseWriter, _ *http.Request) {
	services, err := models.Services().AllG()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(services)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// CreateService creates a new service from JSON request.
func CreateService(w http.ResponseWriter, r *http.Request) {
	var service models.Service
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&service); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := service.InsertG(boil.Infer()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := proxy.LoadServices(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//proxy.Swapper.Swap(proxy.LoadRouter())

	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(utils.Response{Status: "OK"}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// UpdateService updates a service from JSON request, and remove it from redis to update the cache
func UpdateService(w http.ResponseWriter, r *http.Request) {
	var service models.Service
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&service); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if _, err := service.UpdateG(boil.Infer()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := state.Manager.RemoveService(&service); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
