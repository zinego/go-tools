package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/skanehira/clipboard-image/v2"
)

func push(filename string) {
	repository, err := git.PlainOpen(imgdCfg.ImgRespository)
	if err != nil {
		fmt.Println("open git repository failed: ", err)
		os.Exit(1)
	}
	w, err := repository.Worktree()
	if err != nil {
		fmt.Println("worktree failed: ", err)
		os.Exit(1)
	}
	err = w.Pull(&git.PullOptions{
		Auth:       newAuth(),
		RemoteName: imgdCfg.RemoteName,
	})
	if err != nil && !strings.Contains(err.Error(), "already up-to-date") {
		fmt.Println("pull repository failed:", err)
		return
	}
	hash, err := w.Add(strings.TrimPrefix(filename, imgdCfg.ImgRespository+"/"))
	if err != nil {
		fmt.Println(hash.String(), err)
		os.Exit(1)
	}
	msg := fmt.Sprintf(time.Now().Format("add image at 2006-01-02T15:04:05"))
	hash, err = w.Commit(msg, &git.CommitOptions{})
	if err != nil {
		fmt.Printf("commit failed: %v\n", err)
		return
	}
	err = repository.Push(&git.PushOptions{
		Auth:       newAuth(),
		RemoteName: imgdCfg.RemoteName,
	})
	if err != nil {
		fmt.Println("save image failed: ", err)
		return
	}
	fmt.Println("save image succeeded. addr: ", fmt.Sprintf("%s/%s/%s", imgdCfg.ImgUrlPrefix, hash.String(), strings.TrimPrefix(filename, imgdCfg.ImgRespository+"/")))
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
	sshKey, _ := os.ReadFile(filePath)
	publicKey, err := ssh.NewPublicKeys("git", []byte(sshKey), "")
	if err != nil {
		fmt.Println("new publicKey failied: ", err)
		os.Exit(1)
	}
	return publicKey
}

func saveToRepository() (fname string) {
	var dir = imgdCfg.ImgRespository
	imgd, err := clipboard.Read()
	if err != nil {
		fmt.Println("read clipboard image failed: ", err)
		os.Exit(1)
	}
	body, err := io.ReadAll(imgd)
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
