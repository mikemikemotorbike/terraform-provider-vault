package codegen

import (
	"github.com/hashicorp/vault/sdk/framework"
	"testing"
)

func TestLastField(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "/transform/alphabet",
			expected: "alphabet",
		},
		{
			input:    "/transform/alphabet/{name}",
			expected: "{name}",
		},
		{
			input:    "/transform/decode/{role_name}",
			expected: "{role_name}",
		},
		{
			input:    "/transit/datakey/{plaintext}/{name}",
			expected: "{name}",
		},
		{
			input:    "/transit/export/{type}/{name}/{version}",
			expected: "{version}",
		},
		{
			input:    "/unlikely",
			expected: "unlikely",
		},
	}
	for _, testCase := range testCases {
		actual := lastField(testCase.input)
		if actual != testCase.expected {
			t.Fatalf("input: %q; expected: %q; actual: %q", testCase.input, testCase.expected, actual)
		}
	}
}

func TestClean(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "alphabet",
			expected: "alphabet",
		},
		{
			input:    "{name}",
			expected: "name",
		},
		{
			input:    "{role_name}",
			expected: "rolename",
		},
		{
			input:    "{name}",
			expected: "name",
		},
		{
			input:    "{version}",
			expected: "version",
		},
		{
			input:    "unlikely",
			expected: "unlikely",
		},
	}
	for _, testCase := range testCases {
		actual := clean(testCase.input)
		if actual != testCase.expected {
			t.Fatalf("input: %q; expected: %q; actual: %q", testCase.input, testCase.expected, actual)
		}
	}
}

func TestValidate(t *testing.T) {
	testCases := []struct {
		input       *templatableEndpoint
		expectedErr string
	}{
		{
			input:       nil,
			expectedErr: "endpoint is nil",
		},
		{
			input:       &templatableEndpoint{},
			expectedErr: "endpoint cannot be blank for &{Endpoint: DirName: ExportedFuncPrefix: PrivateFuncPrefix: Parameters:[] SupportsRead:false SupportsWrite:false SupportsDelete:false}",
		},
		{
			input: &templatableEndpoint{
				Endpoint: "foo",
			},
			expectedErr: "dirname cannot be blank for &{Endpoint:foo DirName: ExportedFuncPrefix: PrivateFuncPrefix: Parameters:[] SupportsRead:false SupportsWrite:false SupportsDelete:false}",
		},
		{
			input: &templatableEndpoint{
				Endpoint: "foo",
				DirName:  "foo",
			},
			expectedErr: "exported function prefix cannot be blank for &{Endpoint:foo DirName:foo ExportedFuncPrefix: PrivateFuncPrefix: Parameters:[] SupportsRead:false SupportsWrite:false SupportsDelete:false}",
		},
		{
			input: &templatableEndpoint{
				Endpoint:           "foo",
				DirName:            "foo",
				ExportedFuncPrefix: "foo",
			},
			expectedErr: "private function prefix cannot be blank for &{Endpoint:foo DirName:foo ExportedFuncPrefix:foo PrivateFuncPrefix: Parameters:[] SupportsRead:false SupportsWrite:false SupportsDelete:false}",
		},
		{
			input: &templatableEndpoint{
				Endpoint:           "foo",
				DirName:            "foo",
				ExportedFuncPrefix: "foo",
				PrivateFuncPrefix:  "foo",
			},
			expectedErr: "",
		},
		{
			input: &templatableEndpoint{
				Endpoint:           "foo",
				DirName:            "foo",
				ExportedFuncPrefix: "foo",
				PrivateFuncPrefix:  "foo",
				Parameters: []*templatableParam{
					{
						OASParameter: &framework.OASParameter{
							Name: "some-param",
							Schema: &framework.OASSchema{
								Type: "foo",
							},
						},
					},
				},
			},
			expectedErr: "unsupported type of foo for some-param",
		},
		{
			input: &templatableEndpoint{
				Endpoint:           "foo",
				DirName:            "foo",
				ExportedFuncPrefix: "foo",
				PrivateFuncPrefix:  "foo",
				Parameters: []*templatableParam{
					{
						OASParameter: &framework.OASParameter{
							Name: "some-param",
							Schema: &framework.OASSchema{
								Type: "string",
							},
						},
					},
				},
			},
			expectedErr: "",
		},
		{
			input: &templatableEndpoint{
				Endpoint:           "foo",
				DirName:            "foo",
				ExportedFuncPrefix: "foo",
				PrivateFuncPrefix:  "foo",
				Parameters: []*templatableParam{
					{
						OASParameter: &framework.OASParameter{
							Name: "foo",
							Schema: &framework.OASSchema{
								Type: "array",
								Items: &framework.OASSchema{
									Type: "string",
								},
							},
						},
					},
				},
			},
			expectedErr: "",
		},
		{
			input: &templatableEndpoint{
				Endpoint:           "foo",
				DirName:            "foo",
				ExportedFuncPrefix: "foo",
				PrivateFuncPrefix:  "foo",
				Parameters: []*templatableParam{
					{
						OASParameter: &framework.OASParameter{
							Name: "foo",
							Schema: &framework.OASSchema{
								Type: "array",
								Items: &framework.OASSchema{
									Type: "object",
								},
							},
						},
					},
				},
			},
			expectedErr: "unsupported array type of object for foo",
		},
	}
	for _, testCase := range testCases {
		shouldErr := testCase.expectedErr != ""
		err := testCase.input.Validate()
		if err != nil {
			if err.Error() != testCase.expectedErr {
				t.Fatalf("input: %+v; expected err: %q; actual: %q", testCase.input, testCase.expectedErr, err)
			}
		} else {
			if shouldErr {
				t.Fatalf("expected an error for %+v", testCase.input)
			}
		}
	}
}

func TestToTemplatableParam(t *testing.T) {
	// TODO
}

func TestCollectParameters(t *testing.T) {
	// TODO
}

func TestToTemplatable(t *testing.T) {
	// TODO
}

func TestWrite(t *testing.T) {
	// TODO
}

func TestNewTemplateHandler(t *testing.T) {
	// TODO
}
