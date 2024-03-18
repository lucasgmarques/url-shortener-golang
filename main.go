package main

import (
    "fmt"
    "net"
    "net/http"
    "strconv"
    "strings"
)

var (
    baseUrl  string
    urlMap   = make(map[string]string)
    urlCount = 0
)

func main() {
    // Get the container's IP address dynamically
    containerIP := getOutboundIP()
    baseUrl = "http://" + containerIP.String() + ":8080/"

    http.HandleFunc("/", redirectToOriginal)
    http.HandleFunc("/shorten", shortenURL)
    fmt.Printf("Server started at %s\n", baseUrl)
    http.ListenAndServe(":8080", nil)
}

func getOutboundIP() net.IP {
    // This function retrieves the IP address of the container's network interface
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP
}

func redirectToOriginal(w http.ResponseWriter, r *http.Request) {
    short := strings.TrimPrefix(r.URL.Path, "/")
    if originalURL, ok := urlMap[short]; ok {
        http.Redirect(w, r, originalURL, http.StatusFound)
        return
    }
    http.NotFound(w, r)
}

func shortenURL(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    originalURL := r.FormValue("url")
    if originalURL == "" {
        http.Error(w, "URL cannot be empty", http.StatusBadRequest)
        return
    }

    shortURL := generateShortURL()
    urlMap[shortURL] = originalURL

    fmt.Fprintf(w, "Shortened URL: %s%s\n", baseUrl, shortURL)
}

func generateShortURL() string {
    urlCount++
    return strconv.Itoa(urlCount)
}

