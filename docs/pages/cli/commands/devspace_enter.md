---
title: "Command - devspace enter"
sidebar_label: devspace enter
---


Open a shell to a container

## Synopsis


```
devspace enter [flags]
```

```
#######################################################
################## devspace enter #####################
#######################################################
Execute a command or start a new terminal in your 
devspace:

devspace enter
devspace enter --pick # Select pod to enter
devspace enter bash
devspace enter -c my-container
devspace enter bash -n my-namespace
devspace enter bash -l release=test
#######################################################
```
## Options

```
  -c, --container string        Container name within pod where to execute command
  -h, --help                    help for enter
  -l, --label-selector string   Comma separated key=value selector list (e.g. release=test)
      --pick                    Select a pod
      --pod string              Pod to open a shell to
```

### Options inherited from parent commands

```
      --debug                 Prints the stack trace if an error occurs
      --kube-context string   The kubernetes context to use
  -n, --namespace string      The kubernetes namespace to use
      --no-warn               If true does not show any warning when deploying into a different namespace or kube-context than before
  -p, --profile string        The devspace profile to use (if there is any)
      --silent                Run in silent mode and prevents any devspace log output except panics & fatals
  -s, --switch-context        Switches and uses the last kube context and namespace that was used to deploy the DevSpace project
      --var strings           Variables to override during execution (e.g. --var=MYVAR=MYVALUE)
```
