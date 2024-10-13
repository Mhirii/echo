package internal

import (
	"fmt"
	"os"
)

func CreateDirectories(path string) error {
	dirs := []string{
		"/internal",
		"/internal/server",
		"/internal/lib",
		"/internal/database",
		"/proto",
		"/cmd",
		"/cmd/api",
	}

	for _, dir := range dirs {
		err := os.Mkdir(path+dir, 0o755)
		if err != nil {
			return err
		}
	}

	return nil
}

func WriteProtoFile(path string, proto Template) error {
	content := proto.GenProtoFile()
	fmt.Println(content)

	p := path + "/proto/" + proto.Pkg + ".proto"

	file, err := os.Create(p)
	if err != nil {
		return err
	}

	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

func WriteFile(path, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
