package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/skanehira/clipboard-image/v2"
)

func push(fname string) {
	respository, err := git.PlainOpen(imgdCfg.ImgRespository)
	if err != nil {
		fmt.Println("open git repository failed: ", err)
		os.Exit(1)
	}
	w, err := respository.Worktree()
	if err != nil {
		fmt.Println("worktree failed: ", err)
		os.Exit(1)
	}
	hash, err := w.Add(strings.TrimPrefix(fname, imgdCfg.ImgRespository+"/"))
	if err != nil {
		fmt.Println(hash.String(), err)
		os.Exit(1)
	}
	msg := fmt.Sprintf(time.Now().Format("add image at 2006-01-02T15:04:05"))
	w.Commit(msg, &git.CommitOptions{})
	err = respository.Push(&git.PushOptions{Auth: newAuth()})
	if err != nil {
		fmt.Println("save image failed: ", err)
		return
	}
	addr := "https://raw.githubusercontent.com/zinego/image/main/" + strings.TrimPrefix(fname, imgdCfg.ImgRespository+"/")
	fmt.Println("save image succeeded. addr: ", addr)
}

const (
	privateKeyMethod = "private_key"
	passwordMethod   = "password"
)

func newAuth() transport.AuthMethod {
	if imgdCfg.Auth.Method == privateKeyMethod {
		return publicKey(imgdCfg.Auth.PrivateKeyPath)
	} else if imgdCfg.Auth.Method == passwordMethod {
		return &ssh.Password{
			User:     imgdCfg.Auth.Username,
			Password: imgdCfg.Auth.Password,
		}
	}
	fmt.Println("err: incorrect method, supported methods: (private_key or password)")
	os.Exit(1)
	return nil
}

func publicKey(filePath string) *ssh.PublicKeys {
	var publicKey *ssh.PublicKeys
	sshKey, _ := ioutil.ReadFile(filePath)
	publicKey, err := ssh.NewPublicKeys("git", []byte(sshKey), "")
	if err != nil {
		fmt.Println("new publicKey failied: ", err)
		os.Exit(1)
	}
	return publicKey
}

func saveToRespository() (fname string) {
	var dir = imgdCfg.ImgRespository
	imgd, err := clipboard.Read()
	if err != nil {
		fmt.Println("read clipboard image failed: ", err)
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(imgd)
	if err != nil {
		fmt.Println("read all clipboard image failed: ", err)
		os.Exit(1)
	}
	fname = fmt.Sprintf("%s/img/%s.png", dir, time.Now().Format("2006_01_02T15_04_05"))
	f, err := os.Create(fname)
	if err != nil {
		fmt.Println("create image failed: ", err)
		os.Exit(1)
	}
	_, err = f.Write(body)
	if err != nil {
		fmt.Println("write image failed: ", err)
		os.Exit(1)
	}
	return fname
}
