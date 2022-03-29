package main

import (
	"fmt"
)

// Introduction
func kgit_introduction() {
  fmt.Println(payload["prompts_welcome_text"])
  fmt.Println(payload["prompts_description_text"])
}

// Topic : Initialize/Reinitialize Git Repo 
func kgit_topic_init_repo() {
  fmt.Println(payload["topic_one_title"])
  fmt.Println(payload["topic_one_summary"])
  fmt.Println(payload["topic_one_description"])
  fmt.Println(payload["topic_one_test"])
  
  if askForConfirmation("Have you managed to finish the test ?") == true {
    if checkFileExists(".git/") {
      fmt.Println("\nCongratulations! You made it! Awesome sauce!")
    }else {
      fmt.Println("\nIt seems you did not pass the test. The directory '.git/' does not seem to exist. Please try again")
    }
  } else {
    fmt.Println("Don't you worry. You can come back anytime and proceed!")
  }
}

// Topic : Create a file and commit in the master branch
func kgit_topic_create_and_commit() {
  fmt.Println(payload["topic_two_title"])
  fmt.Println(payload["topic_two_summary"])
  fmt.Println(payload["topic_two_description"])
  fmt.Println(payload["topic_two_test"])
  
  if askForConfirmation("Have you managed to finish the test ?") == true {
    // 1. Check if we are on the master branch 
    check_one := checkCurrentBranch("* master")

    // 2. Check if file exists 
    check_two := checkFileExists("kairitu.txt")

    // 3. Check if the file has the content we want : I am Kairitu and I love Njamba
    check_three := checkFileContents("kairitu.txt", "I am Kairitu and I love Njamba")

    // 4. Check if the commit was done
    check_four := checkCommitMessage("Initial commit")
    
    // Print the final results
    fmt.Println(displayFinalResults(check_one, check_two, check_three, check_four)) 
  } else {
    fmt.Println("Don't you worry. You can come back anytime and proceed!")
  }
}

// Topic : Create a new branch, make some changes to the file and commit the changes
func kgit_topic_switch_branch_and_commit() {
  fmt.Println(payload["topic_three_title"])
  fmt.Println(payload["topic_three_summary"])
  fmt.Println(payload["topic_three_description"])
  fmt.Println(payload["topic_three_test"])
  
  if askForConfirmation("Have you managed to finish the test ?") == true {
    // 1. Check if we are on the develop branch 
    check_one := checkCurrentBranch("* develop")

    // 2. Check if file exists 
    check_two := checkFileExists("njamba.txt")

    // 3. Check if the file has the content we want : I am Kairitu and I love Njamba and I know he loves me too 
    check_three := checkFileContents("njamba.txt", "I am Kairitu and I love Njamba and I know he loves me too")

    // 4. Check if the commit was done
    check_four := checkCommitMessage("First change in develop commit")
    
    // Print the final results
    fmt.Println(displayFinalResults(check_one, check_two, check_three, check_four)) 
  } else {
    fmt.Println("Don't you worry. You can come back anytime and proceed!")
  }
}

// Topic : Checkout back to master branch and merge the changes
func kgit_topic_merge_branches() {
  fmt.Println(payload["topic_four_title"])
  fmt.Println(payload["topic_four_summary"])
  fmt.Println(payload["topic_four_description"])
  fmt.Println(payload["topic_four_test"])
  
  if askForConfirmation("Have you managed to finish the test ?") == true {
    // 1. Check if we are on the master branch 
    check_one := checkCurrentBranch("* master")

    // 2. Check if merged file exists 
    check_two := checkFileExists("njamba.txt")

    // 3. Check if the merged file has the content we want : I am Kairitu and I love Njamba and I know he loves me too 
    check_three := checkFileContents("njamba.txt", "I am Kairitu and I love Njamba and I know he loves me too")

    // 4. Check if the commit was done
    check_four := checkCommitMessage("First change in develop commit") && checkCommitMessage("Initial commit")
    
    // Print the final results
    fmt.Println(displayFinalResults(check_one, check_two, check_three, check_four)) 
  } else {
    fmt.Println("Don't you worry. You can come back anytime and proceed!")
  }
}

// Main Game Loop
func kgit_gameloop() bool {
  clearMyScreen()
  // Show the welcome message, description
  kgit_introduction()

  // Get user selection on topic to go to 
  user_selection := get_user_selection()
  clearMyScreen()

  // Switch statement to go to topic 
  switch user_selection {
    case 1:
      kgit_topic_init_repo()
      if askForConfirmation("") == false {
        return false
      }
    case 2:
      kgit_topic_create_and_commit()
      if askForConfirmation("") == false {
        return false
      }
    case 3:
      kgit_topic_switch_branch_and_commit()
      if askForConfirmation("") == false {
        return false
      }
    case 4:
      kgit_topic_merge_branches()
      if askForConfirmation("") == false {
        return false
      }
  }

  return true 
}

