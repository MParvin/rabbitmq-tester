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

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "This command will send messages to a queue",
	Long: `To send a message to a queue, use the following command:

rabbitmq-test -c /PATH/TO/CONFIG_FILE send`,
	Run: func(cmd *cobra.Command, args []string) {
		L.Send(viper.GetString("RabbitHost"), viper.GetString("Queue"), viper.GetString("Message"))
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

}
