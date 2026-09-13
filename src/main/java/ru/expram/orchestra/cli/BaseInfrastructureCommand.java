package ru.expram.orchestra.cli;

import ru.expram.orchestra.operation.InfrastructureOperation;

public abstract class BaseInfrastructureCommand {

    public abstract InfrastructureOperation operation();
}
