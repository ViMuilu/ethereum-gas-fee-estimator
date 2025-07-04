package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Ethereum Gas & Fee Estimator backend running on :8080")
	http.ListenAndServe(":8080", nil)
}
