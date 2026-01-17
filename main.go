package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

type Action struct {
	Type     string `yaml:"type"`
	Message  string `yaml:"message,omitempty"`
	Command  string `yaml:"command,omitempty"`
	Username string `yaml:"username,omitempty"`
	URL      string `yaml:"url,omitempty"`
}

type Intent struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Actions     []Action `yaml:"actions"`
}

type IntentSpec struct {
	Intents []Intent `yaml:"intents"`
}

// In-memory user store
var users = make(map[string]bool)

// Execute action based on type
func executeAction(a Action) {
	switch a.Type {
	case "log":
		fmt.Println("[LOG]:", a.Message)
	case "create_user":
		if a.Username != "" {
			users[a.Username] = true
			fmt.Println("[USER]: Created user", a.Username)
			logAction(fmt.Sprintf("Created user: %s", a.Username))
		}
	case "delete_user":
		if a.Username != "" {
			delete(users, a.Username)
			fmt.Println("[USER]: Deleted user", a.Username)
			logAction(fmt.Sprintf("Deleted user: %s", a.Username))
		}
	case "http_get":
		if a.URL != "" {
			client := &http.Client{}
			req, err := http.NewRequest("GET", a.URL, nil)
			if err != nil {
				fmt.Println("[HTTP ERROR]:", err)
				break
			}
			// Set proper header for JSON response
			req.Header.Set("Accept", "application/json")

			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("[HTTP ERROR]:", err)
				break
			}
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println("[HTTP RESPONSE]:", string(body))
		}
	}
}

// Simple log to file
func logAction(entry string) {
	f, err := os.OpenFile("ioba.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Failed to log:", err)
		return
	}
	defer f.Close()
	f.WriteString(entry + "\n")
}

// Execute intent by name
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
	fmt.Println("[WARN]: Intent not found:", intentName)
}

// Load YAML
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

func main() {
	spec := loadIntents("intent.yml")

	// Run all intents
	for _, intent := range spec.Intents {
		executeIntent(spec, intent.Name)
	}

	// Show current users
	usersJSON, _ := json.Marshal(users)
	fmt.Println("Current users:", string(usersJSON))
	fmt.Println("✅ IOBA prototype finished")
}
