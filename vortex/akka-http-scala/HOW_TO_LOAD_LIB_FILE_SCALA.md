Load the separated folder such as vertx-scala into the IntelliJ IDE.

IntelliJ Idea will detect sbt project and then auto importing is trigger and indexing

Sbt version is declared in __project/build__ properties

Sbt will download packages using coursier into the folder __~/.cache/coursier__. Also the sbt download packages command is
__sbt updateClassifiers__

Coursier is the modern, default dependency resolver for sbt v1.3.0_ compared to the older Apache Ivy

