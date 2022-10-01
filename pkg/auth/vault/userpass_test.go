package vault_test

import (
	"github.com/akastav/argocd-terraform-plugin/pkg/auth/vault"
	"github.com/akastav/argocd-terraform-plugin/pkg/helpers"
	"testing"
)

func TestUserPassLogin(t *testing.T) {
	cluster, username, password := helpers.CreateTestUserPassVault(t)
	defer cluster.Cleanup()

	userpass := vault.NewUserPassAuth(username, password, "")

	if err := userpass.Authenticate(cluster.Cores[0].Client); err != nil {
		t.Fatalf("expected no errors but got: %s", err)
	}
}
