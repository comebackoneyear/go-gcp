// Package gmutex implements a global mutex using Google Cloud Storage.
package gmutex

import (
	"context"
	"sync"

	"cloud.google.com/go/storage"
)

// HTTPClient should be set to an http.Client before first use.
// If unset google.DefaultClient will be used.
var StorageClient *storage.Client
var initMtx sync.Mutex

func initClient(ctx context.Context) (err error) {
	initMtx.Lock()
	defer initMtx.Unlock()

	if StorageClient == nil {
		StorageClient, err = storage.NewClient(ctx)
		if err != nil {
			return err
		}
	}
	return err

}
