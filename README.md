# Version
Version package helps to automate process of versioning Go executables by 
using information from git repository and build environment and 
assigning them during compilation using `-X importpath.name=value` flag of [go link tool](https://golang.org/cmd/link/).

## Assigned information

- **version number**
  - the most recent tag
  - if the tag doesn't point to last commit, `+N` is added where N is number of commits since the tag
  - if the working tree is dirty, `-dev` is added
- **commit hash**
  - 7 hexadecimal digits representing abbreviated object name
- **commit timestamp**
  - time of the last commit
- **build user**
  - user that built the executable
- **build host**
  - host that the executable was built on
- **build timestamp**
  - time of building the executable

These are concatenated into long version that looks like:

> v0.6.3+35-dev-0788ff2 (2016/02/05-22:10) built on furby@beast (go1.5.1 linux-amd64) at Sat Feb 06 09:52:22 UTC 2016

## How to use

1. Copy the version package into your project.
2. Use `go build -ldflags "$(go run version/ldflags.go)"` to build executable.
3. Access exported variables to use in the project.

```go
var	(
    printVersion = flag.Bool("version", false, "print version")
)

func main() {
    flag.Parse()
    
    if *printVersion {
        fmt.Println(version.Long)
        os.Exit(0)
    }
}
```
