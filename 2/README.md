# Statefulset

This is a _very_ raw statefulset for Litecoind. There would be many additions I would add, but don't understand the application well enough to tackle them now:

- Healthcheck/liveness checks: How do I konw the application is healthy/running?
- Connectivity - Do these applications need to talk to eachother, or is the use of a Statefulset simply to manage PVC claims more easily?
- I would extract the config used by the daemon into a Secret to prevent injecting secrets in the Docker file
