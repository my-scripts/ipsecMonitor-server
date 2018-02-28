package base

import (
	"bytes"
	"io"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/process"
)

type IpsecService struct {
	Name     string
	ProcName string
	Pid      int
	State    int
}

var IPSEC_SERVICE IpsecService

func GetIpsecService() *IpsecService {
	IPSEC_SERVICE.State, IPSEC_SERVICE.Pid = GetServiceState("/var/run/pluto/pluto.pid")
	IPSEC_SERVICE.Name = "IPSEC"
	IPSEC_SERVICE.ProcName = "ipsec"
	return &IPSEC_SERVICE
}

func GetServiceState(path string) (int, int) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return 2, -1
	}

	pidstr := strings.TrimSpace(string(content))
	pid, err := strconv.Atoi(pidstr)
	if err != nil {
		return 2, -1
	}
	state, _ := process.PidExists(int32(pid))
	if state {
		if IpsecIsWork() {
			return 1, int(pid)
		} else {
			return 0, int(pid)
		}

	}
	return 2, int(pid)
}

func IpsecIsWork() bool {
	cmd1 := exec.Command("netstat", "-nupl")
	cmd2 := exec.Command("grep", "pluto")
	r, w := io.Pipe()
	cmd1.Stdout = w
	cmd2.Stdin = r
	var out bytes.Buffer
	cmd2.Stdout = &out
	cmd1.Start()
	cmd2.Start()
	cmd1.Wait()
	w.Close()
	cmd2.Wait()
	return len(out.String()) != 0
}
