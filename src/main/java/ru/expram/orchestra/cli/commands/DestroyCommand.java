package ru.expram.orchestra.cli.commands;

import picocli.CommandLine.Command;
import ru.expram.orchestra.cli.BaseInfrastructureCommand;
import ru.expram.orchestra.operation.InfrastructureOperation;
import ru.expram.orchestra.operation.OperationMode;
import ru.expram.orchestra.operation.OperationType;

@Command(name = "destroy",
        mixinStandardHelpOptions = true,
        description = "Apply infrastructure components destroy")
public class DestroyCommand extends BaseInfrastructureCommand {

    @Override
    public InfrastructureOperation operation() {
        return new InfrastructureOperation(
                OperationMode.APPLY,
                OperationType.DEPROVISION
        );
    }

    @Override
    public Integer call() throws Exception {
        return 0;
    }
}
