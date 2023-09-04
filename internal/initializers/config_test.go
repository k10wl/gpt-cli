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
			name: "return model and system",
			flags: Flags{
				Model:  "second",
				System: "first",
				Path:   mockDir,
			},
			expected: Settings{
				Model:  "second",
				System: "first system message",
			},
			wantError: false,
		},
		{
			name: "return first model and empty system message in they where not specified by flag",
			flags: Flags{
				Model:  "",
				System: "",
				Path:   mockDir,
			},
			expected: Settings{
				Model:  "first",
				System: "",
			},
			wantError: false,
		},
		{
			name: "error if model does not exist",
			flags: Flags{
				Model:  "non existing",
				System: "",
				Path:   mockDir,
			},
			expected:  Settings{},
			wantError: true,
		},
		{
			name: "error if system does not exist",
			flags: Flags{
				Model:  "first",
				System: "non existing",
				Path:   mockDir,
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
