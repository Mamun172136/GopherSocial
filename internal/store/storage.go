package store

import (
	"context"
	"database/sql"
	"errors"
)


var ErrNotFound=errors.New("record Not Found")
var ErrConflict  = errors.New("resource already exists")

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetByID(context.Context, int64) (*Post, error)
		Delete(context.Context, int64)error
		Update(context.Context, *Post) error
	}
	Users interface {
		Create(context.Context, *User) error
		GetByID(context.Context, int64)(*User, error)
	}
	Comments interface{
		GetByPostID(context.Context, int64) ([]Comment,error)
	}
	Followers interface{
		Follow(context.Context,int64, int64  ) error
		Unfollow(context.Context, int64, int64) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:  &PostsStore{db},
		Users:  &UsersStore{db},
		Comments: &CommentStore{db},
		Followers: &FollowerStore{db},
	}
}