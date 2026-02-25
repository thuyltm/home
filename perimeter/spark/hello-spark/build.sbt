name := "hello-spark"
version := "1.0"
scalaVersion := "2.13.17"
lazy val sparkVersion = "4.1.1"

libraryDependencies ++= Seq(
// Use the "provided" configuration, which will scope your dependent library
  "org.scala-lang.modules" %% "scala-parser-combinators" % "2.3.0" % "provided",
  "org.apache.spark" %% "spark-core" % "4.1.1" % "provided",
  "org.apache.spark" %% "spark-sql" % "4.1.1" % "provided",
  "org.apache.spark" %% "spark-mllib" % "4.1.1" % "provided"
)

import sbtassembly.MergeStrategy
import sbtassembly.PathList

assembly / assemblyMergeStrategy := {
  case PathList("META-INF", _*) => MergeStrategy.discard
  case _                        => MergeStrategy.first
}
