package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"VenusX/api/controllers"
	"VenusX/api/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestDB initializes an in-memory SQLite database for testing purposes
func setupTestDB() {
	testDB, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	testDB.AutoMigrate(&models.Sale{}, &models.SaleItem{})

	// TODO: Assign testDB to your global database variable
	// Example: db.DB = testDB
}

func TestCreateSale_Success(t *testing.T) {
	setupTestDB()
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/sales", controllers.CreateSale)

	mockSales := []models.Sale{
		{
			Total:         100.0,
			PaymentMethod: "Efectivo",
			Status:        "Completado",
			Items: []models.SaleItem{
				{ProductID: 1, Quantity: 2, UnitPrice: 50.0, Subtotal: 100.0},
			},
		},
	}

	jsonValue, _ := json.Marshal(mockSales)

	req, _ := http.NewRequest("POST", "/sales", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Venta creada exitosamente", response["message"])
}
