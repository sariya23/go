package main

import (
	"errors"
	"fmt"
)

func check(a, b int) error {
	if a == 0 && b == 0 {
		return errors.New("Custom error")
	}
	return nil
}

func formattedError(a, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("Both zero")
	} else if a == 1 && b == 1 {
		return fmt.Errorf("Both one")
	}
	return nil
}

func main() {
	err := check(0, 0)
	if err == nil {
		fmt.Println("No error")
	} else {
		fmt.Println(err)
	}

	err2 := formattedError(0, 0)
	if err2 == nil {
		fmt.Println("No error")
	} else {
		fmt.Println(err2.Error())
	}
}
