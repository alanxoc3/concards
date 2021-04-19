# <img src="logo.svg" />

[![Build Status](https://travis-ci.org/alanxoc3/concards.svg?branch=master)](https://travis-ci.org/alanxoc3/concards)
[![Go Report Card](https://goreportcard.com/badge/github.com/alanxoc3/concards)](https://goreportcard.com/report/github.com/alanxoc3/concards)
[![Coverage Status](https://coveralls.io/repos/github/alanxoc3/concards/badge.svg?branch=master)](https://coveralls.io/github/alanxoc3/concards?branch=master)

Turning notes into flashcards, or should I say concards! This is my ongoing
attempt to make flashcards more simple and convenient. Concards provides much
of the functionality of other mainstream flashcard applications, but with a
unique focus on parsing cards embedded within text files.

## Features
- Spaced repetition similar to [SM2](https://www.supermemo.com/english/ol/sm2.htm)!
- [UTF-8](https://en.wikipedia.org/wiki/UTF-8) as a first-class citizen!
- Configure with your favorite editor!
- Undo & Redo support!
- Read from directories or files!
- Reversible cards!
- [Cloze](https://en.wikipedia.org/wiki/Cloze_test) cards!

## Install
Download the latest release from the [release
page](https://github.com/alanxoc3/concards/releases). At the moment, only Linux
and Mac are supported.

You can also build a snapshot from source with the `go` command.
```bash
$ go install github.com/alanxoc3/concards
```

Once installed, you may want to try running concards on this readme!
```bash
$ concards README.md
```

You may also want to review the help command's output.
```bash
$ concards --help
```

## Basic Syntax
You can learn the full flashcard embedding syntax in just a few minutes! Let's
get started.

### Creating a flashcard
To make a flashcard, you must put the flashcard text within a concards block. A
concards block looks like this `#: :#`, where text would be placed between the
two colons. Ex:
```
#: This is a one sided flashcard. :#
```

The text above will produce a one sided flashcard! But flashcards are normally
2 sided, so let's create a new flashcard that separates a question and answer
with the pipe symbol:
```
#: What is a great way to decrease the effects of the forgetting curve?
 | Spending time every day to review previously learned information. :#
```

Any number of sides are supported, so creating a 3 sided flashcard is a piece
of cake:
```
#: What are Newton's 3 laws of motion?
 | 1. An object at rest stays at rest unless acted upon.
 | 2. Force is equal to mass times acceleration.
 | 3. For every action, there is an equal and opposite reaction. :#
```

You can either create new blocks for each card, or you can keep them in the
same block. This creates 2 cards:
```
#: Who published the first flashcards? | Favell Lee Mortimer
#: When were the first flashcards published? | 1834 :#
```

### Reversible Cards
When learning a language, you might find yourself writing a flashcard that
transitions a phrase from language #1 to language #2 and writing another
flashcard that transitions the same phrase from language #2 to language #1.
Concards makes this easier with the  `::` operator
```
#: saluton al la mundo :: hello world :#

Generates these cards:

#: saluton al la mundo | hello world
#: hello world | saluton al la mundo :#
```

If you are learning two languages, you can expand this with an extra `::`:
```
#: spagetoj :: spaghetti :: 意面 :#

Generates these cards:

#: spagetoj | spaghetti | 意面
#: spaghetti | spagetoj | 意面
#: 意面 | spagetoj | spaghetti :#
```

Translating a word from one language to another often results in multiple
definitions. Concards can represent these scenarios more efficiently when
combining the `|` and `::`.
```
#: apricot | almond :: 杏仁 :#

Generates these cards:

#: apricot | 杏仁
#: almond | 杏仁
#: 杏仁 | apricot | almond :#
```

The double colon operator always takes precedence before the pipe operator.

### Cloze Cards
Cloze cards are handy when working with phrases or related facts. In concards,
a cloze is created by putting text within a set of curly braces. Concards will
generate cards from the text in the curly braces and replace the text with an
empty set of curly braces.
```
#: {Hermann Ebbinghaus} published his findings on the forgetting curve in {1885}. :#

Generates these cards:

#: {} published his findings on the forgetting curve in 1885. | Hermann Ebbinghaus
#: Hermann Ebbinghaus published his findings on the forgetting curve in {}. | 1885 :#
```

Cloze nesting is supported:
```
#: {Education is the {kindling of a flame}}, {not the {filling of a vessel}}. :#

Generates these cards:

#: {}, not the filling of a vessel. | Education is the kindling of a flame
#: Education is the {}, not the filling of a vessel. | kindling of a flame
#: Education is the kindling of a flame, {}. | not the filling of a vessel
#: Education is the kindling of a flame, not the {}. | filling of a vessel :#
```

You can replace consecutive curly braces with the colon operator. This
especially makes separation within a single word look nicer.
```
#: {Pneumono:ultra:microscopic:silico:volcano:coniosis} :#

Is the same as:

#: {Pneumono}{ultra}{microscopic}{silico}{volcano}{coniosis} :#

And generates these cards:

#: {}ultramicroscopicsilicovolcanoconiosis | Pneumono
#: Pneumono{}microscopicsilicovolcanoconiosis | ultra
#: Pneumonoultra{}silicovolcanoconiosis | microscopic
#: Pneumonoultramicroscopic{}volcanoconiosis | silico
#: Pneumonoultramicroscopicsilico{}coniosis | volcano
#: Pneumonoultramicroscopicsilicovolcano{} | coniosis :#
```

To group multiple clozes together, use the hash symbol before a set of curly
braces.
```
#: #{Sebastian Leitner} published about the Leitner System in #{1972}. :#

Generates this card:

#: {} published about the Leitner System in {}. | Sebastian Leitner | 1972 :#
```

Cloze groups are different based on the number of hash symbols before the curly
brace. Here is an example with 3 cloze groups:
```
#: ###{Spaced repetition} is an #{evidence-based} learning technique which
   ##{incorporates} increasing time intervals between each ##{review} of a
   flashcard in order to exploit the ###{psychological} #{spacing effect}. :#

Generates these cards:

#: Spaced repetition is an {} learning technique which incorporates increasing
   time intervals between each review of a flashcard in order to exploit the
   psychological {}.
 | evidence-based
 | spacing effect

#: Spaced repetition is an evidence-based learning technique which {}
   increasing time intervals between each {} of a flashcard in order to exploit
   the psychological spacing effect.
 | incorporates
 | review

#: {} is an evidence-based learning technique which incorporates increasing
   time intervals between each review of a flashcard in order to exploit the {}
   spacing effect.
 | Spaced repetition
 | psychological :#
```

Finally, you can combine the cloze syntax with `::` and `|`:
```
#: {新型:冠状:病毒} :: Coronavirus | COVID-19 :#

Generates these cards:

#: {}冠状病毒 | 新型
#: 新型{}病毒 | 冠状
#: 新型冠状{} | 病毒
#: 新型冠状病毒 | Coronavirus | COVID-19
#: Coronavirus | 新型冠状病毒
#: COVID-19 | 新型冠状病毒 :#
```

### Whitespace & Escaping
Concards ignores consecutive whitespace. The following flashcards are
equivalent:
```
#: {Piotr A. Woźniak} created the SM-2 spaced repetition algorithm in {1987}.
#: { Piotr A. Woźniak } created the SM-2 spaced repetition algorithm in { 1987}.
#:{Piotr A. Woźniak }created the SM-2 spaced repetition algorithm in{ 1987}. :#

Generates these cards:

#: {} created the SM-2 spaced repetition algorithm in 1987. | Piotr A. Woźniak
#: Piotr A. Woźniak created the SM-2 spaced repetition algorithm in {}. | 1987 :#
```

Backslash any reserved character or whitespace to include it in the card text:
```
#: Which characters are special in concards?
 | \# \: \| \{ \}

#: Leave my door open just a crack\
Cause I feel like such an insomniac\
Why do I tire of counting sheep?\
When I'm far too tired to fall asleep
 | Fireflies, by Owl City :#
```

## File Structure
Concards follows the XDG standard for config/data file placement.

### Config Directory
The config directory is calculated by following this order of steps until one
succeeds:
1. `concards --config-dir <directory>`
2. `$CONCARDS_CONFIG_DIR`
3. `$XDG_CONFIG_HOME/concards`
4. `$HOME/.config/concards`
5. `./`

#### Hooks
Hooks are currently an experimental feature. Concards hooks works similar to
git hooks. You must place an executable file with a specific name in
`<config-dir>/hooks/`. Hooks that begin with `event-` are meant to be run in
parallel with concards and perform tasks that don't affect concards directly.
Hooks that begin with `hook-` are similar to plugins in that they are meant to
change program behavior.

Here are all the currently supported hooks:
- `hooks/event-review`  - executed right after passing off a card with a pass or fail
- `hooks/event-startup` - executed once if/when concards starts the GUI up successfully

### Data Directory
The data directory is calculated by following this order of steps until one
succeeds:
1. `concards --data-dir <directory>`
2. `$CONCARDS_DATA_DIR`
3. `$XDG_DATA_HOME/concards`
4. `$HOME/.local/share/concards`
5. `./`

Concards has a very simple file structure. This section explains the content of
the meta data files concards writes to after a review session.

#### The Predict File
The predict file is located at `<data-dir>/predict`. This file contains
information needed to make a prediction when you should review a card next.

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

#### The Outcome File
The outcome file is located at `<data-dir>/outcome`. This file contains the
historical outcomes of every time a card has been passed off or failed. It
differs only slightly from the predict file. Here is the corresponding outcome
file for the predict file example above:
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

## Dependencies
Concards currently depends on these libraries:
- [stretchr/testify](https://github.com/stretchr/testify) for unit tests.
- [spf13/cobra](https://github.com/spf13/cobra) for CLI options.
- [nsf/termbox-go](https://github.com/nsf/termbox-go) for the terminal gui.
- [go-cmd/cmd](https://github.com/go-cmd/cmd) for hook support.
- [mattn/go-runewidth](https://github.com/mattn/go-runewidth) to help with
  Asian characters.

Concards wouldn't be where it is today without those repositories & their
contributors, so please check them out too!

## Future Work
The next big things to focus on:
- Improved CLI UI.
- Local Web Server UI.
- Support Unix Piping.
- Customizable Algorithm Plugin System.
