# Basics of the file system and practicing commands

In this lesson, we will apply what you have learnt so far to help you navigate
your computer using the command line.

## A refresher from the last lesson

In the last lesson, we looked at how to formulate commands.
`command` `--flags` `arguments`

Now we will practice using some commands to help us navigate around the file system.

Type `n` or `next` and hit enter to continue.

## Prompt

Before we begin, let's talk a bit more about the `command prompt` where you've
been entering commands.

It displays information about the:

current user = `{{GetUser}}` ← You!
host machine = `{{GetHost}}` ← Your machine
current directory you are in = `{{GetWd}}` ← Where you are on your machine.

The `prompt` is there to give you some additional context about the state of
your system. Importantly for us, it conveniently tells us where we are on our
computers. This will be very helpful as we proceed.

Type `n` or `next` and hit enter to continue.

## Current Working Directory

Where you are at any moment on your computer's file system is known as the
"current working directory" (CWD). As discussed earlier, it is also often
displayed in the prompt.

*Note*: The term "directory" is interchangeable with the term "folder".

However, we can also use a shell command to **print** our **working** **directory** to the screen.
Notice how the command is named by its function:
`print` `working` `directory`
`↓       ↓         ↓`
`p       w         d`

Try it out for yourself, Type `pwd` and hit enter.

> !pwd

## Path

What the `pwd` command actually outputted to you is known as the `path` to your
current working directory. The `path` in this case describes the location of
something on the file system or in other words the path one must take from the
`root` (top level) of the file system to the location in question (in this case
our current working directory).

The file system is a tree structure where you can take different paths from the
root. The root is typically depicted with a `/`. The tree below shows a section
of the file system stemming from the home folder which resides in the root
directory. `/` + `home` = `/home`. There are a lot more paths one can take from
the root but for the sake of clarity let's stay in the user's home directory
for now.

```
/home
`-- cli-student
    `-- tutor  <--- This is your current working directory.
```

Type `n` or `next` and hit enter to continue.

## Important to know!

When you open a terminal the shell program will always be pointing to some
directory, this is what we previously introduced as the "current working
directory". On most systems, this defaults to the user's home directory which is
also sometimes represented with the shorthand `~`.

The prompt looks something like this: `user@host:~`.
If you were to use the `pwd` command the output would look something like: `/home/user/`.

The file system tree is something that can be navigated up and down, but first
one must know where to go. We need to be able to list the files and
directories in our current working directory in order to choose a path to
navigate down.  

We will discuss this in the next step. Type `n` or `next` and hit enter to continue.

## Listing files

Now that we know which folder we're in, a natural follow-up question would be
"What does this folder contain?"

To list to contents of a directory, we use the `ls` command to list files in a directory.

Try it out for yourself, Type `ls` and hit enter.

> !ls

## Listing files with flags

The `ls` command also takes a number of flags. For example the `-a` flag which
stands for *all*. This instructs the ls program to show all the files in the
directory including the hidden ones.

Try it out for yourself, Type `ls -a` and hit enter.

> !ls -a

## More interesting stuff about paths

Notice the two additional results at the top: `.` and `..`. These are shorthand
for **current directory** and **parent directory** (the directory containing
the current) respectively.

This is not to be confused with when a `.` is the first character of a
filename, **e.g.**: `.hiddenfile.txt`. By convention, files starting with a `.`
are "hidden files". These will be hidden by default, both when typing `ls` on
the command line or when using a graphical file browser.

The `ls -a` command displays these hidden "dotfiles" because of the `-a` flag.

Type `n` or `next` and hit enter to continue.

## Changing Directory

Now let's navigate around the file system a bit. To do this, we will utilise the
`cd` command.
`change` `directory`
`↓        ↓`
`c        d`

First use the `ls` command to list where you are and then use the `cd **NAME OF
DIRECTORY**` command to change into the desired directory.

The changing of your directory can be observed by the `pwd` command or by
looking at the `path` displayed in your `prompt`.

*Hint*: Use the command `cd ..` to move up to a directory (parent directory).

Type `n` or `next` and hit enter to continue.

## Creating Directories

Now let's practice creating directories.
To make directories we can use the `mkdir` command.

To create a directory in your current working directory we must supply the
`mkdir` command with a name like:
`mkdir NewDirectory`

This will create a new directory named `NewDirectory` at the current working directory.

Use ls to confirm that your new directory was created.
Type `n` or `next` and hit enter to continue.

## Removing Directories

To remove directories we can use the `rmdir` command.

To remove a directory in your current working directory we must supply the
`rmdir` command with a name like:

`rmdir NewDirectory`

This will remove a directory named `NewDirectory` at the current working directory.

*Note*: This will only work with empty directories, to delete directories with
files inside them, we must use the file deletion command `rm`. We will discuss
this in the next section.

Use ls to confirm that your directory has been deleted.
Type `n` or `next` and hit enter to continue.

## Removing Directories II

To remove directories with contents we can use the `rm` command.

To remove a non-empty directory or regular file in your current working
directory, we must supply the `rm` command with a name like:

*Removing a file*
`rm filename.txt`

*Removing a non-empty directory*
`rm -r directory/`

The `-r` flag instructs the `rm` command to apply the procedure recursively,
which is to say it should delete the directory and everything inside it.

*Warning*: The `rm` command can be dangerous, and its deletions are not
reversible. Be careful to ensure you supply the correct command; otherwise you
may lose data or cause harm to your system.

The command line gives you a lot of power over your computer. However, with
this great power comes great responsibility.

Type `n` or `next` and hit enter to continue.

## Creating files

Creating empty files can be done simply by using the `touch` command.

To create a file in your current working directory, we must supply the
`touch` command with a file name like:
`touch NewFile.txt`

This will create a new file named `NewFile.txt` in the current working directory.

You can use `ls` to confirm this.

To remove this file, simply use the `rm` command like: `rm NewFile.txt`.

When this makes sense, type `n` or `next` and hit enter to continue.

## All Together Now

Let's put everything we have learnt all together now into a full example:

Follow each step one by one:
`mkdir NewDirectory`  # Creates a directory

`ls`    # View your created directory

`cd NewDirectory`   # Go into your new directory

`touch NewFile.txt`    # Create a file

`ls`    # View your created file

`cd ..` # Go back to the parent directory

`rmdir NewDirectory` # Try to delete the directory **(This will fail as the directory is not empty)**

`rm -r NewDirectory` # Delete it the right way!

`ls`    # Confirm that the directory has been deleted.

When this all makes sense, type `n` or `next` and hit enter to continue.

## Great work!

This is the end of the lesson, feel free to continue playing around with
commands we have covered in this lesson:
`ls`, `pwd` and `cd`, `rm`, `rmdir` and `touch`.

When you are done, type `n`, `next` or `quit` to end this lesson.

### Vocabulary

wc, pwd, ls, cd, whoami, uname, echo, curl, man, clear, less, !!, cat, cal, touch, mkdir, rm, rmdir
