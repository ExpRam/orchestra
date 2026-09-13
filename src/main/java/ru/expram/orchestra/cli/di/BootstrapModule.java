package ru.expram.orchestra.cli.di;

import dagger.Module;
import dagger.Provides;
import ru.expram.orchestra.cli.BaseInfrastructureCommand;
import ru.expram.orchestra.cli.OrchestraBootstrap;
import ru.expram.orchestra.cli.commands.OrchestraCommand;

import java.util.Set;

@Module
public final class BootstrapModule {

    @Provides
    static OrchestraBootstrap provideBootstrap(
            OrchestraCommand orchestraCommand,
            Set<BaseInfrastructureCommand> infrastructureCommandSet
    ) {
        return new OrchestraBootstrap(orchestraCommand, infrastructureCommandSet);
    }
}
