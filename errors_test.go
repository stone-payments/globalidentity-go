package globalidentity

import (
	"testing"
	"github.com/fortytw2/leaktest"
)

func TestGlobalIdentityError_Error(t *testing.T) {
	defer leaktest.Check(t)()
	err := GlobalIdentityError([]string{"error01", "error01"})
	if err.Error() != `[]string{"error01", "error01"}` {
		t.FailNow()
	}
}
