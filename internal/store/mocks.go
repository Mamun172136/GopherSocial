package store

import (
	"context"

	
)

func NewMockStore() Storage {
	return Storage{
		Users: &MockUserStore{},
	}
}

type MockUserStore struct {}

func (m *MockUserStore) Create(ctx context.Context, u *User) error {
	return nil
}

func (m *MockUserStore) GetByID(ctx context.Context, userID int64) (*User, error) {
	return &User{ID: userID}, nil
}

func (m *MockUserStore) GetByEmail(context.Context, string) (*User, error) {
	return &User{}, nil
}



func (m *MockUserStore) Delete(ctx context.Context, id int64) error {
	return nil
}