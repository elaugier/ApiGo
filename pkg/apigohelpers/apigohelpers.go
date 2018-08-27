package apigohelpers

import (
	"bytes"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/elaugier/ApiGo/pkg/apigoconfig"
	"github.com/kardianos/osext"
)

//PowershellRun ...
func PowershellRun(PSModule string, PSCmdLet string, args map[string]string) (int, string, error) {
	binaryName := "powershell.exe"
	powershellArgs := []string{
		"-NoProfile",
		"-ExecutionPolicy", "ByPass",
		"-Command", "$PSDefaultParameterValues['*:Encoding'] = 'utf8';$ErrorActionPreference=\\\"Stop\\\";Import-Module -Name {PSModule} 3>$null;{Cmd};if($?){Exit(0);}else{Exit(0);};"}
	for i := 0; i < len(powershellArgs); i++ {
		if !strings.HasPrefix(powershellArgs[i], "-") {
			powershellArgs[i] = "\"" + powershellArgs[i] + "\""
		}
	}
	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}
	var argsinline string
	for key, value := range args {
		argsinline += " " + key + " " + value
	}
	command := strings.Join(powershellArgs, " ")
	powershellCommand := PSCmdLet + " " + argsinline
	config, err := apigoconfig.Get()
	if err != nil {
		log.Fatal(err)
	}
	var pathModule string
	scriptsPath := config.GetString("ScriptsPath")
	volumeName := filepath.VolumeName(scriptsPath)
	if strings.HasPrefix(volumeName, "\\") || strings.HasPrefix(scriptsPath, "/") {
		pathModule = folderPath + scriptsPath + "/"
	}
	powershellModule := pathModule + PSModule
	command = strings.Replace(command, "{Cmd}", powershellCommand, 1)
	command = strings.Replace(command, "{PSModule}", powershellModule, 1)
	e := exec.Command(binaryName, command)
	var stdout, stderr bytes.Buffer
	e.Stdout = &stdout
	e.Stderr = &stderr
	err = e.Run()
	if err != nil {
		log.Printf("e.Run() failed with %s\n", err)
		outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
		log.Printf("stderr => %s", errStr)
		return 1, outStr, err
	}
	outStr := string(stdout.Bytes())
	return 0, outStr, nil

}

//PythonRun ...
func PythonRun(PyEnv string, PyScript string, args map[string]string) (int, string, error) {
	binaryName := "python.exe"
	exec.Command(binaryName, PyScript)
	return 0, "OK", nil
}

//PerlRun ...
func PerlRun(PerlScript string, args map[string]string) (int, string, error) {
	binaryName := "perl.exe"
	exec.Command(binaryName, PerlScript)
	return 0, "OK", nil
}

//RubyRun ...
func RubyRun(RubyScript string, args map[string]string) (int, string, error) {
	binaryName := "ruby.exe"
	exec.Command(binaryName, RubyScript)
	return 0, "OK", nil
}

//CLIRun ...
func CLIRun(Command string, args map[string]string) (int, string, error) {
	binaryName := "cmd.exe"
	exec.Command(binaryName, Command)
	return 0, "OK", nil
}
