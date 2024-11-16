package models

type DataCount struct {
	Total       int64 `json:"total,omitempty"`
	Delivered   int64 `json:"delivered,omitempty"`
	Canceled    int64 `json:"canceled,omitempty"`
	Returned    int64 `json:"returned,omitempty"`
	InWarehouse int64 `json:"inWarehouse,omitempty"`
	InVehicle   int64 `json:"inVehicle,omitempty"`
	Failed      int64 `json:"failed,omitempty"`
}
