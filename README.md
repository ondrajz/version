# Version
Version package helps to automate process of versioning Go executables by 
using information from *git repository* and *build environment* and 
assigning them during compilation using `-X importpath.name=value` flag of [go link tool](https://golang.org/cmd/link/).

## Assigned information

- **version.Number**
  - the most recent annotated tag
  - if the last commit doesn't point to the tag, `+N` is added where N is number of commits since the tag
  - if the working tree is dirty, `-dev` is added
- **version.CommitHash**
  - 7 hexadecimal digits representing abbreviated object name
- **version.CommitStamp**
  - time of the last commit
- **version.BuildUser**
  - user that built the executable
- **version.BuildHost**
  - host that the executable was built on
- **version.BuildStamp**
  - time of building the executable
- **version.Long**
  - all information concatenated into long form

> v0.6.3+35-dev-0788ff2 (2016/02/05-22:10) built on furby@beast (go1.5.1 linux-amd64) at Sat Feb 06 09:52:22 UTC 2016

## Quick start

1. Run `ln -s $GOPATH/src/github.com/TrueFurby/version` in the directory of your main package to create a symlink to the version package. 
This will provide cleaner way for importing package and compiling binary.
2. Use `go build -ldflags "$(sh version/ldflags.sh)"` when building your executable to assign all version information.
3. Now you can refer to the exported variables wherever you need them.

## Usage example

```go
package main

import "yourproject/version"

func main() {
    printVersion = flag.Bool("version", false, "print version")
    flag.Parse()
    
    if *printVersion {
        fmt.Println(version.Long)
        os.Exit(0)
    }
}
```

