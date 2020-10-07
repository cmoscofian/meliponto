package main

import (
	"fmt"
	"time"

	"github.com/cmoscofian/meliponto/src/cli/command"
	"github.com/cmoscofian/meliponto/src/shared/infrastructure/restclient"
	"github.com/cmoscofian/meliponto/src/shared/repositories/rest"
	"github.com/cmoscofian/meliponto/src/shared/util/constant"
)

func buildCommands() []command.Command {
	commands := make([]command.Command, 0)

	defaultClient := restclient.NewRestClientPool(constant.BaseURI, nil, time.Second)

	// Communication layer
	loginClient := rest.NewLogin(defaultClient)
	fetchClient := rest.NewFetchPunch(defaultClient)
	createClient := rest.NewCreatePunch(defaultClient)
	deleteClient := rest.NewDeletePunch(defaultClient)

	fmt.Println(loginClient, fetchClient, createClient, deleteClient)

	return commands
}
