package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cmdm := cmdmanager.New()
		job := prices.NewTaxIncludedPriceJob(fm, taxRate)
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)

		go job.Process(doneChans[index], errorChans[index])
	}

	// Wait for all jobs to complete
	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			fmt.Println(err)
		case <-doneChans[index]:
			fmt.Println("Job done!!")
		}
	}

	fmt.Println("All tax calculations completed.")
}
