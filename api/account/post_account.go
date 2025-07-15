package account

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
)

// @Summary Create Account
// @Description create one new user account
// @Tags account
// @Accept json
// @Produce json
// @Param account body Account true "Account details"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /account [post]
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received on /api/v1/account")
	w.Header().Set("Content-Type", "application/json")

	account := Account{}

	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "JSON invalid"})
		return
	}

	account.Id = uuid.New().String()

	query := `INSERT INTO user_account (id, status, name, age, phone, gender, address)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err = DB.Exec(r.Context(), query,
		account.Id,
		account.Status,
		account.Name,
		account.Age,
		account.Phone,
		account.Gender,
		account.Address,
	)

	if err != nil {
		log.Println("Error inserting into database:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create account"})
		return
	}

	log.Println("Account created successfully:", account)

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(map[string]string{"id": account.Id})
	if err != nil {
		log.Println("Error encoding JSON in response:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}
