package client

import (
	"context"
	"net/http"

	"github.com/hasura/go-graphql-client"
)

var httpClient *http.Client = &http.Client{
	Transport: RoundTripper{rt: http.DefaultTransport},
}

type RoundTripper struct {
	rt http.RoundTripper
}

func (rt RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-hasura-admin-secret", "myadminsecretkey")
	return rt.rt.RoundTrip(req)
}

func Mutation(mutation interface{}, variables map[string]interface{}) error {
	var client = graphql.NewClient("http://localhost:8181/v1/graphql", httpClient)
	return client.Mutate(context.Background(), mutation, variables)
}

func Query(query interface{}, variables map[string]interface{}) error {
	var client = graphql.NewClient("http://localhost:8181/v1/graphql", httpClient)
	return client.Query(context.Background(), query, variables)
}
