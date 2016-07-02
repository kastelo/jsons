package jsons_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/calmh/jsons"
)

func TestReader(t *testing.T) {
	cases := [][2]string{
		{``, ``},
		{`{}`, `{}`},
		{`[{}]`, ` {}\n`},
		{`[{}, {}]`, ` {}\n {}\n`},
		{`[{"hey":{"ho":["lets go"]}}, {"foo":{"baz":["quux"]}}]`, ` {"hey":{"ho":["lets go"]}}\n {"foo":{"baz":["quux"]}}\n`},
		{`[["foo", "bar"], ["baz", "quux"]]`, ` ["foo", "bar"]\n ["baz", "quux"]\n`},
		{`[{"hey":{"ho[":["lets go{"]}}, {"foo":{"baz":["quux"]}}]`, ` {"hey":{"ho[":["lets go{"]}}\n {"foo":{"baz":["quux"]}}\n`},
		{`[{"hey":{"ho[":["lets \"go{"]}}, {"foo":{"baz":["quux"]}}]`, ` {"hey":{"ho[":["lets \"go{"]}}\n {"foo":{"baz":["quux"]}}\n`},
	}

	for i, tc := range cases {
		in := bytes.NewBufferString(tc[0])
		s := jsons.New(in)
		out, err := ioutil.ReadAll(s)
		if err != nil {
			t.Errorf("case %d: %v", i, err)
			continue
		}
		tc[1] = strings.Replace(tc[1], `\n`, "\n", -1)
		if string(out) != tc[1] {
			t.Errorf("mismatch: %q != expected %q", out, tc[1])
		}
	}
}

func ExampleReader() {
	type Object struct {
		Foo string
	}

	input := bytes.NewBufferString(`[{"foo":"bar"}, {"foo":"baz"}]`)
	streamer := jsons.New(input)
	dec := json.NewDecoder(streamer)
	for {
		var res Object
		err := dec.Decode(&res)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Read: %+v\n", res)
	}

	// Output:
	// Read: {Foo:bar}
	// Read: {Foo:baz}
}
