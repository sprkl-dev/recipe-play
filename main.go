package main

import (
	"context"
	"log"
	"os"
	"os/exec"
)

func main() {
	GetCommitsSinceBranch()
}

func GetCommitsSinceBranch() []string {
	ctx := context.Background()
	// git reflog show --no-abbrev --grep-reflog="Created" $(git rev-parse --abbrev-ref HEAD)
	// create command
	args := []string{"reflog", "show", "--no-abbrev", "$(git rev-parse --abbrev-ref HEAD)"}
	cmd := exec.CommandContext(ctx, "git", args...)
	cmd.Dir, _ = os.Getwd()
	out, err := cmd.Output()

	log.Println(out, "err", err)
	// 	dir, _ = filepath.Abs(dir)
	// 	cmd.Dir = dir
	//
	// 	// setup buffer for stderr
	// 	var stderr bytes.Buffer
	// 	cmd.Stderr = &stderr
	//
	// 	out, err := cmd.Output()
	// 	if err != nil {
	// 		return nil, err
	// 	}
	return []string{}
}
