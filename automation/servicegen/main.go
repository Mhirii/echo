package main

import (
	"fmt"
	"os"
	"path/filepath"

	"servicegen/internal"

	"gopkg.in/yaml.v2"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run gen_service.go <path_to_yaml_file>")
		os.Exit(1)
	}

	yamlFilePath := os.Args[1]
	parentDir := filepath.Dir(yamlFilePath)
	fmt.Printf("Parent directory: %s\n", parentDir)

	proto, err := readAndParseYAML(yamlFilePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	err = internal.CreateDirectories(parentDir)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	err = internal.WriteFile(parentDir+"/internal/server/server.go", internal.WriteServer(proto))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	err = internal.WriteFile(parentDir+"/cmd/api/main.go", internal.WriteMainGo(proto))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	err = internal.WriteResources(parentDir)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	err = internal.WriteProtoFile(parentDir, proto)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func readAndParseYAML(filePath string) (internal.Template, error) {
	var proto internal.Template

	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return proto, fmt.Errorf("error reading YAML file: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &proto)
	if err != nil {
		return proto, fmt.Errorf("error unmarshalling YAML file: %v", err)
	}

	return proto, nil
}
