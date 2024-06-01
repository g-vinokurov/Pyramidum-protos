package main

import (
	"io/fs"
	"log"
	"os/exec"
	"path/filepath"
)

func main() {
	targetDir := "proto"

	finalErr := filepath.Walk(targetDir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		cmd := exec.Command(
			"protoc",
			"-I",
			targetDir,
			path,
			"--go_out=./gen/go",
			"--go_opt=paths=source_relative",
			"--go-grpc_out=./gen/go",
			"--go-grpc_opt=paths=source_relative",
		)
		err = cmd.Run()

		if err != nil {
			log.Print(err)
		}

		return nil
	})
	if finalErr != nil {
		log.Fatal(finalErr)
	}
}
