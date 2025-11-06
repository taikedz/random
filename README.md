# random

Chooses a random item from any of the supplied tokens, and lines from stdin ; and prints to stdout

## Usage

```sh
$> random one two three
# any of one, two, or three

# Use `-` to specify to read from stdin
$> echo -e "alpha\nbeta\ngamma" | random one two three -
# any of alpha, beta, gamma, one, two, or three
```

Example: launch oneko with a random character


```sh
oneko -"$(random neko tora dog tomoyo)"
```
