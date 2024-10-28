package pkkRosreestr

type PkkRosreestrLotDto struct {
	CadastreNumber string `json:"cadastreNumber"`
	Total          int    `json:"total"`
	TotalRelation  string `json:"total_relation"`
	Features       []struct {
		Center struct {
			Y float64 `json:"y"`
			X float64 `json:"x"`
		} `json:"center"`
		Extent struct {
			Xmin float64 `json:"xmin"`
			Xmax float64 `json:"xmax"`
			Ymax float64 `json:"ymax"`
			Ymin float64 `json:"ymin"`
		} `json:"extent"`
		Sort  int64 `json:"sort"`
		Type  int   `json:"type"`
		Attrs struct {
			Address      string `json:"address"`
			CategoryType string `json:"category_type"`
			Cn           string `json:"cn"`
			Id           string `json:"id"`
		} `json:"attrs"`
	} `json:"features"`
}
