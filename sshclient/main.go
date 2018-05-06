package main

import (
	"fmt"
	"log"
	"os"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"flag"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatalf("Usage: %s <user> <host:port> <command>", os.Args[0])
	}

	var keyfilePathVar = flag.String("keyfile", "", "the key file path")


	flag.StringVar(&keyfilePathVar, "gopher_type", defaultGopher, usage)
	flag.StringVar(&keyfilePathVar, "g", defaultGopher, usage+" (shorthand)")

	client, session, err := connectToHost(os.Args[1], os.Args[2])
	if err != nil {
		panic(err)
	}
	out, err := session.CombinedOutput(os.Args[3])
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
	client.Close()
}

func PublicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(key)
}


func connectToHostByPublicKey(user, host string, keyfilePath string) (*ssh.Client, *ssh.Session, error) {

	var pass string

	fmt.Print("Password: ")
	fmt.Scanf("%s\n", &pass)

	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			PublicKeyFile(keyfilePath),
		},
	}


	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()


	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		return nil, nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	return client, session, nil
}


func connectToHost(user, host string) (*ssh.Client, *ssh.Session, error) {

	var pass string

	fmt.Print("Password: ")
	fmt.Scanf("%s\n", &pass)


	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(pass)},
	}
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()


	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		return nil, nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	return client, session, nil
}
