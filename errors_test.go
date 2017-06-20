package globalidentity

import "testing"

func TestGlobalIdentityError_Error(t *testing.T) {
	err := GlobalIdentityError([]string{"error01", "error01"})
	if err.Error() != `[]string{"error01", "error01"}` {
		t.FailNow()
	}
}
