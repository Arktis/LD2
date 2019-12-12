package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"encoding/json"
	"os"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var mySigningKey = []byte("captainjacksparrowsayshi")

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Authorized sees this")
	fmt.Println("Endpoint Hit: homePage")
}

type Building struct {
	ID      string `json:"buildingId"`
	Address string `json:"address"`
	//Apartaments []Apartament
}

type Apartament struct {
	ID         string `json:"apartamentID"`
	BuildingID string `json:"buildingID"`
	Num        string `json:"apartamentNumber"`
	//Contracts  Contract
	//contracts []Contract
}

type Contract struct {
	ID           string `json:"contractID"`
	ApartamentID string `json:"apartamentID"`
	BuildingID   string `json:"buildingID"`
	StartDate    string `json:"startDate"`
	EndDate      string `json:"endDate"`
}

//Init books var as a slice(variable lenght array) Book struct
//var books []Book
var buildings []Building
var apartaments []Apartament
var contracts []Contract
var buildingID = 2
var apartamentID = 4
var contractID = 4
var apartamentCount = 0
var contractCount = 0

//var mySigningKey = []byte("supersecure")

// func GenerateJWT() (string error) {
// 	token := jwt.New(jwt.SigningMethodHS256)

// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["authorized"] = true
// 	claims["User"] = "Arturas Naujokas"
// 	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

// 	tokenString, err := token.SignedString(mySigningKey)

// 	if err != nil {
// 		fmt.Errorf("Something went wrong: %s", err.Error())
// 		return "", err
// 	}
// 	return tokenString, nil
// }
// func jwtGen(response http.ResponseWriter, Request *hettp.Request) {

// }

func returnCode200(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

//created successfully
func returnCode201(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}

//deleted successfully
func returnCode204(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
func returnCode400(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
}

func returnCode404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

//Get all buildings
//var getBuildings = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
func getBuildings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if len(buildings) == 0 {
		returnCode404(w, r)
	} else {
		returnCode200(w, r)
		json.NewEncoder(w).Encode(buildings)
	}

} //)

func getBuilding(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params
	//Loop Through books and find with id
	for _, item := range buildings {
		if item.ID == params["id3"] {
			returnCode200(w, r)
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	returnCode404(w, r)
	//json.NewEncoder(w).Encode(&Building{})
}

// Create a Building
//func createBuilding(w http.ResponseWriter, r *http.Request) {
var createBuilding = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var building Building
	_ = json.NewDecoder(r.Body).Decode(&building)
	buildingID = buildingID + 1
	building.ID = strconv.Itoa(buildingID)

	if building.Address != "" {
		returnCode200(w, r)
		buildings = append(buildings, building)
		json.NewEncoder(w).Encode(building)
	} else {
		returnCode400(w, r)
	}

})

//func updateBuilding(w http.ResponseWriter, r *http.Request) {
var updateBuilding = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	found := false
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range buildings {
		if item.ID == params["id3"] {

			var building Building
			_ = json.NewDecoder(r.Body).Decode(&building)
			building.ID = params["id3"]

			if building.Address != "" {
				buildings = append(buildings[:index], buildings[index+1:]...)
				buildings = append(buildings, building)
				returnCode200(w, r)
				json.NewEncoder(w).Encode(building)
			} else {
				returnCode400(w, r)
			}

			found = true
			return
		}
	}
	if found == false {
		returnCode404(w, r)
	}
	json.NewEncoder(w).Encode(buildings)
})

//func deleteBuilding(w http.ResponseWriter, r *http.Request) {
var deleteBuilding = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	found := false
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //for building's id
	for index, item := range buildings {
		if item.ID == params["id3"] {
			found = true
			returnCode204(w, r)
			buildings = append(buildings[:index], buildings[index+1:]...)
			break
		}
	}

	for {
		found2 := false
		for index2, item2 := range apartaments {
			if item2.BuildingID == params["id3"] {
				apartaments = append(apartaments[:index2], apartaments[index2+1:]...)
				found2 = true
				break
			}
		}
		if !found2 {
			break
		}
	}

	for {
		found2 := false
		for index3, item3 := range contracts {
			if item3.BuildingID == params["id3"] {
				contracts = append(contracts[:index3], contracts[index3+1:]...)
				found2 = true
				break
			}
		}
		if !found2 {
			break
		}
	}

	if found == false {
		returnCode404(w, r)
	}
	json.NewEncoder(w).Encode(buildings)

})

func getApartaments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //for building's id
	var apartamentsReturn []Apartament
	found := 0
	//var apartamentsReturn []Apartament;
	if len(apartaments) == 0 {
		returnCode404(w, r)
	} else {

		for _, item := range apartaments {
			// 	//removedApartaments := 0
			if item.BuildingID == params["id3"] {
				apartamentsReturn = append(apartamentsReturn, item)
				found++
			}
		}
		if found > 0 {
			returnCode200(w, r)
			json.NewEncoder(w).Encode(apartamentsReturn)
		} else {
			returnCode404(w, r)
		}

	}
}

func getApartament(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params
	found := false
	//Loop Through apartament ids
	for _, item2 := range apartaments {
		// 	//removedApartaments := 0
		if item2.BuildingID == params["id3"] && item2.ID == params["id2"] {
			returnCode200(w, r)
			found = true
			json.NewEncoder(w).Encode(item2)
			return

		}
	}
	if found == false {
		returnCode404(w, r)
	} else {
		returnCode200(w, r)
		json.NewEncoder(w).Encode(&Apartament{})
	}
}

// Create a Apartament
var createApartament = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//func createApartament(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var apartament Apartament
	_ = json.NewDecoder(r.Body).Decode(&apartament)
	apartamentID = apartamentID + 1
	apartament.ID = strconv.Itoa(apartamentID)
	apartaments = append(apartaments, apartament)
	params := mux.Vars(r)
	apartament.BuildingID = params["id3"]
	if apartament.Num != "" {
		returnCode200(w, r)
		json.NewEncoder(w).Encode(apartament)
	} else {
		returnCode400(w, r)
	}
})

var updateApartament = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//func updateApartament(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	found := false
	params := mux.Vars(r)
	var apartament Apartament
	for index, item := range apartaments {
		if item.ID == params["id2"] {
			apartaments = append(apartaments[:index], apartaments[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&apartament)

			apartaments = append(apartaments, apartament)
			apartament.ID = params["id2"]
			apartament.BuildingID = params["id3"]
			found = true
			json.NewEncoder(w).Encode(apartament)
			return
		}
	}
	if found == false {
		returnCode404(w, r)
	} else if apartament.BuildingID != "" && apartament.Num != "" {
		returnCode200(w, r)
	} else {
		returnCode400(w, r)
	}
	json.NewEncoder(w).Encode(apartaments)
})

var deleteApartament = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//func deleteApartament(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	found := false
	for index, item := range apartaments {
		if item.ID == params["id2"] && item.BuildingID == params["id3"] {
			found = true
			apartaments = append(apartaments[:index], apartaments[index+1:]...)
			break
		}
	}
	for {
		found2 := false
		for index2, item2 := range contracts {
			if item2.ApartamentID == params["id2"] && item2.BuildingID == params["id3"] {
				contracts = append(contracts[:index2], contracts[index2+1:]...)
				found2 = true
				break
			}
		}
		if !found2 {
			break
		}
	}

	// for index2, item2 := range contracts {
	// 	if item2.ApartamentID == params["id2"] && item2.BuildingID == params["id3"] {
	// 		contracts = append(contracts[:index2], contracts[index2+1:]...)
	// 	}
	// }

	if found == false {
		returnCode404(w, r)
	} else {
		returnCode200(w, r)
	}
	json.NewEncoder(w).Encode(apartaments)
})

func getContracts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	found := false
	if len(contracts) == 0 {
		returnCode404(w, r)
	} else {
		var contractsReturn []Contract
		params := mux.Vars(r)
		//for _, item := range buildings {
		// 	//removedApartaments := 0
		//if item.ID == params["id3"] {
		for _, item2 := range contracts {
			if item2.BuildingID == params["id3"] && item2.ApartamentID == params["id2"] {
				contractsReturn = append(contractsReturn, item2)
				found = true
			}
		}
		if found == false {
			returnCode404(w, r)
		} else {
			returnCode200(w, r)
			json.NewEncoder(w).Encode(contractsReturn)
		}
	}
}

func getContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params
	found := false
	//Loop Through apartament ids
	for _, item := range contracts {
		if item.ID == params["id"] && item.ApartamentID == params["id2"] && item.BuildingID == params["id3"] {
			found = true
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	if found == false {
		returnCode404(w, r)
	} else {
		returnCode200(w, r)
		json.NewEncoder(w).Encode(&Contract{})
	}
}

// Create a Apartament
var createContract = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//func createContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var contract Contract
	_ = json.NewDecoder(r.Body).Decode(&contract)
	contractID = apartamentID + 1
	contract.ID = strconv.Itoa(contractID)

	params := mux.Vars(r) //Get params
	contract.BuildingID = params["id3"]
	contract.ApartamentID = params["id2"]
	if contract.EndDate != "" && contract.StartDate != "" {
		returnCode200(w, r)
		contracts = append(contracts, contract)
		json.NewEncoder(w).Encode(contract)
	} else {
		returnCode400(w, r)
	}
})

var updateContract = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//func updateContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	found := false
	var contract Contract
	for index, item := range contracts {
		if item.ID == params["id"] && item.BuildingID == params["id3"] && item.ApartamentID == params["id2"] {
			found = true
			contracts = append(contracts[:index], contracts[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&contract)
			contract.ID = params["id"]
			contract.BuildingID = params["id3"]
			contract.ApartamentID = params["id2"]
			contracts = append(contracts, contract)
			json.NewEncoder(w).Encode(contract)
			return
		}
	}
	if found == false {
		returnCode404(w, r)
	} else if contract.EndDate != "" && contract.StartDate != "" {
		returnCode200(w, r)
	} else {
		returnCode400(w, r)
	}
	json.NewEncoder(w).Encode(contracts)
})

var deleteContract = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//func deleteContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	found := false
	for index, item := range contracts {
		if item.ID == params["id"] && item.BuildingID == params["id3"] && item.ApartamentID == params["id2"] {
			found = true
			contracts = append(contracts[:index], contracts[index+1:]...)
			break
		}
	}
	if found == false {
		returnCode404(w, r)
	} else {
		returnCode200(w, r)
	}
	json.NewEncoder(w).Encode(contracts)
})

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Token")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization Header"))
			return
		}
		// protection against code reading?
		tokenString = strings.Replace(tokenString, "probably checks the beginings of tokens" /*"Bearer*/, "", 1)
		_, err := verifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}
		//name := claims.(jwt.MapClaims)["name"].(string)
		//role := claims.(jwt.MapClaims)["role"].(string)

		//r.Header.Set("name", name)
		//r.Header.Set("role", role)

		next.ServeHTTP(w, r)
	})
}
func getToken() (string, error) {
	signingKey := []byte("mystring")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// "name": name,
		// "role": "redpill",

	})
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

func verifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte("mystring")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}

//Jeigu pasiekia endpoint'a, vykdo dekoruota funkcija
// func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		if r.Header["Token"] != nil {

// 			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
// 				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 					return nil, fmt.Errorf("There was an error")
// 				}
// 				return mySigningKey, nil
// 			})

// 			if err != nil {
// 				fmt.Fprintf(w, err.Error())
// 			}
// 			//Jeigu pasiekia endpointa, tai autorizuotas
// 			if token.Valid {
// 				endpoint(w, r)
// 			}
// 		} else {
// 			fmt.Fprintf(w, "Not Authorized")
// 		}
// 	})
// }

func ConfigureRouter() *mux.Router {
	r := mux.NewRouter()

	//router.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r.HandleFunc("/api/buildings", getBuildings).Methods("GET")
	r.HandleFunc("/api/buildings/{id3}", getBuilding).Methods("GET")
	r.Handle("/api/buildings", authMiddleware(createBuilding)).Methods("POST")
	r.Handle("/api/buildings/{id3}", authMiddleware(updateBuilding)).Methods("PUT")
	r.Handle("/api/buildings/{id3}", authMiddleware(deleteBuilding)).Methods("DELETE")

	//r.HandleFunc("/api/apartaments", getApartaments).Methods("GET")
	r.HandleFunc("/api/buildings/{id3}/apartaments", getApartaments).Methods("GET")
	r.HandleFunc("/api/buildings/{id3}/apartaments/{id2}", getApartament).Methods("GET")
	/* r.HandleFunc("/api/apartaments", createApartament).Methods("POST")
	r.HandleFunc("/api/apartaments/{id}", updateApartament).Methods("PUT")
	r.HandleFunc("/api/apartaments/{id}", deleteApartament).Methods("DELETE")
	*/
	r.Handle("/api/buildings/{id3}/apartaments", createApartament).Methods("POST")
	r.Handle("/api/buildings/{id3}/apartaments/{id2}", updateApartament).Methods("PUT")
	r.Handle("/api/buildings/{id3}/apartaments/{id2}", deleteApartament).Methods("DELETE")

	r.HandleFunc("/api/buildings/{id3}/apartaments/{id2}/contracts", getContracts).Methods("GET")
	r.HandleFunc("/api/buildings/{id3}/apartaments/{id2}/contracts/{id}", getContract).Methods("GET")
	r.Handle("/api/buildings/{id3}/apartaments/{id2}/contracts", createContract).Methods("POST")
	r.Handle("/api/buildings/{id3}/apartaments/{id2}/contracts/{id}", updateContract).Methods("PUT")
	r.Handle("/api/buildings/{id3}/apartaments/{id2}/contracts/{id}", deleteContract).Methods("DELETE")

	return r
}

func handleRequests() {
	//r := mux.NewRouter()

	//Mock data
	var apartamentsB1 []Apartament
	var apartamentsB2 []Apartament
	var contractsA11 Contract
	var contractsA12 Contract
	var contractsA21 Contract
	var contractsA22 Contract
	contractsA11.ID = "1"
	contractsA11.StartDate = "2021-10-20"
	contractsA11.ApartamentID = "1"
	contractsA11.EndDate = "2022-10-20"
	contractsA11.BuildingID = "1"
	contractsA12.ID = "2"
	contractsA12.ApartamentID = "2"
	contractsA12.StartDate = "2025-10-20"
	contractsA12.EndDate = "2026-10-20"
	contractsA12.BuildingID = "1"
	contractsA21.ID = "3"
	contractsA21.ApartamentID = "3"
	contractsA21.StartDate = "2021-10-20"
	contractsA21.EndDate = "2027-10-20"
	contractsA21.BuildingID = "2"
	contractsA22.ID = "4"
	contractsA22.ApartamentID = "4"
	contractsA22.StartDate = "2077-10-20"
	contractsA22.EndDate = "2030-10-20"
	contractsA22.BuildingID = "2"
	apartamentsB1 = append(apartamentsB1, Apartament{ID: "1", BuildingID: "1", Num: "545" /* , Contracts: contractsA11 */})
	apartamentsB1 = append(apartamentsB1, Apartament{ID: "2", BuildingID: "1", Num: "245" /* , Contracts: contractsA12 */})
	apartamentsB2 = append(apartamentsB2, Apartament{ID: "3", BuildingID: "2", Num: "3456" /* , Contracts: contractsA21 */})
	apartamentsB2 = append(apartamentsB2, Apartament{ID: "4", BuildingID: "2", Num: "4987" /* , Contracts: contractsA22 */})
	apartaments = append(apartaments, apartamentsB1[0])
	apartaments = append(apartaments, apartamentsB1[1])
	apartaments = append(apartaments, apartamentsB2[0])
	apartaments = append(apartaments, apartamentsB2[1])
	contracts = append(contracts, contractsA11)
	contracts = append(contracts, contractsA12)
	contracts = append(contracts, contractsA21)
	contracts = append(contracts, contractsA22)

	buildings = append(buildings, Building{ID: "1", Address: "gatve-1" /* , Apartaments: apartamentsB1 */})
	buildings = append(buildings, Building{ID: "2", Address: "gatve-2" /* , Apartaments: apartamentsB2 */})

	// r.HandleFunc("/api/buildings/{id3}", getBuilding).Methods("GET")
	// r.HandleFunc("/api/buildings", createBuilding).Methods("POST")
	// r.HandleFunc("/api/buildings/{id3}", updateBuilding).Methods("PUT")
	// r.HandleFunc("/api/buildings/{id3}", deleteBuilding).Methods("DELETE")

	// //r.HandleFunc("/api/apartaments", getApartaments).Methods("GET")
	// r.HandleFunc("/api/buildings/{id3}/apartaments", getApartaments).Methods("GET")
	// r.HandleFunc("/api/buildings/{id3}/apartaments/{id2}", getApartament).Methods("GET")
	// /* r.HandleFunc("/api/apartaments", createApartament).Methods("POST")
	// r.HandleFunc("/api/apartaments/{id}", updateApartament).Methods("PUT")
	// r.HandleFunc("/api/apartaments/{id}", deleteApartament).Methods("DELETE")
	// */
	// r.HandleFunc("/api/buildings/{id3}/apartaments", createApartament).Methods("POST")
	// r.HandleFunc("/api/buildings/{id3}/apartaments/{id2}", updateApartament).Methods("PUT")
	// r.HandleFunc("/api/buildings/{id3}/apartaments/{id2}", deleteApartament).Methods("DELETE")

	// r.HandleFunc("/api/buildings/{id3}/apartaments/{id2}/contracts", getContracts).Methods("GET")
	// r.HandleFunc("/api/buildings/{id3}/apartaments/{id2}/contracts/{id}", getContract).Methods("GET")
	// r.HandleFunc("/api/buildings/{id3}/apartaments/{id2}/contracts", createContract).Methods("POST")
	// r.HandleFunc("/api/buildings/{id3}/apartaments/{id2}/contracts/{id}", updateContract).Methods("PUT")
	// r.HandleFunc("/api/buildings/{id3}/apartaments/{id2}/contracts/{id}", deleteContract).Methods("DELETE")
	//http.Handle("/", isAuthorized(homePage))
	router := ConfigureRouter()
	log.Fatal(http.ListenAndServe(":9000", handlers.LoggingHandler(os.Stdout, router)))
	//log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	tokenString, err := getToken()
	if err != nil {
		fmt.Println("Error generating token string")
	}
	fmt.Println(tokenString)

	handleRequests()
}
