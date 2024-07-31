# 💳 CloudWalk Assessment
Thanks for allow me to make part of CloudWalk Assessment!

## 👨🏼‍💻 About the code
Here we will talk about some code stuff

### Run
The application can be runned from the root path of the project, that's `cloudwalk-assessment` folder.

The application can be runned using these argument options:
- `go run cw-assessment.go help` will provide all other acceptable arguments available
- `go run cw-assessment.go -t <option>` will provide two different report types based on the `<option>`. The available options are: 
   - **game**: Will provide the expected output in [topic 3.2 of the assessment](https://gist.github.com/cloudwalk-tests/704a555a0fe475ae0284ad9088e203f1#32-report).
   - **mod**: Will provide the expected output in [topic 3.3 of the assessment](https://gist.github.com/cloudwalk-tests/704a555a0fe475ae0284ad9088e203f1#32-report).
- `go run cw-assessment.go` will set automatically the report type to **game**.

There is no configuration file to provide file path, so I added it inside `Assets` folder. If was necessary change the file name and file path, it can be achieved accessing `/cw-logReader/cw-logReader.go:11` and changing the path there.

### Packages used
- [go-funk](https://github.com/thoas/go-funk) package was used become collections more flexible with aggregation, slicing and search operations.
- [orderedmap](https://github.com/iancoleman/orderedmap) package was used to achieve map key order maintenance, since with the default Go map structure, `json.Marshall()` couldn't keep it.
- [testify](https://github.com/stretchr/testify) package was used to generate mocks and assertions in UTs

### Comments (known as summaries in .NET or Javadoc in Java)
Each package, constant, type and function inside the application have it's own comment to explain it's general purpose inside it's scope.

The pattern of documentation was gathered in [official Go Doc Comments page](https://tip.golang.org/doc/comment).

### 🗃️ Backlog
- **Unit Tests**: Using testify I could develop a considering number of UTs, but only basic ones. The general code coverage is over 80%, but I couldn't test all possible edge cases due the considering learning curve about Go and it's UTs. The report about coverage can be found in `coverage` folder inside solution and the report was generated using the command `go test -coverprofile coverage/cover.out -v ./... && go tool cover -html coverage/cover.out -o coverage/cover.html
`.

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