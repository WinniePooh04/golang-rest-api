package models

type FTTxRequestBody struct {
	SWLat               float64 `json:"southwest_lat"`
	SWLng               float64 `json:"southwest_lng"`
	NELat               float64 `json:"northeast_lat"`
	NELng               float64 `json:"northeast_lng"`
	Type                string  `json:"type"`
	Zoom                int     `json:"zoom"`
	LineString          string  `json:"line_string"`
	NodeName            string  `json:"node_name"`
	ShowSearchHierarchy string  `json:"show_search_hierarchy"`
	ShowPolygonCoverage string  `json:"show_polygon_coverage"`
	DeviceStatus        string  `json:"device_status"`
	PortStatus          string  `json:"port_status"`
	BoxType             string  `json:"box_type"`
	Tags                string  `json:"tags"`
	ConnectionTypes     string  `json:"connection_types"`
}
