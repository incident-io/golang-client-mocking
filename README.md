# Mocking external client libraries using context.Context

This code is paired with a blog post:

> [Mocking external client libraries using context.Context](https://incident.io/block/mocking-clients)

The post describes a method for mocking external client libraries which can be
applied to most codebases.

It mocks the client using the context parameter, ensuring all code will receive
the same mock for the duration of an operation, and removing the need to
explicitly stub a mock in code that might get called during a test.

We use ginkgo for our test suite, but the technique is independent of test
framework.
