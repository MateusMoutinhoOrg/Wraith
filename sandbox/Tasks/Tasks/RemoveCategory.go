package tasks

import (
	"errors"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/entries"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/utils"
)

// removeCategoryFields declares what the task accepts.
func removeCategoryFields() []api.Field {
	return []api.Field{
		{Name: CategoryField, Type: api.TextField, Required: true,
			Description: "Display name of the category to remove"},
	}
}

// removeCategoryAction refuses a category anything still points at, and
// otherwise deletes the record.
func removeCategoryAction(args api.HandleActionArgs) error {
	categoryName, err := entries.Text(args.Entries, CategoryField)
	if err != nil {
		return err
	}
	if categoryName == "" {
		return errors.New(CategoryField + " is required")
	}
	categories, err := schemaOf(args, config.CategorySchema)
	if err != nil {
		return err
	}
	record, found := categories.FindByKey(config.NameField, categoryName)
	if !found {
		return errors.New("category not found: " + categoryName)
	}
	// A category still classifying transactions, or still standing as another
	// category's parent, would leave a rendered page pointing at something
	// that is no longer there. The movements it classifies are read off its
	// own nested registry; a parent is only named in another category's
	// packed detail, so that one is still a walk through the registry.
	if transactionCount(record) > 0 {
		return errors.New(categoryName + " still classifies transactions — " +
			"move them to another category first")
	}
	stored, failure := categories.ListAll()
	if failure != nil {
		stored = nil
	}
	for _, category := range stored {
		detail := text(category, config.DetailField)
		if utils.Part(detail, config.CategoryParts, config.CategoryParent) == categoryName {
			return errors.New(categoryName + " is still the parent of " +
				text(category, config.NameField) + " — remove the child first")
		}
	}
	if failure := record.Remove(); failure != nil {
		return errors.New("category could not be removed: " + failure.Message)
	}
	return nil
}

// RemoveCategory returns the task that removes a category from the registry.
// A category still classifying transactions, or still standing as another
// category's parent, is refused: either one would leave a rendered page
// pointing at something that is no longer there.
func RemoveCategory() api.Task {
	return api.Task{
		Name:         "RemoveCategory",
		Description:  "Remove a category nothing is classified under any more",
		Fields:       removeCategoryFields(),
		HandleAction: removeCategoryAction,
	}
}
