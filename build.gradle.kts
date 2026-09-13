plugins {
    id("application")
}

group = "ru.expram.orchestra"
version = "0.0.1-Beta"

application {
    mainClass = "${project.group}.Orchestra"
}

repositories {
    mavenCentral()
}

dependencies {
    // Fw, DI
    implementation("info.picocli:picocli:4.7.7")
    annotationProcessor("info.picocli:picocli-codegen:4.7.7")

    implementation("com.google.dagger:dagger:2.60.1")
    annotationProcessor("com.google.dagger:dagger-compiler:2.60.1")
    // Libs
    implementation("org.projectlombok:lombok:1.18.42")
    annotationProcessor("org.projectlombok:lombok:1.18.42")
    // Tests
    testImplementation(platform("org.junit:junit-bom:6.0.0"))
    testImplementation("org.junit.jupiter:junit-jupiter")
    testRuntimeOnly("org.junit.platform:junit-platform-launcher")
}

tasks.compileJava {
    // https://picocli.info/#_using_build_tools
    options.compilerArgs.add(
        "-Aproject=${project.group}/${project.name}"
    )
}

tasks.test {
    useJUnitPlatform()
}