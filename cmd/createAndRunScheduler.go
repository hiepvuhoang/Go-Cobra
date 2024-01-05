/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

// createAndRunSchedulerCmd represents the createAndRunScheduler command
var createAndRunSchedulerCmd = &cobra.Command{
	Use:   "createAndRunScheduler",
	Short: "Create and trigger a scheduler",
	Long:  `Create a scheduler then triggering it`,
	Run: func(cmd *cobra.Command, args []string) {
		client := &http.Client{}
		var data = strings.NewReader(`{
    "name": "Untitled Test Run",
    "teamId": "430",
    "projectId": "510",
    "testProjectId": 20208,
    "testCloudAgents": [
        {
            "cloudType": "TEST_CLOUD_AGENT",
            "enabledTestCloudTunnel": false,
            "runType": "DESKTOP_BROWSERS",
            "os": "windows",
            "browser": "chrome",
            "browserVersion": "119",
            "browserType": "ALL",
            "metadata": {}
        }
    ],
    "cloudType": "TEST_CLOUD_AGENT",
    "browserType": "ALL",
    "executionProfileId": 41085,
    "configType": "TEST_SUITE",
    "testSuiteCollectionId": 73527,
    "testSuiteId": 177094,
    "ksVersion": "latest",
    "timeOut": 180,
    "executionMode": "SEQUENTIAL",
    "enabledKobitonIntegration": false,
    "enabledTestCloudTunnel": false,
    "baselineCollectionGroupOrder": null,
    "xrayImportReportType": "PUSH_MANUALLY",
    "testRunConfig": {
        "executionType": "G4",
        "executionMetadata": {
            "agentType": "TEST_CLOUD",
            "engine": {
                "type": "G4",
                "version": "latest"
            },
            "config": {
                "testSuiteId": 177094,
                "testSuiteCollectionId": 73527,
                "executionEnvs": [
                    {
                        "platform": "windows",
                        "browser": "chrome",
                        "browserVersion": "119",
                        "framework": "SELENIUM"
                    }
                ]
            }
        },
        "settings": {
            "executionMode": "SEQUENTIAL",
            "enabledKobitonIntegration": false,
            "enabledTestCloudTunnel": false
        },
        "scheduler": {
            "startTime": "2024-01-05T02:33:15.359Z",
            "repeatEnable": false,
            "active": false
        }
    }
}
`)
		req, err := http.NewRequest("POST", "https://testops.qa.katalon.com/api/v1/smart-scheduler/run?verifyRunConfiguration=true", data)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Basic aGllcC52dUBrYXRhbG9uLmNvbTpOdGNoYW5nMTIxMDk5Lg==")
		req.Header.Set("Cookie", "segment-write-key=HkOyi1cO8X1owpp2juM39D4Y6o8AQ3l2")
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", bodyText)
	},
}

func init() {
	rootCmd.AddCommand(createAndRunSchedulerCmd)

	createAndRunSchedulerCmd.PersistentFlags()
}
