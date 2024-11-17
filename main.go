package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// OpenAPI represents the structure of the OpenAPI spec.
type OpenAPI struct {
	OpenAPI     string                 `yaml:"openapi"`
	Info        Info                   `yaml:"info"`
	Servers     []Server               `yaml:"servers"`
	Paths       map[string]interface{} `yaml:"endpoints"`
	Components  Components             `yaml:"components"`
	XContainers []string               `yaml:"x-containers"`
	XDatabase   Database               `yaml:"x-database"`
}

// Info represents the information block in the OpenAPI spec.
type Info struct {
	Description string `yaml:"description"`
	Title       string `yaml:"title"`
	Version     string `yaml:"version"`
}

// Server represents a server in the OpenAPI spec.
type Server struct {
	URL         string `yaml:"url"`
	Description string `yaml:"description"`
}

// Components represents the components section.
type Components struct {
	SecuritySchemes map[string]SecurityScheme `yaml:"securitySchemes"`
}

// SecurityScheme represents a security scheme.
type SecurityScheme struct {
	Type         string `yaml:"type"`
	Scheme       string `yaml:"scheme"`
	BearerFormat string `yaml:"bearerFormat"`
	Description  string `yaml:"description"`
}

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
