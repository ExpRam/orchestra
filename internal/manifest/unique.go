package manifest

import (
	"errors"
	"fmt"

	"github.com/expram/orchestra/internal/errscope"
)

var ErrDuplicateManifest = errors.New("duplicate manifest")

func EnsureUnique(manifests []Manifest) error {
	counts := make(map[Identity]int, len(manifests))

	var duplicates []Identity

	for _, m := range manifests {
		identity := m.Identity()

		counts[identity]++
		if counts[identity] == 2 {
			duplicates = append(duplicates, identity)
		}
	}

	var problems errscope.Problems

	for _, identity := range duplicates {
		problems.Add(errscope.In(
			identity.String(),
			fmt.Errorf("%w: declared %d times", ErrDuplicateManifest, counts[identity]),
		))
	}

	if len(problems) > 0 {
		return Error{Problems: problems}
	}

	return nil
}
