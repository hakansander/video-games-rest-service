package main

import (
	"currency-investment-rest-service/models"
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"time"
)

func (app *application) getGame(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("[getGame] -> invalid id parameter"))

		app.errorHandler(w, err)
		return
	}

	app.logger.Println("id is", id)

	game := models.Game{
		ID:          id,
		Title:       "Red Alert 2",
		Description: "One of the best RTS strategy games",
		ReleaseDate: time.Date(1997, 01, 02, 01, 0, 0, 0, time.Local),
		Rating:      5,
		Price:       9.98,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, game, "game")
}

func (app *application) getAllGames(w http.ResponseWriter, r *http.Request) {

}
