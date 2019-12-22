# Travis CI

Look in the [Travis File](../.travis.yaml).

## Push Images

Following guide for secrets here: https://docs.travis-ci.com/user/docker/#pushing-a-docker-image-to-a-registry

## Secrets

Followed this recommendation for managing secrets with Travis: https://docs.travis-ci.com/user/environment-variables#encrypting-environment-variables

## Future Work

- Not the best rollout of the application. No staging environment, no tests/integration tests.
- No branch strategies implemented
- Needs unit tests as part of the docker packaing
- Should run security scans as part of the CI process to provide updates. Possibly fail rollout when new ones are found and expect human approval.
