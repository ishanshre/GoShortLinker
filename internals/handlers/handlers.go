package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/ishanshre/GoShortLinker/internals/database"
	"github.com/ishanshre/GoShortLinker/internals/helpers"
	"github.com/ishanshre/GoShortLinker/internals/repository"
	"github.com/ishanshre/GoShortLinker/internals/repository/db"
	"github.com/ishanshre/GoShortLinker/internals/types"
)

type Handlers interface {
	ShortenURL(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	DB repository.DbRepository
}

var validate *validator.Validate

func NewHandler(d database.Database) Handlers {
	validate = validator.New()
	return &handler{
		DB: db.NewMongoDbRepo(d, context.Background()),
	}
}

func (h *handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	url := types.Request{}
	if err := json.NewDecoder(r.Body).Decode(&url); err != nil {
		helpers.StatusBadRequest(w, err.Error())
		return
	}
	if err := validate.Struct(url); err != nil {
		helpers.StatusBadRequest(w, err.Error())
		return
	}

	if !helpers.RemoveDomainError(url.LongUrl) {
		helpers.StatusBadRequest(w, "plase enter valid url")
	}
	url.LongUrl = helpers.EnforceHttp(url.LongUrl)
	id := uuid.New().String()[:6]
	shortURL := fmt.Sprintf("%s/%s", os.Getenv("DOMAIN"), id)
	shortURLConfig := types.Url{
		UrlCode:   id,
		ShortUrl:  shortURL,
		LongUrl:   url.LongUrl,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(time.Hour * 24),
	}
	res, err := h.DB.InsertUrl(&shortURLConfig)
	if err != nil {
		helpers.StatusBadRequest(w, err.Error())
		return
	}
	helpers.StatusCreated(w, res)
}
