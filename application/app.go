package main

import (
	"backend-go/mod/domain/repository"
	"backend-go/mod/infrastructure/drivenadapters/zinsearch"
	"backend-go/mod/infrastructure/entrypoints/rest"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("The environment variables were not initialized")
	}

	var wg sync.WaitGroup
	var emailAdapter repository.EmailsRespository = zinsearch.NewZincSearchRepositoryImpl()

	if os.Getenv("ENTRYPOINT_APIREST_ENABLED") == "true" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rest.NewRest(emailAdapter).Run(os.Getenv("ENTRYPOINT_APIREST_PORT"))
		}()
	}
	wg.Wait()
}
