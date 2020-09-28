<!-- @> Who is the coolest person in the world? | You are :D. | Have a great day! <@ -->
# <img src="logo.svg" />

[![Build Status](https://travis-ci.org/alanxoc3/concards.svg?branch=master)](https://travis-ci.org/alanxoc3/concards)
[![Go Report Card](https://goreportcard.com/badge/github.com/alanxoc3/concards)](https://goreportcard.com/report/github.com/alanxoc3/concards)
[![Coverage Status](https://coveralls.io/repos/github/alanxoc3/concards/badge.svg?branch=master)](https://coveralls.io/github/alanxoc3/concards?branch=master)

Turning notes into flashcards, or should I say concards! Concards is my ongoing
attempt to make flashcards simple and easily embeddable into text document
based notes. Concards is much lighter than other flashcard applications such as
such as [Anki](https://apps.ankiweb.net/) or
[Memrise](https://www.memrise.com/), but is also very powerful by following the
[Unix Philosophy](https://en.wikipedia.org/wiki/Unix_philosophy) of "Do one
thing and do it well".

## Features
- Implements the [SM2](https://www.supermemo.com/english/ol/sm2.htm) Repetition Algorithm.
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
It should be super simple. Just use the `go install` command:
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
their contributors, so please check them out too :).

## Usage
The complete syntax of embedding your flashcards into text documents consists
of these keywords:
```
'@>' = Starts a concards block and also starts a question.
'|'  = Separates sides.
':'  = Separates sides, and adds a reversed card.
'<@' = Ends the concards block.
'\'  = Escapes the special tokens above.
```

Here are a few example concards:
```
@> Concards
 : A lightweight embeddable note-taking flashcard program.

@> What does "concards" stand for?
 | Console Cards

@> What does the ":" do in concards?
 | It will add an extra card where the "colon" side is the question and the
   question is the answer.
 | This syntax is especially useful for vocabulary when learning a language and
   can save typing.

@> Can a concard have more than 2 sides?
 | Yes

@> What does a concard look like?
 | \@> It could look like this \| What does a concard look like? \<@
 | \@> It could also look like this \| with multiple \| answers! \<@

@> How do you escape a concard keyword?
 | Put a backslash before it. Your text file would show "\@>", but the app shows "@>".

@> How do you show a backslash then a keyword in the concards ui?
 | To see \\@> in the ui, your text document must have 2 backslashes: \\\@>

@> 你好世界
 : Hello World
 : Greetings World

@> The human who created concards.
 : Alan Morgan
<@
```

The easiest way to understand that syntax is by trying it out! Just run
concards on this `README.md` file and see what happens!
``` bash
$ concards README.md
```

## Advanced Usage
### The Meta Data File
Here is an example meta-data file:
```
3dda75cb44ed447186834541475f32e2 2019-01-01T00:00:00Z 0 sm2 2.5
8525b45f883c05eec46b4f7a88e7f7ef 2020-01-01T00:00:00Z 0 sm2 2.5
```

Here is the same file, but annotated:
```
sha256sum cut in half            | review timestamp      | streak | alg | data
---------------------------------+-----------------------+--------+-----+-----
3dda75cb44ed447186834541475f32e2 | 2019-01-01T00:00:00Z  | 0      | sm2 | 2.5
8525b45f883c05eec46b4f7a88e7f7ef | 2020-01-01T00:00:00Z  | 0      | sm2 | 2.5
```

This file is saved to `$CONCARDS_META`. If that environment variable doesn't
exist, then it is saved to `$HOME/.concards-meta`.
