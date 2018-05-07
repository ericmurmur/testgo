package main

import (
	"fmt"
	"log"
	"os"

	"io/ioutil"
	"flag"
	"errors"

	"golang.org/x/crypto/ssh"
	"os/signal"
	"bufio"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatalf("Usage: %s [-k keyfile] <user> <host:port> <command>", os.Args[0])
	}

	var keyfilePathVar string
	//= flag.String("keyfile", "", "the key file path")

	flag.StringVar(&keyfilePathVar, "keyfile", "", "keyfile path")
	flag.StringVar(&keyfilePathVar, "k", "", " (shorthand)")

	flag.Parse()
	args := flag.Args()

	fmt.Println("flag args num is ", args)

	if len(keyfilePathVar) == 0 {
		// no key file
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

	} else {

		client, session, err := connectToHostByPublicKey(args[0], args[1], keyfilePathVar)
		if err != nil {
			panic(err)
		}

		defer client.Close()


		OpenTerminal(session)
		out, err := session.CombinedOutput(args[2])
		if err != nil {
			panic(err)
		}

		fmt.Println(string(out))

	}



}

func PublicKeyFile(file string, passphrase string) ssh.AuthMethod {

	buffer, err := ioutil.ReadFile(file)

	if err != nil {
		return nil
	}

	//key, err := ssh.ParsePrivateKey(buffer)
	key, err := ssh.ParsePrivateKeyWithPassphrase(buffer, []byte(passphrase))

	if err != nil {
		//return nil
		panic(err)
	}

	return ssh.PublicKeys(key)
}

//
//
//
func connectToHostByPublicKey(user, host string, keyfilePath string) (*ssh.Client, *ssh.Session, error) {

	var pass string

	fmt.Print("Password: ")
	fmt.Scanf("%s\n", &pass)

	authmethod := PublicKeyFile(keyfilePath, pass)

	if authmethod==nil {
		return nil, nil, errors.New("key file failure")
	}

	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			authmethod,
			ssh.Password(pass),
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


func OpenTerminal(session *ssh.Session) {



	// Set IO
	//session.Stdout = ansicolor.NewAnsiColorWriter(os.Stdout)
	//session.Stderr = ansicolor.NewAnsiColorWriter(os.Stderr)

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	in, _ := session.StdinPipe()

	// Set up terminal modes
	// https://net-ssh.github.io/net-ssh/classes/Net/SSH/Connection/Term.html
	// https://www.ietf.org/rfc/rfc4254.txt
	// https://godoc.org/golang.org/x/crypto/ssh
	// THIS IS THE TITLE
	// https://pythonhosted.org/ANSIColors-balises/ANSIColors.html
	modes := ssh.TerminalModes{
		ssh.ECHO:  0, // Disable echoing
		ssh.IGNCR: 1, // Ignore CR on input.
	}

	// Request pseudo terminal
	//if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
	//if err := session.RequestPty("xterm-256color", 80, 40, modes); err != nil {
	if err := session.RequestPty("vt100", 80, 40, modes); err != nil {
		//if err := session.RequestPty("vt220", 80, 40, modes); err != nil {
		log.Fatalf("request for pseudo terminal failed: %s", err)
	}

	// Start remote shell
	if err := session.Shell(); err != nil {
		log.Fatalf("failed to start shell: %s", err)
	}

	// Handle control + C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for {
			<-c
			fmt.Println("^C")
			fmt.Fprint(in, "\n")
			//fmt.Fprint(in, '\t')
		}
	}()

	//var b []byte = make([]byte, 1)

	// Accepting commands
	for {
		reader := bufio.NewReader(os.Stdin)
		str, _ := reader.ReadString('\n')
		fmt.Fprint(in, str)
	}
}