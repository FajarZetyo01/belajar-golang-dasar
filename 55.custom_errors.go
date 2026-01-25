package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validation Error"}
	}
	if id != "fajar" {
		return &notFoundError{"Data not found"}
	}
	//ok
	return nil
}

func main() {
	err := SaveData("fajar", nil)

	if err != nil {
		if validationErr, ok := err.(*validationError); ok {
			fmt.Println("Validation error : ", validationErr.Error())
		} else if notFoundErr, ok := err.(*notFoundError); ok {
			fmt.Println("Not found error : ", notFoundErr.Error())
		} else {
			fmt.Println("Unknown error : ", err.Error())
		}
	} else {
		fmt.Println("Sukses")
	}
}
