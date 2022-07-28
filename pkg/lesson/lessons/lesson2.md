# Getting more familiar with the shell

In this lesson we will practice formulating commands and get more comfortable
interacting with the shell.

## A refresher from the last lesson

Let's recall that commands are typically formulated using this format:

`command` `--flags` `arguments`

On a higher level we can see this as a sentence like:

`verb` `--adjective/adverb` `noun`

Consider `wc -l file.txt` from the last lesson: the word count program `wc` can
be seen as a verb or action, it's `-l` or `--lines` flag can be considered an
adjective, and, finally, the noun would be the file (`file.txt`) to operate on.

Read all together, `wc --lines files.txt`, can be read as 'count all the lines
in file.txt' in English.

It's worth mentioning that some programs allow `--flags` to be specified after
the `arguments`.

Type `n` or `next` and hit enter to continue.

## Sub Commands

Commands or programs may also have several sub-commands
that perform different actions and may have their own `--flags` to further
control those actions.

*e.g.* `apt install -y vim` or `git clone REPO --bare`

We can still apply the same analogy of a sentence as discussed in the previous
step, with a slight addition.
`program` `verb` `--adjective/adverb` `     noun`
`↓         ↓           ↓               ↓  `
`apt   install        -y              vim`

*In English*:

```
→ "APT package manager install vim and confirm any prompts y/n with the answer yes."
```

If all this makes sense, type `n` or `next` and hit enter to continue or type
`p` or `prev` to return to the previous step of the lesson.

## Flags

As previously mentioned, `--flags` are ways we can modify or specify the
behaviour of a program or one of it's sub-commands.

Remember that flags can be specified in a long or short form, *e.g.* `-l` vs
`--lines` which both specify the `wc` program to count the lines in a file.

Commonly used flags in programs often have this pair of long and short forms to
specify the command. The important thing to see here is that there are
different amounts of dashes (`-`) in both commands. Another cool functionality
of short form flags is that they may be stacked together to combine the
behaviours of multiple flags the program may accept, for example:

`wc -wl file.txt`

Here, we stack the two short form flags `-w` and `-l` together as `-wl` This
specifies to the `wc` program to output both the word count (`-w`) and the line
count (`-l`). This is equivalent to running:

`wc --words --lines file.txt`

Flags can also be optional, for example when `wc` is run without flags like:

`wc file.txt`

It would perform its default behaviour and print out the word, line and
character counts.

Let's practice this a little bit. Type `n` or `next` to begin.

## Let's practice

Type the `wc file.txt` command and hit enter.

> !wc file.txt

## Now let's use some flags

Type the `wc -l file.txt` command and hit enter.

> !wc -l file.txt

## Now let's use some flags

Type the `wc -w file.txt` command and hit enter.

> !wc -w file.txt

## Finally, let's stack some flags

Type the `wc -wl file.txt` command and hit enter.

>  !wc -wl file.txt

## Change the order of the flags

Notice the order of the flags in this case does not influence the output.

Type the `wc -lw file.txt` command and hit enter.

>  !wc -lw file.txt

## Awesome work!

This is the end of the lesson, feel free to continue playing around with
commands we have covered in this lesson:
`wc` and `clear`.

When you are done, type `n`, `next` or `quit` to end this lesson.

### Vocabulary

wc, whoami, curl, clear
