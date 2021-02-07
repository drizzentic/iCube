package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type IOU struct {
	gorm.Model `json:"-"`
	Name       string             `json:"name"`
	Owes       map[string]float64 `json:"owes"`
	OwedBy     map[string]float64 `json:"owed_by"`
	Balance    float64            `json:"balance"`
}
type IOUs struct {
	Iou []IOU
}
type User struct {
	gorm.Model
	Name string `json:"user" gorm:"index:idx_name,unique"`
}
type Users struct {
	Users []string `json:"users"`
}

type IOUPayload struct {
	Lender   string  `json:"lender"`
	Borrower string  `json:"borrower"`
	Amount   float64 `json:"amount"`
}

var DB *gorm.DB

func main() {

	DB = Connect()
	r := mux.NewRouter()
	// Urls for accessing the rest api
	r.HandleFunc("/add", CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/users", ListUsers).Methods(http.MethodGet)
	r.HandleFunc("/iou", IOWEYOU).Methods(http.MethodPost)

	// Starts the http server

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	// Listen for http output

	log.Fatal(srv.ListenAndServe())
}

func IOWEYOU(writer http.ResponseWriter, request *http.Request) {

	var iouPayload IOUPayload
	var iouPayloads []IOUPayload
	var user []User

	body, _ := ioutil.ReadAll(request.Body)

	// Read and parse the request JSON

	json.Unmarshal(body, &iouPayload)

	// Create an IOU
	iou := IOUPayload{
		Lender:   iouPayload.Lender,
		Borrower: iouPayload.Borrower,
		Amount:   iouPayload.Amount,
	}

	// Check if the users exists in the database before creating an iou

	rows := DB.Table("users").Where("name", iou.Borrower).Or("name", iou.Lender).Find(&user)

	if rows.RowsAffected < 2 {
		json.NewEncoder(writer).Encode("You cannot create i owe you for users not in system")
		return
	}
	DB.Create(&iou)

	DB.Table("iou_payloads").Find(&iouPayloads)
	writer.Header().Set("Content-Type", "application/json")

	json.Marshal(&iouPayloads)

	// Fetch the user object for the users
	a := fetchUserLedger(iouPayloads, nil)
	json.NewEncoder(writer).Encode(a)
}

func ListUsers(writer http.ResponseWriter, request *http.Request) {
	var users Users
	var user []User
	var iouPayloads []IOUPayload
	//var iou IOU;

	body, _ := ioutil.ReadAll(request.Body)

	json.Unmarshal(body, &users)
	// Retrieve Users from DB

	DB.Table("users").Where("name IN ?", users.Users).Order("name desc").Find(&user)

	// Retrieve IOU payloads for the users
	DB.Table("iou_payloads").Where("borrower IN ?", users.Users).Or("lender IN ?", users.Users).Find(&iouPayloads)

	// Fetch the user object for the users parsed
	a := fetchUserLedger(iouPayloads, user)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(a)
}

func CreateUser(writer http.ResponseWriter, request *http.Request) {
	var user User
	var iouPayloads []IOUPayload
	// Create the tables on first instance
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&IOUPayload{})

	body, _ := ioutil.ReadAll(request.Body)

	json.Unmarshal(body, &user)
	//Return nil when no user object is passed
	if user.Name == "" {
		json.NewEncoder(writer).Encode(nil)
		return
	}
	row := DB.Create(&User{
		gorm.Model{},
		user.Name,
	})
	if row.RowsAffected > 0 {
		var users []User
		b := append(users, user)
		// Retrieve the initial user object for the created user
		DB.Table("iou_payloads").Where("borrower", user.Name).Or("lender", user.Name).Find(&iouPayloads)
		// Fetch the user object for the users parsed// Fetch the user object for the users parsed
		a := fetchUserLedger(iouPayloads, b)
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(a)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(nil)
	}
}

func fetchUserLedger(a []IOUPayload, filteredUsers []User) []IOU {
	var user []User
	var userObject = []IOU{}
	user = filteredUsers
	if filteredUsers == nil {
		DB.Find(&user)
	}
	for i, _ := range user {
		// Generate the User object
		var userLedger IOU
		owees := make(map[string]float64)
		owed := make(map[string]float64)
		owedamount := 0.0
		oweeamount := 0.0
		balance := 0.0
		for k, _ := range a {
			if a[k].Lender == user[i].Name {
				borrower := a[k].Borrower
				oweeamount = oweeamount + a[k].Amount
				balance = balance - owedamount
				owed[borrower] = oweeamount
			} else if a[k].Borrower == user[i].Name {
				lender := a[k].Lender
				owedamount = owedamount + a[k].Amount
				balance = balance + oweeamount
				owees[lender] = owedamount
			}
		}
		userLedger = IOU{
			Name:    user[i].Name,
			Owes:    owees,
			OwedBy:  owed,
			Balance: balancesSum(owed) - balancesSum(owees),
		}
		userObject = append(userObject, userLedger)
	}
	return userObject
}

func balancesSum(owed map[string]float64) float64 {
	balance := 0.0
	for i, _ := range owed {
		balance += owed[i]
	}
	return balance
}
