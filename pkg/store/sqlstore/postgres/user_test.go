package postgres

import (
	"testing"

	"github.com/jewelsama/go-api-tdd/pkg/domain"
)

func TestCreateUser(t *testing.T) {
	pStore := NewPostgresStore()
	user := &domain.User{
		Email:   "test@test.com",
		Pasword: "password",
		Name:    "John Doe",
	}

	createdUser, err := pStore.CreateUser(user)

	if err != nil {
		t.Fatal(err)
	}

	if createdUser.ID == 0 {
		t.Errorf("want id not to be zero")
	}
}
