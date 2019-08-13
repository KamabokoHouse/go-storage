package infrastructure

import (
	"context"

	"cloud.google.com/go/storage"
)

//GcsMethod ...
type GcsMethod interface {
	Get(ctx context.Context) *storage.ObjectIterator
}

//NewGcs return new instance
func NewGcs(client *storage.Client, bucketName string) GcsMethod {
	return &GCS{
		client,
		bucketName,
	}
}

//GCS ...
type GCS struct {
	client     *storage.Client
	bucketName string
}

//Get objects from gcs buckets
func (my *GCS) Get(ctx context.Context) *storage.ObjectIterator {
	return my.client.Bucket(my.bucketName).Objects(ctx, nil)
}
