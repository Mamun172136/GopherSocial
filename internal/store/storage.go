package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)


var ErrNotFound=errors.New("record Not Found")
var ErrConflict  = errors.New("resource already exists")
var  ErrDuplicateEmail  = errors.New("a user with that email already exists")
var ErrDuplicateUsername = errors.New("a user with that username already exists")
var QueryTimeoutDuration = time.Second * 5

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetByID(context.Context, int64) (*Post, error)
		Delete(context.Context, int64)error
		Update(context.Context, *Post) error
	GetUserFeed(context.Context, int64, PaginatedFeedQuery) ([]PostWithMetadata, error)
	}
	Users interface {
		Create(context.Context, *User) error
		GetByID(context.Context, int64)(*User, error)
		GetByEmail(context.Context, string) (*User, error)
	}
	Comments interface{
		GetByPostID(context.Context, int64) ([]Comment,error)
		Create(context.Context, *Comment)error
	}
	Followers interface{
		Follow(context.Context,int64, int64  ) error
		Unfollow(context.Context, int64, int64) error
	}
	Roles interface {
		GetByName(context.Context, string) (*Role, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:  &PostsStore{db},
		Users:  &UsersStore{db},
		Comments: &CommentStore{db},
		Followers: &FollowerStore{db},
		Roles:     &RoleStore{db},
	}
}