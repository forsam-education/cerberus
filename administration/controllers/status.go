package controllers

import (
	"encoding/json"
	"net/http"
)

type clusterStatus struct {
	NodeCount    int `json:"node_count"`
	RequestCount int `json:"request_count"`
}

// Status returns a JSON with current Cerberus cluster status.
func Status(w http.ResponseWriter, _ *http.Request) {
	//nodeCount, err := utils.SharedStateManager.GetCurrentNodesCount()
	//if err != nil {
	//	utils.Logger.StdError(err, nil)
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//requestCount, err := utils.SharedStateManager.GetCurrentRequestsCount()
	//if err != nil {
	//	utils.Logger.StdError(err, nil)
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//status := clusterStatus{
	//	NodeCount:    nodeCount,
	//	RequestCount: requestCount,
	//}
	status := clusterStatus{
		NodeCount:    0,
		RequestCount: 0,
	}
	err := json.NewEncoder(w).Encode(status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
