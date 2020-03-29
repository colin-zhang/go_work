## golang
Golang execute shell command.
```go
func ExeCommand(commandStr string) string {
	commands := "echo hello"
	cmd := exec.CommandContext(nil,"/bin/bash", "-c", commands)
	if output, err := cmd.CombinedOutput(); err != nil {
		return err.Error()
	} else {
		return string(output)
	}
}
```

## etcd

## nginx