If the console throws the below error
```sh
ERROR: /home/thuy/.cache/bazel/_bazel_thuy/74279c9e3277e09badffc975a0b6c8d8/external/gazelle~~go_deps~org_golang_x_sys/unix/BUILD.bazel:3:11: GoCompilePkg external/gazelle~~go_deps~org_golang_x_sys/unix/unix.a failed: (Exit 1): builder failed: error executing GoCompilePkg command (from target @@gazelle~~go_deps~org_golang_x_sys//unix:unix) bazel-out/k8-opt-exec-ST-d57f47055a04/bin/external/rules_go~~go_sdk~main___download_0/builder_reset/builder compilepkg -sdk external/rules_go~~go_sdk~main___download_0 -goroot ... (remaining 527 arguments skipped)

Use --sandbox_debug to see verbose messages from the sandbox and retain the sandbox build root for debugging
compilepkg: error running subcommand external/rules_go~~go_sdk~main___download_0/pkg/tool/linux_amd64/pack: fork/exec external/rules_go~~go_sdk~main___download_0/pkg/tool/linux_amd64/pack: no such file or directory
Target //phase1/echo:echo_lib failed to build
Use --verbose_failures to see the command lines of failed build steps.
```
Upgrader rules_go