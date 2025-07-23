package main

import (
	"database/sql"
	"fmt"
	"log"
  "github.com/gofiber/fiber/v2"
  "github.com/go-sql-driver/mysql"
  "github.com/joho/godotenv"
)


type Car struct {
	id        int    `json:"id"`
	Brand_id     int `json:"brand_id"`
	Brand 	  string `json:"brand"`
	Model_id     int `json:"model_id"`
	Model     string `json:"model"`
	Year      int    `json:"year"`
}

type Model struct {
	id        int    `json:"id"`
	Brand     string    `json:"brand"` 
  Name      string    `json:"name"`
}

var db *sql.DB

func main()  {
envFile, _ := godotenv.Read(".env")
    cfg := mysql.NewConfig()
		cfg.User = envFile["DB_USER"] 
		cfg.Passwd = envFile["DB_PASS"]
    cfg.Net = "tcp"
    cfg.Addr = "127.0.0.1:3306"
    cfg.DBName = "carsdb"
  var err error

  db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
        log.Fatal(err)
    }
	pingErr := db.Ping()
  if pingErr != nil {
		log.Fatal(pingErr)
    }
  fmt.Println("Connected!")
defer db.Close()


	app := fiber.New()

	app.Get("/cars", getCars)
  app.Get("/cars/:id", getCarByID)
	app.Post("/cars", createCar)
	app.Put("/cars/:id", updateCar)
	//app.Delete("/cars/:id", deleteCar)

	app.Get("/Models", getModels)

	log.Println("Servidor escuchando en http://localhost:3000")
	log.Fatal(app.Listen(":3001"))
}

func getModels(c *fiber.Ctx) error {
	rows, err := db.Query(`
	SELECT models.id, brands.name, models.name 
	FROM models
	JOIN brands ON models.brand_id = brands.id
	`)
	if err != nil {
  	return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	models := []Model{}
	for rows.Next() {
		var a Model
		if err := rows.Scan(&a.id, &a.Brand, &a.Name); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		models = append(models, a)
	}
	return c.JSON(models)
}

func getCars(c *fiber.Ctx) error {
	rows, err := db.Query(`
		SELECT cars.id, cars.brand_id, brands.name, cars.model_id, models.name, cars.year
		FROM cars
		JOIN brands ON cars.brand_id = brands.id
		JOIN models ON cars.model_id = models.id
	`)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	cars := []Car{}
	for rows.Next() {
		var a Car
		if err := rows.Scan(&a.id, &a.Brand_id, &a.Brand, &a.Model_id, &a.Model, &a.Year); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		cars = append(cars, a)
	}	
	return c.JSON(cars)
}

func getCarByID(c *fiber.Ctx) error {
 	id := c.Params("id")
	row := db.QueryRow(`
		SELECT cars.id, brands.name, models.name, cars.year
		FROM cars
		JOIN brands ON cars.brand_id = brands.id
		JOIN models ON cars.model_id = models.id
		WHERE autos.id = ?
	`, id)
	var a Car
	if err := row.Scan(&a.id, &a.Brand, &a.Model, &a.Year); err != nil {
		if err == sql.ErrNoRows {
			return c.Status(404).SendString("Auto no encontrado")
		}
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(a)
}

func createCar(c *fiber.Ctx) error {
	var a Car
	if err := c.BodyParser(&a); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	_, err := db.Exec(`
		INSERT INTO cars (brand_id, model_id, year)
		VALUES (?, ?, ?)
	`, a.Brand_id, a.Model_id, a.Year)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(201).SendString("Auto creado")
}

func updateCar(c *fiber.Ctx) error {
	id := c.Params("id")
	var a Car
	if err := c.BodyParser(&a); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	_, err := db.Exec(`
		UPDATE cars
		SET brand_id = ?, model_id = ?, year = ?
		WHERE id = ?
	`, a.Brand_id, a.Model_id, a.Year, id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendString("Auto actualizado")
}
