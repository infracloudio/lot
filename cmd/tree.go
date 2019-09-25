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
	"github.com/arush-sal/lot/pkg/util/procutil"
	"github.com/spf13/cobra"
)

// treeCmd represents the tree command
var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Prints a process tree with process ancestral hierarchy",
	Long:  `lot process tree prints the hierarchical tree of processes`,
	Run: func(cmd *cobra.Command, args []string) {
		procutil.PrintProcessTree()
	},
}

func init() {
	processCmd.AddCommand(treeCmd)

}
