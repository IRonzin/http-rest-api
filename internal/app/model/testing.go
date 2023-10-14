package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@example.com",
		Password: "password",
	}
}

func TestUserBenchmark(b *testing.B) *User {
	return &User{
		Email:    "user@example.com",
		Password: "password",
	}
}
