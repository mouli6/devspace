package remove

/* @Florian adjust to new behaviour
import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/devspace-cloud/devspace/pkg/devspace/config/configutil"
	"github.com/devspace-cloud/devspace/pkg/devspace/config/constants"
	"github.com/devspace-cloud/devspace/pkg/devspace/config/versions/latest"
	"github.com/devspace-cloud/devspace/pkg/util/log"
	"github.com/devspace-cloud/devspace/pkg/util/survey"

	"gotest.tools/assert"
)

type removeSyncTestCase struct {
	name string

	fakeConfig *latest.Config

	args          []string
	answers       []string
	labelSelector string
	localPath     string
	containerPath string
	removeAll     bool

	expectedOutput   string
	expectedPanic    string
	expectConfigFile bool
}

func TestRunRemoveSync(t *testing.T) {
	testCases := []removeSyncTestCase{
		removeSyncTestCase{
			name:          "No devspace config",
			expectedPanic: "Couldn't find a DevSpace configuration. Please run `devspace init`",
		},
		removeSyncTestCase{
			name:          "Specify nothing",
			fakeConfig:    &latest.Config{},
			expectedPanic: "You have to specify at least one of the supported flags",
		},
		removeSyncTestCase{
			name:       "Remove all zero syncs",
			fakeConfig: &latest.Config{},
			removeAll:  true,
		},
	}

	log.SetInstance(&testLogger{
		log.DiscardLogger{PanicOnExit: true},
	})

	for _, testCase := range testCases {
		testRunRemoveSync(t, testCase)
	}
}

func testRunRemoveSync(t *testing.T, testCase removeSyncTestCase) {
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

	for _, answer := range testCase.answers {
		survey.SetNextAnswer(answer)
	}

	isSyncsNil := testCase.fakeConfig == nil || testCase.fakeConfig.Images == nil
	configutil.SetFakeConfig(testCase.fakeConfig)
	if isSyncsNil && testCase.fakeConfig != nil {
		testCase.fakeConfig.Images = nil
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

		rec := recover()
		if testCase.expectedPanic == "" {
			if rec != nil {
				t.Fatalf("Unexpected panic in testCase %s. Message: %s", testCase.name, rec)
			}
		} else {
			if rec == nil {
				t.Fatalf("Unexpected no panic in testCase %s", testCase.name)
			} else {
				assert.Equal(t, rec, testCase.expectedPanic, "Wrong panic message in testCase %s", testCase.name)
			}
		}
		assert.Equal(t, logOutput, testCase.expectedOutput, "Unexpected output in testCase %s", testCase.name)
	}()

	(&syncCmd{
		LabelSelector: testCase.labelSelector,
		LocalPath:     testCase.localPath,
		ContainerPath: testCase.containerPath,
		RemoveAll:     testCase.removeAll,
	}).RunRemoveSync(nil, testCase.args)

	assert.Equal(t, logOutput, testCase.expectedOutput, "Unexpected output in testCase %s", testCase.name)

	err = os.Remove(constants.DefaultConfigPath)
	assert.Equal(t, !os.IsNotExist(err), testCase.expectConfigFile, "Unexpectedly saved or not saved in testCase %s", testCase.name)
}
*/
