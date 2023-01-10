package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ComposerJson struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func main() {
	dirEntries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, dir := range dirEntries {
		if dir.IsDir() {
			// 1. read version

			composerJson, err := os.ReadFile(
				fmt.Sprintf("./%s/composer.json", dir.Name()),
			)
			if err != nil {
				panic(err)
			}

			composerData := &ComposerJson{}
			err = json.Unmarshal(composerJson, &composerData)
			if err != nil {
				panic(err)
			}

			fmt.Println(composerData.Name, composerData.Version)
			//fmt.Println(dir.Name(), composerData.Version)

			// 2 run git tag $(version) in this dir
			//cmd := exec.Command("git", "tag", composerData.Version)
			//cmd.Dir = fmt.Sprintf("./%s", dir.Name())
			//cmd.Stdout = os.Stdout
			//cmd.Start()
		}

	}
}
