# ProfileStatusSyncer

[![Go Report Card](https://goreportcard.com/badge/github.com/XiaoMengXinX/ProfileStatusSyncer)](https://goreportcard.com/report/github.com/XiaoMengXinX/ProfileStatusSyncer)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/XiaoMengXinX/ProfileStatusSyncer/gh)

A tool to synchronize user profile status of GitHub and Netease CloudMusic

## Features

- Synchronize user profile status between GitHub and NeteaseCloud Music
- Automatically maintain your NeteaseCloud Music profile status to avoid it being cancelled after three days

## Quick start

### Preparations

1. GitHub Access Token (Generate
   from [Settings -> Developer settings -> Personal access tokens](https://github.com/settings/tokens))

2. Your NeteaseCloud Music Cookie (Only needs the `MUSIC_U` field)

The NeteaseCloud Music Cookie can be taken from your browser's cookie storage.Or you can also use
the [QuickLogin tool](https://github.com/XiaoMengXinX/Fuck163MusicTasks/releases/tag/v2.1.1) to get it.

### Configuration

#### GitHub Action

1. Fork the [ProfileStatusSyncer](https://github.com/XiaoMengXinX/ProfileStatusSyncer) repository.

2. Go to settings -> secrets -> Actions

3. **Add a secret field `TOKEN` and set the value to your GitHub access token.**

4. **Add a secret field `MUSIC_U` and set the value to your NeteaseCloud Music Cookie.**

Then you can then run the action named syncer manually and check the output log to see if it works.

The action will run automatically at 21:00 UTC time every day by default, if you want to change this setting, please
edit `.github/workflows/syncer.yml` file.

#### Command Line

Download the latest release from [here](https://github.com/XiaoMengXinX/ProfileStatusSyncer/releases/latest) and run the
following command:

```bash
$ chmod +x ./ProfileStatusSyncer
$ GITHUB_TOKEN=<your GitHub access token> MUSIC_U=<your NeteaseCloud Music Cookie> ./ProfileStatusSyncer
```

### Mode Configuration

The tools can run in three modes:

1. Set the environment variable `MODE` to `GitHub2Netease` to sync GitHun profile status to NeteaseCloud Music.

2. Set the environment variable `MODE` to `Netease2GitHub` to sync NeteaseCloud Music profile status to GitHub.

3. Set the environment variable `MODE` to `KeepNeteaseStatus` to maintain the profile status of NeteaseCloud Music.

Notice that if you don't set the `MODE` variable, the tool will run in `GitHub2Netease` mode by default.And if you run
it on GitHub Action, just add a secret field called `MODE` and set its value.