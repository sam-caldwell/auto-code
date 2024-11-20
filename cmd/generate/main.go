package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Database represents the x-database section.
type Database struct {
	Driver string   `yaml:"driver"`
	Server string   `yaml:"server"`
	Tables []string `yaml:"tables"`
}

// resolveRefs resolves $ref links in the parsed YAML.
func resolveRefs(data map[string]interface{}, basePath string) error {
	for key, value := range data {
		switch v := value.(type) {
		case map[string]interface{}:
			// Recursively resolve nested objects
			if ref, ok := v["$ref"].(string); ok {
				resolvedContent, err := loadRef(ref, basePath)
				if err != nil {
					return fmt.Errorf("failed to resolve $ref %s: %v", ref, err)
				}
				data[key] = resolvedContent
			} else {
				if err := resolveRefs(v, basePath); err != nil {
					return err
				}
			}
		case []interface{}:
			// Recursively resolve arrays
			for i, item := range v {
				if m, ok := item.(map[string]interface{}); ok {
					if err := resolveRefs(m, basePath); err != nil {
						return err
					}
					v[i] = m
				}
			}
		}
	}
	return nil
}

// loadRef loads the content of a $ref file.
func loadRef(ref string, basePath string) (map[string]interface{}, error) {
	// Resolve relative path
	refPath := filepath.Join(basePath, ref)
	file, err := os.Open(refPath)
	if err != nil {
		return nil, fmt.Errorf("cannot open ref file: %v", err)
	}
	defer file.Close()

	var content map[string]interface{}
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&content); err != nil {
		return nil, fmt.Errorf("failed to decode ref content: %v", err)
	}
	return content, nil
}

func main() {
	// Open the YAML file
	file, err := os.Open("openapi.yaml")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Parse the YAML
	var openAPI map[string]interface{}
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&openAPI); err != nil {
		fmt.Printf("Error decoding YAML: %v\n", err)
		return
	}

	// Resolve $ref links
	basePath := filepath.Dir("openapi.yaml")
	if err := resolveRefs(openAPI, basePath); err != nil {
		fmt.Printf("Error resolving $ref links: %v\n", err)
		return
	}

	// Print the resolved OpenAPI spec
	fmt.Printf("Resolved OpenAPI YAML:\n%+v\n", openAPI)
}
