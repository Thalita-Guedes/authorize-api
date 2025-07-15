package account

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// @Summary Search account by Id
// @Description Returns data for a specific account by ID
// @Tags account
// @Accept json
// @Produce json
// @Param id path string true "ID of the account"
// @Success 200 {object} Account
// @Failure 500 {object} map[string]string
// @Router /account/{id} [get]
func SearchAccountByID(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received on /api/v1/account/{id}")
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]
	log.Println("ID received:", id)

	account := Account{}

	query := `SELECT id, status, name, age, phone, gender,address FROM user_account WHERE id=$1`

	err := DB.QueryRow(r.Context(), query, id).Scan(
		&account.Id,
		&account.Status,
		&account.Name,
		&account.Age,
		&account.Phone,
		&account.Gender,
		&account.Address,
	)

	if err != nil {
		log.Println("error fetching database:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error internal server"})
		return
	}

	err = json.NewEncoder(w).Encode(account)
	if err != nil {
		log.Println("Error encoding JSON in response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error internal server"})
		return
	}

}
