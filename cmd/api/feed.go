package main

import "net/http"

func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()

	feed, err := app.store.Posts.GetUserFeed(ctx, int64(4))
	if err != nil{
		app.internalServerError(w,r,err)
	}

	err = writeJSON(w,http.StatusOK,feed)
	if err != nil{
		app.internalServerError(w,r,err)
	}
}