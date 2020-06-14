package envsugar

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func keyNormalize(prefix, key string) string {
	prefix = strings.TrimSpace(prefix)

	key = strings.TrimSpace(key)

	if prefix != "" {
		key = fmt.Sprintf("%s_%s", prefix, key)
	}

	return strings.ToUpper(key)
}

// Check verifies if an environment variable is available and valued
func Check(prefix, key string, defaultValue string, required, logPrint bool) error {
	key = keyNormalize(prefix, key)

	if logPrint {
		log.Printf("checking env.var %s ...", key)
	}

	if os.Getenv(key) == `` {
		if defaultValue != `` {
			if err := os.Setenv(key, defaultValue); err != nil {
				log.Println(err.Error())
				return err
			}

			if logPrint {
				log.Println(`set with default value`)
			}

			return nil
		}

		if os.Getenv(key) == `` && required {
			if logPrint {
				fmt.Println("required but not set")
			}

			return fmt.Errorf("required but not set")
		}
	}

	if logPrint {
		fmt.Println("ok!")
	}

	return nil
}

// Directives describes rules for variable evaluation
type Directives struct {
	Name         string
	Required     bool
	DefaultValue string
}

// Check verifies if tow or more environment variables are available and valued
func CheckMany(prefix string, directives []Directives, logPrint bool) error {
	for _, d := range directives {
		if err := Check(prefix, d.Name, d.DefaultValue, d.Required, logPrint); err != nil {
			return err
		}
	}

	return nil
}

// Str returns the env var value as string
func String(prefix, key, defaultValue string) string {
	key = keyNormalize(prefix, key)

	if os.Getenv(key) != `` {
		return os.Getenv(key)
	}

	return defaultValue
}

// StrS returns the env var value as []string
func StrS(prefix, key, separator string, defaultValue []string) []string {
	key = keyNormalize(prefix, key)

	if os.Getenv(key) != `` {
		return strings.Split(os.Getenv(key), separator)
	}

	if len(defaultValue) > 0 {
		return defaultValue
	}

	return []string{}
}

// Int returns the env var value as int
func Int(prefix, key string, defaultValue int) int {
	key = keyNormalize(prefix, key)

	if os.Getenv(key) != `` {
		if i, err := strconv.Atoi(os.Getenv(key)); err == nil {
			return i
		}
	}

	return defaultValue
}

// Int64 returns the env var value as int64
func Int64(prefix, key string, defaultValue int64) int64 {
	key = keyNormalize(prefix, key)

	if os.Getenv(key) != `` {
		if i, err := strconv.ParseInt(os.Getenv(key), 10, 64); err == nil {
			return i
		}
	}

	return defaultValue
}

// IntS returns the env var value as []int
func IntS(prefix, key, separator string, defaultValue []int) []int {
	key = keyNormalize(prefix, key)

	if os.Getenv(key) != `` {
		a := strings.Split(os.Getenv(key), separator)

		is := make([]int, len(a))

		for i, x := range a {
			is[i], _ = strconv.Atoi(x)
		}

		return is
	}

	if len(defaultValue) > 0 {
		return defaultValue
	}

	return []int{}
}

// Float returns the env var value as float64
func Float(prefix, key string, defaultValue float64) float64 {
	key = keyNormalize(prefix, key)

	if os.Getenv(key) != `` {
		if f, err := strconv.ParseFloat(os.Getenv(key), 64); err == nil {
			return f
		}
	}

	return defaultValue
}

// Bool returns the env var value as boolean
func Boolloat(prefix, key string, defaultValue bool) bool {
	key = keyNormalize(prefix, key)

	if os.Getenv(key) != `` {
		if b, err := strconv.ParseBool(os.Getenv(key)); err == nil {
			return b
		}
	}

	return defaultValue
}
