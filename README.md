# pair

A simple command line tool to generate kerning pairs from arguments

## Usage

`pair abcde`

returns:

```
a
aaaaaa babbab caccac daddad eaeeae
b
abaaba bbbbbb cbccbc dbddbd ebeebe
c
acaaca bcbbcb cccccc dcddcd eceece
d
adaada bdbbdb cdccdc dddddd edeede
e
aeaaea bebbeb ceccec dedded eeeeee

```

You can pipe it to the clipboard:

`pair abcde | pbcopy`

And then paste it in your font or text editor.

## Installing

Once you have golang installed you can get it with:

`go get github.com/typographias/pair`
