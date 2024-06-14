# Epha

## Getting Started

TODO:

## Development Setup

You need antlr4 to get started

```bash
(cd /usr/local/lib && sudo curl -O https://www.antlr.org/download/antlr-4.13.1-complete.jar)
```

Add these to your terminal `.bashrc` or `.zshrc`, depending on your terminal.
```bash
export CLASSPATH=".:/usr/local/lib/antlr-4.13.1-complete.jar:$CLASSPATH"
alias antlr4='java -Xmx500M -cp "/usr/local/lib/antlr-4.13.1-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
alias grun='java -Xmx500M -cp "/usr/local/lib/antlr-4.13.1-complete.jar:$CLASSPATH" org.antlr.v4.gui.TestRig'
```