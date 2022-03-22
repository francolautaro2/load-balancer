package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Service of backend and url config
type Service struct {
	url *url.URL // Example https://127.0.0.1:8080
	id  int
}

// List of services
type Services struct {
	Sv []Service // Array of services
}

// User can create a Service Backend
func CreateService(identity int, addr string) Service {
	u, err := url.Parse(addr)
	if err != nil {
		panic(err)
	}
	sn := Service{
		url: u,
		id:  identity,
	}
	return sn
}

// Create List of Services
func (r *Services) AddToList(u Service) {
	r.Sv = append(r.Sv, u)
}

// Return len of services list
func LenServicesList(s Services) int {
	return len(s.Sv)
}

func PrintAllServices(s Services) {
	for i := range s.Sv {
		fmt.Println("id services: ", s.Sv[i].id)
	}
}

// counter of services
// and create Services main
var idx int = 0
var s Services

// Load balancer function
// implementation round robin algorithm
func LoadBalancer(w http.ResponseWriter, r *http.Request) {

	// DEFINE THE SERVICES
	d := CreateService(0, "URL")
	c := CreateService(1, "URL")
	e := CreateService(2, "URL")
	// LIST SERVICES
	s.AddToList(d)
	s.AddToList(c)
	s.AddToList(e)

	lenServices := LenServicesList(s)
	currentURL := s.Sv[idx%lenServices].url

	idx++

	proxy := httputil.NewSingleHostReverseProxy(currentURL)
	proxy.ServeHTTP(w, r)
}

func main() {

	// Listen the main service
	http.HandleFunc("/", LoadBalancer)
	PrintAllServices(s)
	http.ListenAndServe(":8080", nil)

}
