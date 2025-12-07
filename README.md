# 💰 Go Price Calculator Project

## Project Overview

This is a **Go** project designed to calculate **tax-included prices** from a list of raw input values and output the results into structured **JSON** files.

The primary goal of this project is to provide a practical, hands-on demonstration of Go's powerful concurrency model, specifically utilizing **goroutines** and **channels** to run multiple tax calculations in parallel.

---

## ✨ Features and Highlights

* **⚡️ Concurrency:** Demonstrates **goroutines** for parallel job execution (one per tax rate).
* **📡 Synchronization:** Uses **channels** (`doneChan` and `errChan`) to communicate completion or errors back to the main thread.
* **🔄 Structured Control:** Employs the **`select` statement** to non-blockingly wait for messages from multiple channels.
* **📦 Modular Design:** Project is structured into separate packages (`prices`, `fileMananger`, `conversion`) to adhere to Go's best practices.
* **📁 File Handling:** Includes robust logic for reading input data from a text file and writing calculated results to JSON.
* **✅ Error Management:** Propagates errors safely across goroutines using dedicated channels.

---

## 🏗️ Project Structure

The codebase is organized into clear, functional packages:
Price-Calculator-Project/
* |
* ├── main.go               # Orchestrates concurrent jobs and handles results
* ├── prices/
* │   └── prices.go         # Defines TaxIncludedPriceJob struct and core processing logic
* ├── fileMananger/
* │   └── fileMananger.go   # Handles file I/O operations (read input, write JSON output)
* ├── conversion/
* │   └── conversion.go     # Utility for safe string-to-float array conversion
* └── prices.txt            # Input file containing raw prices (one per line)

## 🚀 How It Works

The application processes four fixed tax rates concurrently (0%, 7%, 10%, 15%) to calculate the final price for every value in `prices.txt`.

1.  **Job Creation:** For each tax rate, a **`TaxIncludedPriceJob`** instance is created.
2.  **Concurrent Execution:** In `main.go`, each job is launched as a separate **goroutine** using the following pattern:
    ```go
    go priceJob.Process(doneChan[index], errChan[index])
    ```
3.  **Data Processing:** Each job independently:
    * Loads raw prices via `fileMananger.LoadData()`.
    * Converts string prices to `float64` via `conversion.ToFloat()`.
    * Calculates the tax-included price (e.g., $P_{\text{tax}} = P_{\text{raw}} \times (1 + \text{rate})$).
    * Writes the final results to a unique JSON file (e.g., `result_10.json`).
4.  **Synchronization:** The main function uses a **`select`** block to wait for successful completion messages on `doneChan` or error messages on `errChan` from *any* of the running goroutines, ensuring all jobs finish before exiting.

---

## ⚙️ Example Run

### Prerequisites

You must have **Go (1.18 or higher)** installed.

### Execution

Navigate to the project's root directory and run:

```bash
go run .

### Output
Terminal Logs:
Done!
Done!
Done!
Done!

Generated Files:
Four JSON files will be created in the root directory:
* result_0.json
* result_7.json
* result_10.json
* result_15.json

Example content of result_10.json:
[
  {
    "OriginalPrice": 100.0,
    "TaxIncludedPrice": 110.0
  },
  {
    "OriginalPrice": 55.5,
    "TaxIncludedPrice": 61.05
  }
]

