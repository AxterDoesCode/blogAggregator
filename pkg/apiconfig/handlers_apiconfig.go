package apiconfig

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/AxterDoesCode/blogAggregator/internal/database"
	httphandler "github.com/AxterDoesCode/blogAggregator/pkg/httpHandler"
)

func (cfg *ApiConfig) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	type requestBody struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := requestBody{}
	err := decoder.Decode(&params)
	if err != nil {
		httphandler.RespondWithError(
			w,
			http.StatusInternalServerError,
			"Couldn't decode parameters",
		)
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		httphandler.RespondWithError(w, http.StatusInternalServerError, "Error creating user")
		return
	}

	httphandler.RespondWithJSON(w, http.StatusOK, user)
}