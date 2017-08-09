package globalidentity

import (
	"github.com/fortytw2/leaktest"
	"testing"
)

func TestGlobalIdentityError_Error(t *testing.T) {
	defer leaktest.Check(t)()
	err := GlobalIdentityError([]string{"error01", "error01"})
	if err.Error() != `[]string{"error01", "error01"}` {
		t.FailNow()
	}
}
