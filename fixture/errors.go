package fixture

import "github.com/pkg/errors"

var (
	ErrZeroGenerateTask = errors.New("not have struct for generate")
	ErrPathFile         = errors.New("the path for the file is not specified")
)
