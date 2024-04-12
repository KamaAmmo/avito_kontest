package main

import (
	"avito_app/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) getBanner(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	tagId, err := strconv.Atoi(v.Get("tag_id"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	featureId, err := strconv.Atoi(v.Get("feature_id"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	banner, err := app.banners.Get(tagId, featureId)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.clientError(w, http.StatusNotFound)
		} else {
			app.serverError(w, err)
		}
		return
	}
	data, err := json.MarshalIndent(banner, "", "	")
	if err != nil {
		app.serverError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (app *application) GetAllBanners(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	featureID, tagID, limitID := 0, 0, 0
	var err error
	tag := v.Get("tag_id")
	if tag != "" {
		tagID, err = strconv.Atoi(tag)
		if err != nil {
			app.clientError(w, http.StatusBadRequest)
			return
		}
	}
	feature := v.Get("feature_id")
	if feature != "" {
		featureID, err = strconv.Atoi(feature)
		if err != nil {
			app.clientError(w, http.StatusBadRequest)
			return
		}
	}
	limit := v.Get("limit")
	if limit != "" {
		limitID, err = strconv.Atoi(limit)
		if err != nil {
			app.clientError(w, http.StatusBadRequest)
			return
		}
	}

	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprint(tagID, featureID, limitID)))

}