<!-- #: Who is the coolest person in the world? | You are :D. | Have a great day! :# -->
# <img src="logo.svg" />

[![Build Status](https://travis-ci.org/alanxoc3/concards.svg?branch=master)](https://travis-ci.org/alanxoc3/concards)
[![Go Report Card](https://goreportcard.com/badge/github.com/alanxoc3/concards)](https://goreportcard.com/report/github.com/alanxoc3/concards)
[![Coverage Status](https://coveralls.io/repos/github/alanxoc3/concards/badge.svg?branch=master)](https://coveralls.io/github/alanxoc3/concards?branch=master)

Turning notes into flashcards, or should I say concards! Concards is my ongoing
attempt to make flashcards simple and easily embeddable into text document
based notes. Concards is much lighter than other flashcard applications such as
such as [Anki](https://apps.ankiweb.net/) or
[Memrise](https://www.memrise.com/), but it is also very powerful striving to
do one thing and do it well.

## Features
- Implements a repetition algorithm similar to [SM2](https://www.supermemo.com/english/ol/sm2.htm).
- Reading from multiple files & directories.
- Conveniently edit cards while reviewing them!
- Helpful syntax for adding reversible cards.
- Built with Unicode in mind!
- Undoing/Redoing support.

## Install
Download the latest binary executable from the [release
page](https://github.com/alanxoc3/concards/releases). At the moment, only Linux
and Mac are supported.

### Building From Source
It should be super simple:
``` bash
go install github.com/alanxoc3/concards
```

### Dependencies
This project currently depends on:
- [stretchr/testify](https://github.com/stretchr/testify) for unit tests.
- [alanxoc3/argparse](https://github.com/alanxoc3/argparse) forked from
  [akamensky/argparse](https://github.com/akamensky/argparse) for CLI options.
- [nsf/termbox-go](https://github.com/nsf/termbox-go) for the terminal gui.
- [mattn/go-runewidth](https://github.com/mattn/go-runewidth) to help with
  Asian characters.

Concards wouldn't be where it is today without those open source projects &
their contributors, so please check them out too!

## Usage
The complete syntax of embedding your flashcards into text documents consists
of these symbols:
```
'#:' = Starts a concards block and also starts a question.
'|'  = Separates sides.
':'  = Separates sides, and adds a reversed card.
':#' = Ends the concards block.
'\'  = Escapes the special tokens above.
```

Here are a few example concards:
```
#: Concards
 : A lightweight embeddable note-taking flashcard program.

#: What does "concards" stand for?
 | Console Cards

#: What does the ":" do in concards?
 | It will add an extra card where the "colon" side is the question and the
   question is the answer.
 | This syntax is especially useful for vocabulary when learning a language and
   can save typing.

#: Can a concard have more than 2 sides?
 | Yes

#: What does a concard look like?
 | \#: It could look like this \| What does a concard look like? \:\#
 | \#: It could also look like this \| with multiple \| answers! \:\#

#: How do you escape a concard keyword?
 | Put a backslash before it. Your text file would show "\#:", but the app shows "#:".

#: How do you show a backslash then a keyword in the concards ui?
 | To see \\#: in the ui, your text document must have 2 backslashes: \\\#:

#: 你好世界
 : Hello World
 : Greetings World

#: The human who created concards.
 : Alan Morgan
:#
```

The easiest way to understand that syntax is by trying it out! Just run
concards on this `README.md` file and see what happens!
``` bash
$ concards README.md
```

## Advanced Usage
### The Predict File
The predict file contains information needed to make a prediction when you
should review a card next.

Here is an example predict file:
```
002141b9b9448a257b05da1f2eb78972 2020-08-08T18:00:17Z 2020-08-02T18:00:17Z 2 1 1 sm2
3dda75cb44ed447186834541475f32e2 2020-08-08T18:00:17Z 2020-08-02T18:00:17Z 1 3 -2 sm2
```

Here is that same file, but annotated:
```
sha256sum cut in half            | next timestamp       | previous timestamp   | total yes count | total no count | current streak | spaced repetition algorithm used
---------------------------------+----------------------+----------------------+---+---+----+----
002141b9b9448a257b05da1f2eb78972 | 2020-08-08T18:00:17Z | 2020-08-02T18:00:17Z | 2 | 1 |  1 | sm2
3dda75cb44ed447186834541475f32e2 | 2020-08-08T18:00:17Z | 2020-08-02T18:00:17Z | 1 | 3 | -2 | sm2
```

This file is read from `$CONCARDS_PREDICT`, or `$HOME/.config/concards/predict`
if that environment variable doesn't exist.

### The Outcome File
The outcome file contains the historical outcomes of every time a card has been
passed off or failed. It differs only slightly from the predict file. Here is
the corresponding outcome file for the predict file example above:
```
002141b9b9448a257b05da1f2eb78972 2020-08-08T18:00:17Z 2020-08-02T18:00:17Z 0 0 0 0
002141b9b9448a257b05da1f2eb78972 2020-08-08T18:00:17Z 2020-08-02T18:00:17Z 0 1 -1 1
002141b9b9448a257b05da1f2eb78972 2020-08-08T18:00:17Z 2020-08-02T18:00:17Z 1 1 0 1
3dda75cb44ed447186834541475f32e2 2020-08-08T18:00:17Z 2020-08-02T18:00:17Z 0 0 0 1
3dda75cb44ed447186834541475f32e2 2020-08-08T18:00:17Z 2020-08-02T18:00:17Z 1 0 1 0
3dda75cb44ed447186834541475f32e2 2020-08-08T18:00:17Z 2020-08-02T18:00:17Z 1 1 0 0
3dda75cb44ed447186834541475f32e2 2020-08-08T18:00:17Z 2020-08-02T18:00:17Z 1 2 -1 0
```

You can notice that there are two main differences from the predict file:
- There are usually multiple lines with the same hash.
- The last column is a boolean "pass or fail" instead of an algorithm name.

This file is read from `$CONCARDS_OUTCOME`, or `$HOME/.config/concards/outcome`
if that environment variable doesn't exist.
