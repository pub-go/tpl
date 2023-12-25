## Links
- JDK 11+
- https://github.com/antlr/antlr4/blob/master/doc/getting-started.md
- https://github.com/antlr/antlr4/blob/master/doc/go-target.md

## Download ANTLR4 JAR file
```bash
cd /usr/local/lib
curl -O https://www.antlr.org/download/antlr-4.13.1-complete.jar
```

## Setup ANTLR4
```bash
export JAVA_HOME=/Library/Java/JavaVirtualMachines/jdk-17.0.1.jdk/Contents/Home
export CLASSPATH=".:/usr/local/lib/antlr-4.13.1-complete.jar:$CLASSPATH"
alias antlr4='java -jar /usr/local/lib/antlr-4.13.1-complete.jar'
alias grun='java org.antlr.v4.gui.TestRig'
```

## Generate Go code
```bash
rm *.go *.interp *.tokens
antlr4 -Dlanguage=Go -visitor -package parser *.g4
go mod tidy
```
