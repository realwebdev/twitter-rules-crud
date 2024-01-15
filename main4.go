package main

// import (
// 	"encoding/csv"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// )

// // User struct to represent the user data
// type User struct {
// 	Username string `json:"username"`
// 	// ID        string    `json:"id"`
// 	// Name      string    `json:"name"`
// 	// CreatedAt time.Time `json:"created_at"`
// }

// func jsonToCsv(jsonData []byte) ([][]string, error) {
// 	var usersData struct {
// 		Data []User `json:"data"`
// 	}

// 	err := json.Unmarshal(jsonData, &usersData)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var csvData [][]string

// 	// Append headers
// 	csvData = append(csvData, []string{"Username"})

// 	// Append user data
// 	for _, user := range usersData.Data {
// 		// createdAt := user.CreatedAt.Format("2006-01-02 15:04:05") // Format time as per your requirement
// 		row := []string{user.Username}
// 		csvData = append(csvData, row)
// 	}

// 	return csvData, nil
// }

// func main() {
// 	// Open the JSON file
// 	file, err := os.Open("oncologist-user.json")
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	// Read JSON data from the file
// 	jsonData, err := ioutil.ReadAll(file)
// 	if err != nil {
// 		fmt.Println("Error reading JSON data:", err)
// 		return
// 	}

// 	csvData, err := jsonToCsv(jsonData)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	// Write CSV data to a file
// 	csvFile, err := os.Create("oncologist-username.csv")
// 	if err != nil {
// 		fmt.Println("Error creating file:", err)
// 		return
// 	}
// 	defer csvFile.Close()

// 	writer := csv.NewWriter(csvFile)
// 	defer writer.Flush()

// 	for _, row := range csvData {
// 		err := writer.Write(row)
// 		if err != nil {
// 			fmt.Println("Error writing to CSV:", err)
// 			return
// 		}
// 	}

// 	fmt.Println("Conversion successful. CSV data written to output.csv")
// }
