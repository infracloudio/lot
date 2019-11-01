/*
Copyright © 2019 LoT Authors

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
	"fmt"
	"os"

	"github.com/infracloudio/lot/pkg/util/bpfutil"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var slistCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the out-of-box ready to use eBPF snippets",
	Long:  `lot bpf snippet list will list all of the ready to use eBPF snippets provided out-of-box`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println("Invalid number of arguments")
			fmt.Println("lot bpf snippet list does not take any argument")
			// Refer https://www.tldp.org/LDP/abs/html/exitcodes.html#EXITCODESREF
			os.Exit(126)
		}

		bpfutil.ListSnippets()
	},
}

func init() {
	snippetCmd.AddCommand(slistCmd)
}
