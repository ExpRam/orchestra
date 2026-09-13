package ru.expram.orchestra.cli.commands;

import ru.expram.orchestra.cli.BaseInfrastructureCommand;
import ru.expram.orchestra.operation.InfrastructureOperation;
import ru.expram.orchestra.operation.OperationMode;
import ru.expram.orchestra.operation.OperationType;

public class PlanCommand extends BaseInfrastructureCommand {

    private boolean isDestroy;

    @Override
    public InfrastructureOperation operation() {
        return new InfrastructureOperation(
                OperationMode.PLAN,
                isDestroy ? OperationType.DEPROVISION : OperationType.PROVISION
        );
    }
}
