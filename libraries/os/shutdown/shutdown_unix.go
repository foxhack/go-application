package shutdown

import (
	"github.com/rs/zerolog/log"

	"fmt"
	"github.com/walterjwhite/go-application/libraries/application"
	"github.com/walterjwhite/go-application/libraries/application/logging"
	"github.com/walterjwhite/go-application/libraries/utils/runner"
)

func (r *ShutdownRequest) log() {
	command := r.getShutdownAction()
	log.Info().Msgf("Arguments: %v", command)
}

func (r *ShutdownRequest) getShutdownAction() string {
	if r.ShutdownAction == Reboot {
		return "reboot"
	} else if r.ShutdownAction == Poweroff {
		return "poweroff"
	} else {
		logging.Panic(fmt.Errorf("Unknown option specified: %v\n", r.ShutdownAction))
		return ""
	}
}

func (r *ShutdownRequest) run() {
	command := r.getShutdownAction()
	_, err := runner.Run(application.Context, command)
	logging.Panic(err)
}
