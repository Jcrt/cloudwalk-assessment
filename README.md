# 💳 CloudWalk Assessment
Thanks for allow me to make part of CloudWalk Assessment!

## 👨🏼‍💻 About the code
Here we will talk about some code stuff

### Run
In order to run this code, following the steps above: 
- Navigate to root folder 

The application can be runned using these argument options:
- `go run cw-assessment.go help` will provide all other acceptable arguments available
- `go run cw-assessment.go -t <option>` will provide two different report types based on the `<option>`. The available options are: 
   - **game**: Will provide the expected output in [topic 3.2 of the assessment](https://gist.github.com/cloudwalk-tests/704a555a0fe475ae0284ad9088e203f1#32-report).
   - **mod**: Will provide the expected output in [topic 3.3 of the assessment](https://gist.github.com/cloudwalk-tests/704a555a0fe475ae0284ad9088e203f1#32-report).
- `go run cw-assessment.go` will set automatically the report type to **game**.

### Packages used
- [go-funk](https://github.com/thoas/go-funk) package was used become collections more flexible with aggregation, slicing and search operations.
- [orderedmap](https://github.com/iancoleman/orderedmap) package was used to achieve map key order maintenance, since with the default Go map structure, `json.Marshall()` couldn't keep it.

### Comments (known as summaries in .NET or Javadoc in Java)
Each package, constant, type and function inside the application have it's own comment to explain it's general purpose inside it's scope.

The pattern of documentation was gathered in [official Go Doc Comments page](https://tip.golang.org/doc/comment).

### 🗃️ Backlog
- **Unit Tests**: Due the time limit for this exercise, I couldn't do more than create the empty ut's files. I understand the great importance of this phase of development process and even know that there is some design methods that starts from it (TDD), and in a normal routine it would be done and probably with an reasonably threshold about 80% of coverage. I have splitted, as well, the application in several packages with reduce number of lines to become this task as smoothly as it could be.

## 🕹️ About the game rules
Here we will talk about some game rules stuff

### Game scope
I couldn't understand if a game is considered valid only when we have a score of the match (as occurr in the last log) or if I find the `InitGame` keyword. I took the decision of consider a game all entries after `InitGame`, even though no deaths or no actions had taken during game time.

### MOD Report
In this report, I only included deaths that happened during the match and not included all available MODs in the report.
In order to add it, is necessary to iterate through MeanOfDeath enum and add 0 score value for each one in `func cw_logAnalyzer.getKillsByMean()` 

### Edge Cases
1) When a player changes the name, supposely the log parser will not keep the correspndence between the name in the moment of the event and the aggregated result.
For example, lets say that a player `Doguinho caramelo` entered in the game, killed someone and after got the name `dogão é mau`. In this case, all deaths will be counted as `dogão é mau`'s deaths.

## 🧔🏽‍♂️ About me and my insights
During the project I spent a considering time studing and trying to understand about the language and how the things should be to have success in this assessment. My main references were: 

-  [Official Go documentation](https://tip.golang.org/doc/) to understand the basics of Go
    - [Stringer command tool](https://pkg.go.dev/golang.org/x/tools/cmd/stringer) to generate the String function for required enums
- [Medium](https://medium.com) to read about collections and aggregations in Go
- [Go By Examples](https://gobyexample.com/) for see real examples of code structures
- [Stack Overflow](https://stackoverflow.com/) when above topics won't work
- [ChatGPT](https://chatgpt.com/) when even Stack overflow had not worked 😅