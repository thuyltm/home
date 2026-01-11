
![Architecture](https://grafana.com/media/docs/loki/get-started-flog-v3.png)

- Grafana Loki has a microservices-based architecture and is designed to run as a horizontally scalable, distributed system. The system has multiple components that can run separately and in parallel. 
- The Grafana Loki design compiles the code for all components(s) into a single binary or Docker image. The _-target_ command-line flag controls which component(s) that binary will behave as

Loki's simple scalable deployment mode separates execution paths into read, write and backend targets

![scalable-monolithic-mode](https://grafana.com/docs/loki/latest/get-started/scalable-monolithic-mode.png)
- _-target=write_ The write target is stateful and is controlled by a Kubernetes StatefulSet.
- _-target=read_ The read target is stateless and can be run as a Kubernetes Deployment that can be scaled automatically
- _-target=backend_ The backend target is stateful, and is controlled by a Kubernetes StatefulSet

#### Understand labels
Labels are a crucial part of Loki. They allow Loki to organize and group together log messages into log streams. Each log stream must have at least one label to be stored and queried in Loki.

In Loki, the content of each log line is not indexed. Instead, log entries are grouped into streams which are indexed with labels.

The way that labels are added to logs is configured in the client that you use to send logs to Loki. The specific configuration will be different for each client

For example: Grafana Alloy is used to send logs to Loki
```sh
loki.process "add_new_label" {
    // Extract the value of "level" from the log line and add it to the extracted map as "extracted_level"
    // You could also use "level" = "", which would extract the value of "level" and add it to the extracted map as "level"
    // but to make it explicit for this example, we will use a different name.
    //
    // The extracted map will be covered in more detail in the next section.
    stage.logfmt {
        mapping = {
            "extracted_level" = "level",
        }
    }

    // Add the value of "extracted_level" from the extracted map as a "level" label
    stage.labels {
        values = {
            "level" = "extracted_level",
        }
    }

    forward_to = [loki.relabel.add_static_label.receiver]
}
```

[Link is the detail of how to configure Loki](https://grafana.com/docs/loki/latest/configure/)