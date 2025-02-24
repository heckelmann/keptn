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
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/keptn/keptn/cli/internal"

	"github.com/spf13/cobra"

	versionCheck "github.com/hashicorp/go-version"
	"github.com/keptn/keptn/cli/pkg/config"
	"github.com/keptn/keptn/cli/pkg/credentialmanager"
	"github.com/keptn/keptn/cli/pkg/logging"
	"github.com/keptn/keptn/cli/pkg/version"
)

var (
	// Version information which is passed by ldflags
	Version             string
	keptnReleaseDocsURL string
	versionCheckInfo    string
)

const setVersionCheckMsg = `* To %s the daily version check, please execute:
 - keptn set config AutomaticVersionCheck %s

`

const disableVersionCheckMsg = "To disable this notice, run: '%s set config AutomaticVersionCheck false'"

// KubeServerVersionConstraints the Kubernetes Cluster version's constraints is passed by ldflags
var KubeServerVersionConstraints string

// versionCmd represents the version command
var versionCmd = NewVersionCommand(version.NewVersionChecker())

func NewVersionCommand(vChecker *version.VersionChecker) *cobra.Command {
	versionCmd := &cobra.Command{
		Use:     "version",
		Args:    cobra.NoArgs,
		Short:   "Shows the version of Keptn and Keptn CLI",
		Long:    `Shows the version of Keptn and Keptn CLI, and a note when a new version is available.`,
		Example: `keptn version`,
		Run: func(cmd *cobra.Command, args []string) {
			isLastCheckStale, err := isLastCheckStale()
			if err != nil {
				logging.PrintLog(err.Error(), logging.InfoLevel)
				return
			}

			var cliChecked, keptnChecked bool

			// Keptn CLI
			fmt.Println("\nKeptn CLI version: " + Version)
			if isLastCheckStale {
				cliChecked, _ = vChecker.CheckCLIVersion(Version, false)
			}

			// Keptn
			fmt.Print("Keptn cluster version: ")
			keptnVersion, err := getKeptnServerVersion()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(keptnVersion)
				if isLastCheckStale {
					kvChecker := version.NewKeptnVersionChecker()
					keptnChecked, _ = kvChecker.CheckKeptnVersion(Version, keptnVersion, false)
				}
			}

			if cliChecked || keptnChecked {
				updateLastVersionCheck()
			}

			if err := printDailyVersionCheckInfo(); err != nil {
				logging.PrintLog(err.Error(), logging.InfoLevel)
				return
			}
		},
	}
	return versionCmd
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func getReleaseDocsURL() string {
	if keptnReleaseDocsURL != "" {
		return keptnReleaseDocsURL
	}
	v, err := versionCheck.NewSemver(Version)
	keptnReleaseDocsURL = "0.8.x" //fallback version if provided doc version is invalid
	if err == nil {
		segments := v.Segments()
		if len(segments) == 3 {
			keptnReleaseDocsURL = strconv.Itoa(segments[0]) + "." + strconv.Itoa(segments[1]) + ".x"
		}
	}

	versionCheckInfo = `Daily version check is %s. 
Keptn will%s collect statistical data and will%s notify about new versions and security patches for Keptn. Details can be found at: https://keptn.sh/docs/` + keptnReleaseDocsURL + `/reference/version_check
---------------------------------------------------
`
	return keptnReleaseDocsURL
}

func isLastCheckStale() (bool, error) {
	configMng := config.NewCLIConfigManager("")
	cliConfig, err := configMng.LoadCLIConfig()
	if err != nil {
		return false, err
	}
	return cliConfig.LastVersionCheck == nil || time.Now().Sub(*cliConfig.LastVersionCheck) >= time.Second, nil
}

func printDailyVersionCheckInfo() error {
	configMng := config.NewCLIConfigManager("")
	cliConfig, err := configMng.LoadCLIConfig()
	if err != nil {
		return err
	}
	fmt.Println()
	if cliConfig.AutomaticVersionCheck {
		fmt.Printf(versionCheckInfo, "enabled", "", "")
		fmt.Printf(setVersionCheckMsg, "disable", "false")
	} else {
		fmt.Printf(versionCheckInfo, "disabled", " not", " not")
		fmt.Printf(setVersionCheckMsg, "enable", "true")
	}
	return nil
}

func updateLastVersionCheck() {
	configMng := config.NewCLIConfigManager("")
	cliConfig, err := configMng.LoadCLIConfig()
	if err != nil {
		logging.PrintLog(err.Error(), logging.InfoLevel)
		return
	}
	currentTime := time.Now()
	cliConfig.LastVersionCheck = &currentTime
	if err := configMng.StoreCLIConfig(cliConfig); err != nil {
		logging.PrintLog(err.Error(), logging.InfoLevel)
	}
}

func getKeptnServerVersion() (string, error) {
	var endPoint url.URL
	var apiToken string
	var err error
	if !mocking {
		endPoint, apiToken, err = credentialmanager.NewCredentialManager(assumeYes).GetCreds(namespace)
	} else {
		endPointPtr, _ := url.Parse(os.Getenv("MOCK_SERVER"))
		endPoint = *endPointPtr
		apiToken = ""
	}

	if err != nil {
		return "", errors.New(authErrorMsg)
	}

	api, err := internal.APIProvider(endPoint.String(), apiToken, nil)
	if err != nil {
		return "", err
	}

	metadataData, errMetadata := api.APIV1().GetMetadata()
	if errMetadata != nil {
		if errMetadata.Message != nil {
			return "", errors.New("Error occurred with response code " + strconv.FormatInt(errMetadata.Code, 10) + " with message " + *errMetadata.Message)
		}
		return "", errors.New("received invalid response from Keptn API")
	}
	return metadataData.Keptnversion, nil
}
