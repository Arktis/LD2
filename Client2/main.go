package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	//"time"

	//"github.com/gorilla/handlers"
	jwt "github.com/dgrijalva/jwt-go"
	//"github.com/codegangsta/negroni"
)

var mySigningKey = []byte("mystring")

// func getBuildings(w http.ResponseWriter, r *http.Request) {
// 	validToken, err := getToken()
// 	if err != nil {
// 		fmt.Println("Failed to generate token")
// 	}

// 	client := &http.Client{}

// 	req, _ := http.NewRequest("GET", "http://localhost:9000/api/buildings", nil)
// 	req.Header.Set("Token", validToken)
// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Fprintf(w, "Error: %s", err.Error())
// 	}

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Fprintf(w, string(body))

// }

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	validToken, err := getToken()
// 	if err != nil {
// 		fmt.Println("Failed to generate token")
// 	}

// 	client := &http.Client{}
// 	req, _ := http.NewRequest("GET", "http://localhost:9000/", nil)
// 	req.Header.Set("Token", validToken)
// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Fprintf(w, "Error: %s", err.Error())
// 	}

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Fprintf(w, string(body))
// }

func getToken() (string, error) {
	signingKey := []byte("mystring")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// "name": name,
		// "role": "redpill",

	})
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

// func GenerateJWT() (string, error) {
// 	token := jwt.New(jwt.SigningMethodHS256)

// 	claims := token.Claims.(jwt.MapClaims)

// 	claims["authorized"] = true
// 	claims["client"] = "LoggedIn User"
// 	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

// 	tokenString, err := token.SignedString(mySigningKey)

// 	if err != nil {
// 		fmt.Errorf("Something Went Wrong: %s", err.Error())
// 		return "", err
// 	}

// 	return tokenString, nil
// }

func handleRequests() {
	//http.HandleFunc("/", getBuildings)
	//handlers.HandleFunc("/api/buildings", getBuildings).Methods("GET")
	//r := mux.NewRouter()
	handler := http.NewServeMux()

	handler.HandleFunc("/api/buildings", func(w http.ResponseWriter, r *http.Request) {
		// validToken, err := getToken()
		// if err != nil {
		// 	fmt.Println("Failed to generate token")
		// }

		client := &http.Client{}

		req, _ := http.NewRequest("GET", "http://localhost:9000/api/buildings", nil)
		// req.Header.Set("Token", validToken)
		res, err := client.Do(req)
		// if err != nil {
		// 	fmt.Fprintf(w, "Error: %s", err.Error())
		//}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(w, string(body))
	})
	handler.HandleFunc("/api/buildings/{id:[1-1000]}", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, string(id3))
		// validToken, err := getToken()
		// if err != nil {
		// 	fmt.Println("Failed to generate token")
		// }
		//C

		client := &http.Client{}

		req, _ := http.NewRequest("GET", "http://localhost:9000/api/buildings/1", nil)

		// req.Header.Set("Token", validToken)
		res, err := client.Do(req)
		// if err != nil {
		// 	fmt.Fprintf(w, "Error: %s", err.Error())
		//}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Fprintf(w, string(body))
	})
	handler.HandleFunc("/api/buildings/2", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, string(id3))
		validToken, err := getToken()
		if err != nil {
			fmt.Println("Failed to generate token")
		}
		//params := mux.Vars(r)
		//fmt.Println(r)
		client := &http.Client{}

		req, _ := http.NewRequest("DELETE", "http://localhost:9000/api/buildings/2", nil)
		req.Header.Set("Token", validToken)
		res, err := client.Do(req)
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err.Error())
		}

		//http.NewRequest("DELETE", "http://localhost:9000/api/buildings/"+params["id"], nil)

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(w, string(body))
	})
	// r.Handle("/api/buildings", authMiddleware(createBuilding)).Methods("POST")
	// r.Handle("/api/buildings/{id3}", authMiddleware(updateBuilding)).Methods("PUT")
	// r.Handle("/api/buildings/{id3}", authMiddleware(deleteBuilding)).Methods("DELETE")

	// //r.HandleFunc("/api/apartaments", getApartaments).Methods("GET")
	// r.HandleFunc("/api/buildings/{id3}/apartaments", getApartaments).Methods("GET")
	// r.HandleFunc("/api/buildings/{id3}/apartaments/{id2}", getApartament).Methods("GET")
	// /* r.HandleFunc("/api/apartaments", createApartament).Methods("POST")
	// r.HandleFunc("/api/apartaments/{id}", updateApartament).Methods("PUT")
	// r.HandleFunc("/api/apartaments/{id}", deleteApartament).Methods("DELETE")
	// */
	// r.Handle("/api/buildings/{id3}/apartaments", createApartament).Methods("POST")
	// r.Handle("/api/buildings/{id3}/apartaments/{id2}", updateApartament).Methods("PUT")
	// r.Handle("/api/buildings/{id3}/apartaments/{id2}", deleteApartament).Methods("DELETE")

	// r.HandleFunc("/api/buildings/{id3}/apartaments/{id2}/contracts", getContracts).Methods("GET")
	// r.HandleFunc("/api/buildings/{id3}/apartaments/{id2}/contracts/{id}", getContract).Methods("GET")
	// r.Handle("/api/buildings/{id3}/apartaments/{id2}/contracts", createContract).Methods("POST")
	// r.Handle("/api/buildings/{id3}/apartaments/{id2}/contracts/{id}", updateContract).Methods("PUT")
	// r.Handle("/api/buildings/{id3}/apartaments/{id2}/contracts/{id}", deleteContract).Methods("DELETE")
	err := http.ListenAndServe(":9001", handler)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
	//log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	fmt.Println("mysimpleclient")
	tokenString, err := getToken()
	if err != nil {
		fmt.Println("Error generating token string")
	}
	fmt.Println(tokenString)
	handleRequests()
}
