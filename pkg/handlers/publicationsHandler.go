package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/noskine/pilot_api/internal/usecases"
	"github.com/noskine/pilot_api/pkg/dto"
)

func PublicationsHandler(w http.ResponseWriter, r *http.Request) {
	InputRequest := new(dto.DTORequestPepplos)

	if err := json.NewDecoder(r.Body).Decode(InputRequest); err != nil {
		http.Error(w, fmt.Sprintf("Internal Server Error: %s", err), http.StatusInternalServerError)
	}

	valid := validator.New()

	if err := valid.Struct(InputRequest); err != nil {
		erros := err.(validator.ValidationErrors)
		http.Error(w, fmt.Sprintf("Validator error: %s", erros), http.StatusBadRequest)
		return
	}

	response, err := usecases.NewPublicationsUseCase().CreatePublicationsUseCase(InputRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("Validator error: %s", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Validator error: %s", err), http.StatusBadRequest)
		return
	}
}
