package common

import "syscall"

func EnvString(key, fallback string) string {
	if val, ok := syscall.Gatenv(key); ok {
		return val
	}

	return fallback

}
