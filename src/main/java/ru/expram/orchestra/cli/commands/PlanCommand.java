package ru.expram.orchestra.cli.commands;

import picocli.CommandLine.Command;
import picocli.CommandLine.Option;
import ru.expram.orchestra.cli.BaseInfrastructureCommand;
import ru.expram.orchestra.operation.InfrastructureOperation;
import ru.expram.orchestra.operation.OperationMode;
import ru.expram.orchestra.operation.OperationType;

@Command(name = "plan",
        mixinStandardHelpOptions = true,
        description = "Preview infrastructure changes")
public class PlanCommand extends BaseInfrastructureCommand {

    @Option(names = "--destroy", description = "Look a plan for destroy operation")
    private boolean isDestroy;

    @Override
    public InfrastructureOperation operation() {
        return new InfrastructureOperation(
                OperationMode.PLAN,
                isDestroy ? OperationType.DEPROVISION : OperationType.PROVISION
        );
    }

    @Override
    public Integer call() throws Exception {
        return 0;
    }
}
