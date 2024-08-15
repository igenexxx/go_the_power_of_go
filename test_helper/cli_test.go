package test_helper

import (
	"github.com/rogpeppe/go-internal/testscript"
	"os"
	"path/filepath"
	"testing"
)

type tester struct {
	testdataPath  string
	mainScriptMap map[string]func() int
}

type option func(*tester) error

func New(opts ...option) (*tester, error) {
	c := &tester{}

	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func WithTestDataPath(path string) option {
	if path == "" {
		path = filepath.Join("testdata", "scripts")
	}

	return func(t *tester) error {
		t.testdataPath = path
		return nil
	}
}

func WithScriptMap(mainScriptMap map[string]func() int) option {
	return func(t *tester) error {
		t.mainScriptMap = mainScriptMap
		return nil
	}
}

func (u *tester) RunTestScript(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: u.testdataPath,
	})
}

func (u *tester) TestMain(m *testing.M) {
	os.Exit(
		testscript.RunMain(
			m,
			u.mainScriptMap,
		),
	)
}
