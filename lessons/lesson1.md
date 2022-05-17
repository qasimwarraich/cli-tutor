# Basics of Textual Interaction

This lesson will introduce you to the very basics of command line interaction.
We will introduce you to the tutor program and how to formulate a command.

## Introducton

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
 
## Prompt

This funny looking line below you is called the `prompt`.

It displays information about the:

current user = `{{GetUser}}`← You!
host machine = `{{GetHost}}` ← Your machine
the current directory you are in = `{{GetWd}}` ← Where on you are on your machine.

## Try a Command

Try it out for yourself!
Type the command `whoami`, hit Enter and look at the result.

### Vocabulary
pwd, ls, cd, whoami, uname, echo, curl, man, clear, less, vim
