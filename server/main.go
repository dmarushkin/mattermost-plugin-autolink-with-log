package main

import (
	"github.com/mattermost/mattermost-server/v6/plugin"

	"github.com/dmarushkin/mattermost-plugin-autolink-with-log/server/autolinkplugin"
)

func main() {
	plugin.ClientMain(autolinkplugin.New())
}
