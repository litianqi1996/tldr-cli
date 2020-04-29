# tldr-cli

[Tldr](https://github.com/tldr-pages/tldr) client for golang



## Installation

OS X & Linux:

```bash
brew install litianqi1996/taps/tldr
```



## Usage example

```
tldr [command] 
tldr -u/--update  //update tldr pages from gitrepo
tldr -c/--clean   //clean tldr local repo
```



## Configure

```bash
$HOME/.tldrtmp/tldr.yaml
```

```bash
# set tldr-pages by yourself
gitrepo: https://github.com/tldr-pages/tldr

# set language, default "" means english.  
# languagse: "de", "es", "fr", "hbs", "it", "ja", "ko", "pt_BR", "pt_PT", "ta", "zh"
language: "" 

# no need to pay attention
updatetime: 1588094105
```



## Built With

* [fatih/color](https://github.com/fatih/color)
* [go-git](https://github.com/src-d/go-git)
* [goreleaser](https://goreleaser.com/)
* [figlet](http://www.figlet.org/)

