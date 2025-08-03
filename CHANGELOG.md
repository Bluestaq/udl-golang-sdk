# Changelog

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
