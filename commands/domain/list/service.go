package list

import (
	"context"
	"fmt"
	"go-storage/commands/infrastructure"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

//List ...
type List interface {
	Display() error
}

//NewList return new instance
func NewList(bucketName string) List {
	return &MyList{
		bucketName,
	}
}

//MyList ...
type MyList struct {
	bucketName string
}

//Display objects
func (m *MyList) Display() error {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	fmt.Println(client)
	if err != nil {
		return err
	}
	defer client.Close()

	gcs := infrastructure.NewGcs(client, m.bucketName)

	for {
		object, err := gcs.Get(ctx).Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		fmt.Println(object)
	}
	return nil
}
