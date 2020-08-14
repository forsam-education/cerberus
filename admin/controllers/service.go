package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/forsam-education/cerberus/models"
	"github.com/forsam-education/cerberus/proxy"
	"github.com/forsam-education/cerberus/state"
	"github.com/forsam-education/cerberus/utils"
	"github.com/valyala/fasthttp"
	"github.com/volatiletech/sqlboiler/boil"
	"net/http"
)

// ListServices returns a JSON list of services.
func ListServices(context *fasthttp.RequestCtx) {
	services, err := models.Services().AllG()
	if err != nil {
		context.Response.SetStatusCode(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(context.Response.BodyWriter()).Encode(services)
	if err != nil {
		context.Response.SetStatusCode(http.StatusInternalServerError)
		return
	}
}

// CreateService creates a new service from JSON request.
func CreateService(context *fasthttp.RequestCtx) {
	var service models.Service
	if err := json.Unmarshal(context.Request.Body(), &service); err != nil {
		context.Response.SetStatusCode(http.StatusBadRequest)
		return
	}

	message := fmt.Sprintf("There was an error Bro ! : %s", context.Request.Body())

	utils.Logger.Info(message, nil)
	if err := service.InsertG(boil.Infer()); err != nil {




		context.Response.SetStatusCode(http.StatusInternalServerError)
		return
	}
	if err := proxy.LoadServices(); err != nil {
		context.Response.SetStatusCode(http.StatusInternalServerError)
		return
	}

	context.Response.SetStatusCode(http.StatusCreated)

	if err := json.NewEncoder(context.Response.BodyWriter()).Encode(utils.Response{Status: "OK"}); err != nil {
		context.Response.SetStatusCode(http.StatusInternalServerError)
		return
	}
}

// UpdateService updates a service from JSON request, and remove it from redis to update the cache
func UpdateService(context *fasthttp.RequestCtx) {
	var service models.Service
	if err := json.Unmarshal(context.Request.Body(), &service); err != nil {
		context.Response.SetStatusCode(http.StatusBadRequest)
		return
	}
	if _, err := service.UpdateG(boil.Infer()); err != nil {
		context.Response.SetStatusCode(http.StatusInternalServerError)
		return
	}

	if err := state.Manager.RemoveService(&service); err != nil {
		context.Response.SetStatusCode(http.StatusInternalServerError)
		return
	}

	context.Response.SetStatusCode(http.StatusNoContent)
}
