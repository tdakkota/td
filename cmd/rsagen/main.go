// The rsagen command generates rsa.PublicKey variables from PEM-encoded
// RSA public keys.
package main

import (
	"bytes"
	"crypto/rsa"
	"embed"
	"flag"
	"fmt"
	"go/format"
	"io"
	"io/fs"
	"log"
	"os"
	"text/template"

	"github.com/gotd/td/internal/crypto"
)

//go:embed _template/*.tmpl
var embedFS embed.FS

var funcs = template.FuncMap{
	"fingerprint": func(pubkey *rsa.PublicKey) string {
		return fmt.Sprintf("%08x", uint64(crypto.RSAFingerprint(pubkey)))
	},
	"chunks": func(b []byte, size int) [][]byte {
		var chunks [][]byte
		for len(b) > size {
			chunks = append(chunks, b[:size])
			b = b[size:]
		}
		if len(b) != 0 {
			chunks = append(chunks, b)
		}
		return chunks
	},
	"single": func(keys []*rsa.PublicKey) (*rsa.PublicKey, error) {
		if count := len(keys); count != 1 {
			return nil, fmt.Errorf("expected single key, got %d keys", count)
		}
		return keys[0], nil
	},
}

func main() {
	os.Exit(main1())
}

func main1() int {
	log.SetFlags(log.Llongfile)
	var (
		inPath       = ""
		outPath      = ""
		tplPath      = ""
		tplName      = "main.tmpl"
		pkgName      = "main"
		varName      = "PK"
		singleMode   = false
		formatOutput = true
	)
	flag.StringVar(
		&inPath, "f", inPath,
		"input path (defaults to stdin)",
	)
	flag.StringVar(
		&outPath, "o", outPath,
		"output path (defaults to stdout)",
	)
	flag.StringVar(
		&tplPath, "templates", tplPath,
		"templates directory (defaults to embed)",
	)
	flag.StringVar(
		&tplName, "exec", tplName,
		"template name",
	)
	flag.StringVar(
		&pkgName, "pkg", pkgName,
		"package name",
	)
	flag.StringVar(
		&varName, "var", varName,
		"variable name",
	)
	flag.BoolVar(
		&singleMode, "single", singleMode,
		"emit single key instead of slice",
	)
	flag.BoolVar(
		&formatOutput, "format", formatOutput,
		"run gofmt on output",
	)
	flag.Parse()

	var err error

	var in []byte
	if inPath != "" {
		in, err = os.ReadFile(inPath)
	} else {
		in, err = io.ReadAll(os.Stdin)
	}
	if err != nil {
		log.Printf("read input: %v", err)
		return 1
	}

	keys, err := crypto.ParseRSAPublicKeys(in)
	if err != nil {
		log.Printf("parse public keys: %v", err)
		return 1
	}

	fsys, _ := fs.Sub(embedFS, "_template")
	if tplPath != "" {
		fsys = os.DirFS(tplPath)
	}

	tpl, err := template.New("").Funcs(funcs).ParseFS(fsys, "*.tmpl")
	if err != nil {
		log.Printf("parse templates: %v", err)
		return 1
	}

	buf := bytes.NewBuffer(nil)
	if err := tpl.ExecuteTemplate(buf, tplName, map[string]interface{}{
		"Package":  pkgName,
		"Variable": varName,
		"Keys":     keys,
		"Single":   singleMode,
	}); err != nil {
		log.Printf("execute template: %v", err)
		return 1
	}

	if formatOutput {
		p, err := format.Source(buf.Bytes())
		if err != nil {
			log.Printf("format output: %v", err)
			return 1
		}
		buf = bytes.NewBuffer(p)
	}

	if outPath != "" {
		err = os.WriteFile(outPath, buf.Bytes(), 0o744)
	} else {
		_, err = buf.WriteTo(os.Stdout)
	}
	if err != nil {
		log.Printf("write output: %v", err)
		return 1
	}
	return 0
}
