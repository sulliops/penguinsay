# penguinsay
Penguinsay is a simple command-line utility that prints the Linux mascot, Tux, saying a variety of facts about Linux. The package, however, includes a malicious program that sends the results of a system vulnerability report to an external listening server.

Penguinsay and its associated program, vulnreport, were created for CPSC 4240 at Clemson University in the Fall 2024 semester.

## Components
The penguinsay package is made up of a few different programs and scripts:

1. [penguinsay](penguinsay/) — a series of Python scripts (made executable by a [wrapper script](penguinsay_1.0-1/usr/bin/penguinsay)) that prints facts about Linux coming from the Linux mascot, Tux
2. [vulnreport](vulnreport/) — a pair of Go programs (client and server) that send and receive the results of a system vulnerability report
3. [linux-exploit-suggester](https://github.com/The-Z-Labs/linux-exploit-suggester) — a third-party, open-source system vulnerability detection and reporting script
4. [Debian package scripts](penguinsay_1.0-1/DEBIAN/) — a series of scripts for the Debian package format that enable the installation of a system vulnerability reporter
5. [vulnreport.service](penguinsay_1.0-1/lib/systemd/system/vulnreport.service) — a daemon which runs the vulnreport client binary every time the system boots

## Usage
The penguinsay package is made available in the Debian package format and can be downloaded from the [Releases](https://github.com/sulliops/penguinsay/releases) page. Then, install like so (may require sudo privileges):

```
dpkg -i penguinsay_1.0-1.deb
```

Then, run penguinsay like so:

```
penguinsay
```

--

The [makefile](makefile) includes a number of targets for compiling vulnreport (requires [installing Go](https://go.dev/doc/install)), running the vulnreport server, and packaging the Debian package:

```
make vrclient
make vrserver
make package
make run
```