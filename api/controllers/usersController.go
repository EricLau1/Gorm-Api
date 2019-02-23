package controllers

import (
  "net/http"
  "encoding/json"
  "strconv"
  "api/utils"
  "api/models"
  "github.com/gorilla/mux"
)

func PostUsers(w http.ResponseWriter, r *http.Request) {
  body := utils.BodyParser(r)
  var user models.User
  err := json.Unmarshal(body, &user)
  if err != nil {
    utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
    return
  }
  err = models.NewUser(user)
  if err != nil {
    utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
    return
  }
  utils.ToJson(w, "Usuário adicionado com sucesso!", http.StatusCreated)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
  users := models.GetAll(models.USERS)
  utils.ToJson(w, users, http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, _ := strconv.ParseUint(vars["id"], 10, 64)
  user := models.GetById(models.USERS, uint64(id))
  utils.ToJson(w, user, http.StatusOK)
}

func PutUsers(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, _ := strconv.ParseUint(vars["id"], 10, 32)
  body := utils.BodyParser(r)
  var user models.User
  err := json.Unmarshal(body, &user)
  if err != nil {
    utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
    return
  }
  user.Id = uint32(id)
  rows, err := models.UpdateUser(user)
  if err != nil {
    utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
    return
  }
  utils.ToJson(w, rows, http.StatusOK)
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, _ := strconv.ParseUint(vars["id"], 10, 64)
  _, err := models.Delete(models.USERS, uint64(id))
  if err != nil {
    utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
    return
  }
  w.WriteHeader(http.StatusNoContent)
}