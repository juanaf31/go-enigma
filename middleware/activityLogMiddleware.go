package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func ActivityLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.Header.Get("User-Agent")
		wrt := fmt.Sprintf("%v Accessing path %v with application %v\n", time.Now(), r.RequestURI, userAgent)
		log.Printf("Accessing path %v with application %v\n", r.RequestURI, userAgent)

		file, err := os.OpenFile("../logData/data.txt", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer file.Close()
		if _, err := file.WriteString(wrt); err != nil {
			log.Println(err)
		}
		// Untuk melanjutkan ke middlware selanjutnya
		next.ServeHTTP(w, r)
	})
}
