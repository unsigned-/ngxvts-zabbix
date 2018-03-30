package utils

import (
	"os"
	"os/exec"
	"fmt"
	"github.com/ximply/ngxvts-zabbix/config"
	"strings"
	"bufio"
	"io"
)

func isIgnoreVhost(vhost string, ignores []string) bool {
	for _, ig := range ignores {
		if strings.Compare(vhost, ig) == 0 {
			return true
		}
	}
	return false
}

func DiscoeryNginxServerName() (string, []string) {

	vp := config.NgxVhostsPath()
	if !FileExists(vp) {
		return "", nil
	}

	cmdStr := ""
	tmpFile := "/tmp/ngx_vhosts"
	cmdStr = fmt.Sprintf("cat %s*.conf | grep \"server_name\" | grep -v \"grep\" > %s",
		vp, tmpFile)

	cmd := exec.Command("/bin/bash","-c", cmdStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	cmd.Run()
	cmd.Wait()

	f, err := os.Open(tmpFile)
	if err != nil {
		return "", nil
	}
	defer f.Close()

	ignores := config.NgxVhostsIgnore()
	br := bufio.NewReader(f)
	vl := ""
	for {
		line, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		s := strings.TrimLeft(string(line)," ")
		if !strings.HasPrefix(s, "server_name") || !strings.HasSuffix(s, ";") {
			continue
		}
		s = strings.TrimLeft(s,"server_name")
		s = strings.TrimLeft(s," ")
		s = strings.TrimRight(s,";")
		l := strings.Split(s, " ")
		for _, i := range l {
			is := strings.Replace(i, " ", "", -1)
			if len(is) == 0 {
				continue
			}
			if isIgnoreVhost(is, ignores) {
				continue
			}
			vl += is
			vl += "|"
		}
	}
	vl = strings.TrimRight(vl, "|")
	vlList := strings.Split(vl, "|")

	ngxVhostVar := config.NgxZbxVhostVar()
	if len(vlList) == 0 {
		return "", nil
	}

	m := make(map[string]string)
	retStrList := ""
	ret := "{"
	ret += "\"data\" : ["
	for _, i := range vlList {
		if _, ok := m[i]; ok {
			continue
		}
		m[i] = i
		ret += "{"
		ret += fmt.Sprintf("\"{#%s}\" : \"%s\"", ngxVhostVar, i)
		ret += "},"

		retStrList += i
		retStrList += "|"
	}
	ret = strings.TrimRight(ret, ",")
	retStrList = strings.TrimRight(retStrList, "|")
	ret += "]"
	ret += "}"

	return ret, strings.Split(retStrList, "|")
}

func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

func FileExists(file string) bool {
	_, err := os.Stat(file)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}
