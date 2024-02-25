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

## Project Structure

└── code
    ├── codegami
    │   └── closet_to_zero
    ├── freecodecamp
    │   ├── fibonnaci
    │   └── numbers_cansum
    ├── gnirut
    ├── random
    │   └── binary_tree_mirror
    └── spoj
        ├── spoj_aba12c
        ├── spoj_acido
        ├── spoj_chocopj09
        ├── spoj_obitetri
        ├── spoj_placar
        ├── spoj_rouba
        └── spoj_test
