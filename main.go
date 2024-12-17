package main

import (
	//"database/sql"

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

	fmt.Println(carStore)

	carService := carService.NewService(carStore)
	engineService := engineService.NewService(engineStore)

	carController := carHandler.NewCarHandler(carService)
	engineController := engineHandler.NewEngineHandler(engineService)

	schemaFile := "store/schema.sql"
	if err := executeSchema(db, schemaFile); err != nil {
		log.Fatal("error while executing schema file!")
	}

	router := mux.NewRouter()

	router.HandleFunc("/car/{id}", carController.GetCarById).Methods("GET")
	router.HandleFunc("/engine/{id}", engineController.GetEngineById).Methods("GET")

	port := os.Getenv("DB_PORT")

	fmt.Println(port)

	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
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
