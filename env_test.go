package env

import "testing"
import "os"

func setEnvT(t *testing.T, key string, val string) {
	err := os.Setenv(key, val)
	if err != nil {
		t.Errorf("failed to set env for testcase")
	}
}

// Test to ensure that when no env is foun
func TestGetEnvOrDefaultNoEnv(t *testing.T) {
	testKey := "test"
	testVal := "this is a test default"
	env := GetEnvOrDefault(testKey, testVal)
	if env != testVal {
		t.Errorf("Default value was not returned when it was suppose to")
	}
}

func TestGetEnvOrDefaultEnv(t *testing.T) {
	testKey := "test"
	testVal := "this is a test env"
	setEnvT(t, testKey, testVal)
	env := GetEnvOrDefault(testKey, "")
	if env != testVal {
		t.Errorf("Env value was not returned when it was suppose to")
	}
}
