package ru.expram.orchestra.cli;

import lombok.AllArgsConstructor;
import picocli.CommandLine;
import ru.expram.orchestra.cli.commands.OrchestraCommand;

import java.util.Set;

@AllArgsConstructor
public class OrchestraBootstrap {

    private final OrchestraCommand orchestraCommand;
    private final Set<BaseInfrastructureCommand> infrastructureCommandSet;

    public int execute(String[] args) {
        return createCommandLine().execute(args);
    }

    private CommandLine createCommandLine() {
        CommandLine commandLine = new CommandLine(orchestraCommand);
        infrastructureCommandSet.forEach(commandLine::addSubcommand);
        return commandLine;
    }
}
