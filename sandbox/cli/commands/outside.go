package commands

import "errors"

// errOutside reports a destination that climbs out of the vault. Writing
// outside the folder you are standing in is the one thing a render is never
// allowed to do, whether the path came from the config or from a flag.
func errOutside(dest string) error {
	return errors.New(dest + " is outside the vault — a destination stays inside it")
}
