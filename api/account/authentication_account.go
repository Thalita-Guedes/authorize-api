package account

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"time"
)

// Secret key to sign the token
var MySignatureJWT = []byte("JWT_SECRET")

// @Summary Authentication Account
// @Description Authenticate user and return a JWT token
// @Tags account
// @Accept json
// @Produce json
// @Param credentials body account.Login true "Login credentials"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /authentication [post]
func AuthenticationAccount(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received on /api/v1/authentication")
	w.Header().Set("Content-Type", "application/json")

	login := Login{}

	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"Error": "JSON invalid"})
		return
	}

	AuthenticationUser(w, r, login.Email, login.Senha)

}

// Helper function to authenticate and generate the token
func AuthenticationUser(w http.ResponseWriter, r *http.Request, email, senha string) {

	query := `SELECT email FROM user_authentication WHERE email=$1 AND senha=$2`

	login := Login{}

	err := DB.QueryRow(r.Context(), query, email, senha).Scan(&login.Email)
	if err != nil {
		log.Println("Authentication failed for email:", email)
		log.Println("Database connection error", err)
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid credentials"})
		return
	}

	log.Println("User authenticated successfully:", login.Email)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": login.Email,
		"exp":   time.Now().Add(9 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(MySignatureJWT)
	if err != nil {
		log.Println("Error signing token:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error internal server"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
		"email": login.Email,
	})
}
