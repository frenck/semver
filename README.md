# Semver

[![GitHub Release][releases-shield]][releases]
![Project Stage][project-stage-shield]
![Project Maintenance][maintenance-shield]
[![GitHub Activity][commits-shield]][commits]
[![License][license-shield]](LICENSE.md)

Simple CLI tool to check a semver version against a condition

## About

There were many times I needed a way to check version numbers of programs
in bash scrips. Conclusion: there is no way to do this right!

That is why I have created this little CLI tool to help me (and maybe you) out!

## Installation

Just download a binary release on the [GitHub Releases page][releases].

If you would like to build it yourself, go ahead, the source code is in this
repository.

## Usage

```bash
semver [-q] <condition> <version>
```

A `<condition>` specifies which versions are satisfactory, while `<version>`
specifies the version that needs to be checked against.

Adding `-q` causes to the program to not give any output at all.

The command gives feedback using exit codes:

- **0**: Version is valid for the given condition
- **1**: Version is NOT valid for the given condition
- **128**: Incorrect arguments/parameters.

## Conditions

A `<condition>` is composed of an operator and a version.
The supported operators are:

- `<1.0.0` Less than `1.0.0`
- `<=1.0.0` Less than or equal to `1.0.0`
- `>1.0.0` Greater than `1.0.0`
- `>=1.0.0` Greater than or equal to `1.0.0`
- `1.0.0`, `=1.0.0`, `==1.0.0` Equal to `1.0.0`
- `!1.0.0`, `!=1.0.0` Not equal to `1.0.0`. Excludes version `1.0.0`.

Note that spaces between the operator and the version will be gracefully tolerated.

A `<condition>` can link multiple conditions separated by space:

Conditions can be linked by logical AND:

  - `>1.0.0 <2.0.0` would match between both ranges, so `1.1.1` and `1.8.7` but not `1.0.0` or `2.0.0`
  - `>1.0.0 <3.0.0 !2.0.3-beta.2` would match every version between `1.0.0` and `3.0.0` except `2.0.3-beta.2`

Conditions can also be linked by logical OR:

  - `<2.0.0 || >=3.0.0` would match `1.x.x` and `3.x.x` but not `2.x.x`

AND has a higher precedence than OR. It's not possible to use brackets.

Conditions can be combined with both AND and OR

  - `>1.0.0 <2.0.0 || >3.0.0 !4.2.1` would match `1.2.3`, `1.9.9`, `3.1.1`, but not `4.2.1`, `2.1.1`

## Examples

```bash

$ semver ">=1.0.0" "1.0.2"
1.0.2 meets the condition.
$ echo $?
0

$ semver "<2.0.0 || >=3.0.0" "2.0.2"                                                                                        2.0.2 does not meet the condition!
$ echo $?
1

$ semver -q ">1.0.0 <2.0.0 || >3.0.0 !4.2.1" "1.9.9"
$ echo $?
0

```

## Changelog & Releases

This repository keeps a [change log](CHANGELOG.md). The format of the log
is based on [Keep a Changelog][keepchangelog].

Releases are based on [Semantic Versioning][semver], and use the format
of ``MAJOR.MINOR.PATCH``. In a nutshell, the version will be incremented
based on the following:

- ``MAJOR``: Incompatible or major changes.
- ``MINOR``: Backwards-compatible new features and enhancements.
- ``PATCH``: Backwards-compatible bugfixes and package updates.

## Support

Got questions?

You could [open an issue here][issue] GitHub.

## Contributing

This is an active open-source project. We are always open to people who want to
use the code or contribute to it.

We have set up a separate document containing our
[contribution guidelines](CONTRIBUTING.md).

Thank you for being involved! :heart_eyes:

## Authors & contributors

The original setup of this repository is by [Franck Nijhof][frenck].

For a full list of all authors and contributors,
check [the contributor's page][contributors].

## License

MIT License

Copyright (c) 2017 Franck Nijhof

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

[commits-shield]: https://img.shields.io/github/commit-activity/y/frenck/semver.svg
[commits]: https://github.com/frenck/semver/commits/master
[contributors]: https://github.com/frenck/semver/graphs/contributors
[frenck]: https://github.com/frenck
[issue]: https://github.com/frenck/semver/issues
[keepchangelog]: http://keepachangelog.com/en/1.0.0/
[license-shield]: https://img.shields.io/github/license/frenck/semver.svg
[maintenance-shield]: https://img.shields.io/maintenance/yes/2017.svg
[project-stage-shield]: https://img.shields.io/badge/project%20stage-production%20ready-brightgreen.svg
[releases-shield]: https://img.shields.io/github/release/frenck/semver.svg
[releases]: https://github.com/frenck/semver/releases
[semver]: http://semver.org/spec/v2.0.0.htm
