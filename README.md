# Kube tools

A tool to perform simple Kubernetes related actions

## Commands

### Simple Backup

Example: Backup Jenkins
```
kube-tools simple-backup -n default -l release=jenkins -c jenkins --path /var/jenkins_home --dst s3://nuvo-jenkins-data
```