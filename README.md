# OC-prompt
Interactive CLI tool for OpenShift. Includes autocomplete for valid functions and flags for each function.

# Installation 

1. Install [Golang](https://go.dev/doc/install) and [Openshift-CLI](https://docs.openshift.com/container-platform/4.8/cli_reference/openshift_cli/getting-started-cli.html)
2. Clone this repository 

            git clone https://github.com/Eli-IBM/OC-prompt.git
3. Run the following command
   
            go build -o oc-prompt main.go
4. Copy file into local executables
   
            cp oc-prompt /usr/local/bin/
5. To execute run

            oc-prompt
----------------

## Phase 1 
- [x] Compile a list of all the flags from command arguments
- [x] Compile a list of all the openshift command arguments
- [x] Decide on how to access/store commands/flags
- [x] Parse file for commands/flags 
- [x] create a completer function
## Phase 2
- [ ] Implement Usage instruction for each Openshift Command
- [x] Pull Recommended local files Into AutoComplete
- [ ] Pull Recommended Ports Into Autocomplete

## Phase 3
- [ ] Create Tool to speed up adding new commands/flag
- [ ] Create Easy Interface to add new commands
- [ ] Cache parsed JSON to make runtime faster
- [ ] Pull Local Images Into Autocomplete
- [ ] Implement History

## Optional Features
- [ ] login with ENV Vars

# Testing
Built on: <br />
            Client Version: 4.10.18 <br />
            Server Version: 4.9.33 <br />
            Kubernetes Version: v1.22.8+c02bd9d <br />
            Go Version: go1.18.3 darwin/amd64 <br />
<br />
Libraries Used: <br />
        github.com/c-bata/go-prompt

# Contributors
Architeture - Eli Feinberg <br />
Data Formating - Jeffrey Huang <br />