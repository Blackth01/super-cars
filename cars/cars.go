package cars

import (
	"supercars/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type car struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Year uint32 `json:"year"`
	Price float64 `json:"price"`
	Status uint8 `json:"status"`
}

// Insert a car in the database
func CreateCar(w http.ResponseWriter, r *http.Request) {
	bodyReq, error := ioutil.ReadAll(r.Body)
	if error != nil {
		w.Write([]byte("Error while reading request body"))
		return
	}

	var car car
	if error = json.Unmarshal(bodyReq, &car); error != nil {
		fmt.Println("Error: ",error)
		w.Write([]byte("Error while converting car to struct"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		w.Write([]byte("Error while connecting to the database!"))
		return
	}
	defer db.Close()

	statement, error := db.Prepare("insert into cars (car_name, car_year, car_price) values (?, ?, ?)")
	if error != nil {
		fmt.Println("Error: ",error)
		w.Write([]byte("Error while creating statement!"))
		return
	}
	defer statement.Close()

	insert, error := statement.Exec(car.Name, car.Year, car.Price)
	if error != nil {
		fmt.Println("Error: ",error)
		w.Write([]byte("Error while executing statement!"))
		return
	}

	idInsert, error := insert.LastInsertId()
	if error != nil {
		fmt.Println("Error: ",error)
		w.Write([]byte("Error while getting insert id!"))
		return
	}

	// STATUS CODES

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Congratulations, you're extremely awesome, believe in yourself! Look at what you accomplished: the car was sucessfully created! Id: %d", idInsert)))
}

// Get all cars in the database
func GetCars(w http.ResponseWriter, r *http.Request) {
	db, error := database.Connect()
	if error != nil {
		fmt.Println("Error: ",error)
		w.Write([]byte("Error while connecting to the database!"))
		return
	}
	defer db.Close()

	lines, error := db.Query("select * from cars where car_status=1")
	if error != nil {
		fmt.Println("Error: ",error)
		w.Write([]byte("Error while getting cars"))
		return
	}
	defer lines.Close()

	var cars []car
	for lines.Next() {
		var car car

		if error := lines.Scan(&car.ID, &car.Name, &car.Year, &car.Price, &car.Status); error != nil {
			fmt.Println("Error: ",error)
			w.Write([]byte("Error while scanning car"))
			return
		}

		cars = append(cars, car)
	}

	w.WriteHeader(http.StatusOK)
	if error := json.NewEncoder(w).Encode(cars); error != nil {
		fmt.Println("Error: ",error)
		w.Write([]byte("Error while converting users to JSON"))
		return
	}
}

// Get specific car
func GetCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, error := strconv.ParseUint(params["id"], 10, 32)
	if error != nil {
		w.Write([]byte("Error while converting param to int"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		w.Write([]byte("Error while connecting to the database!"))
		return
	}

	line, error := db.Query("select * from cars where car_id = ?", ID)
	if error != nil {
		w.Write([]byte("Error while getting the car!"))
		return
	}

	var car car
	if line.Next() {
		if error := line.Scan(&car.ID, &car.Name, &car.Year, &car.Price, &car.Status); error != nil {
			w.Write([]byte("Error while scanning car"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if error := json.NewEncoder(w).Encode(car); error != nil {
		w.Write([]byte("Error while converting car to JSON!"))
		return
	}
}

// Update a car from the database
func UpdateCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, error := strconv.ParseUint(params["id"], 10, 32)
	if error != nil {
		w.Write([]byte("Error while converting the param to integer!"))
		return
	}

	bodyReq, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Error while reading request body!"))
		return
	}

	var car car
	if erro := json.Unmarshal(bodyReq, &car); erro != nil {
		w.Write([]byte("Error while converting car to struct"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		w.Write([]byte("Error while connecting to the database!"))
		return
	}
	defer db.Close()

	statement, error := db.Prepare("update cars set car_name = ?, car_year = ?, car_price = ? where car_id = ?")
	if error != nil {
		w.Write([]byte("Error while creating statement!"))
		return
	}
	defer statement.Close()

	if _, error := statement.Exec(car.Name, car.Year, car.Price, ID); error != nil {
		w.Write([]byte("Error while updating car!"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Congratulations, you're extremely awesome, believe in yourself! Look at what you accomplished: the car was sucessfully edited!")))
}

// Remove a car from the database
func DeleteCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, error := strconv.ParseUint(params["id"], 10, 32)
	if error != nil {
		w.Write([]byte("Error while converting the param to int!"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		w.Write([]byte("Error while connecting to the database!"))
		return
	}
	defer db.Close()

	statement, error := db.Prepare("update cars set car_status=0 where car_id = ?")
	if error != nil {
		w.Write([]byte("Error while creating statement!"))
		return
	}
	defer statement.Close()

	if _, error := statement.Exec(ID); error != nil {
		w.Write([]byte("Error while deleting car!"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Congratulations, you're extremely awesome, believe in yourself! Look at what you accomplished: the car %d was sucessfully deleted!", ID)))
}