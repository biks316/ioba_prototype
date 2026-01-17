package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	"gopkg.in/yaml.v3"
)

// Action defines a single action inside an intent
type Action struct {
	Type    string `yaml:"type"` // "log" or "execute"
	Message string `yaml:"message,omitempty"`
	Command string `yaml:"command,omitempty"`
}

// Intent defines a high-level goal
type Intent struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Actions     []Action `yaml:"actions"`
}

// IntentSpec holds all intents
type IntentSpec struct {
	Intents []Intent `yaml:"intents"`
}

// Executes a single action
func executeAction(a Action) {
	switch a.Type {
	case "log":
		fmt.Println("[LOG]:", a.Message)
	case "execute":
		fmt.Println("[EXEC]:", a.Command)
		out, err := exec.Command("bash", "-c", a.Command).Output()
		if err != nil {
			log.Println("Error executing command:", err)
		} else {
			fmt.Println(string(out))
		}
	default:
		fmt.Println("Unknown action type:", a.Type)
	}
}

// Executes an intent by name
func executeIntent(spec IntentSpec, intentName string) {
	for _, intent := range spec.Intents {
		if intent.Name == intentName {
			fmt.Println("Executing Intent:", intent.Name)
			for _, action := range intent.Actions {
				executeAction(action)
			}
			return
		}
	}
	fmt.Println("Intent not found:", intentName)
}

// Helper function: load YAML
func loadIntents(filePath string) IntentSpec {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Failed to read YAML:", err)
	}

	var spec IntentSpec
	if err := yaml.Unmarshal(data, &spec); err != nil {
		log.Fatal("Failed to parse YAML:", err)
	}
	return spec
}

// Starter template: Generate a PoC prototype
func main() {
	// Load intents from YAML
	spec := loadIntents("intent.yml")

	// Example: iterate over all intents and execute
	for _, intent := range spec.Intents {
		executeIntent(spec, intent.Name)
	}

	fmt.Println("✅ IOBA prototype execution finished")
}
