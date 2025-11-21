package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	//MinIO client setup
	endpoint := "localhost:9000"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	useSSL := false
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
	// HTTP handler for streaming the video
	http.HandleFunc("/stream-video", func(w http.ResponseWriter, r *http.Request) {
		bucketName := "videos"
		objectName := "my-video.mp4"
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Minute)
		defer cancel()
		// Get the object from MinIO
		object, err := minioClient.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
		if err != nil {
			http.Error(w, "Failed to get object from MinIO: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer object.Close()
		//Get object info to set Content-Type and Content-Length headers
		objectInfo, err := object.Stat()
		if err != nil {
			http.Error(w, "Failed to get object info: "+err.Error(), http.StatusInternalServerError)
			return
		}
		//Set appropriate headers for video streaming
		w.Header().Set("Content-Type", objectInfo.ContentType)
		w.Header().Set("Content-Length", string(objectInfo.Size))
		w.Header().Set("Accept-Ranges", "bytes")
		//Stream the content directly to the HTTP response writer
		http.ServeContent(w, r, objectName, objectInfo.LastModified, object)
	})
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
