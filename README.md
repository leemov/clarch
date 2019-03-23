# clarch
This is my approach to apply clean architecture concept on golang.
My goal is to keep application business logic as clean as possible.

You can read the Clean Architecture concept proposed by Robert C. Martin here : 
https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

This is working application.
The case of this application is an election application.
Currently, action (Business logic) provided :
1. Join Election 

### To understand how this application works see this diagram
<p align="center">
  <img src="https://github.com/leemov/clarch/blob/master/files/Clean%20Arch%20Diagram.png">
</p>

The main principle of Clean architecture is `Separating Concern` by layering our application .
According to the concept we can implement more than 4 layers ( Entities, Usecase, Handler and Infrastructure ).
But at least 1 layer consist of business logic.

#### Entities : 
on this project, entities goes to `entities` folder.
entities will consist of the model and interface method to CRUD this model.

#### Usecase : 
on this project, usecase goes to `usecase` folder.
usecase is an interface function that consist of contract what func name to call and input output model for the usecase .
there are usecase interface and presenter interface.
usecase will be implemented by interactor object.
any dependencies needed inside the implementation will be injected trough interactor object.

#### Handler (controller/presenter): 
Handler will act as interface adapter.
on this project, it will act as controller and presenter.
handler goes to `handler` folder.
handler responsible to clean up data from request to fit usecase request model. 
interactor method will be called here.
on my approach, i also implement presenter interface that accept response model from usecase response model.

#### Resource implementation
This is the outer most layer.
this layer responsible to specify framework / technology used to implement entities's method.
Technology implementation lies on `infrastructure` folder.

All init will be called on application binary and injected trough controller.


### To run this application : 
go to bin folder.
- for CLI application .
on project folder go to `bin/cli` and run command on your terminal.
`go run cli.go`

- for REST HTTP application .
on project folder , go to `bin/rest` and run command on your terminal.
`go run rest.go`.

### database 
if you want to use database, you can find the structure and sample record for postgreSQL in files folder

