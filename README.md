# raiding-raccoon

[![LICENSE](https://img.shields.io/badge/license-MIT-orange.svg)](LICENSE)
[![Build Status](https://travis-ci.com/HeikoAlexanderWeber/raiding-raccoon.svg?token=jLWKSu6GaoZv38y9JzqL&branch=master)](https://travis-ci.com/HeikoAlexanderWeber/raiding-raccoon)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/c6a5c39e7bd042389370b22a3c98959f)](https://www.codacy.com/app/HeikoAlexanderWeber/raiding-raccoon?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=HeikoAlexanderWeber/raiding-raccoon&amp;utm_campaign=Badge_Grade)

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

* `Go` >= 1.12.0

## Getting started

For setting up the environment, just call `sh scripts/setup.sh`.

## Commands

* `make install` for downloading dependencies (done in `setup.sh`)
* `make format` for formatting the code using `gofmt`
* `make build` for building the program (UNIX: `bin/raiding-raccoon`, WIN32: `bin/raiding-raccoon.exe`)
* `make test` for executing unit tests
* `make cover` for generating an interactive coverage report
* `make run` for running the program

## Dependencies

This project makes use of the following 3rd party dependencies:

| Library                | License                                                                   | Link                                                                                   |
| :--------------------- | :------------------------------------------------------------------------ | :------------------------------------------------------------------------------------- |
| sirupsen/logrus        | [MIT](https://github.com/sirupsen/logrus/blob/master/LICENSE)             | [https://github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)               |
| spf13/pflag            | [BSD3-Clause](https://github.com/spf13/pflag/blob/master/LICENSE)         | [https://github.com/spf13/pflag](https://github.com/spf13/pflag)                       |
| orcaman/concurrent-map | [MIT](https://github.com/orcaman/concurrent-map/blob/master/LICENSE)      | [https://github.com/orcaman/concurrent-map](https://github.com/orcaman/concurrent-map) |
| shabbyrobbe/xmlwriter  | [Apache 2.0](https://github.com/shabbyrobe/xmlwriter/blob/master/LICENSE) | [https://github.com/shabbyrobe/xmlwriter](https://github.com/shabbyrobe/xmlwriter)     |

## Using this program

In order to use this program, go ahead and build an executable. After that is done, call the executable with the `--help` flag.\
This will show all available argumentes for the program.\
The most important flag is the `--start` flag. It defines, which URL the crawler starts at. The output will be in the GraphML format, written to the `STDOUT` pipe. In order to save the output, redirect `STDOUT` to a file (using the `>` operator). A few example calls can be found here:

* `./bin/program --help`
* `./bin/program --start https://github.com/HeikoAlexanderWeber > github.graphml`
* `./bin/program --start https://cassini.de > cassini.graphml`

## Evaluating the output

### GraphML

GraphML files can be parsed by a few programs since it is a standardized format. The author recommends the [Gephi](https://gephi.org/) program.\
See [https://gephi.org/](https://gephi.org/) for more information.\
This program contains tools like [ForceAtlas2](https://medialab.sciencespo.fr/publications/Jacomy_Heymann_Venturini-Force_Atlas2.pdf) simulation for logically grouping the nodes (which can lead to very interesting results).
