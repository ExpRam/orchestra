package ru.expram.orchestra.operation;

public record InfrastructureOperation(
        OperationMode operationMode,
        OperationType operationType
) {
    // OperationMode
    public boolean isPlan() {
        return this.operationMode == OperationMode.PLAN;
    }

    public boolean isApply() {
        return !isPlan();
    }

    // OperationType
    public boolean isProvision() {
        return this.operationType == OperationType.PROVISION;
    }

    public boolean isDeprovision() {
        return !isProvision();
    }
}
