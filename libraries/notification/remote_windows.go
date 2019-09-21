package notification

import (
	"github.com/walterjwhite/go-application/libraries/logging"
	"github.com/walterjwhite/go-application/libraries/runner"
	"gopkg.in/toast.v1"

	"context"
	"os/exec"
	"os/user"
)

type remoteNotification struct {
	Username string
	Server   string

	// acknowledge message, wait for seconds, 0 (do not wait)
	TimeWait uint

	Context context.Context
}

func (n *remoteNotification) Notify(notification Notification) {
	cmd := "msg"
	arguments := make([]string, 0)

	arguments = append(arguments, n.Username)
	arguments = append(arguments, fmt.Sprintf("/Server:%v", n.Server))

	if n.TimeWait > 0 {
		arguments = append(arguments, fmt.Sprintf("/Time:%v", n.TimeWait))
	}

	arguments = append(arguments, format(notification))

	_, err := runner.Run(n.Context, cmd, arguments...)
	logging.Panic(err)
}

func format(notification Notification) string {
	return notification.ApplicationId + ":" + notification.Title + ":" + notification.Description
}

func New(ctx context.Context, targetServer string) Notifier {
	username, err := user.Current()
	logging.Panic(err)

	return &remoteNotification{Username: username, Server: targetServer, Context: ctx}
}