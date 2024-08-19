package postgres

import (
	"auth-service/config"
	pb "auth-service/genproto/kitchen"
	"context"
	"database/sql"
	"log"
	"reflect"
	"testing"
)

var mockKitchenData = &pb.CreateResponse{}

func kitchenDB() *KitchenRepo {
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

	return NewKitchenRepo(db)
}

func TestCreateKitchen(t *testing.T) {
	k := kitchenDB()
	testData := &pb.CreateRequest{
		OwnerId:     "123e4567-e89b-12d3-a456-426614174001",
		Name:        "Test Kitchen",
		Description: "A test kitchen for testing purposes",
		CuisineType: "Test cuisine",
		Address:     "123 Test Street",
		PhoneNumber: "+1234567890",
	}

	createdKitchen, err := k.Create(context.Background(), testData)
	if err != nil {
		t.Fatalf("Create method returned error: %v", err)
	}

	expectedKitchen := &pb.CreateResponse{
		Id:          createdKitchen.Id,
		OwnerId:     testData.OwnerId,
		Name:        testData.Name,
		Description: testData.Description,
		CuisineType: testData.CuisineType,
		Address:     testData.Address,
		PhoneNumber: testData.PhoneNumber,
		Rating:      0,
		CreatedAt:   createdKitchen.CreatedAt,
	}

	if !reflect.DeepEqual(createdKitchen, expectedKitchen) {
		t.Errorf("Unexpected kitchen details. Got %+v, expected %+v", createdKitchen, expectedKitchen)
	}

	mockKitchenData = createdKitchen
}

func TestUpdateKitchen(t *testing.T) {
	k := kitchenDB()

	testData := &pb.NewData{
		Id:          mockKitchenData.Id,
		Name:        mockKitchenData.Name,
		Description: mockKitchenData.Description,
		PhoneNumber: mockKitchenData.PhoneNumber,
	}

	updatedData, err := k.Update(context.Background(), testData)
	if err != nil {
		t.Fatalf("Update method returned error: %v", err)
	}

	expectedData := &pb.UpdatedData{
		Id:          mockKitchenData.Id,
		OwnerId:     mockKitchenData.OwnerId,
		Name:        mockKitchenData.Name,
		Description: mockKitchenData.Description,
		CuisineType: mockKitchenData.CuisineType,
		Address:     mockKitchenData.Address,
		PhoneNumber: mockKitchenData.PhoneNumber,
		Rating:      mockKitchenData.Rating,
		UpdatedAt:   updatedData.UpdatedAt,
	}

	if !reflect.DeepEqual(updatedData, expectedData) {
		t.Errorf("Unexpected updated kitchen data. Got %+v, expected %+v", updatedData, expectedData)
	}
}

func TestReadKitchen(t *testing.T) {
	k := kitchenDB()
	testKitchenID := &pb.ID{Id: mockKitchenData.Id}

	kitchenInfo, err := k.Read(context.Background(), testKitchenID)
	if err != nil {
		t.Fatalf("Read method returned error: %v", err)
	}

	expectedKitchenInfo := &pb.Info{
		Id:          testKitchenID.Id,
		OwnerId:     mockKitchenData.OwnerId,
		Name:        mockKitchenData.Name,
		Description: mockKitchenData.Description,
		CuisineType: mockKitchenData.CuisineType,
		Address:     mockKitchenData.Address,
		PhoneNumber: mockKitchenData.PhoneNumber,
		Rating:      mockKitchenData.Rating,
		TotalOrders: kitchenInfo.TotalOrders,
		CreatedAt:   mockKitchenData.CreatedAt,
		UpdatedAt:   kitchenInfo.UpdatedAt,
	}

	if !reflect.DeepEqual(kitchenInfo, expectedKitchenInfo) {
		t.Errorf("Unexpected kitchen information. Got %+v, expected %+v", kitchenInfo, expectedKitchenInfo)
	}
}

func TestDeleteKitchen(t *testing.T) {
	k := kitchenDB()
	testKitchenID := &pb.ID{Id: mockKitchenData.Id}

	err := k.Delete(context.Background(), testKitchenID)
	if err != nil {
		t.Fatalf("Delete method returned error: %v", err)
	}

	var deletedAt sql.NullTime
	err = k.DB.QueryRow("SELECT deleted_at FROM kitchens WHERE id = $1", testKitchenID.Id).Scan(&deletedAt)
	if err != nil {
		t.Fatalf("Error querying deleted kitchen: %v", err)
	}

	if !deletedAt.Valid {
		t.Errorf("Kitchen with ID %s was not marked as deleted", testKitchenID.Id)
	}
}

func TestFetchKitchens(t *testing.T) {
	k := kitchenDB()

	testPagination := &pb.Pagination{
		Limit:  2,
		Offset: 0,
	}

	kitchens, totalRows, err := k.Fetch(context.Background(), testPagination)
	if err != nil {
		t.Fatalf("Fetch method returned error: %v", err)
	}

	if len(kitchens) != int(testPagination.Limit) {
		t.Errorf("Unexpected number of kitchens fetched. Got %d, expected %d", len(kitchens), testPagination.Limit)
	}

	kn1 := pb.KitchenDetails{
		Id:          "223e4567-e89b-12d3-a456-426614174001",
		Name:        "Taste of Italy",
		CuisineType: "Italian",
		Rating:      4.5,
		TotalOrders: 200,
	}
	kn2 := pb.KitchenDetails{
		Id:          "223e4567-e89b-12d3-a456-426614174002",
		Name:        "Sweet Delights Bakery",
		CuisineType: "Bakery",
		Rating:      4.8,
		TotalOrders: 350,
	}

	expectedKitchens := []*pb.KitchenDetails{&kn1, &kn2}

	if !reflect.DeepEqual(kitchens, expectedKitchens) {
		t.Errorf("Fetched kitchens do not match expected. Got %+v, expected %+v", kitchens, expectedKitchens)
	}

	expectedTotalRows, err := k.CountRows(context.Background())
	if err != nil {
		t.Errorf("Error counting total number of rows")
	}

	if totalRows != expectedTotalRows {
		t.Errorf("Unexpected total number of rows. Got %d, expected %d", totalRows, expectedTotalRows)
	}

}

func TestSearchKitchens(t *testing.T) {
	k := kitchenDB()
	testSearchDetails := &pb.SearchDetails{
		Query:       "pasta",
		CuisineType: "Italian",
		Rating:      4,
		Pagination: &pb.Pagination{
			Limit:  1,
			Offset: 0,
		},
	}

	kitchens, totalRows, err := k.Search(context.Background(), testSearchDetails)
	if err != nil {
		t.Fatalf("Search method returned error: %v", err)
	}

	if len(kitchens) != int(testSearchDetails.Pagination.Limit) {
		t.Errorf("Unexpected number of kitchens fetched. Got %d, expected %d", len(kitchens), testSearchDetails.Pagination.Limit)
	}

	kn := pb.KitchenDetails{
		Id:          "223e4567-e89b-12d3-a456-426614174005",
		Name:        "Pasta Paradise",
		CuisineType: "Italian",
		Rating:      4.30,
		TotalOrders: 180,
	}

	expectedKitchens := []*pb.KitchenDetails{&kn}

	if !reflect.DeepEqual(kitchens, expectedKitchens) {
		t.Errorf("Fetched kitchens do not match expected. Got %+v, expected %+v", kitchens, expectedKitchens)
	}

	expectedTotalRows, err := k.CountRows(context.Background())
	if err != nil {
		t.Errorf("Error counting total number of rows")
	}

	if totalRows != expectedTotalRows {
		t.Errorf("Unexpected total number of rows. Got %d, expected %d", totalRows, expectedTotalRows)
	}
}
