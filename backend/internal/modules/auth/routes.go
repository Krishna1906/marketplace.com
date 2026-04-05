// package auth

// import "github.com/gorilla/mux"

// func RegisterAuthRoutes(router *mux.Router) {
// 	authRouter := router.PathPrefix("/auth").Subrouter()

// 	authRouter.HandleFunc("/login", Login).Methods("POST")
// 	authRouter.HandleFunc("/register", Register).Methods("POST")
// }

package auth

import "github.com/gorilla/mux"

func RegisterAuthRoutes(router *mux.Router) {
	// authRouter := router.PathPrefix("/auth").Subrouter() // ❌ REMOVE THIS

	router.HandleFunc("/login", Login).Methods("POST")
	router.HandleFunc("/register", Register).Methods("POST")
}
