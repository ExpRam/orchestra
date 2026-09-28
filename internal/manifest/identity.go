package manifest

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrEmptyKind       = errors.New("kind must not be empty")
	ErrEmptyAPIVersion = errors.New("api version must not be empty")
	ErrEmptyName       = errors.New("name must not be empty")
)

type Kind string

func NewKind(value string) (Kind, error) {
	if strings.TrimSpace(value) == "" {
		return "", ErrEmptyKind
	}

	return Kind(value), nil
}

func (k Kind) String() string {
	return string(k)
}

type APIVersion string

func NewAPIVersion(value string) (APIVersion, error) {
	if strings.TrimSpace(value) == "" {
		return "", ErrEmptyAPIVersion
	}

	group, version, qualified := strings.Cut(value, "/")
	if qualified && (group == "" || version == "" || strings.Contains(version, "/")) {
		return "", fmt.Errorf("api version %q must be <version> or <group>/<version>", value)
	}

	return APIVersion(value), nil
}

func (v APIVersion) String() string {
	return string(v)
}

type Name string

func NewName(value string) (Name, error) {
	if strings.TrimSpace(value) == "" {
		return "", ErrEmptyName
	}

	return Name(value), nil
}

func (n Name) String() string {
	return string(n)
}

type Type struct {
	Kind       Kind
	APIVersion APIVersion
}

type Identity struct {
	Kind Kind
	Name Name
}

func (i Identity) String() string {
	return fmt.Sprintf("%s %q", i.Kind, i.Name)
}
