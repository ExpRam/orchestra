package yaml

import (
	"github.com/expram/orchestra/internal/contract"
	"github.com/expram/orchestra/internal/manifest"
)

type UnresolvedManifestParser = contract.Parser[manifest.UnresolvedManifest]

type UnresolvedManifestValidator = contract.Validator[document]
