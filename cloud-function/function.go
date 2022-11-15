// Package p contains an HTTP Cloud Function.
package p

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	firestore "cloud.google.com/go/firestore"
	"github.com/eduardouliana/to-do-list/internal/task/infra/database"
	"github.com/eduardouliana/to-do-list/internal/task/usecase"
	"google.golang.org/api/option"
)

func save(repository *database.TaskRepository, r *http.Request) interface{} {
	uc := usecase.SaveTaskUseCase{TaskRepository: repository}

	var inputDTO usecase.SaveTaskInputDTO

	json.NewDecoder(r.Body).Decode(&inputDTO)

	outputDTO, err := uc.Execute(inputDTO)
	if err != nil {
		panic(err)
	}

	return outputDTO
}

func getOne(repository *database.TaskRepository, r *http.Request) interface{} {
	uc := usecase.ReadTaskUseCase{TaskRepository: repository}

	var inputDTO usecase.ReadTaskInputDTO

	json.NewDecoder(r.Body).Decode(&inputDTO)

	outputDTO, err := uc.Execute(inputDTO)
	if err != nil {
		panic(err)
	}

	return outputDTO
}

func getAll(repository *database.TaskRepository, r *http.Request) interface{} {
	uc := usecase.ReadTaskUseCase{TaskRepository: repository}

	outputDTO, err := uc.ExecuteAll()
	if err != nil {
		panic(err)
	}

	return outputDTO
}

func deleteOne(repository *database.TaskRepository, r *http.Request) interface{} {
	uc := usecase.DeleteTaskUseCase{TaskRepository: repository}

	var inputDTO usecase.DeleteTaskInputDTO

	json.NewDecoder(r.Body).Decode(&inputDTO)

	outputDTO, err := uc.Execute(inputDTO)
	if err != nil {
		panic(err)
	}

	return outputDTO
}

func deleteAll(repository *database.TaskRepository, r *http.Request) interface{} {
	uc := usecase.DeleteTaskUseCase{TaskRepository: repository}

	err := uc.ExecuteAll()
	if err != nil {
		panic(err)
	}

	return true
}

func MainFunction(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("/workspace/serverless_function_source_code/serviceAccountKey.json")
	client, err := firestore.NewClient(ctx, "to-do-list-96be7", opt)
	if err != nil {
		panic(fmt.Sprintf("firestore error:%s", err))
	}
	defer client.Close()

	repository := database.NewTaskRepository(client)

	output := save(repository, r)

	output = getOne(repository, r)

	output = getAll(repository, r)

	output = deleteOne(repository, r)

	output = deleteAll(repository, r)

	outputJson, err := json.Marshal(output)
	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, outputJson)

	log.Printf("Output: %v", output)
}
