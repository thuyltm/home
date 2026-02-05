import java.util.jar.Attributes.Name

ThisBuild / scalaVersion := Version.Scala
ThisBuild / resolvers += Resolver.mavenLocal

lazy val root = (project in file("."))
  .settings(
    name         := "vertx-scala",
    version      := "0.1.0-SNAPSHOT",
    organization := "vertx.scala",
    description  := "Eclipse Vert.x is a tool-kit for building reactive applications on the JVM. This app uses Vert.x-Scala and ScalaTest.",
    libraryDependencies ++= Seq(
      Library.vertx_lang_scala,
      Library.vertx_web,
      Library.scala_logging,
      Library.vertx_lang_scala_test % Test,
      Library.scalaTest             % Test,
      Library.logback               % Runtime,
    ),
    mainVerticle := "vertx.scala.myapp.HttpVerticle"
  )
