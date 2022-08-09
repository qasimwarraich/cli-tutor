
# CLI TUTOR

```txt
        _ _       __        __
  _____/ (_)     / /___  __/ /_____  _____
 / ___/ / /_____/ __/ / / / __/ __ \/ ___/
/ /__/ / /_____/ /_/ /_/ / /_/ /_/ / /
\___/_/_/      \__/\__,_/\__/\____/_/

A simple command line tutor application that aims to introduce users to the
    basics of command line interaction.
    Web version is available at https://tutor.chistole.ch

Usage:
  cli-tutor [flags]
  cli-tutor [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  info        Prints information about the tool and log collection
  repo        Prints a url to the git repository for this tool
  version     Print the version number of cli-tutor

Flags:
  -h, --help            help for cli-tutor
  -n, --no-upload-log   Do not send a copy of the log to the developer
  -x, --no-welcome      Do not show welcome message when entering tutor

Use "cli-tutor [command] --help" for more information about a command.
```

[![Docker Image Version (tag latest semver)](https://img.shields.io/docker/v/qasimwarraich/cli-tutor/latest?label=docker)](https://hub.docker.com/r/qasimwarraich/cli-tutor)
[![Website](https://img.shields.io/website?label=Web%20Version&up_color=light%20green&up_message=live&url=https%3A%2F%2Ftutor.chistole.ch)](https://tutor.chistole.ch)
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

### Getting started

To simply try out the tool in a safe dockerized environment visit:

[Link to Web Version](https://tutor.chistole.ch)

### Installation

Clone this repository using:
`git clone https://gitlab.com/qasimwarraich/cli-tutor`

Change directory into the freshly cloned repo.
`cd cli-tutor`

 ---

Once you have cloned and changed directories into the repository you have three
methods to choose from:

#### Method 1: Using Docker (Safest)*

Prerequisites:

- go v1.18 or higher
- docker
- make (optional)

Instructions:

Option 1: Without make

`docker build -t qasimwarraich/cli-tutor .`

Note: You can run the container using:
`docker run -it qasimwarraich/cli-tutor`

---

Option 2: With make

`make docker`

Note: You can run the container using:
`make dockerrun`

---

\* Installation via the Docker method allows the tool to be tested in a safe
containerised environment. This mitigates risk of accidentally modifying files
on your personal machine like in the uncontainerised methods that follow.

--- 

#### Method 2: Using make

Prerequisites:

- go v1.18 or higher
- make

Instructions:

Option 1: Simply run tool without installing:
`make`

Note: This runs the tool in a mode that opts out of log sharing. To run the
tool with log sharing use the `make log` command.

---

Option 2: Just build the binary:

`make build`

This creates a /bin folder in the root of the repository containing the binary.
You can run it by entering:

`./bin/cli-tutor`

Note: In order to clean this up, you can use the `make clean` command.

--- 

Option 3: Actually install the binary:

`sudo make install`

Note: In order to uninstall this from your system use the `sudo make uninstall` command.

---

####  Method 3: Using golang tooling

Prerequisites:

- go v1.18 or higher
- make (optional)

Instructions:

Option 1: Simply run tool without installing:

`go run main.go`
___
Option 2: Actually install the binary:

`go install main.go`

Option 3: Using make:

To just run the tool without installing:
`make`

Note: This runs the tool in a mode that opts out of log sharing. To run the
tool with log sharing use the `make log` command.

To actually install the binary
`make goinstall`
___
Note: In order to uninstall the tool you will need to delete it from your `bin`
folder located under the `$GOPATH`. This is usually located in your home
directory at `~/go/bin/`

e.g. `rm ~/go/bin/cli-tutor` or `make gouninstall`
___

### Usage

Once installed or inside the docker container the tool can be started using `cli-tutor`. 

Note: Currently during the research phase of this tool, the tutor program sends
a copy of a log file back to the developer by default. To opt out of this
feature start the program using the `-n` or `--no-upload-log` flags like:
`cli-tutor -n` or `cli-tutor --no-upload-log`

For more information regarding logging and the tool check out the `cli-tutor
info` and `cli-tutor help` sub-commands.

Once you are in the tutor you are presented with a menu where you can select a
lesson. This menu may be navigated using the arrow keys or j/k for up and down
respectively. Once you pick a lesson you will be dropped into an interactive
shell. Each lesson contains a series of prompts that aim to teach you about the
topic at hand. Tasks can be navigated between using typing `next/prev` or `n/p`
and hitting enter. Some task may require you to issue commands and the tutor
will try to validate your inputs. These commands cannot be skipped using `next
or n`. Additionally, each lesson has a certain list of available commands for
the user to enter. A list of available commands is available by typing
`commands` and hitting enter.


Have Fun!
