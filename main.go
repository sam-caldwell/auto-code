// main.go (Generator Entrypoint)
package main

import (
	"fmt"
	generator "github.com/sam-caldwell/auto-code/generator"
	manifest2 "github.com/sam-caldwell/auto-code/manifest"
	"log"
)

func main() {

	manifest, err := manifest2.LoadManifest("manifest.yaml")

	if err != nil {
		log.Fatal("Failed to load manifest:", err)
	}

	err = generator.GenerateCode(manifest)

	if err != nil {
		log.Fatal("Code generation failed:", err)
	}

	fmt.Println("Code generation complete!")
}
