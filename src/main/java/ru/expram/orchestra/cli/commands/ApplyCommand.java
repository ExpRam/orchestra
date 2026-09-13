package ru.expram.orchestra.cli.commands;

import ru.expram.orchestra.cli.BaseInfrastructureCommand;
import ru.expram.orchestra.operation.InfrastructureOperation;
import ru.expram.orchestra.operation.OperationMode;
import ru.expram.orchestra.operation.OperationType;

public class ApplyCommand extends BaseInfrastructureCommand {

    @Override
    public InfrastructureOperation operation() {
        return new InfrastructureOperation(
                OperationMode.APPLY,
                OperationType.PROVISION
        );
    }
}
