# Look At My Site

## Setup
1. Install [golang](https://golang.org/doc/install)
  - Shortcut for macOS if you have `homebrew` installed.
  
    ``` brew install go ```
    
  - Shortcut for Ubuntu
  
    https://github.com/golang/go/wiki/Ubuntu
    
2. Create a workspace
  - Go requires all your go code to be in the same workspace, grouped according to namespaces.
    
    https://golang.org/doc/code.html
  
  - Set GOPATH to the location of your workspace for your shell.
  
    ``` export GOPATH='/path/to/workspace' ```
    
  - Clone this repo to ```$GOPATH/src/github.com/thngkaiyuan/look-at-your-site```
  
3. Test the build
  1. `cd` into the repo.
  2. Run `make build`, everything should build correctly. A binary named `look-at-my-site` will be generated. 
  3. For testing your code changes, it is easier to use `make serve` which will start the server and listen at port 8080.
  
 ## Todo
 - [ ] Implement crawler
 - [ ] Create separate queues for each checker
 - [ ] Parse `comprehensive` parameter and call `CheckAll` or `CheckBasic` accordingly (basic only checks the root domain with the 3 basic checks whereas "all" checks subdomains and includes CORS and directory listing checkers)
 - [ ] Implement HTTPS checker
 - [ ] Implement HSTS checker
 - [ ] Implement DNS rebinding checker
 - [ ] Implement CSP checker
 - [ ] Fill in proper texts on the landing page
 - [ ] Fill in proper texts for each checker
    - [ ] HTTPS
    - [ ] HSTS
    - [ ] DNS rebinding
    - [ ] CSP
    - [ ] CORS
    - [ ] Directory listing
 
 ## Stretch Goals
 - [ ] Implement CORS checker
 - [ ] Implement directory listing checker
