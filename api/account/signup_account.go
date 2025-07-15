package account

import (
	"encoding/json"
	"log"
	"net/http"
)

// @Summary Create New User 
// @Description Creating a new user record in the database
// @Tags account
// @Accept json
// @Produce json
// @Param account body Login true "Login details"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /signup [post]
func SignupAccount(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received on /api/v1/signup")
	w.Header().Set("Content-Type", "application/json")

	login := Login{}

	err:= json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		log.Println("Error decoding JSON", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string {"error":"JSON invalid"})
		return
	}

	query := `INSERT INTO user_authentication (email, senha) VALUES ($1,$2)`

	_, err = DB.Exec(r.Context(), query, login.Email, login.Senha)

	if err != nil{
		log.Println("Error inserting into database:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create account"})
		return
	}

	log.Println("New user registered:", login.Email)

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(map[string]string{"email": login.Email})
	if err != nil {
		log.Println("Error encoding JSON in response:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	}
