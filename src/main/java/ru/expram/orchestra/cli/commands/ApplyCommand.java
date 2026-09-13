package ru.expram.orchestra.cli.commands;

import picocli.CommandLine.Command;
import ru.expram.orchestra.cli.BaseInfrastructureCommand;
import ru.expram.orchestra.operation.InfrastructureOperation;
import ru.expram.orchestra.operation.OperationMode;
import ru.expram.orchestra.operation.OperationType;

@Command(name = "apply",
        mixinStandardHelpOptions = true,
        description = "Apply infrastructure changes")
public class ApplyCommand extends BaseInfrastructureCommand {

    @Override
    public InfrastructureOperation operation() {
        return new InfrastructureOperation(
                OperationMode.APPLY,
                OperationType.PROVISION
        );
    }

    @Override
    public Integer call() throws Exception {
        return 0;
    }
}
