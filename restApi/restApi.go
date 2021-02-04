package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type IOU struct {
	gorm.Model
	Name string `json:"name"`
	Owes map[string]float64 `json:"owes"`
	OwedBy map[string]float64 `json:"owed_by"`
	Balance string `json:"balance"`
}
type IOUs struct {
	Iou []IOU
}
type User struct {
	gorm.Model
	Name string
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
	r.PathPrefix("/api/v1")
	r.HandleFunc("/add", CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/users", ListUsers).Methods(http.MethodGet)
	r.HandleFunc("/iou", IOWEYOU).Methods(http.MethodPost)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func init() {
	//Connect()

}

func IOWEYOU(writer http.ResponseWriter, request *http.Request) {
	DB.AutoMigrate(&IOUPayload{})
	var iouPayload IOUPayload
	var iouPayloads []IOUPayload

	body, _ := ioutil.ReadAll(request.Body)

	json.Unmarshal(body, &iouPayload)
	iou := IOUPayload{
		Lender:   iouPayload.Lender,
		Borrower: iouPayload.Borrower,
		Amount:   iouPayload.Amount,
	}
	DB.Create(&iou)
	DB.Table("iou_payloads").Find(&iouPayloads)
	writer.Header().Set("Content-Type", "application/json")
	//gerneat ledger
	json.Marshal(&iouPayloads)
	fmt.Print(iouPayloads)
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

	DB.Table("users").Where("name IN ?", users.Users).Order("name desc").Find(&user)

	DB.Table("iou_payloads").Where("borrower IN ?", users.Users).Or("lender IN ?", users.Users).Find(&iouPayloads)
	a:=fetchUserLedger(iouPayloads, user)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(a)
}

func CreateUser(writer http.ResponseWriter, request *http.Request) {
	var user User

	DB.AutoMigrate(&User{})

	body, _ := ioutil.ReadAll(request.Body)

	json.Unmarshal(body, &user)

	row := DB.Create(&User{
		gorm.Model{},
		user.Name,
	})
	if row.RowsAffected > 0 {
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(&user)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(map[int]string{400: "User was not created"})
	}
}

func fetchUserLedger(a []IOUPayload, filteredUsers []User) []IOU {
	var user []User
	var userObject = []IOU{}
	user = filteredUsers
	if filteredUsers  == nil {
		DB.Find(&user)
	}
	for i, _ := range user {
		var userLedger IOU
		owees:= make(map[string]float64)
		owed:= make(map[string]float64)
		owedamount:=0.0
		oweeamount:=0.0
		for k, _ := range a {
			if a[k].Lender == user[i].Name {
				borrower :=a[k].Borrower
				oweeamount=oweeamount + a[k].Amount
				owed[borrower]=oweeamount
			} else if a[k].Borrower == user[i].Name {
				lender :=a[k].Lender
				owedamount=owedamount + a[k].Amount
				owees[lender]= owedamount
			}



		}
		userLedger = IOU{
			Name:    user[i].Name,
			Owes:    owees,
			OwedBy:  owed,
			Balance: "",
		}

		fmt.Print("aye",owees)
		userObject = append(userObject, userLedger)
	}
	//fmt.Printf("2**%d = %d\n", i, v)

	return userObject
}
