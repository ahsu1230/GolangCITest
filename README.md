# GolangCITest

Self-made project to integrate Golang with various Continuous Integration (CI) tools.

CI tools include:
 - Circle.CI (https://circleci.com/)
 - Semaphore (https://semaphoreci.com/)
 - Google Cloud Build (https://cloud.google.com/cloud-build)

 Made sure to test the following:
  - unit tests
  - local packages (self-made module packages)
  - external packages

### Resources
https://github.com/golang/dep
https://medium.com/@masroor.hasan/publishing-a-go-package-to-github-with-circleci-2-0-41c1bde1493b

### Notes
 - Watch out for go version (in mod and deployed in CI)
 - Keep track of what kind of Docker platforms are builds being deployed to
 - Go to Github Settings > Branches > protection rules to setup requirements on waiting for CI checks
