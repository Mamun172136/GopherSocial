package main

import (
	"context"
	"net/http"

	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/social/internal/store"
)

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request){

	user:= getUserFromContext(r)

	if err := writeJSON(w, http.StatusOK,user); err!= nil{
		app.internalServerError(w,r,err)
	}

}

type userKey string
const userCtx userKey ="user"

type FollowUser struct{
	UserID int64 `json:"user_id"`
}

func ( app *application) followUserHandler(w http.ResponseWriter, r *http.Request){
	followerUser := getUserFromContext(r)

	var payload FollowUser
	if err := readJson(w, r, &payload); err !=nil{
		app.badRequestResponse(w,r,err)
		return
	}

	ctx := r.Context()
	err:= app.store.Followers.Follow(ctx, followerUser.ID, payload.UserID)
	if err != nil{
		switch err {
		case store.ErrConflict:
			writeJSONError(w, http.StatusConflict,"conflict")
			return
		default :
			app.internalServerError(w,r,err)
			return
		}
	
	}

	err=writeJSON(w, http.StatusNoContent,nil)
	if err !=nil{
		app.internalServerError(w,r,err)
		return
	}

}

func ( app *application) unfollowUserHandler(w http.ResponseWriter, r *http.Request){
	unfollowerUser := getUserFromContext(r)

	var payload FollowUser
	if err := readJson(w, r, &payload); err !=nil{
		app.badRequestResponse(w,r,err)
		return
	}

	ctx := r.Context()
	err:= app.store.Followers.Unfollow(ctx, unfollowerUser.ID, payload.UserID)
	if err !=nil{
		app.internalServerError(w,r,err)
		return
	}

	err=writeJSON(w, http.StatusNoContent,nil)
	if err !=nil{
		app.internalServerError(w,r,err)
		return
	}
}

func (app *application) userContextMiddleware(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		userID,err := strconv.ParseInt(chi.URLParam(r,"userID"),10,64)
	if err != nil{
		app.badRequestResponse(w,r,err)
		return
	}

	ctx := r.Context()

	user,err := app.store.Users.GetByID(ctx, userID)
	if err != nil{
		switch err {
		case store.ErrNotFound:
			app.notFoundResponse(w,r,err)
			return 

		default:
			app.internalServerError(w,r,err)
			return
		}
		
	}

	
	ctx = context.WithValue(ctx, userCtx, user)

	next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserFromContext(r *http.Request) *store.User{
	user, _ :=r.Context().Value(userCtx).(*store.User)

	return user
}