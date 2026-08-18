//go:build ignore

// This file is an illustrative sample, not part of the build.
package tasks

import (
	"errors"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/entries"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/store"
)

// AddNote returns the task that files one note under a title. The title is
// what a later task addresses the note by, so it must be unique; the body is
// free text and may be left out.
func AddNote() api.Task {
	return api.Task{
		Name:        "AddNote",
		Description: "File a note under a title",
		Fields: []api.Field{
			{Name: "title", Type: api.TextField, Required: true,
				Description: "What the note is called. Must be unique"},
			{Name: "body", Type: api.TextField,
				Description: "The note itself"},
			{Name: "pinned", Type: api.BoolField,
				Description: "Keep it at the top of the list",
				Default:     false},
		},
		HandleAction: func(args api.HandleActionArgs) error {
			// 1. Read and check every field first. Nothing is written until
			//    the last one has passed.
			title, err := name(args.Entries, "title")
			if err != nil {
				return err
			}
			body, err := optionalText(args.Entries, "body")
			if err != nil {
				return err
			}
			pinned, err := entries.Bool(args.Entries, "pinned")
			if err != nil {
				return err
			}
			if store.Exists(args.DataBase, store.NoteSchema, title) {
				return errors.New("a note called " + title + " already exists")
			}

			// 2. Write, through the database and nothing else. Free text
			//    travels packed beside the unique key it belongs to.
			return insert(args, store.NoteSchema, "note "+title, map[string]any{
				store.NameField:   title,
				store.DetailField: store.Pack(title, body),
				store.PinnedField: flag(pinned),
			})
		},
	}
}
