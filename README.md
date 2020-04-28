<!-- @> cnc @ concards <@ -->
# <img src="logo.svg" />

Turning notes into flashcards. Concards is my ongoing attempt to make
flashcards simple and quick to record or embed into a text document. Unlike the
overhead of other flashcard applications, Concards accepts a minimal amount of
rules. This allows tons of freedom in the way you want to create your own
flashcards!

## Install
Install like any other go application.
``` bash
go install github.com/alanxoc3/concards
~/$GOPATH/bin/concards --help
```

## Features
* Implements the [SM2](https://www.supermemo.com/english/ol/sm2.htm) Repetition Algorithm.
* Beautiful terminal gui.
* Supports UTF-8!
* Reading in from multiple files.
* Undoing/Redoing
* Easily editing a card while reviewing your cards.
* And More!!!

## Usage
The file syntax was designed to be very simple and flexible, allowing anyone to
quickly embed flashcards into their text document without extra hassle of other
flashcard apps (like [Anki](https://apps.ankiweb.net/) or
[Memrise](https://www.memrise.com/)).

The syntax to embed your flashcards is like this:
```
@> What is the answer?
 @ This is the answer! <@
```

Wanna try it out? Run concards on this README.md file!
``` bash
cnc README.md
```

### Help Output
```
Usage:
  cnc [OPTION]... [FILE|FOLDER]...

Options:
  -r  --review    Show cards available to be reviewed.
  -m  --memorize  Show cards available to be memorized.
  -d  --done      Show cards not available to be reviewed or memorized.
  -n  --number #  Limit the number of cards in the program to "#".
  -e  --edit      Edit the cards, updating used references.
  -p  --print     Prints all cards, slightly formatted.
  -h  --help      If you need assistance.
  -v  --version   Which version are you on again?
      --editor e  Defaults to $EDITOR

For more details, read the fine man page.
```

### A Simple Concard
```
@> This is a question.
 @ Answer #1
 @ Answer #2
 @ Answer #3 <@
```

### 3 Simple Concards
```
@> This is question #1. @ Answer #1
@> This is question #2. @ Answer #1 @ Answer #2
@> This is question #3. @ Answer #1 @ Answer #2 @ Answer #3 <@
```

## All Special Tokens
All the special tokens that are a part of concards syntax are below. Just add
"@" signs to escape them!
```
@>   = Starts a concards block. Starts a question.
<@   = Ends the concards block.
@    = Separates answers.

@@>  = "@>"
<@@  = "<@"
@@   = "@"

@@@> = "@@>"
<@@@ = "<@@"
@@@  = "@@"

...
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
sha256sum                        | review timestamp      | streak | alg | data
---------------------------------+-----------------------+--------+-----+-----
3dda75cb44ed447186834541475f32e2 | 2019-01-01T00:00:00Z  | 0      | sm2 | 2.5
8525b45f883c05eec46b4f7a88e7f7ef | 2020-01-01T00:00:00Z  | 0      | sm2 | 2.5
```

This file is saved in "$CNC_HOME/.cnc-meta". If there is a git repository
available, this file will be committed every time concards updates it.

### The Config File
This file is located in "$CNC_HOME/.cnc-config".

This is a YAML file.
