package ru.expram.orchestra;

import ru.expram.orchestra.di.DaggerOrchestraComponent;

public class Orchestra {

    static void main(String[] args) {
        System.exit(
                DaggerOrchestraComponent
                        .create()
                        .orchestraBootstrap()
                        .execute(args)
        );
    }
}
