package ru.expram.orchestra.cli;

import picocli.CommandLine;
import ru.expram.orchestra.operation.InfrastructureOperation;

import java.util.List;
import java.util.concurrent.Callable;

public abstract class BaseInfrastructureCommand implements Callable<Integer> {

    @CommandLine.Option(names = {"-f", "--filter"}, split = ",", description = "Filter user-manifests by metadata values")
    protected List<String> filterTokens;

    public abstract InfrastructureOperation operation();
}
