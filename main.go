package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

// The app listens on this port.
const port = 3000

// This prefix is used by the IO ingress for its OAuth-related handlers.
const prefix = "@google"

// This is the path to the IO calling interface that is configured to call Google APIs.
func googleproxy() string {
	p := os.Getenv("GOOGLE_PROXY")
	if p != "" {
		return p
	}
	return "http://localhost:4848"
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", HomeHandler)
	httpd := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	if err := httpd.ListenAndServe(); err != nil {
		log.Fatalf("%s", err)
	}
}

type Profile struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	PictureUrl string `json:"picture"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var profile Profile
	providers := r.Header.Get("proxy-provider")
	if strings.Contains(providers, "google") {
		request, err := http.NewRequest("GET", googleproxy()+"/oauth2/v3/userinfo", nil)
		if err == nil {
			request.Header.Set("proxy-session", r.Header.Get("proxy-session"))
			response, err := http.DefaultClient.Do(request)
			if err == nil {
				body, err := io.ReadAll(response.Body)
				if err == nil {
					err = json.Unmarshal(body, &profile)
				}
			}
		}
	}
	t, err := template.New("home").Parse(home_template)
	if err != nil {
		log.Printf("%s", err)
		return
	}
	err = t.ExecuteTemplate(w, "home", map[string]any{
		"Name":       profile.Name,
		"PictureUrl": profile.PictureUrl,
		"Email":      profile.Email,
		"Prefix":     prefix,
	})
	if err != nil {
		log.Printf("%s", err)
		return
	}
}

const home_template = `
<!DOCTYPE html>
<html>
<head>
<title>Google OAuth Demo</title>
</head>
<body>
<h1>Google OAuth Demo</h1>
<div style="width: 100px; height: 100px; float:left; margin-right:1em; background-color:#EEE;">
{{ if .PictureUrl }}
<img src="{{ .PictureUrl }}"/>
{{ end }}
</div>
{{ if .Name }}
<h2>{{ .Name }}</h2>
<p>{{ .Email }}</p>
<a href="/{{.Prefix}}/signout" class="button">Sign out</a>
{{ else }}
<a href="/{{.Prefix}}/signin" class="button">Sign in</a>
{{ end }}
</body>
</html>
`
