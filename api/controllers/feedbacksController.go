package controllers

import (
  "net/http"
  "encoding/json"
  "strconv"
  "api/utils"
  "api/models"
  "github.com/gorilla/mux"
)

func PostFeedbacks(w http.ResponseWriter, r *http.Request) {
  body := utils.BodyParser(r)
  var feedback models.Feedback
  err := json.Unmarshal(body, &feedback)
  if err != nil {
    utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
    return
  }
  err = models.NewFeedback(feedback)
  if err != nil {
    utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
    return
  }
  utils.ToJson(w, "Você comentou neste post!", http.StatusCreated)
}

func GetFeedbacks(w http.ResponseWriter, r *http.Request) {
  feedbacks := models.GetFeedbacks()
  utils.ToJson(w, feedbacks, http.StatusOK)
}

func GetFeedback(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, _ := strconv.ParseUint(vars["id"], 10, 64)
  feedback := models.GetById(models.FEEDBACKS, uint64(id))
  utils.ToJson(w, feedback, http.StatusOK)
}

func PutFeedbacks(w http.ResponseWriter, r *http.Request) {
  body := utils.BodyParser(r)
  vars := mux.Vars(r)
  id, _ := strconv.ParseUint(vars["id"], 10, 64)
  var feedback models.Feedback
  err := json.Unmarshal(body, &feedback)
  if err != nil {
    utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
    return
  }
  feedback.Id = uint64(id)
  rows, err := models.UpdateFeedback(feedback)
  if err != nil {
    utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
    return
  }
  utils.ToJson(w, rows, http.StatusOK)
}

func DeleteFeedbacks(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, _ := strconv.ParseUint(vars["id"], 10, 64)
  _, err := models.Delete(models.FEEDBACKS, uint64(id))
  if err != nil {
    utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
    return
  }
  w.WriteHeader(http.StatusNoContent)
}