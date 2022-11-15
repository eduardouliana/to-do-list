package database

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/eduardouliana/to-do-list/internal/task/entity"
	"google.golang.org/api/iterator"
)

var COLLECTION = "Task"

type TaskRepository struct {
	FirestoreClient *firestore.Client
}

func NewTaskRepository(firestoreClient *firestore.Client) *TaskRepository {
	return &TaskRepository{FirestoreClient: firestoreClient}
}

func (u *TaskRepository) Save(task *entity.Task) error {
	_, err := u.FirestoreClient.Collection(COLLECTION).Doc(task.Id).Set(context.Background(), task)

	if err != nil {
		return err
	}
	return nil
}

func (u *TaskRepository) Delete(id string) error {
	_, err := u.FirestoreClient.Collection(COLLECTION).Doc(id).Delete(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (u *TaskRepository) Read(id string) (*entity.Task, error) {
	data, err := u.FirestoreClient.Doc(fmt.Sprintf("%s/%s", COLLECTION, id)).Get(context.Background())
	if err != nil {
		return nil, err
	}

	jsonData, err := json.MarshalIndent(data.Data(), "", "  ")
	if err != nil {
		return nil, err
	}

	task := entity.Task{}

	err = json.Unmarshal(jsonData, &task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (u *TaskRepository) ReadAll() ([]*entity.Task, error) {
	iter := u.FirestoreClient.Collection(COLLECTION).Documents(context.Background())
	defer iter.Stop()

	list := []*entity.Task{}

	for {
		doc, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, err
		}

		jsonData, err := json.MarshalIndent(doc.Data(), "", "  ")
		if err != nil {
			return nil, err
		}

		task := entity.Task{}

		err = json.Unmarshal(jsonData, &task)
		if err != nil {
			return nil, err
		}

		list = append(list, &task)
	}

	return list, nil
}

func (u *TaskRepository) DeleteAll() error {
	iter := u.FirestoreClient.Collection(COLLECTION).Documents(context.Background())
	numDeleted := 0
	batch := u.FirestoreClient.Batch()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		batch.Delete(doc.Ref)
		numDeleted++

	}
	// If there are no documents to delete,
	// the process is over.
	if numDeleted > 0 {
		_, err := batch.Commit(context.Background())
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *TaskRepository) Authenticate(task *entity.Task) error {
	//TODO - Implement read on Firestore
	return nil
}
