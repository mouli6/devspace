---
title: "Command - devspace add image"
sidebar_label: image
---


Add an image

## Synopsis

 
```
devspace add image [flags]
```

```
#######################################################
############# devspace add image ######################
#######################################################
Add a new image to your DevSpace configuration

Examples:
devspace add image my-image --image=dockeruser/devspaceimage2
devspace add image my-image --image=dockeruser/devspaceimage2 --tag=alpine
devspace add image my-image --image=dockeruser/devspaceimage2 --context=./context
devspace add image my-image --image=dockeruser/devspaceimage2 --dockerfile=./Dockerfile
devspace add image my-image --image=dockeruser/devspaceimage2 --buildtool=docker
devspace add image my-image --image=dockeruser/devspaceimage2 --buildtool=kaniko
#######################################################
```
## Options

```
      --buildtool string    Specify which engine should build the file. Should match this regex: docker|kaniko
      --context string      The path of the images' context
      --dockerfile string   The path of the images' dockerfile
  -h, --help                help for image
      --image string        The image name of the image (e.g. myusername/devspace)
      --tag string          The tag of the image
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

## See Also

* [devspace add](../../cli/commands/devspace_add)	 - Change the DevSpace configuration
