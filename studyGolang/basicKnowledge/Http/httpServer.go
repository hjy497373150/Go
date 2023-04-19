package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"
)

var (
	mu sync.Mutex
	count int
)

func main() {
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	ip := getIP(r)
    fmt.Println("IP address:", ip)

    netIP := net.ParseIP(ip)
    if netIP != nil {
        if isIPv4(netIP) {
            fmt.Println("This is an IPv4 address")
        } else {
            fmt.Println("This is an IPv6 address")
        }
    }

    fmt.Fprintf(w, "Your IP address is %s and you have visited %d times", ip, count)
}

func getIP(r *http.Request) string {
	// remoteIP, _, _ := net.SplitHostPort(r.RemoteAddr)
    // return remoteIP
	host, _, err := net.SplitHostPort(r.RemoteAddr)
    if err != nil {
        return ""
    }
    ips, err := net.LookupIP(host)
    if err != nil {
        return ""
    }
    for _, ip := range ips {
        if ip.To4() != nil {
            return ip.String()
        } else if ip.To16() != nil {
            return ip.String()
        }
    }
    return ""
}

func isIPv4(ip net.IP) bool {
    return ip.To4() != nil
}