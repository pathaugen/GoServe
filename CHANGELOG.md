
# GoServe! Change Log

**ATTN**: This project uses [semantic versioning](http://semver.org/).

***

## 1.0.0 - YYYY-MM-DD

### Added

* x

### Changed

* x

### Fixed

* go-thrust issue that needed fixing:
  * Original Project: <https://github.com/miketheprogrammer/go-thrust>
  * PatHaugen Fork: <https://github.com/pathaugen/go-thrust>
  * Commit and fix: <https://github.com/miketheprogrammer/go-thrust/pull/79>

* External code with unit test errors need to be branched and fixed:
* birpc:
  * wetsock_test.go:84: websocket client failed to start: malformed ws or wss URL
  * TestWSArg issue
* cli:
  * app_test.go:
  * TestApp_BeforeFunc fail
  * TestApp_OrderOfOperations incorrect usage
  * TestApp_Run_DoesNotOverwriteErrorFromBefore before error
  * TestCommand_Run_DoesNotOverwriteErrorFromBefore before error
* pb:
  * profile_test.go failures
* profile:
  * fail
* yaml-2:
  * Using a non-standard test platform. Can branch and migrate over.

***
