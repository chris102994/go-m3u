package main

import (
	"fmt"
	"github.com/chris102994/go-m3u/pkg/m3u/models"
	"os"
)

func main() {
	var received models.M3U

	m3uData, err := os.ReadFile("test_m3u.m3u")
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	err = models.Unmarshal(m3uData, &received)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	output, err := models.Marshal(&received)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	fmt.Println(string(output))
}
