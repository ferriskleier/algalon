package errorHandler

import "fmt"

func Handle(err error) {
	if err != nil {
		fmt.Println("Error occured:", err)
		return
	}
}
