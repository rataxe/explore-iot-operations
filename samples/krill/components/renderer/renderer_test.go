package renderer

import (
	"testing"

	"github.com/iot-for-all/device-simulation/components/formatter"
	"github.com/iot-for-all/device-simulation/lib/composition"
	"github.com/iot-for-all/device-simulation/lib/environment"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestRenderer(t *testing.T) {
	expected := 1
	expectedArr := []any{expected, expected}
	renderer := New(&composition.MockRenderer{
		OnRender: func(m map[string]any) any {
			require.Equal(t, expected, m[""])
			return expected
		},
	}, &formatter.MockFormatter{
		OnFormat: func(a any) ([]byte, error) {
			require.Equal(t, expectedArr, a)
			return nil, nil
		},
	})

	count := 0

	_, err := renderer.Render(&environment.MockEnvironment{
		OnSet: func(s string, a any) {
			if s == "x" {
				require.Equal(t, count, a)
				count++
			} else {
				require.Equal(t, expected, a)
			}
		}, OnEnv: func() map[string]any {
			return map[string]any{"": expected}
		},
	}, 0, 2)

	require.NoError(t, err)
}