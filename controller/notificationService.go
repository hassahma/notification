// Service methods used by controller to fetch and process data from external api.
package controller

import (
	"fmt"
	"net/http"
	"github.com/hassahma/notification/db"
)

// Invalidates the cache by calling flushall
func saveCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("saving customer")
	db.SaveCustomer(w, r)
}