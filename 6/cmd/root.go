// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"

	"github.com/spf13/cobra"
	"github.com/blang/semver"

	"k8s.io/helm/pkg/proto/hapi/chart"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chart-bump",
	Short: "Bumps the chart version of a 2.X Helm Chart",

	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Chart Bump requires two arguments, a path to a Helm chart, and an argument to determine which part of the SemVer to bump")
			os.Exit(1)
		}

		chartPath := args[0]
		semVerPart := args[1]

		//Load the Helm Chart
		metadata := &chart.Metadata{}

		f, err := os.Open(fmt.Sprintf("%v/Chart.yaml",chartPath))
		if err != nil{
			fmt.Printf("Error reading Chart.yaml file in %v: %v\n", chartPath, err)
			os.Exit(1)
		}
		
		b, err := ioutil.ReadAll(f)
		f.Close()
		if err != nil{
			fmt.Printf("Error Reading Chart.yaml file in %v: %v\n", chartPath, err)
			os.Exit(1)
		}

		err = yaml.Unmarshal(b, metadata)
		semVer, err := semver.Make(metadata.Version)

		switch(semVerPart){
		case "p":
			semVer.Patch++
		case "m":
			semVer.Minor++
			semVer.Patch = 0
		case "M":
			semVer.Major++
			semVer.Minor = 0
			semVer.Patch = 0
		default:
			fmt.Println("Must provide which value to increment:  Major (M), Minor (m), or Patch (p)")
		}
		fmt.Printf("Upgrading Chart version from %v to %v\n", metadata.Version, semVer)

		metadata.Version = semVer.String()

		//save the file
		outBytes, err := yaml.Marshal(metadata)
		if err != nil{
			fmt.Printf("Error marshaling output data: %v\n",err)
			os.Exit(1)
		}
		err = ioutil.WriteFile(fmt.Sprintf("%v/Chart.yaml",chartPath),outBytes, 0644)
		if err != nil{
			fmt.Printf("Error writing updated chart file: %v\n",err)
		}
	 },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

