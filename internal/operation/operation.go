package operation

type Mode uint8

const (
	ModePlan Mode = iota
	ModeApply
)

type Type uint8

const (
	TypeProvision Type = iota
	TypeDeprovision
)

type InfrastructureOperation struct {
	Mode Mode
	Type Type
}

func (op InfrastructureOperation) IsPlan() bool {
	return op.Mode == ModePlan
}

func (op InfrastructureOperation) IsApply() bool {
	return op.Mode == ModeApply
}

func (op InfrastructureOperation) IsProvision() bool {
	return op.Type == TypeProvision
}

func (op InfrastructureOperation) IsDeprovision() bool {
	return op.Type == TypeDeprovision
}
