package controller

import "log"

func ssh(compName, destination string) error {
	config := &ssh.ClientConfig{
		User: "username",
		Auth: []ssh.AuthMethod{
			ssh.Password("insertpasshere"),
		},
		HostKeyCallBack: ssh.FixedHostKey(hostKey),
	}
	client, err := ssh.Dial("tcp", "servername.com:22", config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}
	defer client.Close()
}
