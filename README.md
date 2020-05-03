# Overview

This is a quick golang server meant to serve static (HTML & other) content from a local folder.

## Installation

Clone the repository to a local folder, and make sure you have `go` installed.

## Usage

You can serve content from any local folder:

```bash
go run server.go -d /path/to/your/files
```

If you don't supply a directory path using the `-d` argument, it will by default serve content from the current `.` folder.

The default address to serve content from is `localhost:3000`. You can change the port via the `-p` argument:

```bash
go run server.go -p 1234 -d /path/to/your/files
```

Note that the app will by default log every single request in the console. To change this, edit the `server.go` file and comment the logging instruction in the `logEveryRequest` function.
