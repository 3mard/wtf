
[![All Contributors](https://img.shields.io/badge/all_contributors-60-orange.svg?style=flat-square)](#contributors)
[![All Contributors](https://img.shields.io/badge/all_contributors-53-orange.svg?style=flat-square)](#contributors)
[![Build Status](https://travis-ci.com/wtfutil/wtf.svg?branch=master)](https://travis-ci.com/wtfutil/wtf)
[![Gitter Chat](https://badges.gitter.im/wtfutil/Lobby.svg)](https://gitter.im/wtfutil/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Twitter](https://img.shields.io/badge/follow-on%20twitter-blue.svg)](https://twitter.com/wtfutil)
[![Go Report Card](https://goreportcard.com/badge/github.com/wtfutil/wtf)](https://goreportcard.com/report/github.com/wtfutil/wtf)

# WTF

A personal terminal-based dashboard utility, designed for
displaying infrequently-needed, but very important, daily data.

Follow [on Twitter](https://twitter.com/wtfutil) for news and latest updates.

<p align="center">
<img src="./images/screenshot.jpg" title="screenshot" width="720" height="420" />
</p>

## Quick Start

[Download and run the latest binary](https://github.com/wtfutil/wtf/releases) or install from source:

```bash
go get -u github.com/wtfutil/wtf
cd $GOPATH/src/github.com/wtfutil/wtf
make install
make run
```

**Note:** WTF is _only_ compatible with Go versions **1.9.2** or later. It currently _does not_ compile with `gccgo`.

## Documentation

See [https://wtfutil.com](https://wtfutil.com) for the definitive
documentation. Here's some short-cuts:

* [Installation](http://wtfutil.com/posts/installation/)
* [Configuration](http://wtfutil.com/posts/configuration/)
* [Module Documentation](http://wtfutil.com/posts/modules/)

### Contributing to the Documentation

Documenation now lives in its own repository here: [https://github.com/wtfutil/wtfdocs](https://github.com/wtfutil/wtfdocs).

Please make all additions and updates to documentation in that repository.

## Contributing to the Source Code

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests.

### Adding Dependencies

Dependency management in WTF is handled by [dep](https://github.com/golang/dep). See that page for installation and usage details.

If the work you're doing requires the addition of a new dependency,
please be sure to use `dep` to [vendor your dependencies](https://golang.github.io/dep/docs/daily-dep.html#adding-a-new-dependency).

## Contributors

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore -->
| [<img src="https://avatars0.githubusercontent.com/u/6413?v=4" width="48px;" alt="Chris Cummer"/><br /><sub><b>Chris Cummer</b></sub>](https://twitter.com/senorprogrammer)<br /> | [<img src="https://avatars2.githubusercontent.com/u/3252403?v=4" width="48px;" alt="Anand Sudhir Prayaga"/><br /><sub><b>Anand Sudhir Prayaga</b></sub>](https://github.com/anandsudhir)<br /> | [<img src="https://avatars1.githubusercontent.com/u/34973359?v=4" width="48px;" alt="Hossein Mehrabi"/><br /><sub><b>Hossein Mehrabi</b></sub>](https://github.com/jeangovil)<br /> | [<img src="https://avatars0.githubusercontent.com/u/11779018?v=4" width="48px;" alt="FengYa"/><br /><sub><b>FengYa</b></sub>](https://github.com/Fengyalv)<br /> | [<img src="https://avatars2.githubusercontent.com/u/17337753?v=4" width="48px;" alt="deltax"/><br /><sub><b>deltax</b></sub>](https://fluxionnetwork.github.io/fluxion/)<br /> | [<img src="https://avatars0.githubusercontent.com/u/1319630?v=4" width="48px;" alt="Bill Keenan"/><br /><sub><b>Bill Keenan</b></sub>](https://github.com/BillKeenan)<br /> | [<img src="https://avatars2.githubusercontent.com/u/118081?v=4" width="48px;" alt="June S"/><br /><sub><b>June S</b></sub>](http://blog.sapara.com)<br /> |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| [<img src="https://avatars3.githubusercontent.com/u/16461061?v=4" width="48px;" alt="liyiheng"/><br /><sub><b>liyiheng</b></sub>](https://github.com/XanthusL)<br /> | [<img src="https://avatars2.githubusercontent.com/u/9014288?v=4" width="48px;" alt="baustinanki"/><br /><sub><b>baustinanki</b></sub>](https://github.com/baustinanki)<br /> | [<img src="https://avatars0.githubusercontent.com/u/371475?v=4" width="48px;" alt="lucus lee"/><br /><sub><b>lucus lee</b></sub>](https://github.com/lixin9311)<br /> | [<img src="https://avatars1.githubusercontent.com/u/7537841?v=4" width="48px;" alt="Mike Lloyd"/><br /><sub><b>Mike Lloyd</b></sub>](https://github.com/mxplusb)<br /> | [<img src="https://avatars3.githubusercontent.com/u/10998?v=4" width="48px;" alt="Sergio Rubio"/><br /><sub><b>Sergio Rubio</b></sub>](http://rubiojr.rbel.co)<br /> | [<img src="https://avatars3.githubusercontent.com/u/17374492?v=4" width="48px;" alt="Farhad Farahi"/><br /><sub><b>Farhad Farahi</b></sub>](https://github.com/FarhadF)<br /> | [<img src="https://avatars1.githubusercontent.com/u/634604?v=4" width="48px;" alt="Lasantha Kularatne"/><br /><sub><b>Lasantha Kularatne</b></sub>](http://lasantha.blogspot.com/)<br /> |
| [<img src="https://avatars1.githubusercontent.com/u/823331?v=4" width="48px;" alt="Mark Old"/><br /><sub><b>Mark Old</b></sub>](https://github.com/dlom)<br /> | [<img src="https://avatars0.githubusercontent.com/u/5546718?v=4" width="48px;" alt="flw"/><br /><sub><b>flw</b></sub>](http://flw.tools/)<br /> | [<img src="https://avatars0.githubusercontent.com/u/6024927?v=4" width="48px;" alt="David Barda"/><br /><sub><b>David Barda</b></sub>](https://github.com/davebarda)<br /> | [<img src="https://avatars2.githubusercontent.com/u/4261980?v=4" width="48px;" alt="Geoff Lee"/><br /><sub><b>Geoff Lee</b></sub>](https://github.com/matrinox)<br /> | [<img src="https://avatars3.githubusercontent.com/u/1022918?v=4" width="48px;" alt="George Opritescu"/><br /><sub><b>George Opritescu</b></sub>](http://international.github.io)<br /> | [<img src="https://avatars3.githubusercontent.com/u/497310?v=4" width="48px;" alt="Grazfather"/><br /><sub><b>Grazfather</b></sub>](https://twitter.com/Grazfather)<br /> | [<img src="https://avatars2.githubusercontent.com/u/1691120?v=4" width="48px;" alt="Michael Cordell"/><br /><sub><b>Michael Cordell</b></sub>](http://www.mikecordell.com/)<br /> |
| [<img src="https://avatars2.githubusercontent.com/u/1215497?v=4" width="48px;" alt="Patrick José Pereira"/><br /><sub><b>Patrick José Pereira</b></sub>](http://patrick.ibexcps.com)<br /> | [<img src="https://avatars2.githubusercontent.com/u/1483092?v=4" width="48px;" alt="sherod taylor"/><br /><sub><b>sherod taylor</b></sub>](https://github.com/sherodtaylor)<br /> | [<img src="https://avatars2.githubusercontent.com/u/3062663?v=4" width="48px;" alt="Andrew Scott"/><br /><sub><b>Andrew Scott</b></sub>](http://cogentia.io)<br /> | [<img src="https://avatars1.githubusercontent.com/u/12018440?v=4" width="48px;" alt="Lassi Piironen"/><br /><sub><b>Lassi Piironen</b></sub>](https://github.com/lsipii)<br /> | [<img src="https://avatars0.githubusercontent.com/u/14799210?v=4" width="48px;" alt="BlackWebWolf"/><br /><sub><b>BlackWebWolf</b></sub>](https://github.com/BlackWebWolf)<br /> | [<img src="https://avatars0.githubusercontent.com/u/1894885?v=4" width="48px;" alt="andrewzolotukhin"/><br /><sub><b>andrewzolotukhin</b></sub>](https://github.com/andrewzolotukhin)<br /> | [<img src="https://avatars1.githubusercontent.com/u/8568280?v=4" width="48px;" alt="Leon Stigter"/><br /><sub><b>Leon Stigter</b></sub>](https://retgits.github.io)<br /> |
| [<img src="https://avatars3.githubusercontent.com/u/21756?v=4" width="48px;" alt="Amr Tamimi"/><br /><sub><b>Amr Tamimi</b></sub>](https://tamimi.se)<br /> | [<img src="https://avatars3.githubusercontent.com/u/3717137?v=4" width="48px;" alt="Jagdeep Singh"/><br /><sub><b>Jagdeep Singh</b></sub>](https://jagdeep.me)<br /> | [<img src="https://avatars0.githubusercontent.com/u/889171?v=4" width="48px;" alt="Lineu Felipe"/><br /><sub><b>Lineu Felipe</b></sub>](https://github.com/darkSasori)<br /> | [<img src="https://avatars2.githubusercontent.com/u/159124?v=4" width="48px;" alt="Konstantin"/><br /><sub><b>Konstantin</b></sub>](https://github.com/kvj)<br /> | [<img src="https://avatars2.githubusercontent.com/u/6044920?v=4" width="48px;" alt="Brendan O'Leary"/><br /><sub><b>Brendan O'Leary</b></sub>](http://www.brendanoleary.com)<br /> | [<img src="https://avatars2.githubusercontent.com/u/1226441?v=4" width="48px;" alt="bertl4398"/><br /><sub><b>bertl4398</b></sub>](https://github.com/bertl4398)<br /> | [<img src="https://avatars2.githubusercontent.com/u/6553695?v=4" width="48px;" alt="Ferenc-"/><br /><sub><b>Ferenc-</b></sub>](https://github.com/Ferenc-)<br /> |
| [<img src="https://avatars1.githubusercontent.com/u/952036?v=4" width="48px;" alt="Rohan Verma"/><br /><sub><b>Rohan Verma</b></sub>](http://rohanverma.net)<br /> | [<img src="https://avatars1.githubusercontent.com/u/19293566?v=4" width="48px;" alt="Tim Fitzgerald"/><br /><sub><b>Tim Fitzgerald</b></sub>](https://github.com/fimtitzgerald)<br /> | [<img src="https://avatars2.githubusercontent.com/u/1081051?v=4" width="48px;" alt="Federico Ruggi"/><br /><sub><b>Federico Ruggi</b></sub>](https://github.com/ruggi)<br /> | [<img src="https://avatars2.githubusercontent.com/u/7293328?v=4" width="48px;" alt="Craig Woodward"/><br /><sub><b>Craig Woodward</b></sub>](https://github.com/ctwoodward)<br /> | [<img src="https://avatars3.githubusercontent.com/u/15367484?v=4" width="48px;" alt="ReadmeCritic"/><br /><sub><b>ReadmeCritic</b></sub>](https://twitter.com/ReadmeCritic)<br /> | [<img src="https://avatars0.githubusercontent.com/u/141402?v=4" width="48px;" alt="Eugene"/><br /><sub><b>Eugene</b></sub>](https://github.com/jdevelop)<br /> | [<img src="https://avatars1.githubusercontent.com/u/12983705?s=460&v=4" width="48px;" alt="Kenny Wu"/><br /><sub><b>Kenny Wu</b></sub>](https://github.com/Trinergy)<br /> |
| [<img src="https://avatars0.githubusercontent.com/u/538234?v=4" width="48px;" alt="Renán Romero"/><br /><sub><b>Renán Romero</b></sub>](http://www.romeroruiz.com)<br /> | [<img src="https://avatars1.githubusercontent.com/u/5031240?v=4" width="48px;" alt="Bastian Groß"/><br /><sub><b>Bastian Groß</b></sub>](https://github.com/sticreations)<br /> | [<img src="https://avatars1.githubusercontent.com/u/2496835?v=4" width="48px;" alt="nicholas-eden"/><br /><sub><b>nicholas-eden</b></sub>](https://github.com/nicholas-eden)<br /> | [<img src="https://avatars1.githubusercontent.com/u/279390?v=4" width="48px;" alt="Dan Rabinowitz"/><br /><sub><b>Dan Rabinowitz</b></sub>](https://github.com/danrabinowitz)<br /> | [<img src="https://avatars1.githubusercontent.com/u/6897575?v=4" width="48px;" alt="David Missmann"/><br /><sub><b>David Missmann</b></sub>](https://github.com/dvdmssmnn)<br /> | [<img src="https://avatars2.githubusercontent.com/u/882006?v=4" width="48px;" alt="Mathias Weber"/><br /><sub><b>Mathias Weber</b></sub>](https://github.com/mweb)<br /> | [<img src="https://avatars1.githubusercontent.com/u/32081703?v=4" width="48px;" alt="TheRedSpy15"/><br /><sub><b>TheRedSpy15</b></sub>](https://github.com/TheRedSpy15)<br /> |
| [<img src="https://avatars0.githubusercontent.com/u/9569897?v=4" width="48px;" alt="Harald Nordgren"/><br /><sub><b>Harald Nordgren</b></sub>](https://www.linkedin.com/in/harald-nordgren-44778192)<br /> | [<img src="https://avatars0.githubusercontent.com/u/11583824?v=4" width="48px;" alt="Matei Alexandru Gardus"/><br /><sub><b>Matei Alexandru Gardus</b></sub>](http://stormfirefox1.github.io)<br /> | [<img src="https://avatars2.githubusercontent.com/u/1523955?v=4" width="48px;" alt="Sean Smith"/><br /><sub><b>Sean Smith</b></sub>](https://github.com/Seanstoppable)<br /> | [<img src="https://avatars1.githubusercontent.com/u/1646238?v=4" width="48px;" alt="Halil Kaskavalci"/><br /><sub><b>Halil Kaskavalci</b></sub>](http://kaskavalci.com)<br /> | [<img src="https://avatars2.githubusercontent.com/u/246715?v=4" width="48px;" alt="Johan Denoyer"/><br /><sub><b>Johan Denoyer</b></sub>](http://www.johandenoyer.fr)<br /> | [<img src="https://avatars1.githubusercontent.com/u/593516?v=4" width="48px;" alt="Jelle Vink"/><br /><sub><b>Jelle Vink</b></sub>](https://skymeyer.be)<br /> | [<img src="https://avatars1.githubusercontent.com/u/3997333?v=4" width="48px;" alt="Devin Collins"/><br /><sub><b>Devin Collins</b></sub>](http://imdevinc.com)<br /> |
| [<img src="https://avatars3.githubusercontent.com/u/18333?v=4" width="48px;" alt="Danne Stayskal"/><br /><sub><b>Danne Stayskal</b></sub>](http://danne.stayskal.com/)<br /> | [<img src="https://avatars1.githubusercontent.com/u/2006658?v=4" width="48px;" alt="Max Beizer"/><br /><sub><b>Max Beizer</b></sub>](https://www.maxbeizer.com)<br /> | [<img src="https://avatars1.githubusercontent.com/u/194392?v=4" width="48px;" alt="E:V:A"/><br /><sub><b>E:V:A</b></sub>](http://tinyurl.com/nwmj4as)<br /> | [<img src="https://avatars0.githubusercontent.com/u/1425500?v=4" width="48px;" alt="Gabriel"/><br /><sub><b>Gabriel</b></sub>](https://github.com/GaboFDC)<br /> |
<!-- ALL-CONTRIBUTORS-LIST:END -->

## Acknowledgments

The inspiration for `WTF` came from Monica Dinculescu's
[tiny-care-terminal](https://github.com/notwaldorf/tiny-care-terminal).

Many thanks to <a href="https://lendesk.com">Lendesk</a> for supporting this project by
providing time to develop it.

<p align="center">
<img src="./images/dude_wtf.png?raw=true" title="Dude WTF" width="251" height="201" />
</p>
