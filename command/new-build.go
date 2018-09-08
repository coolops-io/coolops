package command

import (
	"bytes"
	"encoding/json"
	"github.com/coolopsio/coolops/flags"
	"github.com/coolopsio/coolops/info"
	"github.com/urfave/cli"
	"net/http"
)

func CmdNewBuild(c *cli.Context) error {
	buildName := c.Args().Get(0)
	if buildName == "" {
		return cli.NewExitError("name is mandatory", 1)
	}

	projectToken := c.String("token")
	if projectToken == "" {
		return cli.NewExitError("The project's api token is mandatory.\nEither pass it using the --token flag or setting the $COOLOPS_PROJECT_API_TOKEN environment variable.", 1)
	}

	params := c.Generic("param").(*flags.KeyValueFlag).Values
	metadata := c.Generic("metadata").(*flags.KeyValueFlag).Values

	requestData := map[string]interface{}{
		"name":   buildName,
		"params": params,
	}

	if len(metadata) > 0 {
		requestData["metadata"] = metadata
	}

	body, _ := json.Marshal(requestData)

	req, _ := http.NewRequest("POST", "http://api.coolops.io/builds", bytes.NewReader(body))
	req.Header.Add("Authorization", "Token " + projectToken)
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
