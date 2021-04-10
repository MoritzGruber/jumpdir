# jumpdir - a little script to jump to you git projects faster 

## Install

clone the repo and run (you have go to be setup)
```bash
go install
```

add this to you .bashrc or .zshrc
```bash
export JUMPDIR_ROOT="$HOME/git"
export JUMPDIR_IGNORE="node_modules,.git,yarn,nvm,npm,Library,dist,lib,target,build,go/pkg/mod,.next,out,vendor"
j() {
    cd $(jumpdir $1)
}
jc() {
    cd $(jumpdir $1) && code .
}
```

## Usage
```bash
j mynestedproject
jc mynestedproject
```