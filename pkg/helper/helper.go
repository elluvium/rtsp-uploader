package helper

import (
	"log"
	"os"
)

func GetEnvVar(name string) string {
	ev, err := os.LookupEnv(name)
	if !err {
		log.Fatalf("Environment variable %v must be specified!", name)
	}
	return ev
}
