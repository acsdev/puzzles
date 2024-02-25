# Before

Make sure to have docker external volumes created

```shell
make prepare-volumes
```

## Config files

```go.work``` allows to have one workspace with multiple golang independent modules

```.vscode``` keeps a few workpsace configurations

## TIPS

- Before run java from vscode debug engine

  - Run ```mvn clean compile```

- JAVA and SPOJ

  - To run those code in Java at SPOJ please remove the class package, other wise you will get NZEC error.

## Project Structure

```plain

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
