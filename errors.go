package globalidentity

import "fmt"

type GlobalIdentityError []string

func (e GlobalIdentityError) Error() string {
	return fmt.Sprintf("%#v", []string(e))
}
