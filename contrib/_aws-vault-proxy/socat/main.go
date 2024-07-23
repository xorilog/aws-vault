package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"syscall"
)

func doExecSyscall(command string, args []string, env []string) error {
	log.Printf("Exec command %s %s", command, strings.Join(args, " "))

	argv0, err := exec.LookPath(command)
	if err != nil {
		return fmt.Errorf("couldn't find the executable '%s': %w", command, err)
	}

	log.Printf("Found executable %s", argv0)

	argv := make([]string, 0, 1+len(args))
	argv = append(argv, command)
	argv = append(argv, args...)

	return syscall.Exec(argv0, argv, env)
}

func extractPortFromURI() (string, error) {
	u, err := url.Parse(os.Getenv("AWS_CONTAINER_CREDENTIALS_FULL_URI"))
	if err != nil {
		return "", fmt.Errorf("bad AWS_CONTAINER_CREDENTIALS_FULL_URI: %s", err.Error())
	}
	return u.Port(), nil
}

func replaceAWSVaultDynamicPort(port string) []string {
	var args []string
	re := regexp.MustCompile(`AWS_VAULT_DYNAMIC_PORT`)
	for _, v := range os.Args[2:] {
		replaced := re.ReplaceAllString(v, port)
		args = append(args, replaced)
	}
	return args
}

func main() {

	dynamicPort, err := extractPortFromURI()
	if err != nil {
		log.Fatal(err)
	}

	args := replaceAWSVaultDynamicPort(dynamicPort)

	if err := doExecSyscall(os.Args[1], args, os.Environ()); err != nil {
		log.Println(err)
	}
}
