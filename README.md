<!--

    TODO:

    - Add the project to the CircleCI:
      https://circleci.com/setup-project/gh/giantswarm/app-service

    - Import RELEASE_TOKEN variable from template repository for the builds:
      https://circleci.com/gh/giantswarm/app-service/edit#env-vars

    - Change the badge (with style=shield):
      https://circleci.com/gh/giantswarm/app-service/edit#badges
      If this is a private repository token with scope `status` will be needed.

    - Run `devctl replace -i "app-service" "$(basename $(git rev-parse --show-toplevel))" *.md`
      and commit your changes.

    - If the repository is public consider adding godoc badge. This should be
      the first badge separated with a single space.
      [![GoDoc](https://godoc.org/github.com/giantswarm/app-service?status.svg)](http://godoc.org/github.com/giantswarm/app-service)

-->
[![CircleCI](https://circleci.com/gh/giantswarm/app-service.svg?style=shield)](https://circleci.com/gh/giantswarm/app-service)

# app-service

Microservice that is part of the Giant Swarm App Platform.
