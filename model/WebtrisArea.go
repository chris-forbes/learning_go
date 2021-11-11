package model

type WebtrisAreas struct {
	RowCount int `json:"row_count"`
	Areas    []struct {
		ID          string `json:"Id"`
		Name        string `json:"Name"`
		Description string `json:"Description"`
		XLongitude  string `json:"XLongitude"`
		XLatitude   string `json:"XLatitude"`
		YLongitude  string `json:"YLongitude"`
		YLatitude   string `json:"YLatitude"`
	} `json:"areas"`
}


