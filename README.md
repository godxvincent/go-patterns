# Software Design Patterns implemented using golang

This repository has examples of software patterns that were written using golang.
Wikipedia was used as a reference of the list of more common sofrware patterns, Disclaimer (I use wikipedia as a starter point but in the future the big goal is to add more and better sources of information)

## List of patterns

### Creational Patterns

- [Abstract factory](https://en.wikipedia.org/wiki/Abstract_factory_pattern)
- [Builder](https://en.wikipedia.org/wiki/Builder_pattern)
- [Dependency Injection](https://en.wikipedia.org/wiki/Dependency_injection)
- [Factory Method](https://en.wikipedia.org/wiki/Factory_method_pattern)
- [Lazy Initialization](https://en.wikipedia.org/wiki/Lazy_initialization)
- [Multiton](https://en.wikipedia.org/wiki/Multiton_pattern)
- [Object Pool](https://en.wikipedia.org/wiki/Object_pool_pattern)
- [Prototype](https://en.wikipedia.org/wiki/Prototype_pattern)
- [RAII (Resource Acquisition is Initialization)](https://en.wikipedia.org/wiki/Resource_Acquisition_Is_Initialization)
- [Singleton](https://en.wikipedia.org/wiki/Singleton_pattern)

### Structural Patterns

- [Adapter](https://en.wikipedia.org/wiki/Adapter_pattern)
- [Bridge](https://en.wikipedia.org/wiki/Bridge_pattern)
- [Composite](https://en.wikipedia.org/wiki/Composite_pattern)
- [Decorator](https://en.wikipedia.org/wiki/Decorator_pattern)
- Extension object
- [Facade](https://en.wikipedia.org/wiki/Facade_pattern)
- [Flyweight](https://en.wikipedia.org/wiki/Flyweight_pattern)
- [Front Controller](https://en.wikipedia.org/wiki/Front_controller)
- [Marker](https://en.wikipedia.org/wiki/Marker_interface_pattern)
- [Module](https://en.wikipedia.org/wiki/Module_pattern)
- [Proxy](https://en.wikipedia.org/wiki/Proxy_pattern)
- [Twin](https://en.wikipedia.org/wiki/Twin_pattern)

### Behavioral Patterns

- [Blackboard](<https://en.wikipedia.org/wiki/Blackboard_(design_pattern)>)
- [Chaing of responsability](https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern)
- [Command](https://en.wikipedia.org/wiki/Command_pattern)
- [Interpreter](https://en.wikipedia.org/wiki/Interpreter_pattern)
- [Iterator](https://en.wikipedia.org/wiki/Iterator_pattern)
- [Mediator](https://en.wikipedia.org/wiki/Mediator_pattern)
- [Memento](https://en.wikipedia.org/wiki/Memento_pattern)
- [Null object](https://en.wikipedia.org/wiki/Null_Object_pattern)
- [Observer](https://en.wikipedia.org/wiki/Observer_pattern)
- [Servant](https://en.wikipedia.org/wiki/Design_pattern_Servant)
- [Specification](https://en.wikipedia.org/wiki/Specification_pattern)
- [State](https://en.wikipedia.org/wiki/State_pattern)
- [Strategy](https://en.wikipedia.org/wiki/Strategy_pattern)
- [Template method](https://en.wikipedia.org/wiki/Template_method_pattern)
- [Visitor](https://en.wikipedia.org/wiki/Visitor_pattern)

### Concurrency Patterns

- [Active Object](https://en.wikipedia.org/wiki/Active_object)
- [Balking](https://en.wikipedia.org/wiki/Balking_pattern)
- [Binding properties](https://en.wikipedia.org/wiki/Binding_properties_pattern)
- [Compute kernel](https://en.wikipedia.org/wiki/Binding_properties_pattern)
- [Double-checked locking](https://en.wikipedia.org/wiki/Double_checked_locking_pattern)
- [Event-based asynchronous](https://en.wikipedia.org/wiki/Event-Based_Asynchronous_Pattern)
- [Guarded suspension](https://en.wikipedia.org/wiki/Guarded_suspension)
- [Join](https://en.wikipedia.org/wiki/Join-pattern)
- [Lock](<https://en.wikipedia.org/wiki/Lock_(computer_science)>)
- [Messaging design pattern (MDP)](https://en.wikipedia.org/wiki/Messaging_pattern)
- [Monitor object](<https://en.wikipedia.org/wiki/Monitor_(synchronization)>)
- [Reactor](https://en.wikipedia.org/wiki/Reactor_pattern)
- [Read-write lock](https://en.wikipedia.org/wiki/Read/write_lock_pattern)
- [Scheduler](https://en.wikipedia.org/wiki/Scheduler_pattern)
- [Thread pool](https://en.wikipedia.org/wiki/Thread_pool_pattern)
- [Thread-specific storage](https://en.wikipedia.org/wiki/Thread-Specific_Storage)
- Safe Concurrency with Exclusive Ownership
- CPU atomic operation

## How to execute this project

Whoever is interested in this project just need to execute on the root of the project this command
`go run patterns/cmd/main.go -pattern <patternname>`

e.g

`go run patterns/cmd/main.go -pattern builder`

In case, -pattern parameter is not passed the default pattern is singleton.

## Some extra commands

- Execute unit test with coverage report `go test -cover`
- Execute unit test with coverage report as an output file `go test -coverprofile=coverage.out`
- Generate html report from previous output `go tool cover -html=coverage.out`
- Execute unit test recursively `go test ./...`

## Building project

- build go projects with [task](https://taskfile.dev/#/)
- Install task command (go itself is a pre-requisit) `go install github.com/go-task/task/v3/cmd/task@latest`
- The decision of using `go-task/task` as a tool for handling local building is because task is a solution multiplatform based on golang.
- File `Taskfile.yml` has all the available tasks.
- In case of skip a file of unit test mark the files with this sentnces `//+build !test` before package name and adding an empty line and passing the name of the components to be considerd. More info [here](https://pkg.go.dev/cmd/go#hdr-Build_constraints)

# Useful links

- [golang struct validator](github.com/go-playground/validator/v10)
- [Code coverage](https://golangdocs.com/code-coverage-in-golang)
- [Go Mock](https://github.com/golang/mock)
