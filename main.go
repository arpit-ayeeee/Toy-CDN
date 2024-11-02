package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var (
	originMap = map[string]string{
		"localhost:9000": "http://google.com",
	}

	cacheDir = "./cache"
)

// CacheEntry represents the cached response structure
type CacheEntry struct {
	Text string `json:"text"`
}

func main() {
	// Ensure the cache directory exists
	err := os.MkdirAll(cacheDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating cache directory:", err)
		return
	}

	http.HandleFunc("/", proxy)
	fmt.Println("Server running at port 9000")
	err = http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

func proxy(w http.ResponseWriter, r *http.Request) {
	originBase, ok := originMap[r.Host]
	if !ok {
		http.Error(w, fmt.Sprintf("Unknown origin: %s", r.Host), http.StatusNotFound)
		return
	}

	originURL := originBase + r.URL.Path + "?" + r.URL.RawQuery
	urlHash := fmt.Sprintf("%x", md5.Sum([]byte(originURL)))

	// Try to get response from cache
	cachedResponse := cacheGet(urlHash)
	if cachedResponse != nil {
		fmt.Println("Cache hit:", urlHash)
		w.Write([]byte(cachedResponse.Text))
		return
	}

	// If not in cache, fetch from origin
	fmt.Println("Fetching from origin:", originURL)
	resp, err := http.Get(originURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching from origin: %s", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Error reading response body", http.StatusInternalServerError)
			return
		}

		cachePut(urlHash, string(body)) // Cache the response
		w.Write(body)
	} else {
		http.Error(w, fmt.Sprintf("Error from origin: %d", resp.StatusCode), resp.StatusCode)
	}
}

func cacheGet(urlHash string) *CacheEntry {
	cachePath := filepath.Join(cacheDir, urlHash)

	file, err := os.Open(cachePath)
	if err != nil {
		return nil // Cache miss
	}
	defer file.Close()

	var entry CacheEntry
	err = json.NewDecoder(file).Decode(&entry)
	if err != nil {
		fmt.Println("Error decoding cached entry:", err)
		return nil
	}

	return &entry
}

func cachePut(urlHash string, response string) {
	cachePath := filepath.Join(cacheDir, urlHash)

	file, err := os.OpenFile(cachePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening cache file:", err)
		return
	}
	defer file.Close()

	entry := CacheEntry{Text: response}
	err = json.NewEncoder(file).Encode(&entry)
	if err != nil {
		fmt.Println("Error encoding cache entry:", err)
	}
}
