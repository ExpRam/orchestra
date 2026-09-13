package ru.expram.orchestra.di;

import dagger.Component;
import ru.expram.orchestra.cli.OrchestraBootstrap;
import ru.expram.orchestra.cli.di.BootstrapModule;
import ru.expram.orchestra.cli.di.CliCommandsModule;

import javax.inject.Singleton;

@Singleton
@Component(
        modules = {
            BootstrapModule.class,
            CliCommandsModule.class
        }
)
public interface OrchestraComponent {

        OrchestraBootstrap orchestraBootstrap();
}
