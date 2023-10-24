# Server

## Echo

We use echo as the syntax is concise and logging in echo is json which has great compaibility with nosql.

## The Twelve Factor

https://12factor.net/

- Use **declarative** formats for setup automation, to minimize time and cost for new developers joining the project;
- Have a **clean contract** with the underlying operating system, offering **maximum portability** between execution environments;
- Are suitable for **deployment** on modern **cloud platforms**, obviating the need for servers and systems administration;
- **Minimize divergence** between development and production, enabling **continuous deployment** for maximum agility;
- And can **scale up** without significant changes to tooling, architecture, or development practices.

The one we are going to focus is IX. Disposability: Maximize robustness with fast startup and graceful shutdown

## Graceful Shutdown

"Graceful Shutdown" refers to a scenario where, for example, if there is a registration process going on with a database, and during this process, if there's a server shutdown or similar event, the system waits until the registration process is completed before proceeding with the shutdown. This concept is particularly important in scenarios where financial transactions are involved; a shutdown during such processes could lead to serious problems.
- `SIGINT`: Force shutdown application
- `SIGTERM`: Force shutdown but return resources to computer