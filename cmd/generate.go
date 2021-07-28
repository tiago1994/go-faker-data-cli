/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

	"github.com/spf13/cobra"
	"github.com/tiago1994/fake-data/pkg/generate"
)

var typeFile string
var numberOfLines int32

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate file.",
	Long:  "Command to generate file with fata data.",
	Run: func(cmd *cobra.Command, args []string) {
		fileName := generate.GenerateData(typeFile, numberOfLines)
		fmt.Println(fileName + " successfully generated")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&typeFile, "type", "t", "txt", "txt, csv")
	generateCmd.MarkFlagRequired("type")

	generateCmd.Flags().Int32VarP(&numberOfLines, "number", "n", 100, "Number of lines for the faker file")
	generateCmd.MarkFlagRequired("number")
}
