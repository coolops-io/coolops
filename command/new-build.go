package command

import (
	"bytes"
	"encoding/json"
	"github.com/coolops-io/coolops/flags"
	"github.com/coolops-io/coolops/info"
	"github.com/urfave/cli"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type buildData struct {
	name         string
	projectToken string
	params       map[string]string
	metadata     map[string]string
}

func (d *buildData) toJson() ([]byte, error) {
	j := map[string]interface{}{
		"name":   d.name,
		"params": d.params,
	}

	if len(d.metadata) > 0 {
		j["metadata"] = d.metadata
	}

	return json.Marshal(j)
}

func (d *buildData) send() error {
	body, err := d.toJson()
	if err != nil {
		return err
	}

	req, _ := http.NewRequest("POST", "http://api.coolops.io/builds", bytes.NewReader(body))
	req.Header.Add("Authorization", "Token "+d.projectToken)
	req.Header.Add("X-CoolOps-Cli-Version", info.Version)

	r, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	switch r.StatusCode {
	case 401:
		return cli.NewExitError("Invalid project token", 1)
	case 200, 201:
		return nil
	default:
		return cli.NewExitError("Invalid request", 1)
	}
}

func genericData(c *cli.Context) (*buildData, error) {
	projectToken := c.String("token")
	if projectToken == "" {
		return nil, cli.NewExitError("The project's api token is mandatory.\nEither pass it using the --token flag or setting the $COOLOPS_PROJECT_API_TOKEN environment variable.", 1)
	}

	params := c.Generic("param").(*flags.KeyValueFlag).Values
	metadata := c.Generic("metadata").(*flags.KeyValueFlag).Values

	return &buildData{
		projectToken: projectToken,
		params:       params,
		metadata:     metadata,
	}, nil
}

func CmdNewBuild(c *cli.Context) error {
	buildName := c.Args().Get(0)
	if buildName == "" {
		return cli.NewExitError("name is mandatory", 1)
	}

	d, err := genericData(c)

	if err != nil {
		return err
	}

	d.name = buildName

	return d.send()
}

func getCommitMessage(ref string) ([]byte, error) {
	return exec.Command("git", []string{"log", "--format=%B", "-n 1", ref}...).Output()
}

func CmdNewBuildCircleCI(c *cli.Context) error {
	d, err := genericData(c)

	if err != nil {
		return err
	}

	d.name = strings.Replace(os.Getenv("CIRCLE_BRANCH"), "/", "-", -1) + "-" + os.Getenv("CIRCLE_BUILD_NUM")

	pullRequest := os.Getenv("CIRCLE_PULL_REQUEST")
	if pullRequest == "" {
		pullRequest = "_No Pull Request_"
	}

	commitMessage, _ := getCommitMessage(os.Getenv("CIRCLE_SHA1"))

	d.metadata["Pull Request"] = pullRequest
	d.metadata["Last Commit"] = "```" + strings.TrimSpace(string(commitMessage[:])) + "```"

	return d.send()
}

func CmdNewBuildGitlab(c *cli.Context) error {
	d, err := genericData(c)

	if err != nil {
		return err
	}

	d.name = strings.Replace(os.Getenv("CI_COMMIT_REF_SLUG"), "/", "-", -1) + "-" + os.Getenv("CI_JOB_ID")

	jobUrl := os.Getenv("CI_JOB_URL")

	commitMessage, _ := getCommitMessage(os.Getenv("CI_COMMIT_SHA"))

	d.metadata["Job Url"] = jobUrl
	d.metadata["Last Commit"] = "```" + strings.TrimSpace(string(commitMessage[:])) + "```"

	return d.send()
}
