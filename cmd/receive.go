/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	L "github.com/mparvin/rabbitmq-tester/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// receiveCmd represents the receive command
var receiveCmd = &cobra.Command{
	Use:   "receive",
	Short: "This command will receive messages from a queue",
	Long: `To receive a message from a queue, use the following command:

rabbitmq-test -c /PATH/TO/CONFIG_FILE receive`,
	Run: func(cmd *cobra.Command, args []string) {
		L.Receive(viper.GetString("RabbitHost"), viper.GetString("Queue"), viper.GetString("Message"))

	},
}

func init() {
	rootCmd.AddCommand(receiveCmd)

}
