package gcloud

import (
  "testing"
)

func TestClient(t *testing.T) {
  gc, err := New("<test-org-id>", "<test-project-id>")
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