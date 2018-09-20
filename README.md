# Jeknins backup

A backup and restore tool for Jenkins on Kubernetes.

## Commands

### Backup

```
jenkins-backup -n default -l release=jenkins -c jenkins --dst s3://nuvo-jenkins-data
```