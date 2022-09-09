# Helping yourself!

In this lesson, we will help you help yourself. #RTFM

## Man, man's best friend

The most helpful in-built shell assistant you have at the command line (other
than this tutorial software) is the utility `man`. `man` is short for
**manual** and is the shell utility that displays the documentation for
commands you may have access to.

In our previous lessons on using flags you might have asked yourself, "Where am
I supposed to find out what all these flags mean?", the manual is the place!

Type `n` or `next` and hit enter to continue.

## Some words about pagers

Before we take a look at the `man` command let's briefly discuss what a pager
is and how to use one. This is important because when you issue the man command
you will be dropped into a very minimal-looking tool known as a pager.

A pager is a tool that "paginates" output from other programs. That is to say,
it divides it into a series of pages or a scrollable window. This is important
because the manuals of certain programs may be very long and having all that
text dumped into your terminal is not only jarring but also difficult to work
with.

Type `n` or `next` and hit enter to continue.

## Some more words about pagers

The pager you will most likely be exposed to in the wild on the command line is
known as `less`. The `less` command is something you can use to paginate any text
really. For instance, you can use it with regular files like:
`less reallylongtextfile.txt`

There are some important keybindings to know when it comes to `less` to avoid
getting stuck in the program. The most important being `q`, which is how you
quit.

Other keybindings that will be important are:
Moving up: `↑` or `k`
Moving down: `↓` or `j`

Moving one page up: `pageup` or `Ctrl+u`
Moving one page down:`pagedown` or `Ctrl+d`

searching: `/`

And remember, to exit the pager: `q`

Type `n` or `next` and hit enter to continue.

## Let's try it out

Let's try reading the manual for a command we are already familiar with. This
is a great way to learn about the extended capabilities of command line tools.

Try:
`man wc`

**Remember!** Press `q` to quit the pager and return to the tutor, happy reading!

When you hopefully make it back here, Type `n` or `next` and hit enter to continue.

## Another way of getting help

Often programs will have a special flag that prints out usage information or
directly opens up a `man` page for you. This flag is typically a variation of
`-h, --help` or `-?`. 

Occasionally, they are also available as a sub-command that can be used like:
`cli-tutor help`

Type `n` or `next` and hit enter to continue.

## Great work!

This is the end of the lesson, feel free to continue playing around with `man`
and learn about some more commands.

*Hint:* There is even a man page for `man`. Try `man man`

When you are done, type `n`, `next` or `quit` to end this lesson.

### Vocabulary

wc, pwd, ls, cd, whoami, uname, echo, curl, man, clear, less, !!, cat, cal, touch, mkdir, rm, rmdir
