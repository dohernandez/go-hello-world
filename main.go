package main

import (
	"net/http"

	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Product struct {
	ID       int     `json:"-"`
	Discount float64 `json:"discount"`
}

func main() {
	var r chi.Router = chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.PlainText(w, r, "Welcome!")
	})

	r.Get("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		ID := chi.URLParam(r, "id")

		pID, err := strconv.Atoi(ID)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.PlainText(w, r, "500 - Something bad happened!")

			return
		}

		var p Product

		p.ID = pID

		if p.ID%2 == 0 {
			p.Discount = 0.2
		}

		render.JSON(w, r, p)
	})

	if err := http.ListenAndServe(":8085", r); err != nil {
		panic(err)
	}
}
