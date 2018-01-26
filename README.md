# pair

A simple command line tool to generate kerning pairs from arguments

## Usage

### with a string

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

### with a file

If you have a test.txt file with this content

`xyz`

you can do:

`pair test.txt`

it will return:

```
x
xxxxxx yxyyxy zxzzxz
y
xyxxyx yyyyyy zyzzyz
z
xzxxzx yzyyzy zzzzzz
```

## Installing

Once you have golang installed you can get it with:

`go get github.com/typographias/pair`
