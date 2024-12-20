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

	// car handlers
	router.HandleFunc("/car/{id}", carController.GetCarById).Methods("GET")
	router.HandleFunc("/car", carController.CreateCar).Methods("POST")
	router.HandleFunc("/car/{id}", carController.UpdateCar).Methods("PUT")
	router.HandleFunc("/car/{id}", carController.DeleteCar).Methods("DELETE")
	router.HandleFunc("/cars", carController.GetAllCars).Methods("GET")

	// engine handlers
	router.HandleFunc("/engine/{id}", engineController.GetEngineById).Methods("GET")
	router.HandleFunc("/engine", engineController.CreateEngine).Methods("POST")
	router.HandleFunc("/engine/{id}", engineController.UpdateEngine).Methods("PUT")
	router.HandleFunc("/engine/{id}", engineController.DeleteEngine).Methods("DELETE")
	router.HandleFunc("/engines", engineController.GetAllEngines).Methods("GET")

	port := ":8080"

	fmt.Println(port)
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
