package tasks

import (
	"errors"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
)

// RemoveCategory returns the task that removes a category from the registry.
// A category still classifying transactions, still named by a recurrence, or
// still standing as another category's parent is refused: every one of those
// would leave a rendered page pointing at something that is no longer there.
func RemoveCategory() api.Task {
	return api.Task{
		Name:        "RemoveCategory",
		Description: "Remove a category nothing is classified under any more",
		Fields: []api.Field{
			{Name: CategoryField, Type: api.TextField, Required: true,
				Description: "Display name of the category to remove"},
		},
		HandleAction: func(args api.HandleActionArgs) error {
			categoryName, err := name(args.Entries, CategoryField)
			if err != nil {
				return err
			}
			for _, transaction := range records(args, config.TransactionSchema) {
				if detail(transaction, config.TransactionParts, config.TransactionCategory) == categoryName {
					return errors.New(categoryName + " still classifies transactions — " +
						"move them to another category first")
				}
			}
			for _, recurrence := range records(args, config.RecurrenceSchema) {
				if detail(recurrence, config.RecurrenceParts, config.RecurrenceCategory) == categoryName {
					return errors.New(categoryName + " is still named by the recurrence " +
						text(recurrence, config.NameField) + " — remove it first")
				}
			}
			for _, category := range records(args, config.CategorySchema) {
				if detail(category, config.CategoryParts, config.CategoryParent) == categoryName {
					return errors.New(categoryName + " is still the parent of " +
						text(category, config.NameField) + " — remove the child first")
				}
			}
			return remove(args, config.CategorySchema, "category", categoryName)
		},
	}
}
