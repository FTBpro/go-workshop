# Installation

Installing Go is very straightforward on any operationg system. 

1. Download the Go compiler installer: [https://golang.org/dl]
2. Follow the installer and use the defaults.
3. If you have any problems with the installer and google didn't help, contact us. 
4. In order to make sure the go binary is available on your path:
 
```sh
$ go version
```

```
go version go1.12.5 darwin/amd64
```

If you got a similar output, you're ready to Go :)

Go can be developed using a set of standard tools using only the commandline, but ain't nobody got time for that!

## Code Editor

We recommend [Visual Studio Code][vscode] as it is open source and has Go support. (It is also possible to use [goland], there is a free trial for 30 days)

[Download Visual Studio Code][vscode]

### Installing Code Editor Plugins

, visit [VSCode Go Plugin in the Marketplace][vscode plugin].

If you're using another editor, you'll probably find your editor plugin [at the official wiki](plugins).

## Let's Go ☕️

The first task in the workshop is to make an `hello world` app. We'll do it together, so you can kick back and relax 🏖

#### Trivia

If you scroll down at [vscode plugin] page you can see how features like auto complete and refactors are implemented using tools like `godoc`, and `goimports`. These tools are developed by the Go Team, and are used by most of the [plugins] of different text editors.

While this sounds really good and all, in Go version 1.11 there was an [experimental](https://github.com/golang/go/wiki/Modules) feature released that broke **all** of these [tools](https://github.com/golang/go/issues/24661).

This is where GoLand, the Jet Brains product for Go development, really shined. In GoLand, they developed all the features themselves, so it was the *only* Go IDE that works out of the box.

Being an experimental feature, this hasn't affected anyone that wasn't willing to take a risk, but it definitely made a good case to move to GoLand for some people. 

[vscode]: https://code.visualstudio.com
[https://golang.org/dl]: https://golang.org/dl
[plugins]: https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins
[vscode plugin]: https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go
[goland]: https://www.jetbrains.com/go/
