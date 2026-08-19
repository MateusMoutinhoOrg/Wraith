package tasks

import (
	"errors"
	"strings"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps/keepdeps"
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

// addCategoryAction reads every field, checks the parent it was given exists,
// and writes one record into the category registry.
func addCategoryAction(args api.HandleActionArgs) error {
	categoryName, err := entries.Text(args.Entries, CategoryField)
	if err != nil {
		return err
	}
	if categoryName == "" {
		return errors.New(CategoryField + " is required")
	}
	if strings.Contains(categoryName, utils.Separator) {
		return errors.New(CategoryField + " may not contain " + utils.Separator)
	}
	description, err := entries.Text(args.Entries, DescriptionField)
	if err != nil {
		return err
	}
	if description == "" {
		return errors.New(DescriptionField + " is required")
	}
	if strings.Contains(description, utils.Separator) {
		return errors.New(DescriptionField + " may not contain " + utils.Separator)
	}
	revenues, err := entries.Bool(args.Entries, RevenuesField)
	if err != nil {
		return err
	}
	expenses, err := entries.Bool(args.Entries, ExpensesField)
	if err != nil {
		return err
	}
	parent := ""
	if entries.Present(args.Entries, ParentField) {
		parent, err = entries.Text(args.Entries, ParentField)
		if err != nil {
			return err
		}
		if strings.Contains(parent, utils.Separator) {
			return errors.New(ParentField + " may not contain " + utils.Separator)
		}
	}
	categories, reachable := args.DataBase.GetSchema(config.CategorySchema)
	if !reachable {
		return errors.New("the " + config.CategorySchema + " registry is unreachable")
	}
	if parent != "" {
		if _, found := categories.FindByKey(config.NameField, parent); !found {
			return errors.New("parent category not found: " + parent)
		}
		if parent == categoryName {
			return errors.New("a category cannot be its own parent")
		}
	}
	// The registries hold a yes-or-no as a whole number.
	revenuesFlag := int64(0)
	if revenues {
		revenuesFlag = 1
	}
	expensesFlag := int64(0)
	if expenses {
		expensesFlag = 1
	}
	_, failure := categories.NewItem(map[string]any{
		config.NameField:     categoryName,
		config.DetailField:   utils.Pack(categoryName, parent, description),
		config.RevenuesField: revenuesFlag,
		config.ExpensesField: expensesFlag,
	})
	if failure == nil {
		return nil
	}
	if failure.Type == keepdeps.KeyConflict {
		return errors.New("category " + categoryName + " already exists")
	}
	return errors.New("category " + categoryName + " could not be stored: " + failure.Message)
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
