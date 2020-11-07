package gcloud

import (
	"context"
	"net/http"
	"golang.org/x/oauth2/google"
	errors "golang.org/x/xerrors"
)

// Client contains information related to client initialization
type Client struct {
	httpc         *http.Client
	orgID         string
	gcProjectID   string
}

func New(orgID, gcProjectID string) (*Client, error) {
  _, err := google.FindDefaultCredentials(context.Background())
  if err != nil {
		return nil, errors.Errorf("gcloud: finding Application Default credentials", err)
  }
  
  client, err := google.DefaultClient(context.Background())
	if err != nil {
		return nil, errors.Errorf("gcloud: setting default google client", err)
  }
  
  return &Client{
		httpc:        client,
		orgID:        orgID,
		gcProjectID:  gcProjectID,
	}, nil
}