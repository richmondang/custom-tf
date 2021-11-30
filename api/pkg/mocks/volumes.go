package mocks

import (
	"net/http"
	"api/pkg/models"
)

var Volumes = []models.Volume{
    {
        ID:     "123456",
        Name:  "Test_Volume1",
        Description: "APEX Data Storage Services Volume 1",
        Size:   156150,
    },
}