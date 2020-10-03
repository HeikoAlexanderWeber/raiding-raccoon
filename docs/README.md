# raiding-raccoon

[![LICENSE](https://img.shields.io/badge/license-MIT-orange.svg)](LICENSE)

## Stats

[![SonarCloud](https://sonarcloud.io/images/project_badges/sonarcloud-black.svg)](https://sonarcloud.io/dashboard?id=HeikoAlexanderWeber.raiding-raccoon)\
[![Quality gate](https://sonarcloud.io/api/project_badges/quality_gate?project=HeikoAlexanderWeber.raiding-raccoon)](https://sonarcloud.io/dashboard?id=HeikoAlexanderWeber.raiding-raccoon)\
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=HeikoAlexanderWeber.raiding-raccoon&metric=security_rating)](https://sonarcloud.io/dashboard?id=HeikoAlexanderWeber.raiding-raccoon)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=HeikoAlexanderWeber.raiding-raccoon&metric=vulnerabilities)](https://sonarcloud.io/dashboard?id=HeikoAlexanderWeber.raiding-raccoon)\
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=HeikoAlexanderWeber.raiding-raccoon&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=HeikoAlexanderWeber.raiding-raccoon)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=HeikoAlexanderWeber.raiding-raccoon&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=HeikoAlexanderWeber.raiding-raccoon)\
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=HeikoAlexanderWeber.raiding-raccoon&metric=bugs)](https://sonarcloud.io/dashboard?id=HeikoAlexanderWeber.raiding-raccoon)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=HeikoAlexanderWeber.raiding-raccoon&metric=code_smells)](https://sonarcloud.io/dashboard?id=HeikoAlexanderWeber.raiding-raccoon)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=HeikoAlexanderWeber.raiding-raccoon&metric=sqale_index)](https://sonarcloud.io/dashboard?id=HeikoAlexanderWeber.raiding-raccoon)\
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=HeikoAlexanderWeber.raiding-raccoon&metric=coverage)](https://sonarcloud.io/dashboard?id=HeikoAlexanderWeber.raiding-raccoon)



## Prerequisites

* `Go` >= 1.13.0

## Getting started

For setting up the environment, simply call `sh scripts/setup.sh`.

## Commands

* `make install` for downloading dependencies (done in `setup.sh`)
* `make format` for formatting the code using `gofmt`
* `make docker-build` for building the production grade docker image
* `make docker-build-debug` for building the debug docker image
* `make test` for executing unit tests
* `make bench` for executing benchmarks
* `make cover` for generating an interactive coverage report

## Dependencies

This project makes use of the following 3rd party dependencies:

| Library                | License                                                                   | Link                                                                                   |
| :--------------------- | :------------------------------------------------------------------------ | :------------------------------------------------------------------------------------- |
| sirupsen/logrus        | [MIT](https://github.com/sirupsen/logrus/blob/master/LICENSE)             | [https://github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)               |
| spf13/pflag            | [BSD3-Clause](https://github.com/spf13/pflag/blob/master/LICENSE)         | [https://github.com/spf13/pflag](https://github.com/spf13/pflag)                       |
| orcaman/concurrent-map | [MIT](https://github.com/orcaman/concurrent-map/blob/master/LICENSE)      | [https://github.com/orcaman/concurrent-map](https://github.com/orcaman/concurrent-map) |
| shabbyrobbe/xmlwriter  | [Apache 2.0](https://github.com/shabbyrobe/xmlwriter/blob/master/LICENSE) | [https://github.com/shabbyrobe/xmlwriter](https://github.com/shabbyrobe/xmlwriter)     |

## Running the container

Following environment variables are used by the container:

* `RR_START_URL`\
    This variable defines the URL that the crawler starts at. The output will be in the GraphML format, written to the `STDOUT` pipe of the container.
* `RR_REDIS_BACKBONE`\
    Optional. Define a redis backbone to enable caching and optimize RAM usage.

Example: `docker run -e RR_START_URL="https://cassini.de" raiding-raccoon:latest > cassini.graphml`

## Evaluating the output

### GraphML

GraphML files can be parsed by a few programs since it is a standardized format. The author recommends the [Gephi](https://gephi.org/) program.\
See [https://gephi.org/](https://gephi.org/) for more information.\
This program contains tools like [ForceAtlas2](https://medialab.sciencespo.fr/publications/Jacomy_Heymann_Venturini-Force_Atlas2.pdf) simulation for logically grouping the nodes (which can lead to very interesting results).
