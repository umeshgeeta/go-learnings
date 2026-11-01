PROMPT (1):

Explain me th basic Go concepts of compiling & running a go program. What to "package" declaration mean at the top? How many 'main' routines are allowed in package? If have codebase with multiple 'main' routines, how do I organize folderwise? 

Explain me the concept of go module and significance of '.mod' and related files.

RESPONSE (1):

A Go program is a `package main` with exactly one `func main()`. A module (your `go.mod`) is the versioned collection of packages those programs are built from.

## Compiling and running

Your runnable program is `learnings26/main.go`:

```1:5:learnings26/main.go
package main

import "fmt"

func main(){
	fmt.Print("Hello World")
}
```

From the module root (`go-learnings/`), two commands matter:

```bash
go run ./learnings26
```

`go run` compiles the package into a temporary binary, runs it, and deletes that binary. Use this while you are iterating.

```bash
go build -o learnings26 ./learnings26
./learnings26
```

`go build` compiles and leaves an executable behind. The name after `-o` is the binary name; without `-o`, the binary is named after the directory (`learnings26`).

What those commands do:

1. Read `go.mod` so they know this tree is the module `go-learnings`.
2. Compile every package the program imports (`fmt`, and anything of yours it imports).
3. For a `package main`, link those packages into one executable whose entry point is `main.main`.

`go run .` (or `go run ./learnings26`) compiles the whole package in that directory. `go run learnings26/main.go` compiles only the files you list, so once a program is split across files, run the package path.

## What `package` means

Every `.go` file starts with a package clause. A package is one unit of compilation: all `.go` files in the same directory belong to one package and must use the same package name.

| Directory | Clause | Role |
|---|---|---|
| `learnings26/` | `package main` | A program. Built into a binary. |
| `pkg/lessons/` | `package lessons` | A library. Imported by other packages. |

`main` is the only package name that produces a program. Every other name produces a library. Another package imports yours by module path plus directory path, then uses the package name as the prefix:

```go
import "go-learnings/pkg/lessons"

lessons.ConvertToTitle(28)
```

Names that start with a capital letter (`ConvertToTitle`, `NewStack`) are visible to importers. Names that start with a lowercase letter (`baseTo26`) stay inside the package.

The directory name, the package clause, and the last piece of the import path are usually the same (`lessons`). The clause is what you type in code (`lessons.NewStack`). The import path is how the toolchain finds the directory.

## How many `main` functions

One directory is one package, and a `package main` has exactly one `func main()`. That function is the program entry point: the runtime starts a single goroutine and calls it.

You can split that program across several files in the same directory (`main.go`, `helpers.go`, all `package main`). Those files are still one package, so they still share a single `func main`. A second `func main` in that directory is a compile error (`main redeclared`).

`func main` only has meaning in `package main`. A function named `main` in `package lessons` is an ordinary function.

## Several programs in one codebase

Each program gets its own directory, because each directory can hold only one `func main`. The usual layout is a `cmd/` directory with one folder per binary, and shared code in library packages:

```text
go-learnings/
  go.mod                      # module go-learnings
  cmd/
    hello/main.go             # package main  →  binary "hello"
    excel/main.go             # package main  →  binary "excel"
  pkg/
    lessons/                  # package lessons, imported by both programs
```

Your repo already follows that split: `learnings26/` is one program, `pkg/lessons/` is shared library code. A second program is a new folder, for example `cmd/excel/main.go`, that imports `go-learnings/pkg/lessons`.

Build them separately or all at once:

```bash
go build -o bin/hello ./cmd/hello
go build -o bin/excel ./cmd/excel
go build ./...
```

`go build ./...` compiles every package in the module and writes one binary per `package main`, named after that program's directory.

## Modules, `go.mod`, and `go.sum`

A module is a set of packages released and versioned together. It is how Go finds your code and any libraries you depend on. Your module file is:

```1:3:go.mod
module go-learnings

go 1.22.4
```

Create it once at the root:

```bash
go mod init go-learnings
```

Two lines do the work:

- `module go-learnings` is the module path. It is the prefix of every import inside this tree, so `pkg/lessons` is imported as `go-learnings/pkg/lessons`. A module you publish for other people to download uses a path the toolchain can fetch, such as `github.com/you/go-learnings`. A local learning module can use a plain name.
- `go 1.22.4` is the language version this module is written for. The toolchain uses that version's language rules. A `toolchain go1.22.4` line, if you add one, pins the exact toolchain.

`go.sum` appears once you depend on another module. Each line is a checksum of a dependency's source (and of that dependency's `go.mod`). On the next build, Go checks the downloaded code against those checksums so the build uses the same bytes as last time. This repo has no `go.sum` yet because nothing outside the standard library is imported.

Other directives you will meet in `go.mod`:

- `require rsc.io/quote v1.5.2` records a dependency and the version you selected.
- `replace example.com/lib => ../lib` points a dependency at a local checkout.
- `exclude` and `retract` drop a bad version from selection or from your own releases.

Day-to-day commands:

```bash
go get rsc.io/quote@v1.5.2   # add a dependency
go mod tidy                  # add missing requirements, drop unused ones, update go.sum
```

Downloaded modules live in the module cache (`~/go/pkg/mod`), shared across projects. The source stays out of your repo; `go.mod` and `go.sum` are what make the build reproducible.

---

LEARNINGS (1):

Each numbered section is your note. The response under it is in a block quote.

### 1

A package is a set of Go files that are compiled together. A package can have at most one `main`, and if it has that, it is the starting point when the compiled binary is executed.

> **Response**
>
> A package is the set of `.go` files in one directory, compiled as one unit.
>
> The entry point exists only when two things are both true: the clause is `package main`, and exactly one file in that directory defines `func main()`. The runtime calls `main.main` when the binary starts.
>
> A second `func main` in the same package fails to compile (`main redeclared`). Other files in the directory share that one entry point and hold the rest of the program's functions.
>
> A function named `main` inside `package lessons` is an ordinary function. The name `main` is special only in `package main`.

### 2

When a package does not have `main`, it works as a Go library. The compiled binary is referred to in other Go programs.

> **Response**
>
> A package whose clause is a normal name, such as `package lessons`, is a library.
>
> Other programs use it by importing the source, for example `import "go-learnings/pkg/lessons"`. Building the program compiles that library into the same executable. There is no separate library binary to point at. `go build` on `pkg/lessons` only checks that the package compiles and stores the result in the build cache.
>
> A directory declared `package main` with no `func main` is neither a program nor a library. The build fails with `function main is undeclared in the main package`.

### 3

A package is associated with the directory layout. In one directory there can be at most one Go file with a `main` function. Different subdirectories under a root directory can have Go files with their own `main` functions. (When there is no Go file with `main`, it would act as a library.)

> **Response**
>
> One directory is one package. Subdirectories are separate packages, so each subdirectory can be its own program or its own library. Your repo already does this: `learnings26/` is a program, `pkg/lessons/` is a library.
>
> The limit inside one directory is one `func main` for the whole package. The directory can hold many `.go` files (`main.go`, `helpers.go`, and so on). All of them use the same package clause, and only one of those files contains `func main`.
>
> A directory with no `func main` is a library when its clause is a normal name such as `package lessons`.

### 4

All Go files in the package are linked by

```
package <package-name>
```

at the top of those Go files.

> **Response**
>
> The shared `package` line is what declares those files to be one package, and they also have to sit in the same directory. Two files that both say `package lessons` but live in different directories are two packages.
>
> Here "linked" means "members of the same package." Later, when a program is built, the linker joins the compiled packages into one executable. That is a separate step.
>
> One exception: a test file may say `package lessons_test`. `go test` compiles that as a separate package so the test imports `lessons` the way outside code would.

### 5

Methods starting with capital letters can be invoked outside of that package by importing that package, while methods with a lowercase starting character can only be invoked within the package.

> **Response**
>
> The capital-letter rule is right, and it covers every name, not only methods: functions, methods, types, struct fields, constants, and variables.
>
> `ConvertToTitle` and `NewStack` can be used by importers. `baseTo26` can be used from any file in `package lessons`, and from nowhere else.
>
> An importer writes the package name in front of the exported name: `lessons.ConvertToTitle(28)`.

### 6

A module is a group of packages and dependencies, along with the specific versions of those packages needed to make that module. It is defined in the `go.mod` file, typically at the root directory, as:

```
module <module-name>
```

You also indicate what version of Go is needed, followed by dependencies as:

```
require (
  <package name> version
)
```

(Question: how do I define the version of my Go program? Whether it is a package, a module, or a library without `main`?)

> **Response**
>
> A module is the group of packages in one tree that are built and versioned together. `go.mod` at the root defines that group. Your own packages are members of the module; they are not listed as dependencies.
>
> `module go-learnings` is the module path. Every import of your code starts with it, so `pkg/lessons` is imported as `go-learnings/pkg/lessons`.
>
> `go 1.22.4` is the Go language version this module is written for. It is separate from the version of your program.
>
> `require` lists other modules, each with one version. The left-hand side is a module path. Importing a package adds the module that contains that package:
>
> ```
> require rsc.io/quote v1.5.2
> ```
>
> **Version of your own code.** You do not write it in `go.mod`.
>
> - A published module gets its version from a Git tag, such as `v1.4.0`. Every package in the module, including a library with no `main`, shares that one version. Packages are not versioned on their own.
> - For major version 2 and later, the major version is also part of the module path: `module github.com/you/go-learnings/v2`.
> - A program you only build locally often has no tag. Built from a tagged commit, `runtime/debug.ReadBuildInfo()` reports that tag. An untagged local build reports `(devel)`. You can also stamp a string into the binary with `go build -ldflags "-X main.version=v1.4.0"`.

### 7

The `go.sum` file is an associated file which tracks the checksum of a dependency's source along with the version of that source. During a build, Go downloads the source of the dependency (the specified version), computes its checksum, and tests whether it matches or not. If it does not match, it updates `go.mod` and `go.sum` with the new checksum. (Does it always update? Then what is the point of tracking this checksum? How is the checksum computed? What is the checksum in the first place?)

Each line of `go.sum` is

```
<package-name> <version>/go.mod <checksum>
```

(Why is `/go.mod` present on some lines and not on others?)

> **Response**
>
> A checksum is a fingerprint of some bytes. Go's fingerprint is a SHA-256 hash of the module's file names and file contents, written in base64 with an `h1:` prefix. One changed byte produces a different fingerprint. The fingerprint cannot be turned back into the source; it can only answer "are these the same bytes?"
>
> `go.sum` stores the fingerprint already accepted for a module version. A version such as `v1.5.2` is supposed to mean those exact bytes forever.
>
> On a later build, Go computes the fingerprint again and compares it with `go.sum`. A mismatch stops the build. The files stay as they are. That refusal is the reason to record the checksum: a corrupted or swapped download cannot quietly become the dependency.
>
> `go.sum` changes when you deliberately change dependencies, through `go get` or `go mod tidy`. Those commands record the fingerprint of the version just selected. The first time a version is added, Go also checks the public checksum database at `sum.golang.org` before it writes the line.
>
> Each line is a module path, a version, and a hash. A downloaded module usually contributes two lines:
>
> ```
> rsc.io/quote v1.5.2 h1:<hash of the module files>
> rsc.io/quote v1.5.2/go.mod h1:<hash of that version's go.mod only>
> ```
>
> The plain line fingerprints the source. The `/go.mod` line fingerprints only that version's `go.mod`, so Go can check the dependency list before downloading the full source. A module that appears in the version graph, but whose packages your program never compiles, often has only the `/go.mod` line.

### 8

Dependencies can be managed by:

```
go get <package-name>@<package-version>
```

while the command

```
go mod tidy
```

refreshes (adds, deletes, or updates) `go.mod` and `go.sum` for defined as well as needed and not-needed dependencies.

> **Response**
>
> `go get` adds, upgrades, or removes one dependency and updates both files. The argument is a module path, or a package path inside a module; the version is always the module's version.
>
> ```
> go get rsc.io/quote@v1.5.2
> go get rsc.io/quote@latest
> go get rsc.io/quote@none
> ```
>
> `@latest` selects the newest version. `@none` drops that requirement.
>
> `go mod tidy` makes the two files match the imports in this module. A needed dependency is one that some package here imports, including tests; tidy adds any of those that are missing. A not-needed dependency is a `require` line that nothing imports anymore; tidy removes it. Checksums for whatever remains are written to `go.sum`.

### 9

Unless the clause in `file1.go` is changed from `package main` to `package some-pkg-name`, that code cannot be referred to from other Go files. The same code should also run as a standalone program and be callable as a library function.

> **Response**
>
> The package name is the clause, `main`. The directory name is only the default binary name (`go build` in `learnings26/` writes a binary called `learnings26`). Other code never imports that directory under the directory's name.
>
> Files in that same directory are already one package. Another file there with `package main` can call these functions directly. No import, and no change to the clause.
>
> A file in a different directory cannot import `package main`. An import of `go-learnings/learnings26` fails because that directory is a program.
>
> Changing the clause to a normal package name turns the directory into a library and it stops being a program. `func main` in that library is then an ordinary unexported function: the build produces no binary, and other packages still cannot call it, because `main` starts with a lowercase letter. A library function needs a capital letter, such as `ConvertToTitle`. A package name is an identifier, so `some-pkg-name` is invalid; `somepkg` is a legal name.
>
> One directory has one package clause, so the same directory cannot be both the program and the library. Keep the reusable function in a library package, and put a small `func main` in another directory that imports it:
>
> ```
> go-learnings/
>   pkg/title/title.go       // package title
>   cmd/title/main.go        // package main
> ```
>
> ```go
> package title
>
> func ConvertToTitle(columnNumber int) string {
>     // reusable function
> }
> ```
>
> ```go
> package main
>
> import (
>     "fmt"
>     "go-learnings/pkg/title"
> )
>
> func main() {
>     fmt.Println(title.ConvertToTitle(28))
> }
> ```
>
> `go run ./cmd/title` runs it as a program. Any other package in the module calls the same function with `title.ConvertToTitle`.

