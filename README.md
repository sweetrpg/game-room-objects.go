# game-room-objects.go

[![CI](https://github.com/sweetrpg/game-room-objects.go/actions/workflows/ci.yaml/badge.svg)](https://github.com/sweetrpg/game-room-objects.go/actions/workflows/ci.yaml)
[![Coverage](https://img.shields.io/endpoint?url=https://sweetrpg.github.io/game-room-objects.go/coverage-badge.json)](https://sweetrpg.github.io/game-room-objects.go/)
[![License](https://img.shields.io/github/license/sweetrpg/game-room-objects.go.svg)](https://img.shields.io/github/license/sweetrpg/game-room-objects.go.svg)
[![Issues](https://img.shields.io/github/issues/sweetrpg/game-room-objects.go.svg)](https://img.shields.io/github/issues/sweetrpg/game-room-objects.go.svg)
[![PRs](https://img.shields.io/github/issues-pr/sweetrpg/game-room-objects.go.svg)](https://img.shields.io/github/issues-pr/sweetrpg/game-room-objects.go.svg)
[![Dependabot](https://badgen.net/github/dependabot/sweetrpg/game-room-objects.go)](https://badgen.net/github/dependabot/sweetrpg/game-room-objects.go)

Persistence models and API value objects for the Game Room microservice's domain: library, wishlist, table, and visibility. Pure data types - no business logic, no I/O.

## Install

```bash
go get github.com/sweetrpg/game-room-objects.go
```

## Packages

- `models` - persistence-layer structs, each embedding `model-core.go`'s `Auditable`
- `vo` - the matching API-facing value objects, each embedding `model-core.go`'s `AuditableVO`

## Documentation

Package documentation: [pkg.go.dev/github.com/sweetrpg/game-room-objects.go](https://pkg.go.dev/github.com/sweetrpg/game-room-objects.go).

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for the development workflow.

## Models

<a name="#person"></a>
### Person

* Fields
    * `name`: *String*. The name of the person.
    * `tag_ids`: *[Tag]*. An array of tags associated with the author.
    * [Audit fields](https://github.com/sweetrpg/design/blob/master/README.md#audit).

<a name="#personproperty"></a>
### PersonProperty

* Fields
    * `name`: *String*. The name of the property.
    * `type`: *String*. The type of the property's value.
    * `value`: *String*. The value of the property.
    * `person_id`: *Volume*. The identifier of the person associated with the property.
    * [Audit fields](https://github.com/sweetrpg/design/blob/master/README.md#audit).

<a name="#publisher"></a>
### Publisher

* Fields
    * `name`: *String*. The name of the publisher.
    * [Audit fields](https://github.com/sweetrpg/design/blob/master/README.md#audit).

<a name="#review"></a>
### Review

* Fields
    * `title`: *String*. The title of the review.
    * `text`: *String*. The body text of the review.
    * `volume_id`: *Volume*. The volume associated with the review.
    * [Audit fields](https://github.com/sweetrpg/design/blob/master/README.md#audit).

<a name="#studio"></a>
### Studio

* Fields
    * `name`: *String*. The name of the studio.
    * [Audit fields](https://github.com/sweetrpg/design/blob/master/README.md#audit).

<a name="#system"></a>
### System

* Fields
    * `game_system`: *String*. The identifier of the game system.
    * `edition`: *String*. The identifier of the game system's edition.
    * [Audit fields](https://github.com/sweetrpg/design/blob/master/README.md#audit).

<a name="#tag"></a>
### Tag

* Fields
    * `name`: *String*. The name of the tag.
    * `value`: *String*. The optional value of the tag.
    * [Audit fields](https://github.com/sweetrpg/design/blob/master/README.md#audit).

<a name="#contribution"></a>
### Contribution

* Fields
    * `person_id`: *Person*. Identifier of the Person making the contribution
    * `roles`: *[Enum{ContributionType}]*. An array of contribution types for this person.
                       Valid values are: `designer`, `developer`, `writer`, `artist`, `cartographer`, `editor`,
                       `producer`, `consultant`, `director`, `illustrator`, `misc`.
    * [Audit fields](https://github.com/sweetrpg/design/blob/master/README.md#audit).

<a name="#volume"></a>
### Volume

* Fields
    * `name`: *String*. The name of the volume.
    * `contributors`: *[Contribution]*. An array of contributions for the volume, indicating a person
                      and their contribution(s) to the volume.
    * `studio_ids`: *[Studio]*. An array of studios associated with the volume.
    * `publisher_ids`: *[Publisher]*. An array of publishers associated with the volume.
    * `system_id`: *System*. The game system associated with the volume.
    * `review_ids`: *[Review]*. An array of reviews associated with the volume.
    * `tag_ids`: *[Tag]*. An array of tags associated with the volume.
    * [Audit fields](https://github.com/sweetrpg/design/blob/master/README.md#audit).

<a name="#volumeproperty"></a>
### VolumeProperty

* Fields
    * `name`: *String*. The name of the property.
    * `type`: *String*. The type of the property's value.
    * `value`: *String*. The value of the property.
    * `volume_id`: *Volume*. The identifier of the volume associated with the property.
    * [Audit fields](https://github.com/sweetrpg/design/blob/master/README.md#audit).

## Development

1. Create a virtual environment
    ```shell
    python -m venv ~/.virtualenvs/sweetrpg-game-room-objects
    source ~/.virtualenvs/sweetrpg-game-room-objects/bin/activate
    ```
2. Install requirements
    ```shell
    pip3 install -r requirements/dev.txt
    ```

### Requirements

Requirements are organized in a number of role-based files in the `requirements/` directory:

* `dev.{in,txt}` -- for project development
* `pkg.{in,txt}` -- for the package itself
* `docs.{in,txt}` -- to generate documentation
* `tests.{in,txt}` -- to run unit tests

### Updating requirements

To update requirements, edit the appropriate `*.in` file, then run the `update.sh` script in the
same directory. The script will run `pip-compile` to generate the `*.txt` file with the actual
resolved versions and dependencies.

## Documentation

Documentation for this package can be found [here](https://sweetrpg.github.io/shelf-objects).
