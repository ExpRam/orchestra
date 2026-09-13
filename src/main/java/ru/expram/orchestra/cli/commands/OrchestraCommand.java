package ru.expram.orchestra.cli.commands;

import picocli.CommandLine.Command;

@Command(name = "orchestra",
        mixinStandardHelpOptions = true,
        version = "0.0.1-Beta",
        description = "Orchestrator for GitOps. Describe anything with k8s-like manifests")
public class OrchestraCommand {
}
