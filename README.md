# Software Design Patterns implemented using golang

Este repositorio contiene ejemplos escritos en golang de los diferentes patrones de construcción de software.
La página de referencia del listado es de wikipedia, ciertamente no es la mejor fuente pero sirve de punto de inicio. 

## Listado de patrones

### Creacionales

* [Abstract factory](https://en.wikipedia.org/wiki/Abstract_factory_pattern)
* [Builder](https://en.wikipedia.org/wiki/Builder_pattern)
* [Dependency Injection](https://en.wikipedia.org/wiki/Dependency_injection)
* [Factory Method](https://en.wikipedia.org/wiki/Factory_method_pattern)
* [Lazy Initialization](https://en.wikipedia.org/wiki/Lazy_initialization)
* [Multiton](https://en.wikipedia.org/wiki/Multiton_pattern)
* [Object Pool](https://en.wikipedia.org/wiki/Object_pool_pattern)
* [Prototype](https://en.wikipedia.org/wiki/Prototype_pattern)
* [RAII (Resource Acquisition is Initialization)](https://en.wikipedia.org/wiki/Resource_Acquisition_Is_Initialization)
* [Singleton](https://en.wikipedia.org/wiki/Singleton_pattern)

### Estructurales

* [Adapter](https://en.wikipedia.org/wiki/Adapter_pattern)
* [Bridge](https://en.wikipedia.org/wiki/Bridge_pattern)
* [Composite](https://en.wikipedia.org/wiki/Composite_pattern)
* [Decorator](https://en.wikipedia.org/wiki/Decorator_pattern)
* Extension object
* [Facade](https://en.wikipedia.org/wiki/Facade_pattern)
* [Flyweight](https://en.wikipedia.org/wiki/Flyweight_pattern)
* [Front Controller](https://en.wikipedia.org/wiki/Front_controller)
* [Marker](https://en.wikipedia.org/wiki/Marker_interface_pattern)
* [Module](https://en.wikipedia.org/wiki/Module_pattern)
* [Proxy](https://en.wikipedia.org/wiki/Proxy_pattern)
* [Twin](https://en.wikipedia.org/wiki/Twin_pattern)

### Comportamentales

* [Blackboard](https://en.wikipedia.org/wiki/Blackboard_(design_pattern))
* [Chaing of responsability](https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern)
* [Command](https://en.wikipedia.org/wiki/Command_pattern)
* [Interpreter](https://en.wikipedia.org/wiki/Interpreter_pattern)
* [Iterator](https://en.wikipedia.org/wiki/Iterator_pattern)
* [Mediator](https://en.wikipedia.org/wiki/Mediator_pattern)
* [Memento](https://en.wikipedia.org/wiki/Memento_pattern)
* [Null object](https://en.wikipedia.org/wiki/Null_Object_pattern)
* [Observer](https://en.wikipedia.org/wiki/Observer_pattern)
* [Servant](https://en.wikipedia.org/wiki/Design_pattern_Servant)
* [Specification](https://en.wikipedia.org/wiki/Specification_pattern)
* [State](https://en.wikipedia.org/wiki/State_pattern)
* [Strategy](https://en.wikipedia.org/wiki/Strategy_pattern)
* [Template method](https://en.wikipedia.org/wiki/Template_method_pattern)
* [Visitor](https://en.wikipedia.org/wiki/Visitor_pattern)

### Concurrencia

* [Active Object](https://en.wikipedia.org/wiki/Active_object)
* [Balking](https://en.wikipedia.org/wiki/Balking_pattern)
* [Binding properties](https://en.wikipedia.org/wiki/Binding_properties_pattern)
* [Compute kernel](https://en.wikipedia.org/wiki/Binding_properties_pattern)
* [Double-checked locking](https://en.wikipedia.org/wiki/Double_checked_locking_pattern)
* [Event-based asynchronous](https://en.wikipedia.org/wiki/Event-Based_Asynchronous_Pattern)
* [Guarded suspension](https://en.wikipedia.org/wiki/Guarded_suspension)
* [Join](https://en.wikipedia.org/wiki/Join-pattern)
* [Lock](https://en.wikipedia.org/wiki/Lock_(computer_science))
* [Messaging design pattern (MDP)](https://en.wikipedia.org/wiki/Messaging_pattern)
* [Monitor object](https://en.wikipedia.org/wiki/Monitor_(synchronization))
* [Reactor](https://en.wikipedia.org/wiki/Reactor_pattern)
* [Read-write lock](https://en.wikipedia.org/wiki/Read/write_lock_pattern)
* [Scheduler](https://en.wikipedia.org/wiki/Scheduler_pattern)
* [Thread pool](https://en.wikipedia.org/wiki/Thread_pool_pattern)
* [Thread-specific storage](https://en.wikipedia.org/wiki/Thread-Specific_Storage)
* Safe Concurrency with Exclusive Ownership
* CPU atomic operation

## Consideraciones para ejecutar el proyecto



## Comandos adicionales
* Links de referencia https://golangdocs.com/code-coverage-in-golang
* Para poder saber la conbertura de las pruebas unitarias `go test -cover` para generar el reporte de las pruebas unitarias 
* Para establecer el archivo de salida se usa el siguiente comando `go test -coverprofile=coverage.out`
* Para generar un archvio sobre el reporte generado se usa el siguiente comando `go tool cover -html=coverage.out`
* Para correr los test en el package principal y en los subfolders usar este comando `go test ./...`

## Building project
* build go projects with [task](https://taskfile.dev/#/)
* Instalación si ya tienes instalado go correctamente solo es necesario instalarlo así `go install github.com/go-task/task/v3/cmd/task@latest`
* Para este proyecto he decido hacer uso de `go-task/task` una herramienta hecha en go multiplataforma para facilitar la ejecución de comandos de building.
* El archivo `Taskfile.yml` contiene los tasks disponibles.
* Taggear los archivos de go con `//+build !test` antes del package y dejando una linea intermedia y ajustando el comando de test y building podemos excluir o indicar que paquetes hacen parte del paquete. Aqui link de referencia https://pkg.go.dev/cmd/go#hdr-Build_constraints


## Go mock
https://github.com/golang/mock
Mock generado para la interfase del S3 Client 
mockgen -destination=spkg-parquet-s3/services/mocks/mocks3.go -package=mocks github.com/aws/aws-sdk-go/service/s3/s3iface S3API

# Useful links
[golang struct validator](github.com/go-playground/validator/v10)