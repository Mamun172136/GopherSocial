package main

import (
	"net/http"

	"github.com/social/internal/store"
)

func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request){

		fq := store.PaginatedFeedQuery{
		Limit:  20,
		Offset: 0,
		Sort:   "desc",
	
	}

	fq, err := fq.Parse(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(fq); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}



	ctx := r.Context()
	user := getUserFromContext(r)

	feed, err := app.store.Posts.GetUserFeed(ctx, user.ID, fq)
	if err != nil{
		app.internalServerError(w,r,err)
	}

	err = writeJSON(w,http.StatusOK,feed)
	if err != nil{
		app.internalServerError(w,r,err)
	}
}