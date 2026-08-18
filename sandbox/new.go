package lib

import (
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib"
)

// New injects a Deps struct into the library and returns the api.Lib entry
// point. It delegates to the internal lib constructor, which stores the deps
// and the database path on the struct and runs the factories over it, each of
// which fills one function field with a closure reading those deps.
//
// databasePath is the folder the registries are persisted in, relative to
// whatever root the adapter was given. It is required on construction rather
// than settled later, because a library that did not know where its data
// lived could not answer a single question about it. Passing "" takes the
// default, `data`.
func New(d deps.Deps, databasePath string) api.Lib {
	return lib.New(d, databasePath)
}
