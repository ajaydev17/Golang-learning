package cmdmanager

import "fmt"

type CMDManager struct{}

func (cmd CMDManager) ReadLinesFromFile() ([]string, error) {
	fmt.Println("Please enter your prices, Confirm every price with ENTER!!")

	var prices []string

	for {
		fmt.Print("Enter price (or type 'done' to finish): ")
		var input string
		fmt.Scanln(&input)
		if input == "done" {
			break
		}
		prices = append(prices, input)
	}
	return prices, nil
}

func (cmd CMDManager) WriteJsonToFile(data interface{}) error {
	fmt.Println(data)
	return nil
}

func New() CMDManager {
	return CMDManager{}
}
