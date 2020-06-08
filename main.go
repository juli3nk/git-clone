package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/juli3nk/go-git"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

func main() {
	c := new(Config)

	if err := envconfig.Process("git_clone", c); err != nil {
		log.Fatal(err.Error())
	}

	dd, err := repoDir(c.Repo)
	if err != nil {
		log.Fatal(err)
	}

	dstDir := *dd
	if len(c.DestDir) > 0 {
		dstDir = c.DestDir
	}

	// Log level
	if c.Debug {
		log.SetLevel(log.DebugLevel)
	}

	// Change dir
	if err := os.Chdir(c.RootDir); err != nil {
		log.Fatal(err)
	}
	log.Debugf("Changed directory to (%s)", cwd())

	g, err := git.New(c.Repo)
	if err != nil {
		log.Fatal(err)
	}

	if len(c.Username) > 0 && len(c.Password) > 0 {
		log.Debugf("Username: ***%s / Password: ***", c.Username[len(c.Username) - 3:])
		if err := g.SetAuth(c.Username, "token", c.Password); err != nil {
			log.Fatal(err)
		}
		log.Debug("Set Authentication")
	}

	// git clone
	log.Debugf("Cloning repository '%s'", c.Repo)
	if err = g.Clone(dstDir); err != nil {
		log.Fatal(err)
	}
	log.Debugf("Cloned repository '%s'", dstDir)

	// Change dir
	if err := os.Chdir(dstDir); err != nil {
		log.Fatal(err)
	}
	log.Debugf("Changed directory to (%s)", cwd())

	if err := g.Open(); err != nil {
		log.Fatal(err)
	}

	if c.Ref != "master" {
		if err := g.Checkout(c.Ref); err != nil {
			log.Fatal(err)
		}
		log.Debugf("Checkout to reference '%s'", c.Ref)
	}

	if err := os.RemoveAll(".git"); err != nil {
		log.Fatal(err)
	}
}

func cwd() string {
        dir, _ := os.Getwd()

        return dir
}

func repoDir(repo string) (*string, error) {
	split := strings.Split(repo, "/")

	if len(split) > 1 {
		t := strings.Split(split[len(split) - 1], ".git")

		return &t[0], nil
	}

	return nil, fmt.Errorf("Could not get repo name")
}
