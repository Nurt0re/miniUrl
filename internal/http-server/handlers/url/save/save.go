package save

import (
	"errors"
	"log/slog"
	resp "miniUrl/internal/lib/api/response"
	"miniUrl/internal/lib/logger/sl"
	"miniUrl/internal/lib/random"
	"miniUrl/internal/storage"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type URLSaver interface {
	SaveURL(urlToSave string, alias string) (int64, error)
}

type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	resp.Response
	Alias string `json:"alias,omitempty"`
}

// move to config
const aliasLength = 6

func New(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.save.New"

		log = log.With(slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))
			render.JSON(w, r, resp.Error("failed to decode request"))
			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			validatErr := err.(validator.ValidationErrors)
			log.Error("invalid request", sl.Err(err))

			render.JSON(w, r, resp.ValidationError(validatErr))
			return
		}

		alias := req.Alias
		var id int64
		if alias == "" {
			for {
				alias = random.NewRandomString(aliasLength)
				id, err = urlSaver.SaveURL(req.URL, alias)
				if errors.Is(err, storage.ErrURLExists) {
					continue // Generate a new alias if it already exists
				} else if err != nil {
					log.Error("failed to save URL", sl.Err(err))
					render.JSON(w, r, resp.Error("failed to save URL"))
					return
				}
				break // Exit the loop if the alias is unique
			}
		} else {

			id, err = urlSaver.SaveURL(req.URL, alias)
			if errors.Is(err, storage.ErrURLExists) {
				log.Info("url already exists", slog.String("url", req.URL))
				render.JSON(w, r, resp.Error("url already exists"))
				return
			}
			if err != nil {
				log.Error("failed to save URL", sl.Err(err))
				render.JSON(w, r, resp.Error("failed to save URL"))
				return
			}
		}
		log.Info("url saved", slog.Int64("id", id), slog.String("alias", alias))
		render.JSON(w, r, Response{
			Response: resp.OK(),
			Alias:    alias,
		})
	}
}
