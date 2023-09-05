package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	// Execute the git command to get the .git directory path
	gitDirCmd := exec.Command("git", "rev-parse", "--git-dir")
	gitDirBytes, err := gitDirCmd.Output()
	if err != nil {
		log.Fatalf("Error getting git directory: %v", err)
	}

	hooksDir := string(gitDirBytes[:len(gitDirBytes)-1]) + "/hooks" // -1 to remove the trailing newline

	// Content for the commit-msg file
	commitMsgContent := `#!/bin/sh

COMMIT_MSG_FILE=$1
COMMIT_MSG=$(cat $COMMIT_MSG_FILE)

# Example: Check if commit message starts with 'feat:', 'fix:', etc.
if ! echo "$COMMIT_MSG" | grep -qE "^(feat|fix|chore|docs|style|refactor|perf|test): "; then
    echo "ERROR: Commit message does not match expected format."
    exit 1
fi
`

	// Write to the commit-msg file
	err = os.WriteFile(hooksDir+"/commit-msg", []byte(commitMsgContent), 0755)
	if err != nil {
		log.Fatalf("Error writing to commit-msg file: %v", err)
	}

	fmt.Println("commit-msg hook has been installed.")
}
