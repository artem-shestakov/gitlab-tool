# GitLab tool
Simple utili that helps in working with GitLab CI

## Getting started

1. Create a personal access token with api scope
2. Set environment variable GL_TOKEN to get access to your project.
>Set GL_TOKEN in your project ‘Settings → CI/CD → Variables’


## Features
### Send pipeline status notification
Use commands `pipeline` and subcommand `notify`.
For selecting channel (Telegram, Slack) use flag `--channel`
Example:
```shell
gitlab-tool pipeline notify --channel telegram
```
Using docker image:
```yaml
stages:
    - build
    - test
    - notify

...

notify:
  # Use the official docker image.
  image: artemshestakov/gitlab-tool:0.0.2
  stage: notify
  script:
    - gitlab-tool pipeline notify --channel telegram
  when: always
```