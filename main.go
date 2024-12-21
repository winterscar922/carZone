package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/winterscar922/carZone/driver"
	carHandler "github.com/winterscar922/carZone/handler/car"
	engineHandler "github.com/winterscar922/carZone/handler/engine"
	middleware "github.com/winterscar922/carZone/middleware"
	carService "github.com/winterscar922/carZone/service/car"
	engineService "github.com/winterscar922/carZone/service/engine"
	carStore "github.com/winterscar922/carZone/store/car"
	engineStore "github.com/winterscar922/carZone/store/engine"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error while loading env")
	}

	driver.InitDb()
	defer driver.CloseDb()

	db := driver.GetDb()
	carStore := carStore.Open(db)
	engineStore := engineStore.Open(db)

	carService := carService.NewService(carStore)
	engineService := engineService.NewService(engineStore)

	carController := carHandler.NewCarHandler(carService)
	engineController := engineHandler.NewEngineHandler(engineService)

	schemaFile := "store/schema.sql"
	if err := executeSchema(db, schemaFile); err != nil {
		log.Fatal("error while executing schema file!")
	}

	router := mux.NewRouter()

	router.HandleFunc("/login", middleware.LoginHandler).Methods("POST")

	// car handlers
	router.HandleFunc("/car/{id}", middleware.JWTMiddleware(carController.GetCarById)).Methods("GET")
	router.HandleFunc("/car", middleware.JWTMiddleware(carController.CreateCar)).Methods("POST")
	router.HandleFunc("/car/{id}", middleware.JWTMiddleware(carController.UpdateCar)).Methods("PUT")
	router.HandleFunc("/car/{id}", middleware.JWTMiddleware(carController.DeleteCar)).Methods("DELETE")
	router.HandleFunc("/cars", middleware.JWTMiddleware(carController.GetAllCars)).Methods("GET")

	// engine handlers
	router.HandleFunc("/engine/{id}", middleware.JWTMiddleware(engineController.GetEngineById)).Methods("GET")
	router.HandleFunc("/engine", middleware.JWTMiddleware(engineController.CreateEngine)).Methods("POST")
	router.HandleFunc("/engine/{id}", middleware.JWTMiddleware(engineController.UpdateEngine)).Methods("PUT")
	router.HandleFunc("/engine/{id}", middleware.JWTMiddleware(engineController.DeleteEngine)).Methods("DELETE")
	router.HandleFunc("/engines", middleware.JWTMiddleware(engineController.GetAllEngines)).Methods("GET")

	port := ":8080"

	http.ListenAndServe(port, router)
}

func executeSchema(db *sql.DB, schemaFile string) error {
	sqlFile, err := os.ReadFile(schemaFile)

	if err != nil {
		return err
	}

	_, err = db.Exec(string(sqlFile))

	if err != nil {
		return err
	}

	fmt.Println("Executed Schema file successfully!")

	return nil
}
