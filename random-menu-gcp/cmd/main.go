package main

import (
	"context"
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/joho/godotenv"
	"stani.com/menuprocessor"
)

func main() {
	ctx := context.Background()
	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/", function.RandomMenu); err != nil {
		log.Fatalf("funcframework.RegisterHTTPFunctionContext: %v\n", err)
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cloudFunctionPort := os.Getenv("CLOUD_FUNCTION_PORT")
	log.Printf("Using port: %s", cloudFunctionPort)

	if err := funcframework.Start(cloudFunctionPort); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
