// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package cmd_ports

import (
	"context"
	"dagent/client"
	"dagent/grpc/proto"
	"dagent/internal/util"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var stopPortForwardCmd = &cobra.Command{
	Use:   "stop [WORKSPACE_NAME] [PROJECT_NAME] -p [PORT]",
	Short: "Stop forwarding a port",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		conn, err := client.GetConn(nil)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		projectName := ""

		if len(args) == 2 {
			projectName = args[1]
		} else {
			projectName, err = util.GetFirstWorkspaceProjectName(conn, args[0], projectName)
			if err != nil {
				log.Fatal(err)
			}
		}

		client := proto.NewPortsClient(conn)

		stopPortForwardRequest := &proto.StopPortForwardRequest{
			WorkspaceName: args[0],
			Project:       projectName,
			Port:          uint32(portArg),
		}

		_, err = client.StopPortForward(ctx, stopPortForwardRequest)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Port forwarding stopped.")
	},
}

func init() {
	stopPortForwardCmd.Flags().IntVarP(&portArg, "port", "p", 0, "Port to forward")
	stopPortForwardCmd.MarkFlagRequired("port")
}