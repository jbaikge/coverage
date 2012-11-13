package testfeed

import (
	"git.300brand.com/coverage/parser"
	_ "git.300brand.com/coverage/parser/atom"
	_ "git.300brand.com/coverage/parser/rdf"
	_ "git.300brand.com/coverage/parser/rss"
	"testing"
)

func TestAtomParse(t *testing.T) {
	f, err := parser.Parse(Atom, "atom")
	if err != nil {
		t.Error(err)
		return
	}
	if len(f.Articles) != 50 {
		t.Errorf("Invalid number of articles: %d", len(f.Articles))
	}
}

func TestAtomType(t *testing.T) {
	if typ, err := parser.Type(Atom); err != nil || typ != "atom" {
		t.Errorf("Expected atom, got %s", typ)
		t.Error(err)
	}
}

func TestRDFParse(t *testing.T) {
	f, err := parser.Parse(RDF, "rdf")
	if err != nil {
		t.Error(err)
		return
	}
	if len(f.Articles) != 90 {
		t.Errorf("Invalid number of articles: %d", len(f.Articles))
	}
}

func TestRDFType(t *testing.T) {
	if typ, err := parser.Type(RDF); err != nil || typ != "rdf" {
		t.Errorf("Expected rdf, got %s", typ)
		t.Error(err)
	}
}

func TestRSSParse(t *testing.T) {
	f, err := parser.Parse(RSS, "rss")
	if err != nil {
		t.Error(err)
		return
	}
	if len(f.Articles) != 10 {
		t.Errorf("Invalid number of articles: %d", len(f.Articles))
	}
}

func TestRSSType(t *testing.T) {
	if typ, err := parser.Type(RSS); err != nil || typ != "rss" {
		t.Errorf("Expected rss, got %s", typ)
		t.Error(err)
	}
}

func TestInvalidParse(t *testing.T) {
	if _, err := parser.Parse(Atom, "rss"); err == nil {
		t.Error("Expected error when using RSS decoder to parse Atom feed")
	}
}

func TestInvalidType(t *testing.T) {
	data := []byte(`<?xml version="1.0"?><bunk><feed /></bunk>`)
	if typ, err := parser.Type(data); err == nil || typ != "" {
		t.Errorf("Expected blank type and error, got %s", typ)
		t.Error(err)
	}
}
