package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func main() {
	demoURL, err := url.Parse("http://httpforever.com/")
	if err != nil {
		log.Fatal(err)
	}
	// Method 01
	// proxy := httputil.NewSingleHostReverseProxy(demoURL)
	// http.ListenAndServe(":8080", proxy)

	// Method 02
	proxy := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Host = demoURL.Host
		r.URL.Host = demoURL.Host
		r.URL.Scheme = demoURL.Scheme
		r.RequestURI = ""

		s, _, _ := net.SplitHostPort(r.RemoteAddr)
		r.Header.Set("X-Forwarded-For", s)

		resp, err := http.DefaultClient.Do(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err)
			return
		}
		for key, values := range resp.Header {
			for _, value := range values {
				w.Header().Set(key, value)
			}
		}

		done := make(chan bool)
		go func() {
			for {
				select {
				case <-time.Tick(10 * time.Millisecond):
					w.(http.Flusher).Flush()
				case <-done:
					return
				}
			}
		}()

		trailerKeys := []string{}
		for key := range resp.Trailer {
			trailerKeys = append(trailerKeys, key)
		}
		w.Header().Set("Trailer", strings.Join(trailerKeys, ","))

		w.WriteHeader(resp.StatusCode)
		io.Copy(w, r.Body)

		for key, values := range resp.Trailer {
			for _, value := range values {
				w.Header().Set(key, value)
			}
		}

		close(done)
	})
	http.ListenAndServe(":8080", proxy)
}
