package main

import (
	"fmt"

	filemananger "example.com/price-calculator/fileMananger"
	"example.com/price-calculator/prices"
)

func main() {
	taxrates := []float64{0, 0.07, 0.1, 0.15}

	doneChan := make([]chan bool, len(taxrates))
	errChan := make([]chan error, len(taxrates))

	for index, taxrate := range taxrates {
		doneChan[index] = make(chan bool)
		errChan[index] = make(chan error)
		fm := filemananger.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxrate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxrate)
		go priceJob.Process(doneChan[index], errChan[index])
		// if err != nil {
		// 	fmt.Println("Could not process the job")
		// 	fmt.Println(err)
		// }
	}

	for index := range taxrates {
		select {
		case err := <-errChan[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChan[index]:
			fmt.Println("Done!")
		}
	}
}
