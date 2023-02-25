# Dive Into Golang: A Crash Course

![Github stars](https://img.shields.io/github/stars/abcdabcd3899/Dive-Into-Golang.svg)
![github language](https://img.shields.io/badge/language-Golang-green.svg)
![license](https://img.shields.io/github/license/abcdabcd3899/Dive-Into-Golang.svg)
![forks](https://img.shields.io/github/forks/abcdabcd3899/Dive-Into-Golang.svg)

It is the vscode golang project. Please use the **linux** (not macos) environment to compile and run all of the examples.
The more concise the language, the more flexible it becomes.

## How to run the project

In ubuntu, I use the vscode to go build, run, go run and  clean the project.
To make the project easier to use, I created the tasks.json and launch.json files.
Using the keyboard shortcut "command + shift + B," you may choose between the commands
`clean`, `go build`, `run`, and `go run` for jobs in tasks.json.

Each options has the following meaning:

1. go build. It will generate the executable file.
2. run. It will execute the executable file created by the `go build` phase.
3. go run. The go source code will be executed directly in tasks.json.
4. clean. It will remove all executable files.

To debug the various source codes, press the `F5` button.

## Outline of Features

* [Test Demo in VSCode](https://github.com/abcdabcd3899/Dive-Into-Golang/tree/main/test_demo)
* [Basic Grammar: Variable, Loop, Type System, Condition Branch, Pointer and Constant](https://github.com/abcdabcd3899/Dive-Into-Golang/tree/main/basic_grammar_demo)
* [Collections: Array, Slice, Map, Set，Stack, and Queue](https://github.com/abcdabcd3899/Dive-Into-Golang/tree/main/collections_demo)
* [String: Unicode, UTF-8, string and strconv Libraries](https://github.com/abcdabcd3899/Dive-Into-Golang/tree/main/string_demo)
* [Function: First-Citizen, Variable Parameters and Defer](https://github.com/abcdabcd3899/Dive-Into-Golang/tree/main/function_demo)
* Object-Oriented Programming: Encapsolution using Struct, Polymorphism using Interface and Inheriance using Composition
  * [Encapsolution](https://github.com/abcdabcd3899/Dive-Into-Golang/tree/main/encap_demo)
  * [Duck Type Interface: It is represented by a pointer](https://github.com/abcdabcd3899/Dive-Into-Golang/tree/main/interface_demo)
  * [Composition better than Inheritance](https://github.com/abcdabcd3899/Dive-Into-Golang/tree/main/composition_demo)
  * [Interface Polymorphism](https://github.com/abcdabcd3899/Dive-Into-Golang/tree/main/polymorphism_demo) [Interface Polymorphism2](https://github.com/abcdabcd3899/Dive-Into-Golang/tree/main/interface_polymorphism_demo)
  * [Self-Defined Virtual Function Table](https://github.com/abcdabcd3899/Dive-Into-Golang/tree/main/virtual_table_demo)
  * [Empty Interface](https://github.com/abcdabcd3899/Dive-Into-Golang/tree/main/empty_interface_demo)
* [Error Handling](https://github.com/abcdabcd3899/Dive-Into-Golang/tree/main/error_handling_demo)
* [Package](https://github.com/abcdabcd3899/Dive-Into-Golang/tree/main/package_demo)
* Concurrency in Action using Shared Memory
  * [Goroutine](https://github.com/abcdabcd3899/Dive-Into-Golang/tree/main/go_routine_demo)
  * [Shared Memory Protected: Mutex, RWLock, and Atomic Variables](https://github.com/abcdabcd3899/Dive-Into-Golang/tree/main/shared_mem_demo)
  * [Producer and Comsumer](https://github.com/abcdabcd3899/Dive-Into-Golang/tree/main/producer_consumer_demo)
  * [Reader and Writer Problem](https://github.com/abcdabcd3899/Dive-Into-Golang/tree/main/rw_demo)
* Concurrency in Action using Channel
  * [CSP Programming: Basics and Common Errors](https://github.com/abcdabcd3899/Dive-Into-Golang/blob/main/csp_demo/test/main.go)
  * [Select Case](https://github.com/abcdabcd3899/Dive-Into-Golang/blob/main/select_demo/test/main.go)
  * [Using 1 Size Buffered Channel as a Lock](https://github.com/abcdabcd3899/Dive-Into-Golang/blob/main/buffer_to_mutex/test/main.go)
  * [Error 1: Multiple Sender Send to A Channel](https://github.com/abcdabcd3899/Dive-Into-Golang/blob/main/csp_demo/test/1.md)
  * [Error 2: Prohibited Main Thread Blocked](https://github.com/abcdabcd3899/Dive-Into-Golang/blob/main/csp_demo/test/2.md)
  * [Error 3: Range for Loop of Channel](https://github.com/abcdabcd3899/Dive-Into-Golang/blob/main/csp_demo/test/3.md)
  * [Error 3: How to Gracefully Close Channels](https://github.com/abcdabcd3899/Dive-Into-Golang/blob/main/csp_demo/test/4.md)
* [Unsafe Programming: Try not to use](https://github.com/abcdabcd3899/Dive-Into-Golang/blob/main/unsafe_demo/test/main.go)

## In Action

* [A Toy Consistent Hash](https://github.com/abcdabcd3899/Dive-Into-Golang/blob/main/consistent_hash/consistent.go)
* [A Toy Raft](https://github.com/abcdabcd3899/Dive-Into-Golang/blob/main/raft)

## Dependency Management

We use the [go module](https://go.dev/blog/using-go-modules) tool in the project to manage dependencies, which is also recommended by the golang community.
We may use the following commands to generate the `go.mod` file for the various source code folds:

```shell
go mod init abcdabcd3899
```

To remove redundant dependencies, we utilize the following method:

```shell
go mod tidy
```

## Contribute

Please open pull requests if you want to add new features.

## References

1. [Go Spec](https://go.dev/ref/spec)
2. [GO FAQ](https://go.dev/doc/faq)
3. [Package](https://pkg.go.dev/)
4. [Golang Composition Better Than Inheriance](https://hackthology.com/golangzhong-de-mian-xiang-dui-xiang-ji-cheng.html)
5. [Mipsmonsta Blog](https://mipsmonsta.medium.com/)

## Agreement

<img src='https://www.gnu.org/graphics/gplv3-127x51.png' width='127' height='51'/>

More information [document of agreement](/LICENSE)

<img src='https://raw.githubusercontent.com/EyreFree/EFArticles/master/res/cc-by-nc-nd.png' width='145.77' height='51'/>

[Attribution - Non-commercial - No interpretation](http://creativecommons.org/licenses/by-nc-nd/3.0/cn/)
