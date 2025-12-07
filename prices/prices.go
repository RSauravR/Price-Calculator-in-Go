package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	filemananger "example.com/price-calculator/fileMananger"
)

type TaxIncludedPriceJob struct {
	IOMananger        filemananger.FileMananger `json:"-"`
	TaxRate           float64                   `json:"tax_rate"`
	InputPrices       []float64                 `json:"input_prices"`
	TaxIncludedPrices map[string]string         `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOMananger.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringToFloats(lines)
	if err != nil {
		return err
	}

	job.InputPrices = prices

	return nil
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errChan chan error) {

	err := job.LoadData()
	if err != nil {
		// return err
		errChan <- err
		return
	}

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	err = job.IOMananger.WriteResult(job)
	if err != nil {
		// return err
		errChan <- err
		return
	}
	doneChan <- true
}

func NewTaxIncludedPriceJob(fm filemananger.FileMananger, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOMananger:  fm,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
