package tasks

import (
	"errors"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/entries"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/utils"
)

// addCategoryFields declares what the task accepts.
func addCategoryFields() []api.Field {
	return []api.Field{
		{Name: CategoryField, Type: api.TextField, Required: true,
			Description: "Display name, e.g. Food"},
		{Name: DescriptionField, Type: api.TextField, Required: true,
			Description: "What this category classifies"},
		{Name: RevenuesField, Type: api.BoolField, Required: true,
			Description: "true when the category accepts money coming in"},
		{Name: ExpensesField, Type: api.BoolField, Required: true,
			Description: "true when the category accepts money going out"},
		{Name: ParentField, Type: api.TextField,
			Description: "Parent category, making this one a child of it"},
	}
}

// addCategoryAction runs the task against the database it is handed.
func addCategoryAction(args api.HandleActionArgs) error {
	categoryName, err := name(args.Entries, CategoryField)
	if err != nil {
		return err
	}
	description, err := name(args.Entries, DescriptionField)
	if err != nil {
		return err
	}
	revenues, err := entries.Bool(args.Entries, RevenuesField)
	if err != nil {
		return err
	}
	expenses, err := entries.Bool(args.Entries, ExpensesField)
	if err != nil {
		return err
	}
	parent, err := optionalText(args.Entries, ParentField)
	if err != nil {
		return err
	}
	if parent != "" {
		if _, err := requireCategory(args, parent); err != nil {
			return errors.New("parent " + err.Error())
		}
		if parent == categoryName {
			return errors.New("a category cannot be its own parent")
		}
	}
	return insert(args, config.CategorySchema, "category "+categoryName, map[string]any{
		config.NameField:     categoryName,
		config.DetailField:   utils.Pack(categoryName, parent, description),
		config.RevenuesField: flag(revenues),
		config.ExpensesField: flag(expenses),
	})
}

// AddCategory returns the task that adds a category — what a transaction is
// classified under. The two flags are the whole idea:
//
//	revenues: true, expenses: false   money coming in — a salary
//	revenues: false, expenses: true   money going out — groceries
//	revenues: false, expenses: false  a transfer between your own accounts,
//	                                  which is neither income nor expense
//
// The third combination is what makes paying a card bill invisible in your
// expenses: the purchases were already counted on the day they happened.
func AddCategory() api.Task {
	return api.Task{
		Name:         "AddCategory",
		Description:  "Add a category — what a transaction is classified under",
		Fields:       addCategoryFields(),
		HandleAction: addCategoryAction,
	}
}

// flag stores a yes-or-no as the whole number the registries hold it as.
func flag(value bool) int64 {
	if value {
		return 1
	}
	return 0
}
