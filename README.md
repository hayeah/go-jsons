A cute tool to expand json using shell's brace expansion feature.

It might be useful for generating cluster configuration.

# JSON Expander

Suppose you want a stream of json objects that are the same, but differ only by an index number:

```
{
  "a": "1",
}
{
  "a": "2",
}
{
  "a": "3",
}
```

You can start with a [template](http://golang.org/pkg/text/template) like:

```
{ "a": "{{.}}"}
```

Then pipe this template into the `jsons` expander:

```
$ echo '{ "a": "{{.}}"}' | jsons {1..3}
{"a":"1"}
{"a":"2"}
{"a":"3"}
```

Shell expansion also works with characters:

```
$ echo '{ "a": "{{.}}"}' | jsons {a..c}{1..3}
{"a":"a1"}
{"a":"a2"}
{"a":"a3"}
{"a":"b1"}
{"a":"b2"}
{"a":"b3"}
{"a":"c1"}
{"a":"c2"}
{"a":"c3"}
```

# Array Expander

If the template is an item in an array, then the expander produces a JSON array:

```
$ echo '[{ "a": "{{.}}"}]' | jsons {1,3,5}
[{"a":"1"},
{"a":"3"},
{"a":"5"}
]
```