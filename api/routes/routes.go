package routes

import (
  "api/controllers"
  "github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
  r := mux.NewRouter().StrictSlash(true)
  /* Users Routes */
  r.HandleFunc("/users", controllers.PostUsers).Methods("POST")
  r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
  r.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
  r.HandleFunc("/users/{id}", controllers.PutUsers).Methods("PUT")
  r.HandleFunc("/users/{id}", controllers.DeleteUsers).Methods("DELETE")
  /* Posts Routes */
  r.HandleFunc("/posts", controllers.PostPosts).Methods("POST")
  r.HandleFunc("/posts", controllers.GetPosts).Methods("GET")
  r.HandleFunc("/posts/{id}", controllers.GetPost).Methods("GET")
  r.HandleFunc("/posts/{id}", controllers.PutPosts).Methods("PUT")
  r.HandleFunc("/posts/{id}", controllers.DeletePosts).Methods("DELETE")
  /* Feedback Routes */
  r.HandleFunc("/feedbacks", controllers.PostFeedbacks).Methods("POST")
  r.HandleFunc("/feedbacks", controllers.GetFeedbacks).Methods("GET")
  r.HandleFunc("/feedbacks/{id}", controllers.GetFeedback).Methods("GET")
  r.HandleFunc("/feedbacks/{id}", controllers.PutFeedbacks).Methods("PUT")
  r.HandleFunc("/feedbacks/{id}", controllers.DeleteFeedbacks).Methods("DELETE")
  return r
}
