# Shell shortcuts and tricks

In this lesson we will learn about some shell shortcuts and tricks to make you
more productive at the command line.

## Woof, that's a lot of typing!

You may have the feeling that issuing commands is very tedious, time-consuming
and hard to get right. This requirement for commands to be stated perfectly is
one of the biggest obstacles to new command line users. Let's overcome this
hardship and look into some cool in built shell features designed to
make your life a little easier.

Type `n` or `next` and hit enter to continue.

## The Up Arrow ↑ and shell history.

The up arrow `↑` is one of the most helpful shortcuts when using the command line. It allows you to quickly fill your input with the previous command. This means if you slightly mistyped a long command you don't have to rewrite the whole command, you can simply use the up arrow `↑` to go up in your command history and retrieve the command. 

*Note*: You can hit the up arrow as many times as you want and keep going backwards in you command history.

Type `n` or `next` and hit enter to continue.

## More on shell history

The shell history is simply a hidden file that usually lives in the home folder
of your system and is a record of every command you have issued in a particular
shell. If you are using the `bash` shell this file is called `.bash_history`.

The history file maintained by this tutor is called `.tutor_history` and resides in the folder you started this program in. 

`ls -a` should display it and to view its contents in the shell you can use the `cat` command like `cat .tutor_history`.

Try this out!

> !cat .tutor_history

## History search

Sometimes moving sequentially through the command history is too slow and
impractical, especially when we are looking for a command we issued a long time
ago.

To help us with this the shell has an inbuilt reverse history search that you
can trigger by pressing `Ctrl + r`. This will open up a text input prompt where
you can type letters and the command that matches best will be shown in the input
area.

Try it out! Type `Ctrl + r` and search for that `cat .tutor_history` by typing
`ca` the command we issued earlier should autofill, and now you can just press
Enter to execute it.

> !cat .tutor_history

## The Double Bang!!

Another very useful shell shortcut is the double bang `!!`. This shell shortcut
replaces the `!!` with the previous command. This is very useful if something
is forgotten proceeding or following the command.

**e.g.**:

You forgot to supply a flag:

```
ls
**OUTPUT of ls**

!! -a
**OUTPUT OF ls -a**

```

Type `n` or `next` and hit enter to continue.

## The Best for last

By far the biggest efficiency booster is going to be `<Tab>` completion.
`<Tab>` completion allows us to fill in completions to things like commands or
paths to files on our system by using the `<Tab>` key on our computer to make
the selections. There are innumerable use cases for tab completion but let's
try to think of it in simple terms. Up until now if we wanted to change our
directory we first used `ls` to show the contents of our current working
directory, and then we used `cd` and typed the directory name in fully to
navigate to that directory.

With `<Tab>` completion we can simply type `cd<Tab>` to open up the completion
menu. You will see a list of candidate directories of which you can cycle
through by using `<Tab>` again. You can even further filter the results down by
continuing to type letters to narrow the list of candidates.

Give it a try!

When this all makes sense, type `n` or `next` and hit enter to continue.

## Great work!

This is the end of the lesson, feel free to continue playing around with shortcuts we have covered in this lesson:
`↑`, `Ctrl + r`, `!!` and `<Tab>`

When you are done press type `n`, `next` or `quit` to end this lesson.

### Vocabulary

wc, pwd, ls, cd, whoami, uname, echo, curl, man, clear, less, !!, cat, cal, touch, mkdir, rm, rmdir
