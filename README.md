jsons
=====

Package jsons converts JSON arrays into streams of objects.

This is useful if you have a file, HTTP response or similar containing a
large array of JSON objects. Instead of reading it all into memory and
deserializing into a `[]Whatever` you can wrap the reader in a
`jsons.Reader` and then use a `json.Decoder` to repeatedly read and
deserialize a single object into a `Whatever`.

Documentation
-------------

https://godoc.org/github.com/calmh/jsons

License
-------

MIT
