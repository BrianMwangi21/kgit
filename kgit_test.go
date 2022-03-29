package main

import (
	"testing"
)

// Test for Import Config Function
func TestImportConfig(t *testing.T) {
  got := import_config()

  // Check if what we go is > 0
  if !(got > 0) {
    t.Errorf("Got %d, Wanted > 0", got)
  }
}

// Test for Getting Current Directory
func TestGetCurrentDirectory(t *testing.T) {
  got := getCurrentDir()

  // Check if length > 0
  if !(len(got) > 0) {
    t.Errorf("Got %d, Wanted len(%s) > 0", len(got), got)
  }
}

// Test for Getting Git Logs 
func TestGetGitLogs(t *testing.T) {
  got := getGitLogs()

  // Check if length > 0
  if !(len(got) > 0) {
    t.Errorf("Got %d, Wanted len(%s) > 0", len(got), got)
  }
}

// Test for Getting Git Branches 
func TestGetGitBranches(t *testing.T) {
  got := getGitBranches()  

  // Check if length > 0
  if !(len(got) > 0) {
    t.Errorf("Got %d, Wanted len(%s) > 0", len(got), got)
  }
}

// Test for Checking Current Branch
func TestCheckCurrentBranch(t *testing.T) {
  got := checkCurrentBranch("* master")  

  // Check if false
  if !got {
    t.Errorf("Got false, wanted true")
  }
}

// Test for Checking if File Exists
func TestCheckFileExists(t *testing.T) {
  got := checkFileExists("README.md")  

  // Check if false
  if !got {
    t.Errorf("Got false, wanted true")
  }
}

// Test for Checking the file contents
func TestCheckFileContents(t *testing.T) {
  got := checkFileContents("README.md", "# Welcome to k-git") 

  // Check if false
  if !got {
    t.Errorf("Got false, wanted true")
  }
}

type finalResultTestStruct struct {
  arg1, arg2, arg3, arg4 bool 
  expected interface{}
}

// Test for Checking the Final Messages
func TestFinalResults(t *testing.T) {
  import_config()

  var finalResultsTests = []finalResultTestStruct{
      {true, true, true, true, payload["check_all_pass"]},
      {false, true, true, true, payload["check_one_fail"]},
      {true, false, true, true, payload["check_two_fail"]},
      {true, true, false, true, payload["check_three_fail"]},
      {true, true, true, false, payload["check_four_fail"]},
  }

  for _, test := range finalResultsTests{
      if output := displayFinalResults(test.arg1, test.arg2, test.arg3, test.arg4); output != test.expected {
          t.Errorf("Got %v, Wanted %v", output, test.expected)
      }
  }
}
