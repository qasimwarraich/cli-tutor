# Basics of Textual Interaction

This lesson will introduce you to the very basics of command line interaction.
We will introduce you to the tutor program and how to formulate a command.

## Introduction

Welcome to the command line, this may look intimidating but, this  
tutorial will introduce you to the concepts you need to be familiar  
with to thrive at the command line.

When you are ready type `n` or `next` and hit enter to continue.

## What is the Shell?

Textual interaction is a form of interacting with your computer's operating
system using textual input. You, the user, can issue `commands` and the `shell`
will interpret them and produce an output.

The `shell` is a program that wraps your operating system and acts an
intermediary between the user and the operating system. It manages the user's
interaction by accepting input in the form of `commands` and relays output in the
form of produced output, errors or special shell features like shortcuts.

```
                       shell───────────────────────┐
                       │     interpret             │
┌──────┐               │     ▲     │               │
│ user ├─command──────►├─────┘     │       ┌────┐  │
└───▲──┘               │           └──────►│ os │  │
    │                  │                   └──┬─┘  │
    └─────output───────┤                      │    │
                       │◄──output─────────────┘    │
                       │                           │
                       └───────────────────────────┘
```

This graph is an illustration of the interaction cycle between the user and shell.

Type `n` or `next` and hit enter to continue.

## Command Formulation

Commands can generally be supplied to the shell by using the following format:

`command` `--flags` `arguments`

e.g. `wc --lines file.txt `

A `command` is actually a program the shell can run. In the above example the program being called is `wc`, a word counting program.

`--flags` are additional cues that can be provided to the program to alter its
behaviour. These can be optional, but almost all command line programs have
flags to help produce the output the user wants. In the above example we supply
the `--lines` flag to let `wc` know we are only interested in the line count
and not the rest of the information the program can produce.

Finally, we have the `argument` which is the input to the program, in the case of our example, a file named `file.txt`.

Type `n` or `next` and hit enter to continue.

## Try a Command

Try it out for yourself!
Type the command `whoami`, hit Enter and look at the result.

> {{GetUser}}

## A note about getting overwhelmed

If at any point the amount of text on the screen becomes overwhelming the shell
has an inbuilt command called `clear` that clears all the text on the screen.
Helpfully, this tutor program will ensure to re-display the text for the step
you are currently on. Give it a try!

This is the end of the lesson, feel free to continue playing around with commands we have covered in this lesson:
`whoami` and `clear`.

When you are done press type `n`, `next` or `quit` to end this lesson.

### Vocabulary

pwd, ls, cd, whoami, uname, echo, curl, man, clear, less
