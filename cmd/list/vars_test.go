package list

/* @Florian adjust to new behaviour
import (
	"io/ioutil"
	"os"
	"runtime/debug"
	"testing"

	"github.com/devspace-cloud/devspace/pkg/devspace/config/configutil"
	"github.com/devspace-cloud/devspace/pkg/devspace/config/generated"
	"github.com/devspace-cloud/devspace/pkg/devspace/config/versions/latest"
	"github.com/devspace-cloud/devspace/pkg/util/fsutil"
	"github.com/devspace-cloud/devspace/pkg/util/log"
	"github.com/mgutz/ansi"
	"gopkg.in/yaml.v2"

	"gotest.tools/assert"
)

type listVarsTestCase struct {
	name string

	fakeConfig           *latest.Config
	generatedYamlContent interface{}

	expectedOutput string
	expectedPanic  string
}

func TestListVars(t *testing.T) {
	expectedHeader := ansi.Color(" Variable  ", "green+b") + ansi.Color(" Value  ", "green+b")
	testCases := []listVarsTestCase{
		listVarsTestCase{
			name:          "no config exists",
			expectedPanic: "Couldn't find a DevSpace configuration. Please run `devspace init`",
		},
		listVarsTestCase{
			name:                 "generated.yaml not parsable",
			fakeConfig:           &latest.Config{},
			generatedYamlContent: "unparsable",
			expectedPanic:        "yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `unparsable` into generated.Config",
		},
		listVarsTestCase{
			name:           "no vars",
			fakeConfig:     &latest.Config{},
			expectedOutput: "\nInfo No variables found",
		},
		listVarsTestCase{
			name:       "one var",
			fakeConfig: &latest.Config{},
			generatedYamlContent: generated.Config{
				ActiveProfile: "myConf",
				Profiles: map[string]*generated.CacheConfig{
					"myConf": &generated.CacheConfig{},
				},
				Vars: map[string]string{
					"hello": "world",
				},
			},
			expectedOutput: "\n" + expectedHeader + "\n hello      world  \n\n",
		},
	}

	log.SetInstance(&testLogger{
		log.DiscardLogger{PanicOnExit: true},
	})

	for _, testCase := range testCases {
		testListVars(t, testCase)
	}
}

func testListVars(t *testing.T, testCase listVarsTestCase) {
	logOutput = ""

	dir, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Fatalf("Error creating temporary directory: %v", err)
	}

	wdBackup, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting current working directory: %v", err)
	}
	err = os.Chdir(dir)
	if err != nil {
		t.Fatalf("Error changing working directory: %v", err)
	}

	defer func() {
		//Delete temp folder
		err = os.Chdir(wdBackup)
		if err != nil {
			t.Fatalf("Error changing dir back: %v", err)
		}
		err = os.RemoveAll(dir)
		if err != nil {
			t.Fatalf("Error removing dir: %v", err)
		}
	}()

	configutil.SetFakeConfig(testCase.fakeConfig)
	generated.ResetConfig()

	if testCase.generatedYamlContent != nil {
		content, err := yaml.Marshal(testCase.generatedYamlContent)
		assert.NilError(t, err, "Error parsing configs.yaml to yaml in testCase %s", testCase.name)
		fsutil.WriteToFile(content, generated.ConfigPath)
	}

	defer func() {
		rec := recover()
		if testCase.expectedPanic == "" {
			if rec != nil {
				t.Fatalf("Unexpected panic in testCase %s. Message: %s. Stack: %s", testCase.name, rec, string(debug.Stack()))
			}
		} else {
			if rec == nil {
				t.Fatalf("Unexpected no panic in testCase %s", testCase.name)
			} else {
				assert.Equal(t, rec, testCase.expectedPanic, "Wrong panic message in testCase %s. Stack: %s", testCase.name, string(debug.Stack()))
			}
		}
		assert.Equal(t, logOutput, testCase.expectedOutput, "Unexpected output in testCase %s", testCase.name)
	}()

	(&varsCmd{}).RunListVars(nil, []string{})
}
*/
