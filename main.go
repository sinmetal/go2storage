package main

import (
	"context"
	"log"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

func main() {
	bucket := os.Getenv("BUCKET")
	log.Printf("Bucket=%s\n", bucket)

	ctx := context.Background()

	for {
		now := time.Now().Format("2006-01-02-15-04-05")
		log.Printf("go2storage next file name %s\n", now)
		write(ctx, bucket, now, []byte(now))
		time.Sleep(30 * time.Second)
	}

	log.Println("DONE")
}

func write(ctx context.Context, bucket string, object string, body []byte) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	b := client.Bucket(bucket)
	w := b.Object(object).NewWriter(ctx)

	_, err = w.Write(body)
	if err != nil {
		log.Fatalf("Failed to Object Write. err=%+v\n", err)
	}

	if err := w.Close(); err != nil {
		log.Fatalf("Failed to Object Write Close. err=%+v\n", err)
	}
}
