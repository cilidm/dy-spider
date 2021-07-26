package request


type LayerListForm struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

type SpiderAddForm struct {
	Ulr    string `json:"ulr" form:"url"`
	ByYear int    `json:"by_year" form:"by_year"`
	Down   int    `json:"down" form:"down"`
}

type SpiderListForm struct {
	LayerListForm
	UserName string `json:"user_name" form:"user_name"`
	Info     string `json:"info" form:"info"`
	Url      string `json:"url" form:"url"`
}
