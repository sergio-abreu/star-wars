package planets

import "fmt"

func NewErrClimateNotFound(climate string) ErrClimateNotFound {
	return ErrClimateNotFound{climate: climate}
}

type ErrClimateNotFound struct {
	climate string
}

func (err ErrClimateNotFound) Error() string {
	return fmt.Sprintf("climate %q not found", err.climate)
}
