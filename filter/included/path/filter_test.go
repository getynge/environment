package path

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestF_Filter(t *testing.T) {
	f := f{}

	type test struct {
		Name   string
		Env    string
		Input  string
		Expect map[string]bool
	}

	common := test{
		Name:  "basicAppend",
		Env:   "/does/not/exist",
		Input: "/big/fish/small/potatoes",
		Expect: map[string]bool{
			"/big/fish/small/potatoes": true,
			"/does/not/exist":          true,
		},
	}

	tests := []test{
		common,
		{
			Name:   "suffixRemoval",
			Env:    common.Env,
			Input:  "/big/fish/small/potatoes:$PATH",
			Expect: common.Expect,
		},
		{
			Name:   "prefixRemoval",
			Env:    common.Env,
			Input:  "$PATH:/big/fish/small/potatoes",
			Expect: common.Expect,
		},
		{
			Name:  "middleRemoval",
			Env:   common.Env,
			Input: "/big/fish/small/potatoes:$PATH:/friendly/spider",
			Expect: map[string]bool{
				"/big/fish/small/potatoes": true,
				"/friendly/spider":         true,
				"/does/not/exist":          true,
			},
		},
		{
			Name:  "emptyEnv",
			Env:   "",
			Input: "/big/fish/small/potatoes",
			Expect: map[string]bool{
				"/big/fish/small/potatoes": true,
			},
		},
		{
			Name:  "emptyInput",
			Env:   common.Env,
			Input: "",
			Expect: map[string]bool{
				"/does/not/exist": true,
			},
		},
	}

	t.Run("nonPath", func(t *testing.T) {
		key, value := "VARIABLE", "VALUE"
		keyOut, valueOut, err := f.Filter(key, value)
		if err != nil {
			t.Fatalf("error filtering non-path: %s", err.Error())
		}

		if keyOut != key {
			t.Fatalf("test failed, expected (key) %s to equal (keyout) %s", key, keyOut)
		}

		if valueOut != value {
			t.Fatalf("test failed, expected (value) %s to equal (valueOut) %s", value, valueOut)
		}
	})

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if err := os.Setenv("PATH", tt.Env); err != nil {
				t.Fatalf("error setting path: %s", err.Error())
			}
			_, result, err := f.Filter("PATH", tt.Input)
			if err != nil {
				t.Fatalf("error filtering path: %s", err.Error())
			}
			outSet := setFromArray(strings.Split(result, ":"))

			if !reflect.DeepEqual(tt.Expect, outSet) {
				t.Logf("error expected sets to be equal")
				t.Logf("inSet:")
				for k := range tt.Expect {
					t.Logf("\t\t%s", k)
				}
				t.Logf("outSet:")
				for k := range outSet {
					t.Logf("\t\t%s", k)
				}
				t.FailNow()
			}
		})
	}
}

func setFromArray(arr []string) (set map[string]bool) {
	set = make(map[string]bool)
	for _, v := range arr {
		set[v] = true
	}
	return set
}
