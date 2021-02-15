package app

import (
	"fmt"
	"io"
	"net"
	"net/url"
	"os/exec"
	"time"

	"github.com/onsi/gomega/gexec"
	"github.com/pkg/errors"
)

func NewExecutableApp(path string) *ExecutableApp {
	appExecPath, err := gexec.Build(path)
	if err != nil {
		panic(errors.Wrap(err, "failed to compile the package"))
	}

	return &ExecutableApp{
		path: appExecPath,
	}
}

type ExecutableApp struct {
	path string
}

func (a *ExecutableApp) Start(outWriter, errWriter io.Writer, env map[string]string) (*Executable, error) {
	var ok bool
	var err error
	var host, port string

	if host, ok = env["HOST"]; !ok {
		panic(errors.New("could not find env 'HOST'"))
	}
	if port, ok = env["PORT"]; !ok {
		panic(errors.New("could not find env 'PORT'"))
	}

	executable := exec.Command(a.path)
	executable.Stderr = errWriter
	executable.Stdout = outWriter

	for ev, value := range env {
		executable.Env = append(executable.Env, fmt.Sprintf("%s=%s", ev, value))
	}

	if err = executable.Start(); err != nil {
		return nil, errors.Wrap(err, "failed to start app")
	}

	return &Executable{
		cmd:  executable,
		host: host,
		port: port,
	}, nil

}

type Executable struct {
	cmd  *exec.Cmd
	host string
	port string
}

func (e *Executable) URL() string {
	url, err := url.Parse(fmt.Sprintf("http://%s:%s", e.host, e.port))
	if err != nil {
		panic(errors.Wrapf(err, "failed to create url from host %q and port %q", e.host, e.port))
	}

	return url.String()
}

func (e *Executable) Kill() error {
	return e.cmd.Process.Kill()
}

func (e *Executable) Netcat() error {
	addr := net.JoinHostPort(e.host, e.port)

	conn, err := net.DialTimeout("tcp", addr, 10*time.Second)
	if err != nil {
		return errors.Wrap(err, "failed to connect")
	}
	defer conn.Close()
	return nil
}
