dbt __transforms__ raw warehouse data into trusted data products. You write simple SQL select statements, and dbt handles the heavy lifting by __creating modular, maintainable data models__ that power analytics, operations, and AI.

You can use dbt to
- __Centralize and modularize__ your analytics code
- Apply sofware engineering best practices like __version control, testing, modularity, CI/CD__, and documentation to analytics workflows
- Build __idempotent transformations__ that are safe to rerun and produce consistent result

As a dbt user, your main focus will be on __writing models (select queries) that reflect core business logic__--there's no need to write boilerplate code to create tables or views, or to define the order of execution of your models. Instead, dbt handles __turning these models into objects__ in your warehouse for you

Focusing on modules allows you to reduce bugs, standardize analytics logic