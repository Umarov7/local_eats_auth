package postgres

import (
	"auth-service/config"
	pba "auth-service/genproto/auth"
	pbu "auth-service/genproto/user"
	"context"
	"log"
	"reflect"
	"testing"
)

var (
	mockUser        = &pba.RegisterResponse{}
	mockUserProfile = &pbu.Profile{}
)

func userDB() *UserRepo {
	db, err := ConnectDB(&config.Config{
		DB_HOST:     "localhost",
		DB_PORT:     "5432",
		DB_USER:     "postgres",
		DB_NAME:     "local_eats_auth",
		DB_PASSWORD: "root",
	})
	if err != nil {
		log.Fatal("could not connect to postgres")
	}

	return NewUserRepo(db)
}

func TestCreate(t *testing.T) {
	u := userDB()
	testUser := &pba.RegisterRequest{
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: "password123",
		FullName: "Test User",
		UserType: "customer",
	}

	createdUser, err := u.Create(context.Background(), testUser)
	if err != nil {
		t.Fatalf("Create method returned error: %v", err)
	}

	expectedUser := &pba.RegisterResponse{
		Username: "testuser",
		Email:    "testuser@example.com",
		FullName: "Test User",
		UserType: "customer",
	}

	if createdUser.Username != expectedUser.Username ||
		createdUser.Email != expectedUser.Email ||
		createdUser.FullName != expectedUser.FullName ||
		createdUser.UserType != expectedUser.UserType {
		t.Errorf("Unexpected user details. Got %+v, expected %+v", createdUser, expectedUser)
	}

	mockUser = createdUser
}

func TestUpdate(t *testing.T) {
	u := userDB()
	newInfo := &pbu.NewInfo{
		Id:          mockUser.Id,
		FullName:    mockUser.FullName,
		Address:     "456 Avenue, City",
		PhoneNumber: "+9876543210",
	}

	updatedDetails, err := u.Update(context.Background(), newInfo)
	if err != nil {
		t.Fatalf("Update method returned error: %v", err)
	}

	expectedDetails := &pbu.Details{
		Id:          mockUser.Id,
		Username:    mockUser.Username,
		Email:       mockUser.Email,
		FullName:    mockUser.FullName,
		UserType:    mockUser.UserType,
		Address:     newInfo.Address,
		PhoneNumber: newInfo.PhoneNumber,
		UpdatedAt:   updatedDetails.UpdatedAt,
	}

	if !reflect.DeepEqual(updatedDetails, expectedDetails) {
		t.Errorf("Unexpected updated details. Got %+v, expected %+v", updatedDetails, expectedDetails)
	}

	mockUserProfile = &pbu.Profile{
		Id:          expectedDetails.Id,
		Username:    expectedDetails.Username,
		Email:       expectedDetails.Email,
		FullName:    expectedDetails.FullName,
		UserType:    expectedDetails.UserType,
		Address:     expectedDetails.Address,
		PhoneNumber: expectedDetails.PhoneNumber,
		CreatedAt:   mockUser.CreatedAt,
		UpdatedAt:   expectedDetails.UpdatedAt,
	}
}

func TestRead(t *testing.T) {
	u := userDB()
	testUserID := &pbu.ID{Id: mockUser.Id}

	profile, err := u.Read(context.Background(), testUserID)
	if err != nil {
		t.Fatalf("Read method returned error: %v", err)
	}

	expectedProfile := mockUserProfile

	if !reflect.DeepEqual(profile, expectedProfile) {
		t.Errorf("Unexpected profile details. Got %+v, expected %+v", profile, expectedProfile)
	}
}

func TestDelete(t *testing.T) {
	u := userDB()
	testUserID := &pbu.ID{Id: mockUserProfile.Id}

	err := u.Delete(context.Background(), testUserID)
	if err != nil {
		t.Fatalf("Delete method returned error: %v", err)
	}

}
