# OC-prompt
Interactive CLI tool for OpenShift. Includes autocomplete for valid functions and flags for each function.

## Phase 1 
- [ ] Compile a list of all the flags from command arguments
- [x] Compile a list of all the openshift command arguments
- [x] Decide on how to access/store commands/flags
- [x] Parse file for commands/flags 
- [x] create a completer function
## Phase 2
- [ ] Implement Usage instruction for each Openshift Command
- [ ] Pull Recommended local files Into AutoComplete
- [ ] Pull Recommended Ports Into Autocomplete
- [ ] login with ENV Vars

## Phase 3
- [ ] Create Tool to speed up adding new commands/flag
- [ ] Create Easy Interface to add new commands
- [ ] Cache parsed JSON to make runtime faster
- [ ] Pull Local Images Into Autocomplete
- [ ] Implement History

# Testing
Built on: <br />
        Client Version: 4.10.18 <br />
        Server Version: 4.9.33 <br />
        Kubernetes Version: v1.22.8+c02bd9d <br />

# Contributors
Architeture - Eli Feinberg <br />
Data Formating - Jeffrey Huang <br />