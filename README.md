# reposiTree

![](https://github.com/chez-shanpu/reposiTree/workflows/go_test/badge.svg)

## What is reposiTree
This CLI tool convert repositories into tree-structured data and calculate their similarity.

Download from [here](https://github.com/chez-shanpu/reposiTree/releases)

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