
## Dependency Management
1. Vendor directory
    * You can add your dependency packages in vendor directory and upload it to source control. Go will use the packages from this vendor directory instead of using the one available in GOPATH
2. Dep (**Obsolete with golang 1.2**)
    * **Install**: go get -u github.com/golang/dep/cmd/dep
    * **Initialize the project**: dep init
    * **Add dependencies**: dep ensure -add \<package1> \<package2>
    * **Update dependencies**: dep ensure -update \<package1> \<package2>
    * **Update all dependencies**: dep ensure -update
    * **Sync dependencies**: dep ensure
3. Go Modules (**Recommended**)
    * A module is a collection of related Go packages that are versioned together as a single unit
    * Modules record precise dependency requirements and create reproducible builds
    * modules must be semantically versioned in the form v(major).(minor).(patch), such as v0.1.0, v1.2.3 or v3.0.1
    * A module is dfined by a tree of Go source files with a go.mod file in the tree's root directory. Module source code may be located outside of GOPATH
    * Uses minimum version selection (MVS) algorithm
    * **Initialize a module**: 'go mod init' initialises a module with a file go.mod
    * 'go get' will automatically update the go.mod file
    * To upgrade to the latest version for all transitive dependencies of the curren module:
        ```
        go get -u
        ```
    * If the package is one which can break past functionality then also change the package for e.g
        ``` 
        banking/pkg/corebanking v1.0.0
        banking/pkg/corebanking/v2 v2.0.0
        ```

### Athens
Athens is a goproxy server and a CDN which is available on cloud from Azure, AWS, GCP

* Settng up Athens Proxy
    ```
    $ mkdir -p $(go env GOPATH)/src/github.com/gomods
    $ cd $(go env GOPATH)/src/github.com/gomods
    $ git clone https://github.com/gomods/athens.git
    $ cd athens
    $ GO111MODULE=on go ./cmd/proxy -config_file=./config.dev.toml &
    [1]25243
    INFO[0000] Starting application at 127.0.0.1:3000
    ```
* Running Athens Proxy
    ```
    export GO111MODULE=on
    export GOPROXY=http://127.0.0.1:3000
    ```