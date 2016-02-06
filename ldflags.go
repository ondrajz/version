// +build ignore

package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"regexp"
	"strings"
	"time"
)

var abbrev = regexp.MustCompile("-([0-9]{1,3})-g[0-9a-f]{5,10}")

func main() {
	pkg := fmt.Sprintf("%s/version", importPath())
	flags := []string{
		fmt.Sprintf("-X %s.Number=%s", pkg, baseVersion()),
		fmt.Sprintf("-X %s.CommitHash=%s", pkg, commitHash()),
		fmt.Sprintf("-X %s.CommitStamp=%s", pkg, commitStamp()),
		fmt.Sprintf("-X %s.BuildUser=%s", pkg, buildUser()),
		fmt.Sprintf("-X %s.BuildHost=%s", pkg, buildHost()),
		fmt.Sprintf("-X %s.BuildStamp=%d", pkg, time.Now().Unix()),
	}
	ldflags := strings.Join(flags, " ")
	fmt.Println(ldflags)
}

func importPath() string {
	return run("go", "list")
}

func baseVersion() string {
	version := run("git", "describe", "--always", "--dirty=-dev")
	return abbrev.ReplaceAllString(version, "+$1")
}

func commitHash() string {
	return run("git", "show", "-s", "--format=%h")
}

func commitStamp() string {
	return run("git", "show", "-s", "--format=%ct")
}

func buildUser() string {
	u, err := user.Current()
	if err != nil {
		return "<unknown>"
	}
	return strings.Replace(u.Username, " ", "_", -1)
}

func buildHost() string {
	h, err := os.Hostname()
	if err != nil {
		return "<unknown>"
	}
	return strings.Replace(h, " ", "_", -1)
}

func run(cmd string, args ...string) string {
	out, err := exec.Command(cmd, args...).CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		os.Exit(1)
	}
	return string(bytes.TrimSpace(out))
}
