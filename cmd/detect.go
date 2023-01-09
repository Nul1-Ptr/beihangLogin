/*
Copyright © 2021 Izumiko <yosoro@outlook.com>

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
	"beihangLogin/util"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// detectCmd represents the detect command
var detectCmd = &cobra.Command{
	Use:   "detect",
	Short: "Detect online status and try to login",
	//	Long: `A longer description that spans multiple lines and likely contains examples
	//and usage of using your command. For example:
	//
	//Cobra is a CLI library for Go that empowers applications.
	//This application is a tool to generate the needed files
	//to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		isOnline := true
		res, err := util.Status()
		if err != nil {
			fmt.Println("Get online info failed:", err)
			isOnline = false
		}
		if res["error"] != "ok" {
			isOnline = false
		}
		if isOnline {
			fmt.Println("User is online. Skip")
		} else {
			fmt.Println("User is not online. Try to login...")
			// try to log in for 4 times
			for i := 0; i < 4; i++ {
				err = util.Login(username, password)
				if err != nil {
					fmt.Println("Login failed:", err)
					time.Sleep(2 * time.Second)
				} else {
					fmt.Println("Login success!")
					break
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(detectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// detectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// detectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	detectCmd.Flags().StringVarP(&username, "username", "u", "", "Username of your account. (required)")
	detectCmd.Flags().StringVarP(&password, "password", "p", "", "Password of your account. (required)")
	_ = detectCmd.MarkFlagRequired("username")
	_ = detectCmd.MarkFlagRequired("password")
}
