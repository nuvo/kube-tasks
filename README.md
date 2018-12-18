[![Release](https://img.shields.io/github/release/nuvo/kube-tasks.svg)](https://github.com/nuvo/kube-tasks/releases)
[![Travis branch](https://img.shields.io/travis/nuvo/kube-tasks/master.svg)](https://travis-ci.org/nuvo/kube-tasks)
[![Docker Pulls](https://img.shields.io/docker/pulls/nuvo/kube-tasks.svg)](https://hub.docker.com/r/nuvo/kube-tasks/)
[![Go Report Card](https://goreportcard.com/badge/github.com/nuvo/kube-tasks)](https://goreportcard.com/report/github.com/nuvo/kube-tasks)
[![license](https://img.shields.io/github/license/nuvo/kube-tasks.svg)](https://github.com/nuvo/kube-tasks/blob/master/LICENSE)

# Kube tools

A tool to perform simple Kubernetes related actions

## Commands

### Simple Backup

Example: Backup Jenkins
```
kube-tasks simple-backup -n default -l release=jenkins -c jenkins --path /var/jenkins_home --dst s3://nuvo-jenkins-data
```

### Wait for Pods
Example: Cassandra Repairer init container
```
kube-tasks wait-for-pods -n prod -l release=prod-cassandra --replicas=3
```

### Execute
Example: Cassandra Repairer
```
kube-tasks execute -n prod -l release=prod-cassandra --command "nodetool repair --full"
```