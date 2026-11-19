# command_line_arguments_01: os.Args

`os.Args` holds the command line, **starting with the program's own path**:

```text
$ ./greet Ada Bob
os.Args = ["./greet", "Ada", "Bob"]
```

The program's logic is in `run`, which takes the arguments and output streams as parameters so
tests can call it directly. `main` just connects it to the real process.

```text
$ greet Ada Bob
Hello, Ada!
Hello, Bob!
$ greet
usage: greet NAME...      (on stderr, exit status 2)
```

The program also greets itself, and never prints the usage message.

## Your task

Fix the program.

Reference: https://gobyexample.com/command-line-arguments
