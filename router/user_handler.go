package router

import (
	"GithubRepository/go_anime_api/db"
	"GithubRepository/go_anime_api/model"
	"GithubRepository/go_anime_api/utils"
	"net/http"
)

func insertUserTokenHandler(w http.ResponseWriter, r *http.Request) {
	var clientTokenRequest model.Token
	err := utils.ParseRequestBody(r, &clientTokenRequest)
	if err != nil {
		utils.WriteResponse(w, nil, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	/*
		need credential to identify row, unfortunately credential means user have to sign up
		so i have to user raw token lmao
	*/
	// hashedToken, err := utils.HashPassword(clientTokenRequest.UserToken)
	// if err != nil {
	// 	utils.WriteResponse(w, nil, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	insertedRows, err := db.SaveUserToken(clientTokenRequest.UserToken)
	if err != nil {
		utils.WriteResponse(w, nil, err.Error(), http.StatusInternalServerError)
		return
	}

	if insertedRows < 1 {
		utils.WriteResponse(w, nil, "Token already exist", http.StatusOK)
		return
	}

	utils.WriteResponse(w, nil, "Success inserting user token", http.StatusOK)
}
