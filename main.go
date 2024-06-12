package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sync"
)

// Struct untuk memuat URL, method, headers, body, query params, dan count dari file JSON
type Config struct {
	Url         string                 `json:"url"`
	Method      string                 `json:"method"`
	Headers     map[string]string      `json:"headers"`
	Body        map[string]interface{} `json:"body"`
	QueryParams map[string]string      `json:"query_params"`
	Count       int                    `json:"count"`
}

// Fungsi untuk membaca konfigurasi dari file JSON
func readConfigFromFile(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

// Fungsi untuk mengakses URL dan menampilkan responsnya
func fetchUrl(wg *sync.WaitGroup, config *Config) {
	defer wg.Done()

	// Menyiapkan request body jika ada
	var requestBody io.Reader
	if len(config.Body) > 0 {
		bodyBytes, err := json.Marshal(config.Body)
		if err != nil {
			fmt.Printf("Error marshalling request body: %s\n", err)
			return
		}
		requestBody = bytes.NewBuffer(bodyBytes)
	}

	// Menyiapkan URL dengan query parameters jika ada
	reqUrl, err := url.Parse(config.Url)
	if err != nil {
		fmt.Printf("Error parsing URL %s: %s\n", config.Url, err)
		return
	}

	query := reqUrl.Query()
	for key, value := range config.QueryParams {
		query.Add(key, value)
	}
	reqUrl.RawQuery = query.Encode()

	// Membuat request
	req, err := http.NewRequest(config.Method, reqUrl.String(), requestBody)
	if err != nil {
		fmt.Printf("Error creating request: %s\n", err)
		return
	}

	// Menambahkan headers jika ada
	for key, value := range config.Headers {
		req.Header.Add(key, value)
	}

	// Mengirim request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error fetching %s: %s\n", reqUrl.String(), err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response from %s: %s\n", reqUrl.String(), err)
		return
	}
	fmt.Printf("Response from %s: %s\n", reqUrl.String(), string(body))
}

func main() {
	// Membaca konfigurasi dari file JSON
	config, err := readConfigFromFile("target.json")
	if err != nil {
		fmt.Printf("Error reading configuration from file: %s\n", err)
		os.Exit(1)
	}

	var wg sync.WaitGroup
	for i := 0; i < config.Count; i++ {
		wg.Add(1)
		go fetchUrl(&wg, config)
	}
	wg.Wait()
	fmt.Println("All requests have been fetched.")
}
