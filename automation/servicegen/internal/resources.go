package internal

import "os"

func WriteResources(path string) error {
	resources := []string{
		"Makefile",
		".air.toml",
		".env.example",
		".gitignore",
		".goreleaser.yml",
		"docker-compose.yml",
	}
	for _, resource := range resources {
		err := writeResourceFile(path, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeResourceFile(path, resourceName string) error {
	currentDir := os.Getenv("PWD")
	from := currentDir + "/resources/" + resourceName
	err := copy(from, path+"/"+resourceName)
	if err != nil {
		return err
	}
	return nil
}

func copy(from, dest string) error {
	filecontent, err := os.ReadFile(from)
	if err != nil {
		return err
	}

	err = os.WriteFile(dest, filecontent, 0o644)
	if err != nil {
		return err
	}
	return nil
}
