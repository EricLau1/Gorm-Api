package controllers

import (
  "net/http"
  "encoding/json"
  "strconv"
  "api/utils"
  "api/models"
  "github.com/gorilla/mux"
)

func PostPosts(w http.ResponseWriter, r *http.Request) {
  body := utils.BodyParser(r)
  var post models.Post
  err := json.Unmarshal(body, &post)
  if err != nil {
    utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
    return
  }
  err = models.NewPost(post)
  if err != nil {
    utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
    return
  }
  utils.ToJson(w, "Post adicionado com sucesso!", http.StatusCreated)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
  posts := models.GetPosts()
  utils.ToJson(w, posts, http.StatusOK)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, _ := strconv.ParseUint(vars["id"], 10, 64)
  post := models.GetById(models.POSTS, uint64(id))
  utils.ToJson(w, post, http.StatusOK)
}

func PutPosts(w http.ResponseWriter, r *http.Request) {
  body := utils.BodyParser(r)
  vars := mux.Vars(r)
  id, _ := strconv.ParseUint(vars["id"], 10, 64)
  var post models.Post
  err := json.Unmarshal(body, &post)
  if err != nil {
    utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
    return
  }
  post.Id = uint32(id)
  rows, err := models.UpdatePost(post)
  if err != nil {
    utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
    return
  }
  utils.ToJson(w, rows, http.StatusCreated)
}

func DeletePosts(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, _ := strconv.ParseUint(vars["id"], 10, 64)
  _, err := models.Delete(models.POSTS, uint64(id))
  if err != nil {
    utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
    return
  }
  w.WriteHeader(http.StatusNoContent)
}