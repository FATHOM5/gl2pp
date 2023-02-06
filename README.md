# :white_sun_cloud: gl2pp

`gl2pp` allows you to export the issues in a group iteration from gitlab in a
format that [PlanningPoker.com](https://planningpoker.com/) understands.

You will need a gitlab.com (or self-hosted instance) with either a Premium or Ultimate
license, as this tool exports issues scoped by a group's iteration. [More
Info](https://about.gitlab.com/pricing/).

---

# Quick Start Guide

    go install github.com/FATHOM5/gl2pp
    gl2pp --version

## Getting Help

Display a list of all subcommands and global flags:

    gl2pp help

Display the help text for a given subcommand:

    gl2pp help {subcommand}

e.g. `gl2pp help whoami`

## Configuration

To save time, set the following ENV vars, which will be used as the default
values for the global flags:

  - `GITLAB_BASE_URL` (defaults to: https://gitlab.com/)
  - `GITLAB_TOKEN` (no default)
    - Create a [Personal Access Token](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html)

## Return the current logged in user

    gl2pp whoami

## List the gitlab groups you have access to

    gl2pp list-groups

_Make note of the {GROUP_ID}. You'll need this in the other commands._

## List the group iterations using the "group id" from list-groups

    gl2pp list-group-iterations --group-id={GROUP_ID}

_Make note of the {ITERATION_ID}. You'll need this in the other commands._

## List issues from the iteration

    gl2pp list-group-issues --group-id={GROUP_ID} --iteration-id={ITERATION_ID}

_Make note of the {ITERATION_ID}. You'll need this in the other commands._

## Export the issues for PlanningPoker.com

    gl2pp list-group-issues \
        --group-id={GROUP_ID} \
        --iteration-id={ITERATION_ID} \
        --output planningpoker.csv

---

## Want to hack on this?

The following tools are required in your development environment:

  - Install [go v1.19+](https://go.dev/)
  - Install [mmake](https://github.com/tj/mmake), and alias it to `make`
  - Install [upx](https://upx.github.io/)

### Getting Started

    git clone https://github.com/FATHOM5/gl2pp.git
    cd gl2pp
    make init

### Build It

    make build

### Test It

    make test

### Install It

    make install

### Uninstall It

    make Uninstall

### Other Commands

    make

---

