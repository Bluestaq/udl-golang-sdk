# Changelog

## 0.1.0-alpha.18 (2025-12-18)

Full Changelog: [v0.1.0-alpha.17...v0.1.0-alpha.18](https://github.com/Bluestaq/udl-golang-sdk/compare/v0.1.0-alpha.17...v0.1.0-alpha.18)

### Chores

* add float64 to valid types for RegisterFieldValidator ([1db9e63](https://github.com/Bluestaq/udl-golang-sdk/commit/1db9e63c0ca409d20c569c3af569adb3a09274f7))

## 0.1.0-alpha.17 (2025-12-11)

Full Changelog: [v0.1.0-alpha.16...v0.1.0-alpha.17](https://github.com/Bluestaq/udl-golang-sdk/compare/v0.1.0-alpha.16...v0.1.0-alpha.17)

### Features

* **api:** api update ([30354da](https://github.com/Bluestaq/udl-golang-sdk/commit/30354daaa86ec11988065cf73bbd3bc445ab44b2))
* **api:** bumps to v1.37.0 of UDL API ([8ef8ac2](https://github.com/Bluestaq/udl-golang-sdk/commit/8ef8ac2bc39f652d2dfae1cc8d9f93c1c2f06ecb))
* **encoder:** support bracket encoding form-data object members ([4b39232](https://github.com/Bluestaq/udl-golang-sdk/commit/4b3923288838c51425b7a122262a4a2c2c054ab1))


### Bug Fixes

* **mcp:** correct code tool API endpoint ([67c5a09](https://github.com/Bluestaq/udl-golang-sdk/commit/67c5a09d61249deff622be816a6d2386c7c8c18d))
* rename param to avoid collision ([f9c8f3a](https://github.com/Bluestaq/udl-golang-sdk/commit/f9c8f3af417db61f7f0af1165937e6e076be3b7b))


### Chores

* elide duplicate aliases ([4e39bb9](https://github.com/Bluestaq/udl-golang-sdk/commit/4e39bb9119e5429f5f24e794159bfe3bb24713e9))
* **internal:** codegen related update ([fdb8e02](https://github.com/Bluestaq/udl-golang-sdk/commit/fdb8e027be08c13606da139897b511738d8a0173))

## 0.1.0-alpha.16 (2025-11-18)

Full Changelog: [v0.1.0-alpha.15...v0.1.0-alpha.16](https://github.com/Bluestaq/udl-golang-sdk/compare/v0.1.0-alpha.15...v0.1.0-alpha.16)

### Bug Fixes

* **client:** correctly specify Accept header with */* instead of empty ([1be215e](https://github.com/Bluestaq/udl-golang-sdk/commit/1be215ebd359cdadd451baa491d7b55f70031748))

## 0.1.0-alpha.15 (2025-11-11)

Full Changelog: [v0.1.0-alpha.14...v0.1.0-alpha.15](https://github.com/Bluestaq/udl-golang-sdk/compare/v0.1.0-alpha.14...v0.1.0-alpha.15)

### Chores

* bump gjson version ([c858fdc](https://github.com/Bluestaq/udl-golang-sdk/commit/c858fdc6b87f6bedac0ad1ff3b6bf3f13faffa03))

## 0.1.0-alpha.14 (2025-11-07)

Full Changelog: [v0.1.0-alpha.13...v0.1.0-alpha.14](https://github.com/Bluestaq/udl-golang-sdk/compare/v0.1.0-alpha.13...v0.1.0-alpha.14)

### Bug Fixes

* remove readonly parameters from request params ([0eb48d8](https://github.com/Bluestaq/udl-golang-sdk/commit/0eb48d8ddae231cf79a42c53f82016514074d7fd))


### Chores

* **internal:** grammar fix (it's -&gt; its) ([44aa3e3](https://github.com/Bluestaq/udl-golang-sdk/commit/44aa3e3030a8f74a4b4fe42d4b8a368540d6b28d))

## 0.1.0-alpha.13 (2025-10-21)

Full Changelog: [v0.1.0-alpha.12...v0.1.0-alpha.13](https://github.com/Bluestaq/udl-golang-sdk/compare/v0.1.0-alpha.12...v0.1.0-alpha.13)

### Features

* **api:** api update ([6f8407c](https://github.com/Bluestaq/udl-golang-sdk/commit/6f8407cff01b8fa9c9e24db6eae7d81e48de50e2))
* **api:** manual updates ([5d27149](https://github.com/Bluestaq/udl-golang-sdk/commit/5d271492fe2ad7f905488bbe409de60f84459324))

## 0.1.0-alpha.12 (2025-09-25)

Full Changelog: [v0.1.0-alpha.11...v0.1.0-alpha.12](https://github.com/Bluestaq/udl-golang-sdk/compare/v0.1.0-alpha.11...v0.1.0-alpha.12)

### Features

* **api:** adding obs correlation, staging data for emitters, and user auth endpoint ([05a59c2](https://github.com/Bluestaq/udl-golang-sdk/commit/05a59c28b067f4489d7f3e03c84b5a64bdcdf36a))
* **api:** api update ([57db327](https://github.com/Bluestaq/udl-golang-sdk/commit/57db32778248642a92a5cc9d080e2f2b7d0303eb))


### Bug Fixes

* bugfix for setting JSON keys with special characters ([5d98629](https://github.com/Bluestaq/udl-golang-sdk/commit/5d9862950fae20e61f5b78e83e4d071fd456bf53))

## 0.1.0-alpha.11 (2025-09-19)

Full Changelog: [v0.1.0-alpha.10...v0.1.0-alpha.11](https://github.com/Bluestaq/udl-golang-sdk/compare/v0.1.0-alpha.10...v0.1.0-alpha.11)

### Bug Fixes

* use slices.Concat instead of sometimes modifying r.Options ([fd9196a](https://github.com/Bluestaq/udl-golang-sdk/commit/fd9196a3247b8a66ff98d6bfc1347b9b97d44a8d))


### Chores

* bump minimum go version to 1.22 ([929c908](https://github.com/Bluestaq/udl-golang-sdk/commit/929c90874ca609546432e936195a2f3157c5a2be))
* do not install brew dependencies in ./scripts/bootstrap by default ([636ac09](https://github.com/Bluestaq/udl-golang-sdk/commit/636ac09ecc10fe50bf559797923899998c975036))
* update more docs for 1.22 ([55d41e6](https://github.com/Bluestaq/udl-golang-sdk/commit/55d41e632751b15f588914ffc3ea04fc4c408a76))

## 0.1.0-alpha.10 (2025-09-17)

Full Changelog: [v0.1.0-alpha.9...v0.1.0-alpha.10](https://github.com/Bluestaq/udl-golang-sdk/compare/v0.1.0-alpha.9...v0.1.0-alpha.10)

### Features

* **api:** removing old routes ([e97b439](https://github.com/Bluestaq/udl-golang-sdk/commit/e97b439ca5749730b4e332ea89496bd5249e3a8e))
* **api:** Support for latest UDL release ([246c7fd](https://github.com/Bluestaq/udl-golang-sdk/commit/246c7fd715f8dfa963cfb7f251c4981b8418e36e))

## 0.1.0-alpha.9 (2025-09-11)

Full Changelog: [v0.1.0-alpha.8...v0.1.0-alpha.9](https://github.com/Bluestaq/udl-golang-sdk/compare/v0.1.0-alpha.8...v0.1.0-alpha.9)

### Features

* **api:** api update ([f68e56e](https://github.com/Bluestaq/udl-golang-sdk/commit/f68e56ec18673cdbb794771fb8bca87c8990554a))
* **api:** api update ([ffef04f](https://github.com/Bluestaq/udl-golang-sdk/commit/ffef04fd570ab0f6873d580662608c59d81a13df))
* **api:** manual updates ([f7c0d79](https://github.com/Bluestaq/udl-golang-sdk/commit/f7c0d79ffe858cce4146b299b8ead6de5ad4368e))


### Bug Fixes

* **internal:** unmarshal correctly when there are multiple discriminators ([7c3970f](https://github.com/Bluestaq/udl-golang-sdk/commit/7c3970f457d11f6d4b8cd03df423853c9cd1ca43))
* remove null from release please manifest ([2157769](https://github.com/Bluestaq/udl-golang-sdk/commit/2157769bd435128b6b33a480658b940f3943e216))
* use release please annotations on more places ([e563f9f](https://github.com/Bluestaq/udl-golang-sdk/commit/e563f9fd246d855a17a133db247b7306a87e1c5b))

## 0.1.0-alpha.8 (2025-09-02)

Full Changelog: [v0.1.0-alpha.7...v0.1.0-alpha.8](https://github.com/Bluestaq/udl-golang-sdk/compare/v0.1.0-alpha.7...v0.1.0-alpha.8)

## 0.1.0-alpha.7 (2025-08-28)

Full Changelog: [v0.1.0-alpha.6...v0.1.0-alpha.7](https://github.com/Bluestaq/udl-golang-sdk/compare/v0.1.0-alpha.6...v0.1.0-alpha.7)

### Bug Fixes

* close body before retrying ([b730154](https://github.com/Bluestaq/udl-golang-sdk/commit/b7301543c09aa212647a7a036d492c5e8b318996))

## 0.1.0-alpha.6 (2025-08-12)

Full Changelog: [v0.1.0-alpha.5...v0.1.0-alpha.6](https://github.com/Bluestaq/udl-golang-sdk/compare/v0.1.0-alpha.5...v0.1.0-alpha.6)

### Features

* **api:** manual updates ([a02e95c](https://github.com/Bluestaq/udl-golang-sdk/commit/a02e95c9f1c651869410b9f00689aeb668a05ec8))
* **api:** manual updates ([0413b89](https://github.com/Bluestaq/udl-golang-sdk/commit/0413b89ef74257fc23c939f34d33bf0bf672d2a4))
* **api:** manual updates ([159d4ab](https://github.com/Bluestaq/udl-golang-sdk/commit/159d4abe6f2cc33ad077e95067fd52502027bb06))
* **api:** manual updates ([cd8cfd1](https://github.com/Bluestaq/udl-golang-sdk/commit/cd8cfd129fd3e04afbd5683e7f5f0b16edad0b2f))
* **client:** support optional json html escaping ([4af63e3](https://github.com/Bluestaq/udl-golang-sdk/commit/4af63e316c5ce62802ffed56f751a52b67757d2c))


### Chores

* **internal:** update comment in script ([eac9138](https://github.com/Bluestaq/udl-golang-sdk/commit/eac9138c5aa0f79996cf2e5cedce9d383fb4fb67))
* update @stainless-api/prism-cli to v5.15.0 ([68a4368](https://github.com/Bluestaq/udl-golang-sdk/commit/68a43685cdbe00f4ecbdd9390d7dc4136b4a7056))

## 0.1.0-alpha.5 (2025-08-03)

Full Changelog: [v0.1.0-alpha.4...v0.1.0-alpha.5](https://github.com/Bluestaq/udl-golang-sdk/compare/v0.1.0-alpha.4...v0.1.0-alpha.5)

## 0.1.0-alpha.4 (2025-08-03)

Full Changelog: [v0.1.0-alpha.3...v0.1.0-alpha.4](https://github.com/Bluestaq/udl-golang-sdk/compare/v0.1.0-alpha.3...v0.1.0-alpha.4)

### Features

* **api:** support for Bearer Auth ([1a9c45c](https://github.com/Bluestaq/udl-golang-sdk/commit/1a9c45c4ffbe51fe29afae641b1f8270e323c805))

## 0.1.0-alpha.3 (2025-07-31)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/Bluestaq/udl-golang-sdk/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### ⚠ BREAKING CHANGES

* **api:** move multiple models to shared resources

### Features

* add Kafka offset pagination ([8923e40](https://github.com/Bluestaq/udl-golang-sdk/commit/8923e40b7e850d98f6d9ad630ded627684485b49))
* **api:** define additional models to reduce duplication ([d96e6a9](https://github.com/Bluestaq/udl-golang-sdk/commit/d96e6a9efa3fe57ece96d6be3977a1ccc7e3a8f8))
* **api:** manual changes ([df5d954](https://github.com/Bluestaq/udl-golang-sdk/commit/df5d95477dfd333bfce06019f056d2962f45ff9a))
* **client:** support file upload requests ([0b77eba](https://github.com/Bluestaq/udl-golang-sdk/commit/0b77eba67b960fbf10f16f0dc31edb66e1a8fd19))


### Chores

* **api:** move multiple models to shared resources ([9e0e66f](https://github.com/Bluestaq/udl-golang-sdk/commit/9e0e66f6e8cc034bc6ff49745dd65ccaa7a4eb4d))

## 0.1.0-alpha.2 (2025-07-22)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/Bluestaq/udl-golang-sdk/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Bug Fixes

* **client:** process custom base url ahead of time ([3741953](https://github.com/Bluestaq/udl-golang-sdk/commit/374195331572d71197a2b9a28d1335bbb08ab1c8))

## 0.1.0-alpha.1 (2025-07-16)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/Bluestaq/udl-golang-sdk/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### ⚠ BREAKING CHANGES

* **client:** rename resp package
* **client:** improve core function names

### Features

* **api:** api update ([cd6e067](https://github.com/Bluestaq/udl-golang-sdk/commit/cd6e067c3618271d427acafb27d353cc2e2b38a0))
* **api:** manual updates ([96f0d7a](https://github.com/Bluestaq/udl-golang-sdk/commit/96f0d7a3dce384a2f0fb2cd796e9a2eca358743e))
* **api:** manual updates ([d1c3cab](https://github.com/Bluestaq/udl-golang-sdk/commit/d1c3cab588537e9b877afdcbc064c3da927b1025))
* **client:** add debug log helper ([eebc984](https://github.com/Bluestaq/udl-golang-sdk/commit/eebc9842651de27442217cbbd671ee03c55170e1))
* **client:** add escape hatch for null slice & maps ([3b298a3](https://github.com/Bluestaq/udl-golang-sdk/commit/3b298a3b55e596707f2511476887f2f7f65bdcbc))
* **client:** add support for endpoint-specific base URLs in python ([aba3aeb](https://github.com/Bluestaq/udl-golang-sdk/commit/aba3aebfde6b1ca7db69ad02591276f0ac22e971))
* **client:** allow overriding unions ([51dd90b](https://github.com/Bluestaq/udl-golang-sdk/commit/51dd90b72b9816145f4720ba267b40fd700f2628))
* **client:** experimental support for unmarshalling into param structs ([2992428](https://github.com/Bluestaq/udl-golang-sdk/commit/2992428e98873a78d94fc71846639349d1fb3b76))
* **client:** rename resp package ([133833a](https://github.com/Bluestaq/udl-golang-sdk/commit/133833a8868ea8f43a04397b7adf27b44a3fdd5f))


### Bug Fixes

* **client:** cast to raw message when converting to params ([cd63183](https://github.com/Bluestaq/udl-golang-sdk/commit/cd63183a104b347b1612213d0904831c802b11cd))
* **client:** clean up reader resources ([2f5c8a9](https://github.com/Bluestaq/udl-golang-sdk/commit/2f5c8a99ee8642499f653f09ade42a42b0d4715a))
* **client:** correctly set stream key for multipart ([470d3ef](https://github.com/Bluestaq/udl-golang-sdk/commit/470d3ef6cdbf3a9b01a3de0f7f26a7253597afd6))
* **client:** correctly update body in WithJSONSet ([77b27e0](https://github.com/Bluestaq/udl-golang-sdk/commit/77b27e035059138f564c6f833f13f2021c03aa84))
* **client:** don't panic on marshal with extra null field ([26c5d03](https://github.com/Bluestaq/udl-golang-sdk/commit/26c5d03352e63ed9f836e9956b99e5fa40fc0eba))
* **client:** improve core function names ([08a5222](https://github.com/Bluestaq/udl-golang-sdk/commit/08a522266a6b533a6cb266c09050ad04b1500b71))
* **client:** unmarshal responses properly ([adb4e10](https://github.com/Bluestaq/udl-golang-sdk/commit/adb4e10d8217781f68dbe773bdfc8e3e4e5566b8))
* correct unmarshalling of root body params ([db54543](https://github.com/Bluestaq/udl-golang-sdk/commit/db54543dbca67a9210bba546fbe553f62b6de3f2))
* don't try to deserialize as json when ResponseBodyInto is []byte ([deef2d8](https://github.com/Bluestaq/udl-golang-sdk/commit/deef2d8eb372e2ab7545c41c6604c283fe04db94))
* fix error ([88dc228](https://github.com/Bluestaq/udl-golang-sdk/commit/88dc2281d1cacd5fd5fc32ee4299007825de9311))
* **pagination:** check if page data is empty in GetNextPage ([337114b](https://github.com/Bluestaq/udl-golang-sdk/commit/337114bef1d71acb097055284d15e8ee87bfc29a))


### Chores

* **ci:** enable for pull requests ([14e3543](https://github.com/Bluestaq/udl-golang-sdk/commit/14e35438bed784d5bb1bd9de8517ab2c4955fd9b))
* **ci:** only run for pushes and fork pull requests ([36256d1](https://github.com/Bluestaq/udl-golang-sdk/commit/36256d1f3fd4f275865e64773533d7a8bf92995b))
* configure new SDK language ([d19bdba](https://github.com/Bluestaq/udl-golang-sdk/commit/d19bdbad643838a1de56f7c32c028b9ceb77468f))
* **docs:** grammar improvements ([253bcc9](https://github.com/Bluestaq/udl-golang-sdk/commit/253bcc9834b5a818f75716dcc4a7657035424000))
* **docs:** update respjson package name ([e454fe1](https://github.com/Bluestaq/udl-golang-sdk/commit/e454fe14663a0b0874b868df40242bb5a6e59686))
* fix documentation of null map ([a882dc0](https://github.com/Bluestaq/udl-golang-sdk/commit/a882dc0fe302c35daf1ccd4a2097384cd034b912))
* improve devcontainer setup ([00ecdb6](https://github.com/Bluestaq/udl-golang-sdk/commit/00ecdb60faf0716da8bdccc33e4ae6dba3d081c6))
* lint tests in subpackages ([4baa91e](https://github.com/Bluestaq/udl-golang-sdk/commit/4baa91ee261201eb23827ce36720d964c4356d5c))
* make go mod tidy continue on error ([9c91464](https://github.com/Bluestaq/udl-golang-sdk/commit/9c91464273466dede8fdd8f343fb0d97b1ca247b))
* update SDK settings ([c4a69d1](https://github.com/Bluestaq/udl-golang-sdk/commit/c4a69d11ce76591f85d940271d914329ade98076))
* update SDK settings ([c431c23](https://github.com/Bluestaq/udl-golang-sdk/commit/c431c23f7b9c2fb5de8acc753676e1be6fd60054))
