package ru.expram.orchestra.cli.di;

import dagger.Module;
import dagger.Provides;
import dagger.multibindings.IntoSet;
import ru.expram.orchestra.cli.BaseInfrastructureCommand;
import ru.expram.orchestra.cli.commands.ApplyCommand;
import ru.expram.orchestra.cli.commands.DestroyCommand;
import ru.expram.orchestra.cli.commands.PlanCommand;

@Module
public final class CliCommandsModule {

    @Provides
    @IntoSet
    static BaseInfrastructureCommand providePlanCommand() {
        return new PlanCommand();
    }

    @Provides
    @IntoSet
    static BaseInfrastructureCommand provideApplyCommand() {
        return new ApplyCommand();
    }

    @Provides
    @IntoSet
    static BaseInfrastructureCommand provideDestroyCommand() {
        return new DestroyCommand();
    }
}
