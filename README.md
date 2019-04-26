# gosf-sample-app
GoLang SocketIO API sample server written using GOSF

**This SocketIO API server was written using GOSF.  [Learn more about GOSF](https://github.com/ambelovsky/gosf)**

## File Structure

```txt
 scripts
 -- setpath.sh
 -- getdeps.sh
 -- build.sh
 -- run.sh
```

Scripts are fairly self-explanatory, but they only work on *nix OS's.
 - setpath.sh will do its best to set default GOPATH information.
 - getdeps.sh stores all known package dependencies and will make fetching them easy.
 - build.sh calls both setpath.sh and getdeps.sh before attempting to build a binary and copy config files.
 - run.sh calls build.sh and runs the resulting application.

```txt
 config
 -- server.json
 -- server-secure.json
```

Configuration files get copied to the output directory with the binary during the build process.
 - server.json can be referenced inside your application if you only need a standard socket server
 - server-secure.json can be referenced inside your application if you need a secure socket server

Config files are loaded in main.go.  Be sure to double-check the content of the config files to adjust settings.

```txt
 src
 -- app
 ----- main.go
 ----- plugins.go
 ----- routes.go
 ----- controller-*.go
```

The app directory contains all of your application files.  This is where you'll do most of your development.
 - main.go configures and starts the server
 - plugins.go is responsible for loading plugin packages
 - routes.go connects socket.io endpoints to controllers
 - controller-*.go contain methods that are called by connected clients based on the endpoint configuration in routes.go

## Original Author

- [Aaron Belovsky](https://github.com/ambelovsky)

## License

MIT
