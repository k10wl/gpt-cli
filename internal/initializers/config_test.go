package initializers

import (
	"path"
	"testing"
)

func TestLoad(t *testing.T) {
	mockDir := path.Join("..", "..", "testdata")
	mockedConfig, err := readConfigFile(&Flags{Path: mockDir})
	if err != nil {
		t.Error("failed to load mocked config")
	}

	testCases := []struct {
		name      string
		flags     Flags
		expected  Settings
		wantError bool
	}{
		{
			name: "return model and assistant",
			flags: Flags{
				Model:     "second",
				Assistant: "first",
				Path:      mockDir,
			},
			expected: Settings{
				Model:     "second",
				Assistant: "first assistant message",
			},
			wantError: false,
		},
		{
			name: "return first model and empty assistant message in they where not specified by flag",
			flags: Flags{
				Model:     "",
				Assistant: "",
				Path:      mockDir,
			},
			expected: Settings{
				Model:     "first",
				Assistant: "",
			},
			wantError: false,
		},
		{
			name: "error if model does not exist",
			flags: Flags{
				Model:     "non existing",
				Assistant: "",
				Path:      mockDir,
			},
			expected:  Settings{},
			wantError: true,
		},
		{
			name: "error if assistant does not exist",
			flags: Flags{
				Model:     "first",
				Assistant: "non existing",
				Path:      mockDir,
			},
			expected:  Settings{},
			wantError: true,
		},
	}

	for _, tc := range testCases {
		setting, err := loadSettings(mockedConfig, &tc.flags)

		if tc.wantError != (err != nil) {
			t.Errorf("expected: "+tc.name+"\nactual: %+v", err)

			return
		}

		if *setting != tc.expected {
			t.Errorf("expected: "+tc.name+"\nactual: %+v", *setting)
		}
	}
}
