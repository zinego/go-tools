package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/zinego/go-tools/utils/log"
)

const publicKeyPath = "~/.ssh/id_rsa"

func upload(dir, fname string) {
	respository, err := git.PlainOpen("../image")
	if err != nil {
		log.Errorf("open git repository failed: %v", err)
		panic(err)
	}
	w, err := respository.Worktree()
	if err != nil {
		log.Errorf("worktree failed: %v", err)
		panic(err)
	}
	w.Add(fname)
	msg := fmt.Sprintf(time.Now().Format("add image at 2006-01-02T15:04:05"))
	w.Commit(msg, &git.CommitOptions{})
	log.Infof("push result: %v", respository.Push(&git.PushOptions{Auth: publicKey(publicKeyPath)}))
}

func publicKey(filePath string) *ssh.PublicKeys {
	var publicKey *ssh.PublicKeys
	sshKey, _ := ioutil.ReadFile(filePath)
	publicKey, err := ssh.NewPublicKeys("git", []byte(sshKey), "")
	if err != nil {
		panic(err)
	}
	return publicKey
}
