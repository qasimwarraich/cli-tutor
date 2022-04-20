
# CLI TUTOR

## What is this?

Despite the arguably dated appearance, difficult learning curve and practical
non-existence in the general personal computing space, Command Line Interfaces
(CLIs) have more than stood the test of time in the software development world.
There are a multitude of extremely popular tools and applications that
primarily focus on the command line as an interaction medium. Some examples
include version control software like `git`, compilers and interpreters for
programming languages, package managers and various core utilities that are
popular in areas such as software development, scripting and system
administration.

As mentioned before, the use of the command line as an interaction paradigm has
effectively disappeared from a mainstream personal computer usage perspective.
This contributes greatly to the intimidation factor and learning difficulty for
those interested in getting into software engineering or system administration.
This unfamiliarity, paired with the inevitability of usage of CLIs in the
development space highlights a need to make the command line more accessible to
new users for whom text-based interaction with their computer is an alien
concept. In recent years, interactive learning aides utilising tools such as
sandboxed environments have been gaining in popularity and have the potential
to be a suitable medium for learning command line basics through actual usage,
examples and practice.

## The goals of this tool

This tool aims to determine whether an interactive learning method may ease
the introduction into command line interfaces for novice users, particularly
mitigating the 'scare factor' experienced by first-time users. We do so by
creating a forgiving CLI with the goal of teaching topics such as shell
scripting basics and Unix-like core utility usage through the use of
interactive examples. We draw inspiration from the `vimtutor` utility shipping
alongside the popular terminal based text editor Vim. The proposed tool shall
allow for opt-in analytics that are sent back to a data collection service for
the purpose of learning which mistakes are most commonly made, and to improve
the tool accordingly. To validate the tool and answer our research questions, a
user study will be conducted, most likely with bachelor students at the
University of Zurich. A secondary goal is to embed the learning tool into a
prototypical web application in order to make it more accessible and portable.
