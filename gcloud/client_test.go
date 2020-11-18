package gcloud

import (
  "testing"
)

const testScope = "https://www.googleapis.com/auth/cloud-platform"

func TestClient(t *testing.T) {

  gc, err := New("<test-org-id>", "<test-project-id>", testScope)
  if err != nil {
		t.Errorf("gclient: Setting up Client" + "credentials: %w", err)
  }

  if gc.orgID == "" {
    t.Errorf("gclient: Client missing organzation id ")
  }

  if gc.gcProjectID == "" {
    t.Errorf("gclient: Client missing project id")
  }
}