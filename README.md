# Before

Make sure to have docker external volumes created

```
make prepare-volumes
```

## Config files

```go.work``` allows to have one workspace with multiple golang independent modules

```.vscode``` keeps a few workpsace configurations
- the property ```java.project.sourcePaths``` allows to have diferent directories as source folders under the java project


## JAVA and SPOJ

To run those code in Java at SPOJ please remove the class package, other wise you will get NZEC error.