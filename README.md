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
* **|
* **├── main.go               # Orchestrates concurrent jobs and handles results
* **├── prices/
* **│   └── prices.go         # Defines TaxIncludedPriceJob struct and core processing logic
* **├── fileMananger/
* **│   └── fileMananger.go   # Handles file I/O operations (read input, write JSON output)
* **├── conversion/
* **│   └── conversion.go     # Utility for safe string-to-float array conversion
* **└── prices.txt            # Input file containing raw prices (one per line)
