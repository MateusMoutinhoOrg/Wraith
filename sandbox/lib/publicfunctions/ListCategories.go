package publicfunctions

import (
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/category"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/store"
)

// ListCategoriesFactory fills api.Lib.ListCategories with a closure that
// reads every stored category back out of the injected database.
func ListCategoriesFactory(l *api.Lib) func() []api.Category {
	return func() []api.Category {
		categories, ok := store.Categories(l.Deps)
		if !ok {
			return nil
		}
		records, err := categories.ListAll()
		if err != nil {
			return nil
		}
		listed := make([]api.Category, 0, len(records))
		for _, record := range records {
			listed = append(listed, category.New(l.Deps, record))
		}
		return listed
	}
}
