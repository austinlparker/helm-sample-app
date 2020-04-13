package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/lightstep/opentelemetry-exporter-go/lightstep"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/key"
	"go.opentelemetry.io/otel/api/trace"
	"go.opentelemetry.io/otel/plugin/othttp"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func getEnvOrDefault(key, defaultValue string) string {
	value, found := os.LookupEnv(key)
	if !found {
		value = defaultValue
	}
	return value
}

func initTracer() {
	lExporter, err := lightstep.NewExporter(
		lightstep.WithAccessToken(getEnvOrDefault("LS_KEY", "your_access_token")),
		lightstep.WithServiceName(getEnvOrDefault("SERVICE_NAME", "web-service")))

	tp, err := sdktrace.NewProvider(sdktrace.WithConfig(sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
		sdktrace.WithSyncer(lExporter))
	if err != nil {
		log.Fatal(err)
	}
	global.SetTraceProvider(tp)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	trace.SpanFromContext(ctx).SetAttributes(key.String("service.version", getEnvOrDefault("SERVICE_VERSION", "1.1")))
	n := rand.Intn(200) // n will be between 0 and 200
	if n%2 == 0 {
		fmt.Printf("Sleeping %d s...\n", (n / 2))
		time.Sleep(time.Duration(n/2) * time.Second)
	} else {
		fmt.Printf("Sleeping %d ms...\n", n)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}
	fmt.Fprintf(w, "I am a GO application running inside Docker.")
}

func main() {
	initTracer()
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Basic web server is starting on port 8080...")

	mux := http.NewServeMux()

	mux.Handle("/", othttp.NewHandler(http.HandlerFunc(indexHandler), "root", othttp.WithPublicEndpoint()))
	http.ListenAndServe(":8080", mux)
}
