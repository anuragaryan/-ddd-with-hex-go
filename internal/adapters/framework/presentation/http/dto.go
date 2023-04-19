package http

type createListRequest struct {
	Name string `json:"name"`
}

type list struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Items []item `json:"items"`
}

type getListResponse struct {
	Lists list `json:"list"`
}

type getListsResponse struct {
	Lists []list `json:"lists"`
}

type createItemRequest struct {
	Item string `json:"item"`
}

type item struct {
	ID   string `json:"id"`
	Item string `json:"item"`
}
