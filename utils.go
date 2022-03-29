package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

// askForConfirmation uses Scanln to parse user input. A user must type in "yes" or "no" and
// then press enter. It has fuzzy matching, so "y", "Y", "yes", "YES", and "Yes" all count as
// confirmations. If the input is not recognized, it will ask again. The function does not return
// until it gets a valid response from the user. Typically, you should use fmt to print out a question
// before calling askForConfirmation. E.g. fmt.Println("WARNING: Are you sure? (yes/no)")
func askForConfirmation(custom_message string) bool {
  fmt.Print("\n")
  if len(custom_message) > 0 { 
    fmt.Printf("%s (yes/no) => ", custom_message)
  } else {
    fmt.Print("CONFIRM: Do you wish to proceed ? (yes/no) => ")
  }
  var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	nokayResponses := []string{"n", "N", "no", "No", "NO"}
	if containsString(okayResponses, response) {
		return true
	} else if containsString(nokayResponses, response) {
		return false
	} else {
		fmt.Println("Please type yes or no and then press enter:")
		return askForConfirmation(custom_message)
	}
}

// posString returns the first index of element in slice.
// If slice does not contain element, returns -1.
func posString(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
}

// containsString returns true iff slice contains element
func containsString(slice []string, element string) bool {
	return !(posString(slice, element) == -1)
}

// Global config payload
var payload map[string]interface{}

// Import Config Function
func import_config() int {
  // Let's first read the `config.json` file
  content, err := ioutil.ReadFile("./config.json")
  if err != nil {
    log.Fatal("Error when opening file: ", err)
  }

  // Now let's unmarshall the data into `payload`
  err = json.Unmarshal(content, &payload)
  if err != nil {
    log.Fatal("Error during Unmarshal(): ", err)
  }

  return len(payload)
}

// Function to get the user selection when it comes to the topics
func get_user_selection() int {
  // Show topics and get the user selection 
  fmt.Println(payload["prompts_topics"])
  var user_selection int   
  fmt.Print("\nEnter the number of the topic you want to learn => ")
  _, err := fmt.Scan(&user_selection)
	if err != nil {
		log.Fatal(err)
	}
  fmt.Printf("\nYou selected %d. Let's gooooo!\n\n", user_selection)

  return user_selection
}

// Function to get the current directory
func getCurrentDir() string {
  output, err := exec.Command("pwd").Output()

  if err != nil {
      log.Fatal(err)
  }

  return strings.TrimSpace(string(output))
}

// Function to clear the screen, let's see 
func clearMyScreen() {
  fmt.Print("\033[H\033[2J")
}

// Function to get git logs
func getGitLogs() string {
  output, err := exec.Command("git", "log", "--oneline").Output()

  if err != nil {
      log.Fatal(err)
  }

  return strings.TrimSpace(string(output))
}

// Function to get git branches
func getGitBranches() string {
  output, err := exec.Command("git", "branch").Output()

  if err != nil {
      log.Fatal(err)
  }

  return strings.TrimSpace(string(output))
}

// Function to check if file exists
func checkCurrentBranch(branchname string) bool {
  if strings.Contains(getGitBranches(), branchname) == false {
    return false 
  }

  return true 
}

// Function to check if file exists
func checkFileExists(filename string) bool {
  path_to_file := getCurrentDir() + "/" + filename 

  if _, err := os.Stat(path_to_file); os.IsNotExist(err) {
    return false 
  }
  
  return true
}

// Function to check if contents match 
func checkFileContents(filename string, requested_text string) bool {
  content, err := ioutil.ReadFile(filename)

  if err != nil {
    log.Fatal(err)
  }
  
  if strings.Contains(strings.TrimSpace(string(content)), requested_text) == false {
    return false     
  }

  return true
}

// Function to check if commit was done
func checkCommitMessage(commit string) bool {
  if strings.Contains(getGitLogs(), commit) == false {
    return false
  }

  return true
}

// Function to return the check response if they are 4 
func displayFinalResults(check_one bool, check_two bool, check_three bool, check_four bool) interface{} {
  if !check_one {
    return payload["check_one_fail"] 
  }

  if !check_two {
    return payload["check_two_fail"]
  }

  if !check_three {
    return payload["check_three_fail"]
  }

  if !check_four {
    return payload["check_four_fail"]
  }

  return payload["check_all_pass"]
}
