package operation

type OperationMode uint8

const (
	ModePlan OperationMode = iota
	ModeApply
)

type OperationType uint8

const (
	TypeProvision OperationType = iota
	TypeDeprovision
)

type InfrastructureOperation struct {
	Mode OperationMode
	Type OperationType
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
