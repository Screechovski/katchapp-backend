package helper

import "os"

func IsDev() bool {
	mode, modeExists := os.LookupEnv("MODE")

	return (modeExists && mode == "dev")
}
