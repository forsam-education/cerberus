package proxy

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/forsam-education/cerberus/errors"
	"github.com/forsam-education/cerberus/models"
	"github.com/forsam-education/cerberus/state"
	"github.com/forsam-education/cerberus/utils"
	"github.com/valyala/fasthttp"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"net"
	"strings"
	"time"
)

var proxyClient = &fasthttp.Client{}

func findServiceByPath(path []byte) *models.Service {
	if service, ok := services[string(path)]; ok {
		if service.ExpirationTime.After(time.Now()) {
			return service.Service
		}
		delete(services, string(path))
	}

	return findServiceInRedis(path)
}

func findServiceInRedis(path []byte) *models.Service {
	redisService, err := state.Manager.FindServiceByPath(path)

	if err != nil {
		utils.Logger.StdErrorCritical(err, nil)
		return nil
	}

	if redisService == nil {
		return findServiceInDb(path)
	}

	cacheService(redisService)
	return redisService
}

func findServiceInDb(path []byte) *models.Service {
	dbService, err := models.Services(qm.Where("service_path = ?", string(path))).OneG()
	if err != nil {
		if err != sql.ErrNoRows {
			utils.Logger.StdErrorCritical(err, nil)
		}
		return nil
	}
	err = state.Manager.AddService(dbService)
	if err != nil {
		utils.Logger.StdErrorCritical(err, nil)
		return nil
	}
	cacheService(dbService)

	return dbService
}

func buildRequestHost(service *models.Service) string {
	if !service.TargetPort.Valid {
		return service.TargetHost
	}

	return fmt.Sprintf("%s:%d", service.TargetHost, service.TargetPort.Uint)
}

func tranformRequestURI(path []byte, service *models.Service) []byte {
	var newPath []byte
	if service.TargetPath.Valid {
		newPath = []byte(service.TargetPath.String)
	} else {
		newPath = []byte("/")
	}
	servicePathLength := len(getFirstPathPart(path))

	newPath = append(newPath, path[servicePathLength:]...)

	return newPath
}

func getFirstPathPart(path []byte) []byte {
	tmpURI := []byte("/")
	return append(tmpURI, bytes.Split(path[1:], []byte("/"))[0]...)
}

func proxify(ctx *fasthttp.RequestCtx) {
	req := &ctx.Request
	res := &ctx.Response

	path := ctx.Path()
	method := ctx.Method()

	service := findServiceByPath(getFirstPathPart(path))

	// No service for this path.
	if service == nil {
		handleServiceNotFound(ctx)
		return
	}

	// Method not allowed for this service.
	if !strings.Contains(service.Methods, string(method)) {
		handleMethodNotAllowed(ctx)
		return
	}

	// Replace service path with target path in request.
	req.URI().SetPathBytes(tranformRequestURI(path, service))

	if clientIP, _, err := net.SplitHostPort(ctx.RemoteAddr().String()); err == nil {
		req.Header.Add(utils.HeaderXForwardedFor, clientIP)
	}

	// Set host to service host and port.
	req.SetHost(buildRequestHost(service))

	// Send request the the service and set response to the service response.
	if err := proxyClient.Do(req, res); err != nil {
		return
	}
}

func handleServiceNotFound(ctx *fasthttp.RequestCtx) {
	ctx.Response.Reset()
	ctx.Response.SetStatusCode(fasthttp.StatusNotFound)
	ctx.Response.Header.Set(utils.HeaderContentType, "application/json; charset=utf-8")
	response := errors.BuildErrorResponseBody(errors.ServiceNotFound, string(getFirstPathPart(ctx.Path())), nil)
	encodedResponse, _ := json.Marshal(response)
	ctx.Response.SetBody(encodedResponse)
}

func handleMethodNotAllowed(ctx *fasthttp.RequestCtx) {
	ctx.Response.Reset()
	ctx.Response.SetStatusCode(fasthttp.StatusMethodNotAllowed)
	ctx.Response.Header.Set(utils.HeaderContentType, "application/json; charset=utf-8")
	response := errors.BuildErrorResponseBody(errors.MethodNotAllowed, string(getFirstPathPart(ctx.Path())), utils.ResponseExtraData{"method": string(ctx.Method())})
	encodedResponse, _ := json.Marshal(response)
	ctx.Response.SetBody(encodedResponse)
}
