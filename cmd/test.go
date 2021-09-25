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
	"github.com/seanlan/lazy/generator"
	"github.com/spf13/cobra"
	"os/exec"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use: "test",
	Run: func(cmd *cobra.Command, args []string) {

		dbStr := "root:q145145145@tcp(127.0.0.1:3306)/mutual?parseTime=true&loc=Local&charset=utf8mb4&collation=utf8mb4_unicode_ci"
		tableName := "mutual"
		outPath := "./sqlmodel"
		packageName := "sqlmodel"
		tmplPath := "./templates"

		g := generator.NewGormGenerator(dbStr, tableName, packageName, tmplPath, outPath)
		g.Gen()
		osCmd := exec.Command("gofmt", "-s", "-d", "-w", ".")
		osCmd.Run()
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
