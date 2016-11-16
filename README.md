# Look At My Site

![image](https://cloud.githubusercontent.com/assets/10496851/20352327/d4e984d6-ac51-11e6-86ca-3ef018e32a45.png)


## About

Look At My Site is a free service that checks your domain (and subdomains if a comprehensive scan is chosen) against several best security practices for websites. The current checks supported by our service are:
- Support for HTTPS
- Use of HTTP Strict Transport Security (HSTS)
- Use and correct configuration of Content-Security-Policy
- Implementation of mechanisms against DNS rebinding attacks

We hope to encourage the adoption of best security practices through the providence of this free service.

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
 - [x] Implement crawler
 - [x] Create separate queues for each checker
 - [x] Parse `comprehensive` parameter and call `CheckAll` or `CheckBasic` accordingly (basic only checks the root domain with the 3 basic checks whereas "all" checks subdomains and includes CORS and directory listing checkers)
 - [x] Implement HTTPS checker
 - [x] Implement HSTS checker
 - [x] Implement DNS rebinding checker
 - [x] Implement CSP checker
 - [x] Fill in proper texts on the landing page
 - [x] Fill in proper texts for each checker
    - [x] HTTPS
    - [x] HSTS
    - [x] DNS rebinding
    - [x] CSP
 
 ## Stretch Goals
 - [x] Queueing of scan requests
 - [x] Make sure that our site is safe against all the scanned attacks
 - [x] Auto focus on input textbox upon document ready
 - [x] Prevent submits while a scan result is loading
