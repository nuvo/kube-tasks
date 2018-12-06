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
kube-tasks wait-for-pods -n fda -l release=prod-cassandra --replicas=3
```

### Execute
Example: Cassandra Repairer
```
kube-tasks execute -n fda -l release=prod-cassandra --command "nodetool repair --full"
```