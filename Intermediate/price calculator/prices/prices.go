package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
	FileManager       iomanager.IOManager `json:"-"`
}

func (job *TaxIncludedPriceJob) LoadPriceData() error {
	// Read lines from file
	lines, err := job.FileManager.ReadLinesFromFile()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	// Convert lines to float64 and store in InputPrices
	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println("Error converting prices:", err)
		return err
	}

	job.InputPrices = prices

	return nil
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) {
	err := job.LoadPriceData()

	if err != nil {
		fmt.Println("Error loading price data:", err)
		errorChan <- err
		return
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncluededPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncluededPrice)
	}

	job.TaxIncludedPrices = result

	// Write result to JSON file
	err = job.FileManager.WriteJsonToFile(job)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		errorChan <- err
		return
	}
	doneChan <- true
}

func NewTaxIncludedPriceJob(fm iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		FileManager: fm,
		InputPrices: []float64{10.0, 20.0, 30.0},
	}
}
