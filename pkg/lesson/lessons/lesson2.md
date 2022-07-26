# Getting more familiar with the shell

In this lesson we will practice formulating commands and get more comfortable
interacting with the shell.

## A refresher from the last lesson

Let's recall that we mentioned commands are typically formulated using this format:
`command` `--flags` `arguments`

On a higher level we can see this specification of commands as a sentence like:
`verb` `--adjective/adverb` `noun`

*e.g.* `wc -l file.txt` from the last lesson. The word count program `wc` can be
seen as a verb or action. It's `-l or --lines` flag can be considered an adjective
in this context. Finally, the noun would be the file (`file.txt`) to operate on.

Read all together, `wc --lines files.txt`, can be read as 'count all the lines
in file.txt' in English.

It's worth mentioning that the `--flags` can also in some cases be specified
after the `arguments`.

Type `n` or `next` and hit enter to continue.

## Sub Commands

Commands or programs may also have several sub-commands
that perform different actions and may have their own `--flags` to further
control those actions.

*e.g.* `go test -v main_test.go` or `git clone REPO --bare`

We can still apply the same analogy of a sentence as discussed in the previous
step, with a slight addition.
`program` `verb` `--adjective/adverb` `     noun`
`↓         ↓       ↓                    ↓  `
`go       test    -v                 main_test.go`

*In English*:

```
→ "Go compiler run the testing package with verbose output on file main_test.go"

```
If all this makes sense, type `n` or `next` and hit enter to continue or type
`p` or `prev` to return to the previous step of the lesson.

## Flags

As previously mentioned, `--flags` are ways we can modify or specify the
behaviour of a program or one of it's sub-commands.

As you may have noticed in the examples sometimes flags are specified in a long
or short form.

*e.g.* `-l` vs `--lines` which both specify the `wc` program to count the lines
in a file.

Commonly used flags in programs often have this pair of long and short forms to
specify the command. The important thing to see here is that there are
different amounts of dashes (`-`) in both commands. Another cool functionality
of short form flags is that they may be stacked together to combine the
behaviours of multiple flags the program may accept.

*e.g.* `wc -wl files.txt` in this example we stack the two short form flags
`-w` and `-l` together as `-wl` This specifies to the `wc` program to output
both the word count (`-w`) and the line count (`-l`).

Flags can also be optional, for example when `wc` is run without flags like:
`wc file.txt`.

`wc` would then perform its default behaviour and print out the word, line and
character counts.

## Let's try this out

Type the `wc file.txt` command and hit enter.

> 4  26 147 file.txt

## Now let's use some flags

Type the `wc -l file.txt` command and hit enter.

> 4 file.txt

## Now let's use some flags

Type the `wc -w file.txt` command and hit enter.

> 26 file.txt

## Finally, let's stack some flags

Type the `wc -wl file.txt` command and hit enter.

>  4  26 file.txt

## Change the order of the flags

Notice the order of the flags in this case does not influence the output.

Type the `wc -lw file.txt` command and hit enter.

>  4  26 file.txt

## Awesome work!

This is the end of the lesson, feel free to continue playing around with commands we have covered in this lesson:
`wc` and `clear`.

When you are done press type `n`, `next` or `quit` to end this lesson.

### Vocabulary

wc, pwd, ls, cd, whoami, uname, echo, curl, man, clear, less
