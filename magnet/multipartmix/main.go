package main

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"time"
)

const boundary = "my-custom-boundary"

func main() {
	http.HandleFunc("/multipart/mixed", handleMultipartStream)
	http.HandleFunc("/multipart/mixed-replace", handleMultipartXMixedReplace)
	http.HandleFunc("/download", downloadFile)
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleMultipartXMixedReplace(w http.ResponseWriter, r *http.Request) {
	//Set the Content-Type header to indicate multipart/mixed and specify the boundary
	w.Header().Set("Content-Type", fmt.Sprintf("multipart/x-mixed-replace; boundary=%s", boundary))
	//Create a multipart writer that writes to the HTTP response writer
	mpWriter := multipart.NewWriter(w)
	mpWriter.SetBoundary(boundary) //Ensure the writer uses the same boundary
	partHeader := make(textproto.MIMEHeader)
	partHeader.Add("Content-Type", "image/jpg")
	for count := range 5 {
		imageFile, err := os.Open(fmt.Sprintf("magnet/images/capture_%d.jpg", count))
		if err != nil {
			http.Error(w, "Error opening image file: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer imageFile.Close()
		imagePart, err := mpWriter.CreatePart(partHeader)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//Copy the image content to the part writer
		_, err = io.Copy(imagePart, imageFile)
		if err != nil {
			log.Printf("Error copy image data: %v", err)
		}
		time.Sleep(10 * time.Second)
	}
	mpWriter.Close()
}

func downloadFile(w http.ResponseWriter, r *http.Request) {

	// Construct the full path to the file on the server
	filePath := "magnet/images/capture_0.jpg"

	// Check if the file exists
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, fmt.Sprintf("Error checking file: %v", err), http.StatusInternalServerError)
		return
	}

	// Set the Content-Disposition header to "attachment" to force download
	// The "filename" parameter specifies the name the downloaded file will have
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", "capture_0.jpg"))

	// Optionally, set the Content-Type header based on the file type
	// For example, if it's a PDF: w.Header().Set("Content-Type", "application/pdf")
	// You can use http.DetectContentType for more dynamic detection if needed.

	// Serve the file
	http.ServeFile(w, r, filePath)
}

func handleMultipartStream(w http.ResponseWriter, r *http.Request) {
	//Set the Content-Type header to indicate multipart/mixed and specify the boundary
	w.Header().Set("Content-Type", fmt.Sprintf("multipart/mixed; boundary=%s", boundary))
	//Create a multipart writer that writes to the HTTP response writer
	mpWriter := multipart.NewWriter(w)
	mpWriter.SetBoundary(boundary) //Ensure the writer uses the same boundary
	//Part 1: Text content
	part1, err := mpWriter.CreatePart(map[string][]string{
		"Content-Type": {"text/plain"},
	})

	if err != nil {
		http.Error(w, "Failed to create part 1", http.StatusInternalServerError)
		return
	}
	io.WriteString(part1, "This is the first part of the multipart stream")
	//Part 2: Json content
	part2, err := mpWriter.CreatePart(map[string][]string{
		"Content-Type": {"application/json"},
	})
	if err != nil {
		http.Error(w, "Failed to create part 2", http.StatusInternalServerError)
		return
	}
	io.WriteString(part2, `{"key":"value","number":123}`)
	//Part 3: Simulate a file stream
	part3, err := mpWriter.CreatePart(map[string][]string{
		"Content-Type": {"application/octet-stream"},
	})
	if err != nil {
		http.Error(w, "Failed to create part 3", http.StatusInternalServerError)
		return
	}
	//Write some dummy binary data
	for i := range 5 {
		part3.Write([]byte(fmt.Sprintf("Binary data chunk %d\n", i)))
		time.Sleep(100 * time.Millisecond)
	}
	//Add an image part
	imageFile, err := os.Open("magnet/images/capture_0.jpg")
	if err != nil {
		http.Error(w, "Error opening image file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer imageFile.Close()
	imagePart, err := mpWriter.CreatePart(map[string][]string{
		"Content-Type":        {"image/jpeg"},
		"Content_Disposition": {`attachment; filename="image.jpg"`},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//Copy the image content to the part writer
	_, err = io.Copy(imagePart, imageFile)
	if err != nil {
		log.Printf("Error copy image data: %v", err)
	}
	mpWriter.Close()
}
