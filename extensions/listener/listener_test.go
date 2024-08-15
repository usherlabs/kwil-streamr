package listener

import (
	"testing"

	"github.com/kwilteam/kwil-streamr/extensions/resolution"
	"github.com/stretchr/testify/require"
)

func Test_ParseEvent(t *testing.T) {
	type testcase struct {
		name    string
		params  map[string]string
		obj     map[string]any
		want    []*resolution.ParamValue
		wantErr bool
	}

	tests := []testcase{
		{
			name: "simple",
			params: map[string]string{
				"param1": "key1",
			},
			obj: map[string]any{
				"key1": 1,
			},
			want: []*resolution.ParamValue{
				{
					Param: "param1",
					Value: "1",
				},
			},
		},
		{
			name: "nested",
			params: map[string]string{
				"param1": "key1.key2",
			},
			obj: map[string]any{
				"key1": map[string]any{
					"key2": 2,
				},
			},
			want: []*resolution.ParamValue{
				{
					Param: "param1",
					Value: "2",
				},
			},
		},
		{
			name: "nested array",
			params: map[string]string{
				"param1": "key1.key2",
			},
			obj: map[string]any{
				"key1": map[string]any{
					"key2": []any{3, 2},
				},
			},
			want: []*resolution.ParamValue{
				{
					Param:      "param1",
					ValueArray: []string{"3", "2"},
					IsArray:    true,
				},
			},
		},
		{
			name: "non-existent field",
			params: map[string]string{
				"param1": "key1.key2",
			},
			obj: map[string]any{
				"key1": map[string]any{
					"key3": 3,
				},
			},
			wantErr: true,
		},
		{
			name: "array of objects",
			params: map[string]string{
				"param1": "key1",
			},
			obj: map[string]any{
				"key1": []any{
					map[string]any{
						"key2": 2,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "reference a field that is an object",
			params: map[string]string{
				"param1": "key1",
			},
			obj: map[string]any{
				"key1": map[string]any{
					"key2": 2,
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseEvent(tt.params, tt.obj)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.Nil(t, err)

			require.EqualValues(t, tt.want, got)
		})
	}
}

func TestSetConfig(t *testing.T) {
	type testcase struct {
		name         string
		data         map[string]string
		errorMessage string
		wantErr      bool
	}

	tests := []testcase{
		{
			name: "valid",
			data: map[string]string{
				"node":             "ws://example.com",
				"api_key":          "",
				"max_reconnects":   "5",
				"stream":           "test_stream",
				"target_db":        "0x1A58f48A0369656015D6BE305a3716F84F979A86:dimo_weather",
				"target_procedure": "procedure_name",
				"input_mappings":   "param1:key1,param2:key2.key2.1",
			},
			wantErr: false,
		},
		{
			name: "invalid data_type for max_reconnects",
			data: map[string]string{
				"node":             "ws://example.com",
				"api_key":          "",
				"max_reconnects":   "x",
				"stream":           "test_stream",
				"target_db":        "0x1A58f48A0369656015D6BE305a3716F84F979A86:dimo_weather",
				"target_procedure": "procedure_name",
				"input_mappings":   "param1:key1,param2:key2.key2.1",
			},
			wantErr: true,
			errorMessage: "invalid max_reconnects config: strconv.ParseInt: parsing \"x\": invalid syntax",
		},
		{
			name: "invalid structure for target_db",
			data: map[string]string{
				"node":             "ws://example.com",
				"api_key":          "",
				"max_reconnects":   "5",
				"stream":           "nnn",
				"target_db":        "deployer_address:db_name",
				"target_procedure": "procedure_name",
				"input_mappings":   "param1:key1,param2:key2.key2.1",
			},
			wantErr: true,
			errorMessage: "invalid deployer address in target_db config: encoding/hex: invalid byte: U+0070 'p'",
		},
		{
			name: "invalid structure for target_db",
			data: map[string]string{
				"node":             "ws://example.com",
				"api_key":          "",
				"max_reconnects":   "5",
				"stream":           "nnn",
				"target_db":        "0x1A58f48A0369656015D6BE305a3716F84F979A86:db_name",
				"target_procedure": "procedure_name",
				"input_mappings":   "param1,key11,",
			},
			wantErr: true,
			errorMessage: "invalid input mapping: param1",
		},
		{
			name: "invalid prop",
			data: map[string]string{
				"api_key": "12345",
			},
			errorMessage: "missing required Streamr node URL config",
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := (&listenerConfig{}).setConfig(tt.data)
			if tt.wantErr {
				require.Error(t, err)
				require.Equal(t, tt.errorMessage, err.Error())
				return
			}
			require.Nil(t, err)
		})
	}
}
