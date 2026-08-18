package commands

import (
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
)

// Tasks lists every action this brain can perform, straight out of the
// registry the binary carries — so the list can never name a task that does
// not exist, or miss one that does.
func Tasks(l *api.Lib) int {
	l.Deps.Printf("%s", config.TaskListHeader)
	for _, declared := range l.Tasks {
		l.Deps.Printf(config.ListEntry, declared.Name, declared.Description)
	}
	return api.ExitOk
}

// Visualizations lists every renderer this brain carries, the same way.
func Visualizations(l *api.Lib) int {
	l.Deps.Printf("%s", config.VisualizationListHeader)
	for _, declared := range l.Visualizations {
		l.Deps.Printf(config.ListEntry, declared.Name, declared.Description)
	}
	return api.ExitOk
}
