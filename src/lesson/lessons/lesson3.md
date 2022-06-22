# Basics of the file system and practicing commands

In this lesson we will apply what you have learnt so far to help you navigate
your computer using the command line.

## A refresher from the last lesson

In the last lesson we looked at how to formulate commands.
`command` `--flags` `arguments`

Now we will practice using some commands to help us navigate around the file system.

Type `n` or `next` and hit enter to continue.

## Prompt

Before we begin, let's explain this funny looking line that has been following you where you type.
It is called the `prompt`.

It displays information about the:

current user = `{{GetUser}}`← You!
host machine = `{{GetHost}}` ← Your machine
the current directory you are in = `{{GetWd}}` ← Where on you are on your machine.

The `prompt` is there to give you some additional context about the state of
your system. Importantly for us, it conveniently tells us where we are on our
computer. This will be very helpful as we proceed.

Type `n` or `next` and hit enter to continue.

## Present Working Directory

Where you are at any moment on your computer's file system is known as the
"current working directory". As discussed before it is also often displayed in
the prompt.

However, we can also use a shell command to **print** our **working** **directory** to the screen.
Notice how the command is named by its task:
`print` `working` `directory`
`↓       ↓         ↓`
`p       w         d`

Try it out for yourself, Type `pwd` and hit enter.

> {{GetWd}}

## Listing files

Now that we know where we are a natural follow up question would be, "What else do
we have here?"

To list to contents of a directory we use the `ls` command to list files in a directory.

Try it out for yourself, Type `ls` and hit enter.

> !ls

## Listing files

The `ls` command also takes a number of flags. For example the `-a` flag which
stands for *all*. This instructs the ls program to show all the files in the
directory including the hidden ones.

Try it out for yourself, Type `ls -a` and hit enter.

> !ls -a

## Awesome work!

This is the end of the lesson, feel free to continue playing around with commands we have covered in this lesson:
`ls` and `cd`.

When you are done press type `n`, `next` or `quit` to end this lesson.

### Vocabulary

wc, pwd, ls, cd, whoami, uname, echo, curl, man, clear, less, !!, cat, cal
