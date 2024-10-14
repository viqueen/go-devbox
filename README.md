## go-devbox

tools for working on golang projects

### setup

- install it

```bash
./setup.sh binary
```

- build all

```bash
./build.sh all
```

### usage

#### goscan

- scan a go module

```bash
goscan mod "github.com/your/module"
```

- scan a go module for a given set of checks

```bash
goscan mod "github.com/your/module" --checks http,os
```

- scan dependencies of a go module

```bash
goscan deps
```

- scan dependencies of a go module for a given set of checks

```bash
goscan deps --checks net,exec
```
