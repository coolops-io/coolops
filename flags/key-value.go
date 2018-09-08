package flags

import (
  "strings"
  "github.com/urfave/cli"
)

type KeyValueFlag struct{
  Values map[string]string
}

func (k *KeyValueFlag) Set(flag string) error {
  if flag == "" {
    return nil
  }
  
  if k.Values == nil {
    k.Values = make(map[string]string)
  }

  parts := strings.SplitN(flag, "=", 2)
  if len(parts) != 2 {
    return cli.NewExitError("should have the {key}={value} format", 1)
  }

  key := parts[0]
  value := parts[1]

  k.Values[key] = value
  return nil
}

func (k *KeyValueFlag) String() string {
  return ""
}
