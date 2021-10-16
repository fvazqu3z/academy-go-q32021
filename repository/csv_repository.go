package repository

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"wizeline/common"
	"wizeline/model"

)

type repo struct{

}

//NewRepository method to create a new user repository
func NewRepository() UserRepository {
	return &repo{}
}

//GetUsersFromDataSource read all users from data source
func (r *repo) GetUsersFromDataSource(data [][]string) ([]model.User, error) {
	// convert csv lines to array of structs
	var users []model.User
	for i, line := range data {
		if i == 0 { // omit header line
			continue
		}
		var record model.User
		var readerError error
		record.UserID,readerError = strconv.Atoi(line[0])
		if readerError != nil {
			return nil, errors.New("The UserID must be a numeric value ")
		}
		record.Name = line[1]
		record.Email = line[2]
		record.Phone = line[3]
		users = append(users, record)
	}
	return users, nil
}

//GetUser get a specific user
func (r *repo) GetUser(userID int) (model.User, error) {
	panic("implement me")
}

//SaveUsers save users to csv file
func (r *repo) SaveUsers(data [][]string) (int, error) {
	panic("implement me")
}

//GetUsers retrieve all users
func (r *repo) GetUsers() (string, error){
	// open file
	f, err := os.Open(common.CSVFileNameInput)
	if err != nil {
		log.Println(err)
		return "", errors.New("error opening file " + common.CSVFileNameInput)
	}

	//close the file at the end of the program
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)

	//Read CSV file using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Println(err)
		return "", errors.New("error reading file " + common.CSVFileNameInput)
	}

	userData, err := r.GetUsersFromDataSource(data)
	if err != nil {
		log.Println(err)
		return "", errors.New("error parsing data " + common.CSVFileNameInput)
	}
	jsonData, err := json.MarshalIndent(userData, "", "  ")
	if err != nil {
		log.Println(err)
		return "", errors.New("error transforming csv to json " + common.CSVFileNameInput)
	}
	return string(jsonData), nil
}

//SaveUsersToDataSource method to save user to data source
func (r *repo) SaveUsersToDataSource(users []model.User) error {
	// open file
	f, err := os.Create(common.CSVFileNameOutput)
	if err != nil {
		log.Println(err)
		return errors.New("error opening file " + common.CSVFileNameOutput)
	}
	fmt.Println("File opened")

	//close the file at the end of the program
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)

	var separator string
	var endLine string
	var header string
	separator = ","
	endLine = "\n"

	header = "UserID,Name,Email,Cellphone"
	header += endLine
	_ , err = f.WriteString(header)
	if err != nil {
		log.Println(err)
		return errors.New("error writing headers to file " + common.CSVFileNameOutput)
	}

	for _, user := range users {
		var row string
		row = strconv.Itoa(user.UserID)
		row += separator
		row += user.Name
		row += separator
		row += user.Email
		row += separator
		row += user.Phone
		row += endLine
		_ , err = f.WriteString(row)

		if err != nil {
			log.Println(err)
			return errors.New("error writing " + user.Name + " to file " + common.CSVFileNameOutput)
		}
	}

	err = f.Sync()
	if err != nil {
		log.Println(err)
		return errors.New("error synchronizing the file " + common.CSVFileNameOutput)
	}

	return nil

}