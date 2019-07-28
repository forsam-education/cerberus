package controllers

import (
	"encoding/json"
	"github.com/forsam-education/cerberus/state"
	"github.com/forsam-education/cerberus/utils"
	"net/http"
)

type clusterStatus struct {
	NodeCount    int `json:"node_count"`
	RequestCount int `json:"request_count"`
}

<<<<<<< HEAD
// Status returns a JSON with current Cerberus cluster status.
func Status(w http.ResponseWriter, _ *http.Request) {
	nodeCount, err := utils.SharedStateManager.GetCurrentNodesCount()
	if err != nil {
		utils.Logger.StdError(err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	requestCount, err := utils.SharedStateManager.GetCurrentRequestsCount()
	if err != nil {
		utils.Logger.StdError(err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	status := clusterStatus{
		NodeCount:    nodeCount,
		RequestCount: requestCount,
	}
	err = json.NewEncoder(w).Encode(status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
