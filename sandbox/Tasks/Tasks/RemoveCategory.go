package tasks

import (
	"errors"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps/keepdeps"
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
	// A category still classifying transactions, or still standing as another
	// category's parent, would leave a rendered page pointing at something
	// that is no longer there.
	for _, transaction := range removeCategoryRecords(args, config.TransactionSchema) {
		detail := removeCategoryText(transaction, config.DetailField)
		if utils.Part(detail, config.TransactionParts, config.TransactionCategory) == categoryName {
			return errors.New(categoryName + " still classifies transactions — " +
				"move them to another category first")
		}
	}
	for _, category := range removeCategoryRecords(args, config.CategorySchema) {
		detail := removeCategoryText(category, config.DetailField)
		if utils.Part(detail, config.CategoryParts, config.CategoryParent) == categoryName {
			return errors.New(categoryName + " is still the parent of " +
				removeCategoryText(category, config.NameField) + " — remove the child first")
		}
	}
	categories, reachable := args.DataBase.GetSchema(config.CategorySchema)
	if !reachable {
		return errors.New("the " + config.CategorySchema + " registry is unreachable")
	}
	record, found := categories.FindByKey(config.NameField, categoryName)
	if !found {
		return errors.New("category not found: " + categoryName)
	}
	if failure := record.Remove(); failure != nil {
		return errors.New("category could not be removed: " + failure.Message)
	}
	return nil
}

// removeCategoryRecords returns every record of one registry. An unreachable
// registry reads as no records: a task that cannot see one has nothing to
// refuse over.
func removeCategoryRecords(args api.HandleActionArgs, schemaName string) []keepdeps.SchemaItem {
	instance, reachable := args.DataBase.GetSchema(schemaName)
	if !reachable {
		return nil
	}
	stored, failure := instance.ListAll()
	if failure != nil {
		return nil
	}
	return stored
}

// removeCategoryText reads a text field of a stored record, as "" when the
// field is absent or holds something else.
func removeCategoryText(record keepdeps.SchemaItem, field string) string {
	value, err := record.Get(field)
	if err != nil {
		return ""
	}
	stored, ok := value.(string)
	if !ok {
		return ""
	}
	return stored
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
