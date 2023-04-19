package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/anuragaryan/ddd-with-hex-go/internal/ports"
)

type Configuration func(h *Handler) error

type Handler struct {
	apiPort ports.APIPort
}

func NewHandler(cfgs ...Configuration) (*Handler, error) {
	h := &Handler{}

	for _, cfg := range cfgs {
		err := cfg(h)
		if err != nil {
			return nil, err
		}
	}
	return h, nil
}

func WithService(ap ports.APIPort) Configuration {
	return func(h *Handler) error {
		h.apiPort = ap
		return nil
	}
}

func (h *Handler) CreateList(w http.ResponseWriter, r *http.Request) {
	var payload createListRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		responseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.apiPort.CreateList(payload.Name); err != nil {
		responseJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseJSON(w, http.StatusOK, nil)
}

func (h *Handler) GetList(w http.ResponseWriter, r *http.Request) {
	listID := chi.URLParam(r, "listID")

	l, err := h.apiPort.GetList(listID)
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	var responseListItems []item

	listItems, err := l.ListItems()
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	for _, listItem := range listItems {
		responseListItems = append(responseListItems, item{
			ID:   listItem.ID,
			Item: listItem.Text,
		})
	}

	responseList := list{
		ID:    l.ID,
		Name:  l.Name,
		Items: responseListItems,
	}

	responseJSON(w, http.StatusOK, getListResponse{
		Lists: responseList,
	})
}

func (h *Handler) GetLists(w http.ResponseWriter, _ *http.Request) {
	ll, err := h.apiPort.GetLists()
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	var responseList []list
	for _, l := range ll {
		responseList = append(responseList, list{
			ID:   l.ID,
			Name: l.Name,
		})
	}

	responseJSON(w, http.StatusOK, getListsResponse{
		Lists: responseList,
	})

}

func (h *Handler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var payload createItemRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		responseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	listID := chi.URLParam(r, "listID")

	if err := h.apiPort.AddItemToList(listID, payload.Item); err != nil {
		responseJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseJSON(w, http.StatusOK, nil)
}
