package controllers

import (
	"VenusX/api/db"
	"VenusX/api/models"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetSales(c *gin.Context) {
	sales, err := db.GetAllSales()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al obtener ventas"})
		return
	}
	c.JSON(200, sales)
}

func CreateSales(c *gin.Context) {
	var newSales []models.Sale

	if err := c.ShouldBindJSON(&newSales); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Datos inválidos"})
		return
	}

	itemsJSON, _ := json.Marshal(raw["items"])

	// Determine payment method from payments array or single method
	method := ""
	if payments, ok := raw["payments"]; ok {
		if arr, ok := payments.([]interface{}); ok && len(arr) > 0 {
			parts := make([]string, 0)
			for _, p := range arr {
				if pm, ok := p.(map[string]interface{}); ok {
					if m, ok := pm["method"].(string); ok {
						parts = append(parts, m)
					}
				}
			}
			method = strings.Join(parts, "+")
		}
	}
	if method == "" {
		if m, ok := raw["method"].(string); ok {
			method = m
		}
	}

	newSale := models.Sale{
		Items:     string(itemsJSON),
		Subtotal:  getFloat(raw, "subtotal"),
		TaxRate:   getFloat(raw, "taxRate"),
		TaxAmount: getFloat(raw, "taxAmount"),
		Discount:  getFloat(raw, "discount"),
		Total:     getFloat(raw, "total"),
		Method:    method,
		Cashier:   getString(raw, "cashier"),
		Status:    getString(raw, "status"),
		Comment:   getString(raw, "comment"),
	}

	err := db.CreateSale(&newSale)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al guardar la venta"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Venta creada exitosamente", "id": newSale.ID})
}

func getFloat(data map[string]interface{}, key string) float64 {
	if v, ok := data[key]; ok {
		switch val := v.(type) {
		case float64:
			return val
		case int:
			return float64(val)
		}
	}
	return 0
}

func getString(data map[string]interface{}, key string) string {
	if v, ok := data[key]; ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func DeleteSale(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = db.DeleteSaleByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al borrar la venta"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Venta eliminada correctamente"})
}
