package main

import (
	"fmt"
	"net/http"
)

type Service struct{}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.Host, r.URL.Path )
}

func main()  {
	server := http.Server{
		Addr:":8080",
		Handler: &Service{},
	}
	server.ListenAndServe()
}