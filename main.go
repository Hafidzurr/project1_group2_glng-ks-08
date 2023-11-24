package main

import (
	"log"
	"net/http"

	"github.com/Hafidzurr/project1_group2_glng-ks-08/database"
	_ "github.com/Hafidzurr/project1_group2_glng-ks-08/docs" // Import the generated docs
	"github.com/Hafidzurr/project1_group2_glng-ks-08/router"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	db := database.Init("host=localhost user=postgres password= dbname=todo sslmode=disable")
	defer db.Close()

	r := router.Router()

	// Serve Swagger UI documentation
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
