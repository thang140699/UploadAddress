package api

import (
	"WeddingUtilities/database/repository"
	"WeddingUtilities/model"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type AddressHandler struct {
	AddressRepository repository.AddressRepository
}

func getByPage(start string, end string, limit string, page string, arr []model.Address) ([]model.Address, error) {
	st, errST := strconv.Atoi(start)
	ed, errED := strconv.Atoi(end)
	lm, errLM := strconv.Atoi(limit)
	pg, errPG := strconv.Atoi(page)
	if errST != nil {
		st = 0
	}
	if errED != nil {
		ed = len(arr)
	}
	if errLM != nil {
		lm = len(arr)
	}
	if errPG != nil {
		pg = 1
	}
	divideResult := len(arr) / lm
	surplus := len(arr) % lm
	if surplus == 0 && pg > divideResult || surplus != 0 && pg > divideResult+1 {
		return nil, errors.New("don't have record")
	}
	if surplus != 0 && pg == divideResult+1 {
		return arr[lm*(pg-1) : ed], nil
	}
	return arr[st:ed], nil

}

func (h *AddressHandler) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	query := r.URL.Query()
	start := query.Get("start")
	end := query.Get("end")
	limit := query.Get("limit")
	page := query.Get("page")
	Address, err := h.AddressRepository.All()
	if err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_READ_FAILED, ResponseBody{
			Message: "unable to get User from server",
		})
	}
	results, err := getByPage(start, end, limit, page, Address)
	if err != nil {
		WriteJSON(w, http.StatusNotFound, ResponseBody{
			Message: err.Error(),
			Code:    http.StatusNotFound,
		})
		return
	}
	WriteJSON(w, http.StatusOK, results)

}

func (h *AddressHandler) GetByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("ID")
	Address, err := h.AddressRepository.FindByID(id)
	if err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_READ_FAILED, ResponseBody{
			Message: "unable to get Address by id : " + id,
		})
	}
	WriteJSON(w, http.StatusOK, Address)
}
func (h *AddressHandler) GetByCodeName(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	CodeName := p.ByName("CodeName")
	Address, err := h.AddressRepository.FindByCodeName(CodeName)
	if err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_READ_FAILED, ResponseBody{
			Message: "unable to get Address by codename : " + CodeName,
		})
	}
	WriteJSON(w, http.StatusOK, Address)
}

func (h *AddressHandler) GetByName(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("Name")
	Address, err := h.AddressRepository.FindByName(name)
	if err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_READ_FAILED, ResponseBody{
			Message: "unable to get Address by name : " + name,
		})
	}
	WriteJSON(w, http.StatusOK, Address)
}

func (h *AddressHandler) GetByDivisiontype(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	divisionType := p.ByName("Divisiontype")
	Address, err := h.AddressRepository.FindByDivisionType(divisionType)
	if err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_READ_FAILED, ResponseBody{
			Message: "unable to get Address by divisionType : " + divisionType,
		})
	}
	WriteJSON(w, http.StatusOK, Address)
}

func (h *AddressHandler) GetByPhoneCode(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	phoneCode := p.ByName("PhoneCode")
	Address, err := h.AddressRepository.FindByPhoneCode(phoneCode)
	if err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_READ_FAILED, ResponseBody{
			Message: "unable to get phoneCode by phoneCode : " + phoneCode,
		})
	}
	WriteJSON(w, http.StatusOK, Address)
}

// convert string to int
func (h *AddressHandler) GetByLevel(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	level := p.ByName("level")
	intLevel, err := strconv.Atoi(level)

	if err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_READ_FAILED, ResponseBody{
			Message: "usdf : " + level,
		})
		return
	}
	Address, err := h.AddressRepository.FindByLevel(intLevel)
	if err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_READ_FAILED, ResponseBody{
			Message: "unable to get Level by level : " + level,
		})
		return
	}
	WriteJSON(w, http.StatusOK, Address)
}
func (h *AddressHandler) GetByParentId(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	parentId := p.ByName("ParentId")
	Address, err := h.AddressRepository.FindByParentId(parentId)
	if err != nil {
		WriteJSON(w, HTTP_ERROR_CODE_READ_FAILED, ResponseBody{
			Message: "unable to get Parent by parentId : " + parentId,
		})
	}
	WriteJSON(w, http.StatusOK, Address)
}
