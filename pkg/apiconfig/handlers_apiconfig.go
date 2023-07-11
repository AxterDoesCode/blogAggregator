package apiconfig

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/AxterDoesCode/blogAggregator/internal/database"
	httphandler "github.com/AxterDoesCode/blogAggregator/pkg/httpHandler"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

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

func (cfg *ApiConfig) HandleGetUserByApiKey(
	w http.ResponseWriter,
	r *http.Request,
	user database.User,
) {
	httphandler.RespondWithJSON(w, http.StatusOK, user)
}

func (cfg *ApiConfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("Authorization")
		if apiKey == "" {
			httphandler.RespondWithError(w, http.StatusInternalServerError, "Api Key doesn't exist")
			return
		}
		user, err := cfg.DB.GetUser(r.Context(), apiKey)
		if err != nil {
			httphandler.RespondWithError(
				w,
				http.StatusInternalServerError,
				"Error fetching user from database",
			)
			return
		}
		handler(w, r, user)
	}
}

func (cfg *ApiConfig) HandleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type requestParams struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := requestParams{}
	err := decoder.Decode(&params)
	if err != nil {
		httphandler.RespondWithError(
			w,
			http.StatusInternalServerError,
			"Error decoding request body",
		)
		return
	}
	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		httphandler.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf(("%v"), err))
		return
	}
	httphandler.RespondWithJSON(w, http.StatusOK, feed)
}

func (cfg *ApiConfig) HandleGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		httphandler.RespondWithError(
			w,
			http.StatusInternalServerError,
			"Error getting feeds from database",
		)
		return
	}
	httphandler.RespondWithJSON(w, http.StatusOK, feeds)
}
