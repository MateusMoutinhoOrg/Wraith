package commands

import (
	"strconv"
	"strings"

	visualizations "github.com/MateusMoutinhoOrg/Wraith/sandbox/Visualization/Visualization"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/utils"
)

// Start creates a vault: a task file and a visualization config, copied from
// the defaults compiled into the binary. It is the first command anyone runs,
// and it never overwrites — a file already on disk is reported and left
// exactly as it is.
//
// Its three dashboard flags — `--prev-months`, `--future-months` and
// `--current-month` — are written into the `args:` of the DashBoard entry of
// the config it creates. They are a starting point, not a setting: what they
// write is a line in a file you own from then on, and editing that line is
// how it is changed afterwards.
func Start(l *api.Lib, quiet bool) int {
	args, code := dashBoardArgs(l)
	if code != api.ExitOk {
		return code
	}
	existed := l.Deps.IoLib.Exist(l.TaskPath) || l.Deps.IoLib.Exist(l.VisualizationPath)
	options := api.StartOptions{
		VisualizationArgs: map[string]map[string]any{visualizations.DashBoardName: args},
	}
	if err := l.Start(options); err != nil {
		return Failure(l, err)
	}
	if quiet {
		return api.ExitOk
	}
	if existed {
		l.Deps.Printf(config.AlreadyStarted, l.TaskPath)
		return api.ExitOk
	}
	l.Deps.Printf(config.Started, l.TaskPath, l.VisualizationPath)
	return api.ExitOk
}

// dashBoardArgs reads the three dashboard flags into the args block the
// created config carries, falling back to the interface's defaults for the
// ones that were not given. A flag that was given and cannot be read is a
// usage error rather than a silent fallback: a vault is created once, and a
// mistyped horizon is worth stopping for.
func dashBoardArgs(l *api.Lib) (map[string]any, int) {
	previous, code := monthCount(l, config.PrevMonthsFlag, config.StartPrevMonths)
	if code != api.ExitOk {
		return nil, code
	}
	ahead, code := monthCount(l, config.FutureMonthsFlag, config.StartFutureMonths)
	if code != api.ExitOk {
		return nil, code
	}
	current, code := month(l, config.CurrentMonthFlag)
	if code != api.ExitOk {
		return nil, code
	}
	return map[string]any{
		visualizations.PrevMonthsArg:   previous,
		visualizations.FutureMonthsArg: ahead,
		visualizations.CurrentMonthArg: current,
	}, api.ExitOk
}

// monthCount reads one `--flag <count>` as a whole number of months, falling
// back to the given default when the flag is absent.
func monthCount(l *api.Lib, flag string, fallback int) (int64, int) {
	value, err := l.Deps.VerbLib.GetStringOption([]string{flag}, 0)
	if err != nil || strings.TrimSpace(value) == "" {
		return int64(fallback), api.ExitOk
	}
	count, parseErr := strconv.ParseInt(strings.TrimSpace(value), 10, 64)
	if parseErr != nil || count < 0 {
		return 0, UsageError(l, config.InvalidCount, value, flag)
	}
	return count, api.ExitOk
}

// month reads one `--flag <YYYY-MM>`, falling back to the month today falls
// in — read off the injected clock, so the created vault names the month it
// was created in rather than leaving it to be worked out later.
func month(l *api.Lib, flag string) (string, int) {
	value, err := l.Deps.VerbLib.GetStringOption([]string{flag}, 0)
	if err != nil || strings.TrimSpace(value) == "" {
		return utils.MonthText(utils.MonthOf(utils.DateOf(l.Deps.Now()))), api.ExitOk
	}
	parsed, parseErr := utils.ParseMonth(value)
	if parseErr != nil {
		return "", UsageError(l, config.InvalidMonth, value, flag)
	}
	return utils.MonthText(parsed), api.ExitOk
}
