package router

import (
	"GithubRepository/go_anime_api/model"
	"GithubRepository/go_anime_api/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func userTokenHandler(w http.ResponseWriter, r *http.Request) {
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

	hashedPwd, err := utils.HashPassword(clientTokenRequest.UserToken)
	if err != nil {
		utils.WriteResponse(w, nil, err.Error(), http.StatusInternalServerError)
		return
	}

	// save hashedPwd to user table
	fmt.Println(hashedPwd)
}
