package controller

import (
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func Connect(compName, destination string) error {
	config := &ssh.ClientConfig{
		User: "username",
		Auth: []ssh.AuthMethod{
			ssh.Password("insertpasshere"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", "servername.com:22", config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}
	session, err := client.NewSession()
	if err != nil {
		client.Close()
		log.Fatalf("Failed to create session: %s", err)
	}
	output, err := session.CombinedOutput("ls -l")
	if err != nil {
		log.Fatalf("Failed to run: %s", err)
	}
	fmt.Println(string(output))
	return nil
}
