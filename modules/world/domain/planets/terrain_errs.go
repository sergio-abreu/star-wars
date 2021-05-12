package planets

import "fmt"

func NewErrTerrainNotFound(terrain string) ErrTerrainNotFound {
	return ErrTerrainNotFound{terrain: terrain}
}

type ErrTerrainNotFound struct {
	terrain string
}

func (err ErrTerrainNotFound) Error() string {
	return fmt.Sprintf("terrain %q not found", err.terrain)
}
