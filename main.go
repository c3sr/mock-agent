package main

import (
	"fmt"
	"github.com/c3sr/dlframework"
	cmd "github.com/c3sr/dlframework/framework/cmd/server"
	"io/ioutil"
	"mock-agent/mockagent"
	"os"
)

func main() {
	fmt.Println("Starting up Mock-Agent...")

	framework := mockagent.FrameworkManifest
	mockWorkHandler := MockWork{}
	rootCmd, err := cmd.NewMqWorkerServer(framework, mockWorkHandler)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println("Clean Exit.")
}

type MockWork struct{
}
func (m MockWork) ProcessJob(incomingRequest string, manifest dlframework.FrameworkManifest) (string, error) {
	fmt.Println("New Job : " + incomingRequest)

	filenameToUse := os.Getenv("MOCK_FILE")
	if filenameToUse == "" {
		filenameToUse = "classification.json"
	}
	data, err := ioutil.ReadFile("mock-replies/" + filenameToUse)
	if err != nil {
		return `{"error":"unable to open mock file to send a real reply"}`, nil
	}

	return string(data), nil
}