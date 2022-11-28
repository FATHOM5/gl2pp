# :white_sun_cloud: tools/gl

GitLab cli for automating various tasks relevant to our agile teams.

---

## Getting Started

Complete [onboarding](https://github.com/fathom5/codex/wiki/Onboarding), then run:

    git clone https://github.com/fathom5/skeleton-go-cli.git ${F5_DIR}/tools/gl
    cd ${F5_DIR}/tools/gl
    make build

---

## Return the current logged in user

bin/gl whoami

## List the groups assigned to that user

bin/gl list-groups

## List the group iterations using the "group id" from list-groups

bin/gl list-group-iterations --group-id=478 (from list groups)

## List the group issues using the "group id" and the "iid" from list-group-iterations

bin/gl list-group-issues --group-id=478 --iid=114 (from list group iterations)

## Output issues to a .CSV file

bin/gl list-group-issues --group-id=478 --iteration-id=114 --output filenameyouchoose.csv (with filename of your choice)

## Directions

After saving the build files to your local through a git clone from https://gitlab.fathom5.work/tools/gl, store the directory within the path of your local environment variables.

Navigate to 'GL' in the terminal and run "make build"

Running "gl/bin whoami" returns the current logged in user

Once confirming the correct user is logged in via the returned keys, the client manages all the API interactions with GitLab
