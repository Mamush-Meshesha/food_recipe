package client

import (
	"context"
	"net/http"
	"os"

	"github.com/hasura/go-graphql-client"
	"github.com/joho/godotenv"
)

var httpClient *http.Client

func init() {
		err := godotenv.Load()
		if err != nil {
			panic("Error loading .env file")
		}

	adminSecret := os.Getenv("ADMIN_SECRET")
	if adminSecret == "" {
		panic("ADMIN_SECRET is not set in .env file")
	}

	transport := NewRoundTripper(http.DefaultTransport, adminSecret)
	httpClient = &http.Client{
		Transport: transport,
	}
}

type RoundTripper struct {
	rt          http.RoundTripper
	adminSecret string
}

func (rt RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("x-hasura-admin-secret", rt.adminSecret)
	return rt.rt.RoundTrip(req)
}

func NewRoundTripper(rt http.RoundTripper, adminSecret string) *RoundTripper {
	return &RoundTripper{
		rt:          rt,
		adminSecret: adminSecret,
	}
}

func Mutation(mutation interface{}, variables map[string]interface{}) error {
	var client = graphql.NewClient(os.Getenv("HASURA_GRAPHQL_URL"), httpClient)
	return client.Mutate(context.Background(), mutation, variables)
}

func Query(query interface{}, variables map[string]interface{}) error {
	var client = graphql.NewClient(os.Getenv("HASURA_GRAPHQL_URL"), httpClient)
	return client.Query(context.Background(), query, variables)
}
