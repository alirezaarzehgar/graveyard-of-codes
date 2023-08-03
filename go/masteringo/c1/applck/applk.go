package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"time"
)

const (
	lockfile = "/tmp/app.lock"
	locktime = time.Hour
	runwait  = time.Minute * 5
)

func removeLines(filename string, delines []int) error {
	inp, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var output []string
	lines := strings.Split(string(inp), "\n")

main_loop:
	for i := 0; i < len(lines); i++ {
		for _, l := range delines {
			if l == i {
				continue main_loop
			}
		}

		output = append(output, lines[i])
	}

	out := strings.Join(output, "\n")
	ioutil.WriteFile(filename, []byte(out), 0644)
	return nil
}

func lockExpired(appname string) error {
	var expire int64 = 0
	var delines []int

	f, err := os.Open(lockfile)
	if err != nil {
		_, err := os.Create(lockfile)
		if err != nil {
			return err
		}

		return nil
	}
	defer f.Close()

	scnr := bufio.NewScanner(f)
	scnr.Split(bufio.ScanLines)
	for i := 0; scnr.Scan(); i++ {
		line := strings.Split(scnr.Text(), " ")
		if len(line) != 2 {
			log.Println("Invalid>", scnr.Text())
			continue
		}

		exp, err := strconv.ParseInt(line[1], 10, 64)
		if err != nil {
			log.Println(err)
			continue
		}

		if time.Now().Unix() > exp {
			delines = append(delines, i)
		} else if appname == line[0] {
			expire = exp
		}
	}

	if err := removeLines(lockfile, delines); err != nil {
		return err
	}

	if expire == 0 {
		return nil
	}

	expTime := time.Unix(expire, 0).Format("15:04:05")
	return errors.New("This application is locked until " + expTime)
}

func lockApp(appname string) error {
	f, err := os.OpenFile(lockfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	expireAt := time.Now().Add(locktime).Unix()
	f.WriteString(fmt.Sprintf("%s %d\n", appname, expireAt))
	return nil
}

func runTimeout(appname string, args []string) {
	if strings.HasPrefix(appname, "app.") {
		appname = appname[4:]
	}

	cmd := exec.Command(appname, args...)
	go func() {
		cmd.Output()
	}()

	time.Sleep(runwait)
	cmd.Process.Kill()
}

func main() {
	appname := path.Base(os.Args[0])

	if err := lockExpired(appname); err != nil {
		log.Fatal(err)
	}

	if err := lockApp(appname); err != nil {
		log.Fatal(err)
	}

	runTimeout(appname, os.Args[1:])
}
