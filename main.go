package main

import (
	"context"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	GetCommitsSinceBranch()
}

func GetCommitsSinceBranch() []string {
	ctx := context.Background()
	// git reflog show --no-abbrev --grep-reflog="Created" $(git rev-parse --abbrev-ref HEAD)
	// create command
	//args := []string{"reflog", "show", "--no-abbrev", "$(git rev-parse --abbrev-ref HEAD)"}
	args := []string{"log", "--format=%H,%dMOISHE"}
	cmd := exec.CommandContext(ctx, "git", args...)
	cmd.Dir, _ = os.Getwd()
	out, err := cmd.Output()
	splitStr := strings.Split(string(out), "MOISHE")

	log.Println(splitStr)
	log.Println(string(out), "err", err)
	return []string{}
}
