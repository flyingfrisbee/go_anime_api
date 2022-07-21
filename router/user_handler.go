package router

import (
	"GithubRepository/go_anime_api/db"
	"GithubRepository/go_anime_api/model"
	"GithubRepository/go_anime_api/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func insertUserTokenHandler(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteResponse(w, nil, "Error when reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var clientTokenRequest model.ClientTokenRequest
	err = json.Unmarshal(jsonBytes, &clientTokenRequest)
	if err != nil {
		utils.WriteResponse(w, nil, "Error when parsing json to struct", http.StatusBadRequest)
		return
	}

	validate := validator.New()
	err = validate.Struct(clientTokenRequest)
	if err != nil {
		utils.WriteResponse(w, nil, err.Error(), http.StatusBadRequest)
		return
	}
	/*need credential to identify row, unfortunately credential means user have to sign up
	  so i have to user raw token lmao*/
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
