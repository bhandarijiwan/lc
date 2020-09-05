package main

import (
	"net/http"
)

// Where we wanna see how it scales
// Each Worker picks up a sync request task
// Each task is basically involves

// Cache
// Queue
// FileSystem
// Save Everything back into S3
// thumbnail generation
// Zip, PreZipped Files
// Download Endpoint

func handleTest(w http.ResponseWriter, req *http.Request) {

}

func main() {
	http.HandleFunc("/test", handleTest)
	http.ListenAndServe(":8888", nil)
}
