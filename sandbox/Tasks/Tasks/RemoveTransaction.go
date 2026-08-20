package tasks

import (
	"errors"
	"strconv"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps/keepdeps"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/entries"
)

// removeTransactionFields declares what the task accepts.
func removeTransactionFields() []api.Field {
	return []api.Field{
		{Name: IdField, Type: api.NumberField, Required: true,
			Description: "The id of the transaction to remove, as the statement shows it"},
	}
}

// removeTransactionAction finds the stored movement by its id, takes it out
// of the account and the category that index it, and removes it.
func removeTransactionAction(args api.HandleActionArgs) error {
	id, err := entries.Whole(args.Entries, IdField)
	if err != nil {
		return err
	}
	ledger, err := schemaOf(args, config.TransactionSchema)
	if err != nil {
		return err
	}
	record, found := ledger.FindById(id)
	if !found {
		return errors.New("transaction not found: " + strconv.FormatInt(id, 10))
	}
	// The index goes first: a link pointing at a movement that is no longer
	// there would be read as a movement, while a movement no registry indexes
	// any more is simply about to be removed.
	if err := removeTransactionUnindex(args, record); err != nil {
		return err
	}
	if failure := record.Remove(); failure != nil {
		return errors.New("transaction " + strconv.FormatInt(id, 10) +
			" could not be removed: " + failure.Message)
	}
	return nil
}

// removeTransactionUnindex takes a movement's id out of the account and the
// category it points at. A registry that no longer holds the record it names
// has nothing to take out.
func removeTransactionUnindex(args api.HandleActionArgs, record keepdeps.SchemaItem) error {
	owners := []struct {
		schemaName string
		idField    string
	}{
		{config.AccountSchema, config.AccountID},
		{config.CategorySchema, config.CategoryID},
	}
	for _, owner := range owners {
		instance, err := schemaOf(args, owner.schemaName)
		if err != nil {
			return err
		}
		indexer, found := instance.FindById(number(record, owner.idField))
		if !found {
			continue
		}
		if err := unlinkTransaction(indexer, record.Id); err != nil {
			return err
		}
	}
	return nil
}

// RemoveTransaction returns the task that removes a transaction from the ledger.
func RemoveTransaction() api.Task {
	return api.Task{
		Name:         "RemoveTransaction",
		Description:  "Remove a transaction from the ledger",
		Fields:       removeTransactionFields(),
		HandleAction: removeTransactionAction,
	}
}
