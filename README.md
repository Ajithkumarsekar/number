# number
[![Go Reference](https://pkg.go.dev/badge/github.com/ajithkumarsekar/number.svg)](https://pkg.go.dev/github.com/ajithkumarsekar/number)

Go package for every numbers operations. nah! Just learning about go modules :)

### What is Go module?
Go module is a collection of packages
### Different Go module versioning
| version | Description                                                                                                                                                                                                                                                                                                                                        |
|---------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| v0      | The initial, unstable version. A v0 version does not make any stability guarantees, so nearly all projects should start with v0 as they refine their public API                                                                                                                                                                                    |
| v1      | The first stable version. Once you are absolutely sure your module’s API is stable, you can release v1.0.0. A v1 major version communicates to users that no incompatible changes will be made to the module’s API. They can upgrade to new v1 minor and patch releases, and their code should not break                                           |
| v2+     | A new major version. Breaks compatibility with the previous version. A new sub-folder must be created with v2 and all the source code should be copied with renamed import paths. eg. https://github.com/googleapis/gax-go. Module users should explicitly go get with version number at the end of the module name like `go get rsc.io/quote/v3@latest` | 

[refer go modules series](https://go.dev/blog/publishing-go-modules)

Golang import compatibility rule:
> If an old package and a new package have the same import path,
the new package must be backwards compatible with the old package.
### why go.sum?
### go get order of priority?

### What is GOPROXY?

### What is GOPRIVATE?

## Reference 
- [official](https://go.dev/ref/mod)
- [Talk: Go with Versions - GopherConSG 2018](https://youtu.be/F8nrpe0XWRg)
- [go.dev create module](https://go.dev/doc/tutorial/create-module)
- [Developing and publishing modules](https://go.dev/doc/modules/developing#workflow)

