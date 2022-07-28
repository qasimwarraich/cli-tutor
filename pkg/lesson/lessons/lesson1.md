# Basics of Textual Interaction

This lesson will introduce you to the very basics of command line interaction.
We will introduce you to the tutor program and how to formulate a command.

## Introduction

Welcome to the command line interface (or "CLI" for short). It may look
intimidating for now, but this tutorial will familiarise you with the concepts
you need to thrive on the command line.

Below, you can see the `command prompt`, where you can type commands. When you
are ready, type `n` or `next` and hit enter to continue.

## What is the Shell?

Essentially, the `shell` is a command line interface that let's you "talk
directly to your computer" using text messages. Instead of browsing a hierarchy
of menus and dialogs like in a conventional, graphical user interface, the
shell lets you access any setting or functionality of your computer directly
through dedicated `commands`. The `shell` will interpret these commands and
produce appropriate output.

More formally, the `shell` is a program that wraps your operating system (hence
its name) and acts as an intermediary between you, the user, and the operating
system. It manages the user's interaction by accepting input in the form of
`commands` and relays responses in the form of produced output, errors or
special shell features like shortcuts.

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

## How to issue commands

Commands supplied to the shell usually adhere to the following pattern:

`command` `--flags` `arguments`

A `command` is the name of a program the shell can run, while `--flags` and
`arguments` are ways to tell the program what to do specifically. Most programs
accept multiple flags and/or arguments.

`--flags` are pre-defined cues that alter the programs behaviour. You can
imagine these as "switches" that turn certain options on or off.

`arguments` are additional pieces of information supplied as free text, which
the program will have to interpret.

Here's an example for a command called `wc` (short for "word count"):

`wc --lines file.txt`

Issuing this command will count the lines in a file called "file.txt". Feel
free to try out this command right now by typing it into the command prompt and
hitting enter. You will see the output of the command just below the prompt and
above the current lesson text. You could also try different flags, like
`--words` or `--chars`, or even all three at the same time. You should see that
"file.txt" contains 4 lines, 26 words and 147 characters.

By the way, `--flags` sometimes have a shorthand version: `wc -l` is the same as
`wc --lines`.

Don't worry about making mistakes. You will always be led back to this lesson.

When you're ready, type `n` or `next` and hit enter to continue.

## Try another Command

Type the command `whoami`, hit Enter and look at the result.

> !whoami

## Let's try something more fun

As mentioned before, the shell is just an intermediary between you and your
operating system. This means you can do essentially anything you want here in
the shell, including accessing the internet, for example.

Let's try this out. Type the command `curl wttr.in/Bern` to instruct the `curl` tool
to access the website `wttr.in/Bern`. You should receive a weather forecast.

When you are done, type `n` or `next` to continue.

## A note about getting overwhelmed

If at any point the amount of text on the screen becomes overwhelming the shell
has an inbuilt command called `clear` that clears all the text on the screen.
Helpfully, this tutor program will ensure to re-display the text for the step
you are currently on. Give it a try by typing `clear` and hitting enter!

This is the end of the lesson, feel free to continue playing around with
commands we have covered in this lesson:
`whoami` and `clear`.

When you are done, type `n`, `next` or `quit` to end this lesson.

### Vocabulary

whoami, curl, clear, wc
