# Getting more familiar with the shell

In this lesson we will practice formulating commands and get more comfortable
interacting with the shell.

## A refresher from the last lesson

Let's recall that we mentioned commands are typically formulated using this format:
`command` `--flags` `arguments`

On a higher level we can see this specification of commands as a sentence like:
`verb` `--adjective/adverb` `noun`

e.g. `wc -l file.txt` form the last lesson. The word count program `wc` can be
seen as a verb or action. It's `-l or --lines` flag can be considered an adjective
in this context. Finally, the noun would be the file (`file.txt`) to operate on.

Read all together, `wc --lines files.txt`, can be read as 'count all the lines
in file.txt' in English.

It's worth mentioning that the `flags` can also in some cases be specified
after the arguments.

Type `n` or `next` and hit enter to continue.

## Sub Commands

Commands or programs may also have several sub-commands
that perform different actions and may have their own `flags` to further
control those actions.

e.g. `go test -v main_test.go` or `git clone REPO --bare`

We can still apply the same analogy of a sentence as discussed in the previous
step, with a slight addition.
`program` `verb` `--adjective/adverb` `noun`

`go        test   -v                   main_test.go`

*In English*:

```
--> "Go compiler run the testing package with verbose output on file main_test.go"

```

If all this makes sense, type `n` or `next` and hit enter to continue or type
`p` or `prev` to return to the previous step of the lesson.

## Prompt

This funny looking line below you is called the `prompt`.

It displays information about the:

current user = `{{GetUser}}`← You!
host machine = `{{GetHost}}` ← Your machine
the current directory you are in = `{{GetWd}}` ← Where on you are on your machine.

### Vocabulary

pwd, ls, cd, whoami, uname, echo, curl, man, clear, less,
