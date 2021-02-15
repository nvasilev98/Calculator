package env

import (
	"errors"
	"os"
)

func ReadEnvVariable(envVariable string, data *string) error {
	variable, ok := os.LookupEnv(envVariable)
	if !ok {
		return errors.New("Missing an environment variable")
	}
	*data = variable
	return nil
}
