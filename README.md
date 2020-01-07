# reposiTree

![](https://github.com/chez-shanpu/reposiTree/workflows/go_test/badge.svg)

## What is reposiTree
This CLI tool convert repositories into tree-structured data and calculate their similarity.

## Installation
```bash
# mac
$ export OS_NAME=darwin
# linux
$ export OS_NAME=linux
# windows
$ export OS_NAME=windows

# mac or linux
$ curl -LO https://github.com/chez-shanpu/reposiTree/releases/download/${VERSION}/repotr-${OS_NAME}-amd64
# windows
$ curl -LO https://github.com/chez-shanpu/reposiTree/releases/download/${VERSION}/repotr-${OS_NAME}-amd64.exe
```

## Usage
### tree make
```bash
$ repotr tree make \
    --repository-path <path to repository root dir> \
    --language <repository's main language> \
    --output <tree-structured datafile name>
```

### tree compare
```bash
$ repotr tree compare <path to datafile1> <path to datafile2>
```