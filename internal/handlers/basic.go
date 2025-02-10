package handlers

import (
	"log"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[internal/handlers/basic.go: RootHandler has been called]")
	w.Header().Set("Content-Type", "text/html")
	htmlContent := `
<!DOCTYPE html>
<html lang="en">
<head>
</head>
<body>
<h1>
<p>Hello There</p>
</h1>
</body>
</html>
`
	w.WriteHeader(http.StatusOK)
	data := []byte(htmlContent)
	_, err := w.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}
