package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID           string
	OrginalUrl   string
	Shorturl     string
	Creationdate time.Time
}

var urlDB = make(map[string]URL)

func shorter(Originalurl string) string {
	hasher := md5.New()
	hasher.Write([]byte(Originalurl))
	// fmt.Println("hashed data: ", hasher)

	data := hasher.Sum(nil)
	// fmt.Println("Data :", data)

	hash := hex.EncodeToString(data)
	// fmt.Println("Encoded to String :", hash)
	return hash[:8]
}

func GetUrl(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not Found")
	}
	return url, nil
}

func CreateURL(OriginalURL string) string {
	shorturl := shorter(OriginalURL)
	id := shorturl
	urlDB[id] = URL{
		ID:           id,
		OrginalUrl:   OriginalURL,
		Shorturl:     shorturl,
		Creationdate: time.Now(),
	}
	// fmt.Println("URL DB :")
	// fmt.Println(urlDB[id])
	return shorturl
}

func DefaultPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Server is Live")
	fmt.Fprintf(w, "Hello There")
}

func Shortenerhandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Server is Live")
	type data struct {
		URL string `json:"url"`
	}
	var Data data
	err := json.NewDecoder(r.Body).Decode(&Data)
	if err != nil {
		http.Error(w, "Invalid request Body", http.StatusBadRequest)
		return
	}

	shorturl := CreateURL(Data.URL)
	// fmt.Fprintf(w, "%s", shorturl)

	Responce := struct {
		Shortlink string `json:"shorten"`
	}{Shortlink: shorturl}
	// fmt.Println("Response:", Responce)

	w.Header().Set("Content-type", "application/json")
	error23 := json.NewEncoder(w).Encode(Responce)
	if error23 != nil {
		http.Error(w, "Error occured", http.StatusBadRequest)
		return
	}
}

func RedirectFunction(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	link, error111 := GetUrl(id)
	if error111 != nil {
		http.Error(w, "URL not found", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, link.OrginalUrl, http.StatusFound)
}

func main() {
	// fmt.Println("Link Shortner")
	// Originalurl := "https://www.linkedin.com/in/gagandeep-singh-3281ba220/"
	// CreateURL(Originalurl)

	// Making Handler
	http.HandleFunc("/", DefaultPage)
	http.HandleFunc("/shorten", Shortenerhandler)
	http.HandleFunc("/redirect/", RedirectFunction)

	// Making Server
	fmt.Println("Starting Server on 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error Found", err)
		return
	}
}
