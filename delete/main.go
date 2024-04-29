package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Mock data - a map to simulate a database
var data = map[string]string{
    "1": "Item 1",
    "2": "Item 2",
    "3": "Item 3",
}

// Handler for deleting an item
func deleteItem(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    
    // Check if the item exists
    if _, ok := data[id]; ok {
        delete(data, id)
        w.WriteHeader(http.StatusNoContent) // 204 No Content
    } else {
        http.Error(w, "Item not found", http.StatusNotFound) // 404 Not Found
    }
}

func main() {
    r := mux.NewRouter()
    // Setting up the DELETE endpoint
    r.HandleFunc("/item/{id}", deleteItem).Methods("DELETE")
    
    log.Println("Server starting on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", r))
}
