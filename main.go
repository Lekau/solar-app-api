package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)
func createDummyDB() error {
	db, err := sql.Open("sqlite3", "./dummyData.db")
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE users (
			id INTEGER PRIMARY KEY,
			name TEXT,
			email TEXT,
			battery_percentage REAL,
			battery_health TEXT,
			location TEXT,
			hours_on_battery INTEGER,
			on_battery_power BOOLEAN
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	for i := 1; i <= 10; i++ {
		_, err = db.Exec(`
			INSERT INTO users (name, email, battery_percentage, battery_health, location, hours_on_battery, on_battery_power)
			VALUES (?, ?, ?, ?, ?, ?, ?)
		`, fmt.Sprintf("User %d", i), fmt.Sprintf("user%d@example.com", i), 95.0, "Good", "Edenvale, Gauteng", 8, true)
		if err != nil {
			return fmt.Errorf("failed to insert data: %w", err)
		}
	}
	return nil
}


func ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Yes the server is running",
	})
}

func getInformation(c *gin.Context) {
	// Get the user ID from the URL parameter
	userID := c.Param("id")

	// Open the SQLite database file
	db, err := sql.Open("sqlite3", "./mydatabase.db")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to open database",
		})
		return
	}
	defer db.Close()

	// Query the database for the user with the specified ID
	var (
		name              string
		batteryPercentage float64
		batteryHealth     string
		location          string
		hoursOnBattery    int
		onBatteryPower    bool
	)
	row := db.QueryRow(`
		SELECT name, battery_percentage, battery_health, location, hours_on_battery, on_battery_power
		FROM users
		WHERE id = ?
	`, userID)
	err = row.Scan(&name, &batteryPercentage, &batteryHealth, &location, &hoursOnBattery, &onBatteryPower)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	// Create a JSON response with the user's information
	response := gin.H{
		"id":                userID,
		"name":              name,
		"battery_percentage": batteryPercentage,
		"battery_health":     batteryHealth,
		"location":          location,
		"hours_on_battery":  hoursOnBattery,
		"on_battery_power":  onBatteryPower,
	}

	// Send the JSON response
	c.JSON(http.StatusOK, response)
}


func main() {
	createDummyDB()
	route := gin.Default()
	route.GET("/", ping)
	route.GET("/getInformation", getInformation)
	route.Run("localhost:5001")
}