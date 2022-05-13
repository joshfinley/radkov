# radkov

Rad as fuck Escape From Tarkov radar. 

![Alt Text](./tarkov-radar-demo.gif)

## Disclaimer

This is not directly useable as a game cheat. It's written to directly read process memory using Windows APIs. Additionally, it's written for the SPT-AKI launcher, and its likely that the real game would require different offsets to access player info. Expect to be banned for trying to run this against the real game.

## Project Layout
```

├───pkg     
│   ├───rkpb                -> protocol buffer definition
│   ├───tarkov              -> tarkov game utilities (e.g. read player pos)
│   │   └───tarkov_test     -> READ THESE FIRST to understand tarkov package
│   ├───unity               -> unity game utilities
│   │   └───unity_test      
│   └───winutil             -> utilities for reading process memory
│       └───winutil_test
├───radarsrv                -> radar web application backend 
│   ├───cmd                 -> contains main.go e.g. `go run radarsrv/cmd/main.go`
│   └───radarapp            -> radar web application backend library
│       └───static          -> HTML, CSS
└───tkovmon                 -> main.go for tarkov game memory monitor
```

## Dependencies

- go >= 1.17.6 (lower versions not tested but may work)
- protocol buffers compiler 3.19.4 (only needed if you need to rebuild the rkpb package)
- make (optional)
- go packages/modules: run: `go mod tidy`


## Running

1. start the web server backend:

```
go run radarsrv/cmd/main.go
```

This will listen on port 1337 and 80. 1337 is the gRPC server, 80 is a temporary frontend REST API used to make requests for player positions in a web application.

2. start the game memory monitor (`tkovmon`):

```
go run tkovmon/main.go`
```

This will start the game memory monitor. It will wait for the game to launch and the game session to become available.

## Testing and Development

### setting up a development environment

Go development with Vistual Studio Code is incredibly easy. If you haven't set up a development environment in VSCode for Go before, and you have Go installed, it should prompt you to install various tools. Click install for any of these and it will install plugins for code linting, auto-formatting, auto-completion, debugging, and more.

### running tests

While you can run all tests for a given package using the `go test` command, if the Go tools are installed for VSCode, it is much easier to just open any test file in vscode and use the 'run test | deug test` options that the Go tools provide. Using these, you can set breakpoints and debug individual tests, line by line. 

The test functions in these files make for perfect examples for using the packages they test. Additionally, if making changes or improvements to the Go code, you can drive your development using tests and the debugger instead of invoking `go run <whatever>` and scratching your head when the build or run fails.

## Building
### protocol buffer package (rkpb)

If changes are made to the `rkpb.proto` file, it will need to be recompiled:

```
make
# OR run the 'protoc' command listed inside the makefile
```

### standalone binaries

Unless distributing actual binaries for `radarsrv` and `tkovmon`, there should be no reason to run a standalone build. If this is needed for some reason, however, use the `go build` command.



