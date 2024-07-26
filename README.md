# 💳 CloudWalk Assessment
Thanks for allow me to make part of CloudWalk Assessment!
I hope the project be nice!



## 👨🏼‍💻 About the code
Here we will talk about some code stuff

### Packages
- The package [go-funk](https://github.com/thoas/go-funk) was used become collections more flexible with aggregation, slicing and search operations.

### Comments (known as summaries in .NET or Javadoc in Java)
Each package, constant, type and function inside the application have it's own comment to explain it's general purpose inside it's scope.

The pattern of documentation was gathered in [official Go Doc Comments page](https://tip.golang.org/doc/comment).

### 🗃️ Backlog
- UTs
- Order the output of the reports because Go don't keep map's order
- Create an option to the user can call both reports by command line using a kind of -t in command line argument

## 🕹️ About the game rules
Here we will talk about some game rules stuff

### Edge Cases
1) When a player changes the name, supposely the log parser will not keep the correspndence between the name in the moment of the event and the aggregated result.
For example, lets say that a player `Doguinho caramelo` entered in the game, killed someone and after got the name `dogão é mau`. In this case, all deaths will be counted as `dogão é mau`'s deaths.

## 🧔🏽‍♂️ About me and my insights
During the project I spent a considering time studing and trying to understand about the language and how the things should be to have success in this assessment. My main references were: 

-  [Official Go documentation](https://tip.golang.org/doc/) to understand the basics of Go
    - [Stringer command tool](https://pkg.go.dev/golang.org/x/tools/cmd/stringer) to generate the String function for required enums
- [Medium](https://medium.com) to read about collections and aggregations in Go
- [Stack Overflow](https://stackoverflow.com/) when above topics won't work
- [ChatGPT](https://chatgpt.com/) when even Stack overflow had worked 😅