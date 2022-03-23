# femo

## Usage
```
$ cat ./sample.json
[
    {
        "title": "foo",
        "id": "aaa",
        "value": "zzz"
    },
    {
        "title": "bar",
        "id": "iii",
        "value": "xxx"
    },
    {
        "title": "baz",
        "id": "uuu",
        "value": "vvv"
    }
]
```

```
$ femo -src ./sample.json
> foo
> (paste) zzz
```
