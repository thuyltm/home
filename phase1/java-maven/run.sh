bazel build //phase1/java-maven:java-maven-lib
bazel build //phase1/java-maven:java-maven
bazel run //phase1/java-maven:java-maven
bazel test //phase1/java-maven:tests
bazel run //phase1/java-maven:image_load