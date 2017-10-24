package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/ChristianAEDev/worky-server/data"
)

func main() {
	log.Printf("Starting server")

	port := "9090"
	log.Printf("API endpoint starting on port %v", port)

	router := mux.NewRouter()

	//Setup the endpoints
	initEndpointsV1(router)

	methods := []string{"GET", "POST", "PUT", "DELETE"}
	headers := []string{"Content-Type"}
	// Startup the endpoint
	http.ListenAndServe(":"+port,
		loggingMiddleware(handlers.CORS(handlers.AllowedMethods(methods), handlers.AllowedHeaders(headers))(router)))
}

// loggingMiddleware takes a http.Handler as an argument and allows us to chain handlers. In an
// anonymous function we implement our middleware function. At the end of this function we call
// the handler which will be chained.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Log the request
		log.Printf("[%s] %q", r.Method, r.URL.String())

		// Chain the handler
		next.ServeHTTP(w, r)
	})
}

// initEndpointsV1 initializes the endpoints used for version 1.
func initEndpointsV1(router *mux.Router) {
	apiV1 := router.PathPrefix("/api/v1").Subrouter()

	// Setup the API endpoint
	apiV1.HandleFunc("/tasks", GetTasks).Methods("GET")

}

// GetTasks loads all tasks from the database and sends them to the caller
func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks = data.GetTasksDummy()

	tasksJSON, err := data.ToJSON(tasks)
	if err != nil {
		log.Warningf("Error creating JSON representation of a slice of tasks")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	log.Infof("Sending %v tasks", len(tasks))
	w.Write(tasksJSON)
}
