package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/handlers"
)

func GetReverseProxyTarget(resolve bool) *url.URL {
	// check if we are runing under Linux
	u, err := url.Parse(os.Getenv("AWS_CONTAINER_CREDENTIALS_FULL_URI"))
	if err != nil {
		log.Fatalln("Bad AWS_CONTAINER_CREDENTIALS_FULL_URI:", err.Error())
	}
	if resolve {
		u.Host = "host.docker.internal:" + u.Port()
	} else {
		// Using the default name and port for the socat service.
		u.Host = "socat-tcp-to-unix:20000"
	}
	return u
}

func addAuthorizationHeader(authToken string, next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Header.Add("Authorization", authToken)
		next.ServeHTTP(w, r)
	}
}

func lookup() bool {
	envTime := os.Getenv("AWS_CONTAINER_LOOKUP_TIMEOUT")
	if len(envTime) == 0 {
		envTime = "2s"
	}

	timeout, err := time.ParseDuration(envTime)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, fmt.Sprintf("given timeout: %s is invalid.", timeout))
		os.Exit(1)
	}

	// Create a new context
	// With a deadline of the timeout set earlier
	ctx := context.Background()
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, timeout)

	// capture ctrl+c and stop query
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		for sig := range c {
			// sig is a ^C, handle it
			fmt.Printf("recieved: %s, stopping\n", sig.String())
			cancel()
			os.Exit(1)
		}
	}()

	host := "host.docker.internal"
	var resolver net.Resolver
	ip, err := resolver.LookupIP(ctx, "ip", host)
	if err == nil {
		fmt.Printf("host.docker.internal resolves to : %s", ip)
		return true
	}

	return false
}

func main() {

	target := GetReverseProxyTarget(lookup())
	authToken := os.Getenv("AWS_CONTAINER_AUTHORIZATION_TOKEN")
	log.Printf("reverse proxying target:%s auth:%s\n", target, authToken)

	handler := handlers.LoggingHandler(os.Stderr,
		addAuthorizationHeader(authToken,
			httputil.NewSingleHostReverseProxy(target)))

	_ = http.ListenAndServe(":80", handler)
}
