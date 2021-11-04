package main

import (
	"fmt"
	"github.com/AtlasInsideCorp/UTMStackCloudWazuhApi"
	"os"
)

func main() {
	endpoint := "https://localhost:55000/"
	user := "wazuh"
	pass := "wazuh"

	client, err := wazuh.New(endpoint, wazuh.WithBasicAuth(user, pass))

	agents, err := client.GetAllAgents()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%-20v%-20v\n", "ID", "NAME")
	for _, agent := range *agents {
		if agent.Status == "Active" {
			fmt.Printf("%-20v%-20v\n", agent.ID, agent.Name)
		}
	}
}
