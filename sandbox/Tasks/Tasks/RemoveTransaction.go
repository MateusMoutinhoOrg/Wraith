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

// removeTransactionAction finds the stored movement by its id and removes it.
func removeTransactionAction(args api.HandleActionArgs) error {
	id, err := entries.Whole(args.Entries, IdField)
	if err != nil {
		return err
	}
	ledger, reachable := args.DataBase.GetSchema(config.TransactionSchema)
	if !reachable {
		return errors.New("the " + config.TransactionSchema + " registry is unreachable")
	}
	stored, failure := ledger.ListAll()
	if failure != nil {
		stored = nil
	}
	record := keepdeps.SchemaItem{}
	found := false
	for _, candidate := range stored {
		if candidate.Id == id {
			record = candidate
			found = true
			break
		}
	}
	if !found {
		return errors.New("transaction not found: " + strconv.FormatInt(id, 10))
	}
	if failure := record.Remove(); failure != nil {
		return errors.New("transaction " + strconv.FormatInt(id, 10) +
			" could not be removed: " + failure.Message)
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
