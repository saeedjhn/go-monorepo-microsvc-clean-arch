package bcrypt_test

import (
	"testing"

	"github.com/saeedjhn/go-monorepo-microsvc-clean-arch/pkg/security/bcrypt"
)

func TestBcrypt(t *testing.T) {
	t.Log(bcrypt.Generate("pass", 32))
}
