# Changelog

All notable changes to this project will be documented in this file. See [standard-version](https://github.com/conventional-changelog/standard-version) for commit guidelines.

## [0.12.0](https://github.com/keptn/keptn/compare/0.11.0...0.12.0) (2022-01-17)


### Features

*  Added problem title in sequence endpoint ([#4206](https://github.com/keptn/keptn/issues/4206)) ([#6258](https://github.com/keptn/keptn/issues/6258)) ([130f3d4](https://github.com/keptn/keptn/commit/130f3d4b3ff886716e4caea790644844f5fd86c4))
* Add helm value for configuration-service version selection ([#6387](https://github.com/keptn/keptn/issues/6387)) ([6e85346](https://github.com/keptn/keptn/commit/6e853469d18e8d6e55964fd60361fb8a030b3b07))
* Add warning for missing upstream ([#6433](https://github.com/keptn/keptn/issues/6433)) ([7a25226](https://github.com/keptn/keptn/commit/7a2522682ce7fd06985a9c53c503e67fa4874d9f))
* Added unit tests for webhook-service and improved error messages ([#6064](https://github.com/keptn/keptn/issues/6064)) ([5b4516e](https://github.com/keptn/keptn/commit/5b4516e7835b02757499c36f946f31ebf3ff5653))
* **bridge:** Add event picker to webhook configuration ([#5260](https://github.com/keptn/keptn/issues/5260)) ([a3c30cc](https://github.com/keptn/keptn/commit/a3c30ccb978d47890dc9d0c7caf706e5cad69f65))
* **bridge:** Add hint for tree list select component and change texts ([#5576](https://github.com/keptn/keptn/issues/5576)) ([0707ff7](https://github.com/keptn/keptn/commit/0707ff7dce5da6fd748985627623a011f3e66a42))
* **bridge:** Add hint that secrets are shared among all projects ([#5832](https://github.com/keptn/keptn/issues/5832)) ([0e40acf](https://github.com/keptn/keptn/commit/0e40acf0ea19ca1912cfcc12cd0008c2b02a527f))
* **bridge:** Add validation for payload if it contains specific characters ([#5950](https://github.com/keptn/keptn/issues/5950)) ([5d5b388](https://github.com/keptn/keptn/commit/5d5b3880742eef9eda49ec7769f14d7e921c397a))
* **bridge:** Consider abort state of sequence ([#6215](https://github.com/keptn/keptn/issues/6215)) ([f86f2f2](https://github.com/keptn/keptn/commit/f86f2f2fa920f7a7dacbcf16fb4f1e6d607d2ba9))
* **bridge:** Feature flags for bridge server ([#6073](https://github.com/keptn/keptn/issues/6073)) ([6351a58](https://github.com/keptn/keptn/commit/6351a58213db91bbe1fff551d8131218af8d961b))
* **bridge:** Introduce keptn color scheme [#5081](https://github.com/keptn/keptn/issues/5081) ([#6577](https://github.com/keptn/keptn/issues/6577)) ([9915466](https://github.com/keptn/keptn/commit/9915466adbb639dd40688791b23e62b86f2469f9))
* **bridge:** Login via OpenID ([#6076](https://github.com/keptn/keptn/issues/6076)) ([#6077](https://github.com/keptn/keptn/issues/6077)) ([8762c83](https://github.com/keptn/keptn/commit/8762c830607c8ff010bb9f3182bf4b29aaf124b5))
* **bridge:** Read problem title from GET /sequence endpoint ([#6388](https://github.com/keptn/keptn/issues/6388)) ([a4086c8](https://github.com/keptn/keptn/commit/a4086c8c0a76d2bbee39898bf11d7eee635f4434)), closes [#5526](https://github.com/keptn/keptn/issues/5526)
* **bridge:** Show dialog to prevent data loss on unsaved form ([#6526](https://github.com/keptn/keptn/issues/6526)) ([c7e7273](https://github.com/keptn/keptn/commit/c7e72732980af6755ee1514d362331d3afad6071))
* **bridge:** Use textarea for webhook url config and adapt styles ([#5706](https://github.com/keptn/keptn/issues/5706)) ([d2a8509](https://github.com/keptn/keptn/commit/d2a850914e2099205403530f4b4118d70b7748ba))
* **cli:** Add missing upstream warning during keptn upgrade ([#6434](https://github.com/keptn/keptn/issues/6434)) ([4867fa5](https://github.com/keptn/keptn/commit/4867fa52cc14f59051b1ef403774feb83373e205))
* **cli:** remove legacy code responsible for shipyard file upgrade from version 0.1.* to 0.2.0 ([#6270](https://github.com/keptn/keptn/issues/6270)) ([8b67626](https://github.com/keptn/keptn/commit/8b6762648791294fc0afd8e56ab3190d1b4994d7))
* Create integration test for resource-service ([#6430](https://github.com/keptn/keptn/issues/6430)) ([220f1ed](https://github.com/keptn/keptn/commit/220f1edbf86fd2a0a22654bcc507d709c51bca9f))
* Enable Nats queue groups in helm chart ([#4519](https://github.com/keptn/keptn/issues/4519)) ([#6062](https://github.com/keptn/keptn/issues/6062)) ([9c493c7](https://github.com/keptn/keptn/commit/9c493c70e0af69c0ce8ec89af82b55ce18af00f3))
* Finalize graceful shutdown ([#4522](https://github.com/keptn/keptn/issues/4522)) ([#6079](https://github.com/keptn/keptn/issues/6079)) ([b5c5d8d](https://github.com/keptn/keptn/commit/b5c5d8da52fa4ba448baf0e5ed9f31fa7460b320))
* graceful shutdown for jmeter and helm service ([#4522](https://github.com/keptn/keptn/issues/4522)) ([#5973](https://github.com/keptn/keptn/issues/5973)) ([41df113](https://github.com/keptn/keptn/commit/41df1138cb2e0bbbc7e32345a99244fc0079dcdc))
* Improve error reporting for CLI trigger cmd ([#6516](https://github.com/keptn/keptn/issues/6516)) ([4904c19](https://github.com/keptn/keptn/commit/4904c19d3829d490910f1241f5ac616fac255932))
* Introduce new-configuration-service ([#6400](https://github.com/keptn/keptn/issues/6400)) ([447f7a0](https://github.com/keptn/keptn/commit/447f7a02cfc735400da24e2044acbcc15542e7c8))
* introduce swappable logger in `go-sdk` ([#6284](https://github.com/keptn/keptn/issues/6284)) ([12d222b](https://github.com/keptn/keptn/commit/12d222b906e8a0899e038419ea550c3da9c6d263))
* **jmeter-service:** Improve error reporting for JMeter-service ([#6511](https://github.com/keptn/keptn/issues/6511)) ([c7d8224](https://github.com/keptn/keptn/commit/c7d8224e3a9aaa9898f10496368a9963aad88160))
* **lighthouse-service:** Add compared value to payload ([#5496](https://github.com/keptn/keptn/issues/5496)) ([#6194](https://github.com/keptn/keptn/issues/6194)) ([f5af13c](https://github.com/keptn/keptn/commit/f5af13c6ba48250ab71432c8edc7aac5666a4ecf))
* Resource service endpoint handler implementation([#6346](https://github.com/keptn/keptn/issues/6346)) ([#6461](https://github.com/keptn/keptn/issues/6461)) ([e19ed7b](https://github.com/keptn/keptn/commit/e19ed7b25fba4a52eedfa2b8dc68dea5c7864c5b)), closes [#6448](https://github.com/keptn/keptn/issues/6448) [#6448](https://github.com/keptn/keptn/issues/6448)
* Resource service first working version ([#6346](https://github.com/keptn/keptn/issues/6346)) ([#6517](https://github.com/keptn/keptn/issues/6517)) ([00f81f1](https://github.com/keptn/keptn/commit/00f81f1186704bf9d115064f4dd95cb1f1ac42f2))
* **resource-service:** Add common interface for interacting with Git/Secrets ([#6346](https://github.com/keptn/keptn/issues/6346)) ([#6411](https://github.com/keptn/keptn/issues/6411)) ([14af1d8](https://github.com/keptn/keptn/commit/14af1d881b4385db91bfbe352c4b36e1b076a8f9))
* **resource-service:** Complete implementation of Service ([#6530](https://github.com/keptn/keptn/issues/6530)) ([a91c116](https://github.com/keptn/keptn/commit/a91c116acb997429dbab429bab8a0bf352817422))
* **resource-service:** Improve git implementation ([#6346](https://github.com/keptn/keptn/issues/6346)) ([#6510](https://github.com/keptn/keptn/issues/6510)) ([2f31d44](https://github.com/keptn/keptn/commit/2f31d44d1fd443746940039f00a497f00d02ad2e))
* **resource-service:** Improve git implementation and testing ([#6346](https://github.com/keptn/keptn/issues/6346)) ([#6529](https://github.com/keptn/keptn/issues/6529)) ([91c5417](https://github.com/keptn/keptn/commit/91c54173b782692128cec8a7e9c0ac1c0e449270))
* **resource-service:** Resource service git implementation ([#6346](https://github.com/keptn/keptn/issues/6346)) ([#6438](https://github.com/keptn/keptn/issues/6438)) ([b6b4ac2](https://github.com/keptn/keptn/commit/b6b4ac26fa3b95f1337193de0592b245918bfe06)), closes [#6462](https://github.com/keptn/keptn/issues/6462) [#6462](https://github.com/keptn/keptn/issues/6462) [#6462](https://github.com/keptn/keptn/issues/6462) [#6462](https://github.com/keptn/keptn/issues/6462) [#6462](https://github.com/keptn/keptn/issues/6462) [#6462](https://github.com/keptn/keptn/issues/6462) [#6462](https://github.com/keptn/keptn/issues/6462) [#6462](https://github.com/keptn/keptn/issues/6462) [#6462](https://github.com/keptn/keptn/issues/6462) [#6462](https://github.com/keptn/keptn/issues/6462) [#6462](https://github.com/keptn/keptn/issues/6462) [#6462](https://github.com/keptn/keptn/issues/6462) [#6462](https://github.com/keptn/keptn/issues/6462)
* **shipyard-controller:** Add .triggered events to `lastEventTypes` property of materialized view ([#6220](https://github.com/keptn/keptn/issues/6220)) ([#6235](https://github.com/keptn/keptn/issues/6235)) ([dca04ea](https://github.com/keptn/keptn/commit/dca04eaa4538aab26700da180edfee42c60dcdd8))
* **shipyard-controller:** Allow to filter sequence states by multiple keptnContext IDs ([#6056](https://github.com/keptn/keptn/issues/6056)) ([#6093](https://github.com/keptn/keptn/issues/6093)) ([8a21919](https://github.com/keptn/keptn/commit/8a2191994f279d4ca5c4a1cc7f28dbb9c6c74213))
* **shipyard-controller:** handle sigterm ([#6051](https://github.com/keptn/keptn/issues/6051)) ([adfba40](https://github.com/keptn/keptn/commit/adfba40e4b21f01f0a806b5d525d3944305e6ca3))
* **shipyard-controller:** introduced sequence aborted state ([#6214](https://github.com/keptn/keptn/issues/6214)) ([02ab54b](https://github.com/keptn/keptn/commit/02ab54bfd949f720b41d7f71e3ff0aff06b754c5))
* validate shipyard.yaml when updating project ([#6222](https://github.com/keptn/keptn/issues/6222)) ([499352d](https://github.com/keptn/keptn/commit/499352d322a64b4a0207fb0be48d093198f6dcc1))


### Bug Fixes

* added delay and logging to graceful shutdown ([#6485](https://github.com/keptn/keptn/issues/6485)) ([#6486](https://github.com/keptn/keptn/issues/6486)) ([313db7f](https://github.com/keptn/keptn/commit/313db7ffacd9a202fa96eb95d80a3c08b3dc4eb5))
* Backup and restore integration test ([#6224](https://github.com/keptn/keptn/issues/6224)) ([7d622f8](https://github.com/keptn/keptn/commit/7d622f8f577d721a642094cb3de541b7a88a8671))
* **bridge:** Allow server inline script for base href ([#6248](https://github.com/keptn/keptn/issues/6248)) ([adebbbb](https://github.com/keptn/keptn/commit/adebbbbde0a8dbd0fa39d0e54766a6b4f34e0389))
* **bridge:** Fix problem with redirect and headers on cluster ([7407bcd](https://github.com/keptn/keptn/commit/7407bcdcfa996583d6e028afe4adf10950fd21ab))
* **bridge:** fix showing error message for OAUTH ([#6294](https://github.com/keptn/keptn/issues/6294)) ([6120087](https://github.com/keptn/keptn/commit/61200870c3f75ac0471b90640962dcb2a1d6f6ba))
* **bridge:** Fixed bridge server test ([#6314](https://github.com/keptn/keptn/issues/6314)) ([2d59f64](https://github.com/keptn/keptn/commit/2d59f6477ebe105f14989e5566f2fe181bb65fbb))
* **bridge:** Fixed bridge server tests ([#6261](https://github.com/keptn/keptn/issues/6261)) ([9f02adc](https://github.com/keptn/keptn/commit/9f02adc49695fada972039ac310f04e8c8907e69))
* **bridge:** Fixed environment screen update issues ([#6271](https://github.com/keptn/keptn/issues/6271)) ([0d5ff40](https://github.com/keptn/keptn/commit/0d5ff4002f9066f7ade3addd22da530c38b803f3))
* **bridge:** Fixed incorrect deployment link title ([#6304](https://github.com/keptn/keptn/issues/6304)) ([f237520](https://github.com/keptn/keptn/commit/f23752077c53e06204b14475b4f82b4403e9d913))
* **bridge:** Fixed removal of sequences if project endpoint of bridge server responds before projects endpoint of shipyard ([#6183](https://github.com/keptn/keptn/issues/6183)) ([8153fea](https://github.com/keptn/keptn/commit/8153fea4d707b6102c83292907a1301b6c7c4404))
* **bridge:** Remove .event suffix from payload variables in webhook ([#6396](https://github.com/keptn/keptn/issues/6396)) ([f67e5da](https://github.com/keptn/keptn/commit/f67e5daaa3315728d17a049a577d086ffffe0c2d))
* **bridge:** Update services on project change ([#6252](https://github.com/keptn/keptn/issues/6252)) ([65d4437](https://github.com/keptn/keptn/commit/65d4437d39b9889e28c1895e5e045a1af555d191))
* **cli:** Added rollback events to generated spec ([#3722](https://github.com/keptn/keptn/issues/3722)) ([#6161](https://github.com/keptn/keptn/issues/6161)) ([15ff2c6](https://github.com/keptn/keptn/commit/15ff2c63525a17fecac2e6ab222b8630a06007db))
* **cli:** Fix error handling during helm installation ([#6437](https://github.com/keptn/keptn/issues/6437)) ([#6583](https://github.com/keptn/keptn/issues/6583)) ([88c418b](https://github.com/keptn/keptn/commit/88c418b8ab56fc24a86d36163fd853a6ffabc1fe))
* **cli:** Print error message if service does not exist with `trigger delivery` ([#6351](https://github.com/keptn/keptn/issues/6351)) ([1d994a4](https://github.com/keptn/keptn/commit/1d994a4c1541879d3e6d50e4c4ca7b8a5c3c1c05))
* **cli:** project without upstream is defined as project without ([#6584](https://github.com/keptn/keptn/issues/6584)) ([aaf0a61](https://github.com/keptn/keptn/commit/aaf0a61f853c088ea82bc7de2eabde6a28560032))
* **cli:** Set CLI context before attempting to check for K8s context change ([#6449](https://github.com/keptn/keptn/issues/6449)) ([#6458](https://github.com/keptn/keptn/issues/6458)) ([3c2236d](https://github.com/keptn/keptn/commit/3c2236de055fc9f40183ce3e8c74c54efa1b81f5))
* **cli:** wrong handling of HTTPS in auth command ([#6268](https://github.com/keptn/keptn/issues/6268)) ([fa8fd1c](https://github.com/keptn/keptn/commit/fa8fd1cae9e81b351c378fbe38e7032f60d4c59f))
* **configuration-service:** Create tmpdir for unarchiving in /data/config ([#6329](https://github.com/keptn/keptn/issues/6329)) ([#6331](https://github.com/keptn/keptn/issues/6331)) ([a1f04af](https://github.com/keptn/keptn/commit/a1f04af05fae9366ec976e9b23239682901e74a6))
* **configuration-service:** Creating projects from empty upstream ([#6398](https://github.com/keptn/keptn/issues/6398)) ([#6399](https://github.com/keptn/keptn/issues/6399)) ([dc8337e](https://github.com/keptn/keptn/commit/dc8337e57605bd0760724ab128510810584d84b5))
* **configuration-service:** Fix permission issues for configuration service ([#6315](https://github.com/keptn/keptn/issues/6315)) ([#6317](https://github.com/keptn/keptn/issues/6317)) ([#6321](https://github.com/keptn/keptn/issues/6321)) ([61d9914](https://github.com/keptn/keptn/commit/61d99147389966c4db51e8eff0c30a347ca1e951))
* **configuration-service:** Make check for helm chart archives more strict ([#6447](https://github.com/keptn/keptn/issues/6447)) ([#6457](https://github.com/keptn/keptn/issues/6457)) ([babb3cd](https://github.com/keptn/keptn/commit/babb3cd11c76fe00562e2e83ddb8edcfb591b461))
* Dependencies cleanup ([#6369](https://github.com/keptn/keptn/issues/6369)) ([a38507b](https://github.com/keptn/keptn/commit/a38507b5d2fca643b69748f004835c5174fdfa88))
* Dependencies in lighthouse and remediation services ([#6368](https://github.com/keptn/keptn/issues/6368)) ([3f1646c](https://github.com/keptn/keptn/commit/3f1646cabf7460fef572e9ef3ebe3d0b3cc13cbc))
* Disable gitea installation on k3d ([#6408](https://github.com/keptn/keptn/issues/6408)) ([#6409](https://github.com/keptn/keptn/issues/6409)) ([cd984d4](https://github.com/keptn/keptn/commit/cd984d4f33b8e23df886e244045c7f7d9ba276ad))
* **distributor:** forcing restart if integration is malformed ([#6309](https://github.com/keptn/keptn/issues/6309)) ([#6363](https://github.com/keptn/keptn/issues/6363)) ([308261e](https://github.com/keptn/keptn/commit/308261e20b63160030c461eab36794f96c66dc62))
* fix graceful shutdown in sdk ([#6234](https://github.com/keptn/keptn/issues/6234)) ([a8db696](https://github.com/keptn/keptn/commit/a8db696c0f1df2eca3ef6cd283c3d5337ffd3d3d))
* Fix uniform integration test ([#6171](https://github.com/keptn/keptn/issues/6171)) ([#6174](https://github.com/keptn/keptn/issues/6174)) ([e55c398](https://github.com/keptn/keptn/commit/e55c3982b819de799fe1139eb4092aa72ff4a8d8))
* Graceful shutdown failing test ([#6462](https://github.com/keptn/keptn/issues/6462)) ([#6427](https://github.com/keptn/keptn/issues/6427)) ([4a28d73](https://github.com/keptn/keptn/commit/4a28d731d09c073137c388a1f003f6abfec511e8))
* Increase Bridge memory limits to avoid OOM ([#6562](https://github.com/keptn/keptn/issues/6562)) ([7f8d1a5](https://github.com/keptn/keptn/commit/7f8d1a5168ee3962ea3e6c61cbe1ca792b76736b))
* **installer:** Disable nats config reloader per default ([#6316](https://github.com/keptn/keptn/issues/6316)) ([#6318](https://github.com/keptn/keptn/issues/6318)) ([#6322](https://github.com/keptn/keptn/issues/6322)) ([d9263cf](https://github.com/keptn/keptn/commit/d9263cf574239bd518c517125d9e5ca80bf0e73f))
* **installer:** Remove obsolete openshift-route-service ([#6272](https://github.com/keptn/keptn/issues/6272)) ([#6389](https://github.com/keptn/keptn/issues/6389)) ([508dc25](https://github.com/keptn/keptn/commit/508dc2506414dd3bdfa1f6a290de8d0f1085294b))
* **installer:** Remove unneeded helm chart values ([#6419](https://github.com/keptn/keptn/issues/6419)) ([e5e508e](https://github.com/keptn/keptn/commit/e5e508e08dcba5d2c99d442c5a1beb406e0c16d5))
* **installer:** Use previous fsGroup per default, provide option to execute init container ([#6385](https://github.com/keptn/keptn/issues/6385)) ([#6386](https://github.com/keptn/keptn/issues/6386)) ([91eca02](https://github.com/keptn/keptn/commit/91eca02be99c02dd7df8f5b686b2d9e0368ba0b0))
* **lighthouse-service:** Lighthouse now fails if SLI fails  ([#6096](https://github.com/keptn/keptn/issues/6096)) ([#6281](https://github.com/keptn/keptn/issues/6281)) ([218cc39](https://github.com/keptn/keptn/commit/218cc390846aa79e2fc661a68a6006211960db95))
* **lighthouse-service:** Modified criteria example in SLO ([#6106](https://github.com/keptn/keptn/issues/6106)) ([#6404](https://github.com/keptn/keptn/issues/6404)) ([5b7bd19](https://github.com/keptn/keptn/commit/5b7bd198475703390ae55d21b72f9fcf83cebc76))
* minor fix in integration tests + added configuration-service securityContext ([#6540](https://github.com/keptn/keptn/issues/6540)) ([00cfe26](https://github.com/keptn/keptn/commit/00cfe26b0c0199a07a5728192c79c660c8711ce7))
* **mongodb-datastore:** Ensure MongoDB Client is not nil before retrieving database ([#6251](https://github.com/keptn/keptn/issues/6251)) ([#6255](https://github.com/keptn/keptn/issues/6255)) ([#6257](https://github.com/keptn/keptn/issues/6257)) ([fbaf0f0](https://github.com/keptn/keptn/commit/fbaf0f0b364f1c20c3303ffdba81b7465421ab12))
* **remediation-service:** add problemTitle to event message ([#5719](https://github.com/keptn/keptn/issues/5719)) ([c7d09d8](https://github.com/keptn/keptn/commit/c7d09d8a945c2a693fd52bc3dd9e339cafedae9e))
* Remove deprecated commands from CLI ([#6435](https://github.com/keptn/keptn/issues/6435)) ([d1625a7](https://github.com/keptn/keptn/commit/d1625a70512538f6593f44c67312f45fa97d8be5))
* Remove hardcoded namespace reference in installer ([#6286](https://github.com/keptn/keptn/issues/6286)) ([5396d6d](https://github.com/keptn/keptn/commit/5396d6d9030625dc0f23607abf01a9d879abd5be))
* Removed path issue within tests ([#6523](https://github.com/keptn/keptn/issues/6523)) ([#6525](https://github.com/keptn/keptn/issues/6525)) ([4295e2e](https://github.com/keptn/keptn/commit/4295e2e42abe8d1d3ffb7d8f97cf4d821c059874))
* Stabilize BackupRestore integration test ([#6344](https://github.com/keptn/keptn/issues/6344)) ([6fbd8cb](https://github.com/keptn/keptn/commit/6fbd8cb13f47ebce9fa85bdb4c047e2bea8527be))
* **statistics-service:** migrate data containing dots in keys ([#6266](https://github.com/keptn/keptn/issues/6266)) ([663c2bc](https://github.com/keptn/keptn/commit/663c2bc58d98aee11c9bee0952b6f04cd36314a0))
* **statistics-service:** migration of keptn service execution data ([#6324](https://github.com/keptn/keptn/issues/6324)) ([766a8e3](https://github.com/keptn/keptn/commit/766a8e335586917665bce8d2c73e498fe407cf81))
* Unit test in shipyard-controller ([#6370](https://github.com/keptn/keptn/issues/6370)) ([491a19a](https://github.com/keptn/keptn/commit/491a19a96c50a760868be63a5c66c35f2a9becda))
* Update dependencies ([#6381](https://github.com/keptn/keptn/issues/6381)) ([65a229a](https://github.com/keptn/keptn/commit/65a229aa31b71afab0b05789652c1a98309335e3))
* Update error messages ([#6197](https://github.com/keptn/keptn/issues/6197)) ([d43188e](https://github.com/keptn/keptn/commit/d43188e701e071877bc7cc97aaf20ff51f0a3ae9))
* Update go.sum of distributor ([#6367](https://github.com/keptn/keptn/issues/6367)) ([fc2b60a](https://github.com/keptn/keptn/commit/fc2b60a17dbe1486dc34afdf8fed5db562d6e07e))
* Update the JMeter Service to JMeter 5.4.2 ([#6405](https://github.com/keptn/keptn/issues/6405)) ([ccef405](https://github.com/keptn/keptn/commit/ccef4050a7a2497e6da01bf7d2f0632e72206a20))
* **webhook-service:** Disallow requests to loopback addresses ([#6361](https://github.com/keptn/keptn/issues/6361)) ([e7f814e](https://github.com/keptn/keptn/commit/e7f814e8bd6f975ff4a10c2bb7b056daaeae016f))


### Refactoring

* **bridge:** Move secret picker in own component ([#5733](https://github.com/keptn/keptn/issues/5733)) ([#6099](https://github.com/keptn/keptn/issues/6099)) ([a54f6a7](https://github.com/keptn/keptn/commit/a54f6a7ea6508eac4e27af61877c04cd6a52aa30))
* **bridge:** Replace data service mock with api service mock ([#5093](https://github.com/keptn/keptn/issues/5093)) ([101e472](https://github.com/keptn/keptn/commit/101e4728c2c272658497fe9c4023398757f703b7))


### Docs

* Add Keptn versioning and version compatibility document ([#5489](https://github.com/keptn/keptn/issues/5489)) ([c6e8a5c](https://github.com/keptn/keptn/commit/c6e8a5cda2e17ac6626448669f5572196f8d5511))
* **configuration-service:** Update API documentation ([#6008](https://github.com/keptn/keptn/issues/6008)) ([76f9ef2](https://github.com/keptn/keptn/commit/76f9ef2d24cc0a86e0e73d465cf6a6921b4de0cf))
* Update Integration Tests Developer documentation ([#6548](https://github.com/keptn/keptn/issues/6548)) ([d34b70c](https://github.com/keptn/keptn/commit/d34b70c4d5f0dd63909146d1b205c8056a910e1d))


### Performance

* **shipyard-controller:** Remove DB connection locking ([#6326](https://github.com/keptn/keptn/issues/6326)) ([690ce1c](https://github.com/keptn/keptn/commit/690ce1cacfe77529ff4189463ff5f3cd37496181))


### Other

*  update go_utils to safe version ([#6289](https://github.com/keptn/keptn/issues/6289)) ([f331482](https://github.com/keptn/keptn/commit/f3314822fdf91ef07da9cd1267c1ecc5bf656734))
* Add [@oleg-nenashev](https://github.com/oleg-nenashev) to the list of contributors ([#6256](https://github.com/keptn/keptn/issues/6256)) ([6817795](https://github.com/keptn/keptn/commit/6817795a3121ba651d52c8f0d42a23a830da36c3))
* **bridge:** Revert PR [#6341](https://github.com/keptn/keptn/issues/6341) ([#6585](https://github.com/keptn/keptn/issues/6585)) ([71c1e19](https://github.com/keptn/keptn/commit/71c1e19a45a2f706f69bcb6f75d77dc34faf871d))
* Bump JMeter to latest version ([307abf9](https://github.com/keptn/keptn/commit/307abf9142d52d97c2a5373c5cf6f61273d711f1))
* Correct example lighthouse criteria ([#6160](https://github.com/keptn/keptn/issues/6160)) ([#6406](https://github.com/keptn/keptn/issues/6406)) ([2d432eb](https://github.com/keptn/keptn/commit/2d432eb08105102750047cfa1e04beab71c3ae82)), closes [#6106](https://github.com/keptn/keptn/issues/6106)
* **distributor:** Upgrade go-utils, use thread safe fake EventSender in unit tests ([#6153](https://github.com/keptn/keptn/issues/6153)) ([da6fef0](https://github.com/keptn/keptn/commit/da6fef0c4bea6f26957c27757fd3832aace973fb))
* **helm-service:** Remove `service.create.finished` subscription ([#6181](https://github.com/keptn/keptn/issues/6181)) ([dc21c46](https://github.com/keptn/keptn/commit/dc21c46613625c4ff392db8d60b08b7cfb574387))
* Promote [@oleg-nenashev](https://github.com/oleg-nenashev) to maintainers ([#6522](https://github.com/keptn/keptn/issues/6522)) ([40e2deb](https://github.com/keptn/keptn/commit/40e2deb0fb13827efe26bd374dd67ade631e78f6))
* **secret-service:** updated README.md ([#6156](https://github.com/keptn/keptn/issues/6156)) ([d600e55](https://github.com/keptn/keptn/commit/d600e556236888dc32bab9c670d7565cab0dc9b0))
* update affiliation ([#6521](https://github.com/keptn/keptn/issues/6521)) ([642410f](https://github.com/keptn/keptn/commit/642410f4cc275fcc1900cc8ac6f354cee1b8a723))
* Update contributor lists ([#6450](https://github.com/keptn/keptn/issues/6450)) ([809532b](https://github.com/keptn/keptn/commit/809532b0f49fce2c4abf84e06ec1b4df62d1fbc6))

## [0.11.0](https://github.com/keptn/keptn/compare/0.10.1...0.11.0) (2021-11-24)


### ⚠ BREAKING CHANGES

* MongoDB was updated from 3.6 to 4.4, also the custom helm chart was switched out for the Bitnami MongoDB Helm Chart. This means that a manual database migration is needed to preserve data during the keptn upgrade process! Steps to upgrade keptn with the manual migration can be found on the [Keptn Upgrade page](https://keptn.sh/docs/0.11.x/operate/upgrade/).

### Features

* Added context with cancel function to sdk ([#4552](https://github.com/keptn/keptn/issues/4552)) ([#5972](https://github.com/keptn/keptn/issues/5972)) ([d21e682](https://github.com/keptn/keptn/commit/d21e68230a789ac29dd47bf2b284928ded89a464))
* added probes for readiness and liveness ([#5303](https://github.com/keptn/keptn/issues/5303))  ([#5534](https://github.com/keptn/keptn/issues/5534)) ([6899ee7](https://github.com/keptn/keptn/commit/6899ee7dd6f23724c4c4dc6e16c4a218cbf2453c)), closes [#5533](https://github.com/keptn/keptn/issues/5533)
* **bridge:** Add 404 page ([#4983](https://github.com/keptn/keptn/issues/4983)) ([#6004](https://github.com/keptn/keptn/issues/6004)) ([aa7b4fa](https://github.com/keptn/keptn/commit/aa7b4fa0810da9c40f96996cb9a2fe34f8b66929))
* **bridge:** Add checkbox to set the `sendFinished` flag ([#5735](https://github.com/keptn/keptn/issues/5735)) ([#5989](https://github.com/keptn/keptn/issues/5989)) ([89598f8](https://github.com/keptn/keptn/commit/89598f8aaff5543171910c54b5dceffece4a1029))
* **bridge:** Prevent cut off of evaluation board ([#5279](https://github.com/keptn/keptn/issues/5279)) ([1ae06a1](https://github.com/keptn/keptn/commit/1ae06a13a53ab21df183b858544f3ced115897ba))
* **cli:** created user warning about changed database model in keptn 0.11.* ([#6071](https://github.com/keptn/keptn/issues/6071)) ([7f3447c](https://github.com/keptn/keptn/commit/7f3447c8eab6476ba5b484bae9f1344457f5ed02))
* handle time consistently ([#4788](https://github.com/keptn/keptn/issues/4788)) ([#5971](https://github.com/keptn/keptn/issues/5971)) ([e284d72](https://github.com/keptn/keptn/commit/e284d72268f8fa5b68faf00e23896dd731205b73))
* **lighthouse-service:** Added SIGTERM for lighthouse handlers ([#5304](https://github.com/keptn/keptn/issues/5304)) ([#5558](https://github.com/keptn/keptn/issues/5558)) ([ca9742c](https://github.com/keptn/keptn/commit/ca9742cc6a841bcb59ade85789e2155fbc7d8693))
* Switch mongoDB image to bitnami mongoDB chart ([#4801](https://github.com/keptn/keptn/issues/4801)) ([b3dabd6](https://github.com/keptn/keptn/commit/b3dabd6297bd0ddd0b6e2e5815c53919892045c2))


### Bug Fixes

* Adapt log level of SDK logs ([#5920](https://github.com/keptn/keptn/issues/5920)) ([#5921](https://github.com/keptn/keptn/issues/5921)) ([d314008](https://github.com/keptn/keptn/commit/d314008f3026c90252533fdf1c8ebe46538f9e42))
* **api:** Remove multiple types in event model ([#5948](https://github.com/keptn/keptn/issues/5948)) ([#5957](https://github.com/keptn/keptn/issues/5957)) ([30d5556](https://github.com/keptn/keptn/commit/30d5556e2da731a3d517c14ba36895c4bbaff11a))
* **approval-service:** Fall back to manual strategy when there is no result available ([#6012](https://github.com/keptn/keptn/issues/6012)) ([#6017](https://github.com/keptn/keptn/issues/6017)) ([9617814](https://github.com/keptn/keptn/commit/9617814ecbd5dcc853861d5b9939ab04c3d4a772))
* **bridge:** Add empty state to sequence-view ([#5084](https://github.com/keptn/keptn/issues/5084)) ([#5693](https://github.com/keptn/keptn/issues/5693)) ([b7c10df](https://github.com/keptn/keptn/commit/b7c10dfe0bccce073ef62f81806e0edf1166465a))
* **bridge:** Correctly show warning state ([#6003](https://github.com/keptn/keptn/issues/6003)) ([9a21d19](https://github.com/keptn/keptn/commit/9a21d19a698127801c8dbf9ca35d6e0f0e0531f3))
* **bridge:** don't log err (contains the x-token), only log err.message ([#6047](https://github.com/keptn/keptn/issues/6047)) ([#6052](https://github.com/keptn/keptn/issues/6052)) ([3eea6e3](https://github.com/keptn/keptn/commit/3eea6e37e8682d0a580313b0462cd02447925fcc))
* **bridge:** Fix integration curl commands for api ([#5941](https://github.com/keptn/keptn/issues/5941)) ([d76eccc](https://github.com/keptn/keptn/commit/d76ecccbcacc2b6db3ccbb06aec50942303d2bf7))
* **bridge:** Fixed missing problem title and decode of remediation config ([#6053](https://github.com/keptn/keptn/issues/6053)) ([ea0c53f](https://github.com/keptn/keptn/commit/ea0c53f5b429327c3780ffc6fd2b97522d25b44f))
* **bridge:** Fixed overwriting of data in environment screen ([#5841](https://github.com/keptn/keptn/issues/5841)) ([74a9a3d](https://github.com/keptn/keptn/commit/74a9a3ded3184ebe057fd918c7c1b23fd25d86c2))
* **bridge:** Fixed wrong weight of SLI ([#5987](https://github.com/keptn/keptn/issues/5987)) ([e536dbc](https://github.com/keptn/keptn/commit/e536dbc9bd8602deab4a1037a0f8bd775b7b16b7))
* **bridge:** Possible fix for flaky clicks in UI tests ([#5909](https://github.com/keptn/keptn/issues/5909)) ([58c5deb](https://github.com/keptn/keptn/commit/58c5deb78b5467d3e2b304447fcbed8b0589f984))
* **bridge:** Remove inline script for base url and upgrade-insecure-requests header part ([#6019](https://github.com/keptn/keptn/issues/6019)) ([b2e9960](https://github.com/keptn/keptn/commit/b2e9960a42e2b14f49a888619395855ede3a7c44))
* **bridge:** Show right event type ([#5828](https://github.com/keptn/keptn/issues/5828)) ([316d117](https://github.com/keptn/keptn/commit/316d117aa8c4c6d9f3f66d746be702bda6c53f02))
* **bridge:** Take SLI-weight out of SLO-file ([#5782](https://github.com/keptn/keptn/issues/5782)) ([f961ce1](https://github.com/keptn/keptn/commit/f961ce1fc126130047c0f5971ef7683c5e49a50f))
* **bridge:** Use helmet middlewares to prevent XSS ([8a58fb3](https://github.com/keptn/keptn/commit/8a58fb30389f4433dc049e7da37c03814a712d4b))
* **cli:** Make sure the release version is set in command descriptions ([#5762](https://github.com/keptn/keptn/issues/5762)) ([#5888](https://github.com/keptn/keptn/issues/5888)) ([24110c0](https://github.com/keptn/keptn/commit/24110c067ea8a698b5651d9b777a620765009c53))
* **cli:** problem with missing http(s) in endpoint flag during keptn auth ([#6039](https://github.com/keptn/keptn/issues/6039)) ([e4164db](https://github.com/keptn/keptn/commit/e4164db156cf0d767032d37cc59bd57091b6be10))
* **configuration-service:** changed bad order of extracting and adding resources to services ([#6006](https://github.com/keptn/keptn/issues/6006)) ([35605b7](https://github.com/keptn/keptn/commit/35605b7618c23271eadde447f7b8f12119f8db8b))
* **configuration-service:** Completely replace previous helm chart directory when updating ([#6050](https://github.com/keptn/keptn/issues/6050)) ([#6058](https://github.com/keptn/keptn/issues/6058)) ([74eefdf](https://github.com/keptn/keptn/commit/74eefdf5c12716c466f6ff64f6de4451a41ac650))
* **configuration-service:** Fix order of extracting and adding files to the repo ([#6041](https://github.com/keptn/keptn/issues/6041)) ([#6045](https://github.com/keptn/keptn/issues/6045)) ([4a3bf22](https://github.com/keptn/keptn/commit/4a3bf22bd560f0492e98ba19ea427162b9b549df))
* **distributor:** Fix message filtering in distributor ([#6074](https://github.com/keptn/keptn/issues/6074)) ([#6075](https://github.com/keptn/keptn/issues/6075)) ([602eb37](https://github.com/keptn/keptn/commit/602eb37609f65bc418db165a0679e2c0ab3edd78))
* **distributor:** fix subscription handling after message broker reconnect ([#5823](https://github.com/keptn/keptn/issues/5823)) ([49b1051](https://github.com/keptn/keptn/commit/49b1051fa31f8860fd490d4d5e1c20ffb6465048))
* **distributor:** Sanitized logs and cleaned up forwarder lifecycle ([#6036](https://github.com/keptn/keptn/issues/6036)) ([be5adb5](https://github.com/keptn/keptn/commit/be5adb5b77d10a15a8fa13afa5674460b52bd929))
* **distributor:** Set default timeout of Uniform API requests to 5s ([#6011](https://github.com/keptn/keptn/issues/6011)) ([#6015](https://github.com/keptn/keptn/issues/6015)) ([d89cab9](https://github.com/keptn/keptn/commit/d89cab9b11ebd83cc73f9c63d94ff23ae170818e))
* Fix bug where approval and remediation service would not run through unit tests anymore ([495654c](https://github.com/keptn/keptn/commit/495654cd0c59677527bd22e9764eecafc8b38c5f))
* Fix bug where DCO check always fails on dependabot PRs ([6a4b58d](https://github.com/keptn/keptn/commit/6a4b58d29a5da25206d021780aea4d8d6b1a762d))
* Fix multiple issues found by Sonatype Lift static analysis ([#5934](https://github.com/keptn/keptn/issues/5934)) ([dd93b4e](https://github.com/keptn/keptn/commit/dd93b4ea1f261435e5b995f10c5005aab2be06ca))
* Fix sub-project change detection for build-everything and master builds ([db808d6](https://github.com/keptn/keptn/commit/db808d68ee1cc02a104674078ad3b09d9ac5cc12))
* Fix version not showing up anymore in API ([#5783](https://github.com/keptn/keptn/issues/5783)) ([1eea3f9](https://github.com/keptn/keptn/commit/1eea3f94e62b7ded029c54f36673bb577fee1977))
* Fixed bug where MongoDB would not come up in airgapped setup ([#5939](https://github.com/keptn/keptn/issues/5939)) ([079a6b4](https://github.com/keptn/keptn/commit/079a6b43c33c80f1739a8c2d8985b4c73337f517))
* Handle upstream not found ([#5977](https://github.com/keptn/keptn/issues/5977)) ([#5994](https://github.com/keptn/keptn/issues/5994)) ([77240d4](https://github.com/keptn/keptn/commit/77240d4fa9e3f3f24629319f007051c1cc387a0c))
* **shipyard-controller:** cleanup uniform subscriptions when service is deleted ([#5725](https://github.com/keptn/keptn/issues/5725)) ([#5766](https://github.com/keptn/keptn/issues/5766)) ([d95f7a6](https://github.com/keptn/keptn/commit/d95f7a665f9ffaa8929fa65d84a0c634591c0a9c))
* **shipyard-controller:** migrate and avoid mongodb keys containing dots ([#6065](https://github.com/keptn/keptn/issues/6065)) ([5259bcf](https://github.com/keptn/keptn/commit/5259bcfd534690c596383f4eae93576e292f8b03))
* **shipyard-controller:** removed error shadowing ([#6048](https://github.com/keptn/keptn/issues/6048)) ([04416da](https://github.com/keptn/keptn/commit/04416da01497d561300583770d01914ec3af20c8))
* **shipyard-controller:** Store `lastEventTypes` only for events that belong to a sequence controlled by the shipyard controller ([#5309](https://github.com/keptn/keptn/issues/5309)) ([#5777](https://github.com/keptn/keptn/issues/5777)) ([ee27c62](https://github.com/keptn/keptn/commit/ee27c62f07e04791892c90b39ade410cc345034e))
* Update auto-update pipelines to follow keptns semantic PR guidelines ([#5931](https://github.com/keptn/keptn/issues/5931)) ([280fa4e](https://github.com/keptn/keptn/commit/280fa4ef4829e4d9a3968b7049befe7d3fd87304))
* **webhook-service:** Avoid .finished.finished/.started.finished events ([#5954](https://github.com/keptn/keptn/issues/5954)) ([#6000](https://github.com/keptn/keptn/issues/6000)) ([fbe01a8](https://github.com/keptn/keptn/commit/fbe01a88a642904eaecf507d22f9615a560f124f))
* **webhook-service:** invalid conversion of eventType ([#5998](https://github.com/keptn/keptn/issues/5998)) ([67dba55](https://github.com/keptn/keptn/commit/67dba55d8fed3e978ce596a02073ff059581e165))


### Docs

* Add release notes for 0.10.0 ([d748251](https://github.com/keptn/keptn/commit/d7482517e74b8b12038c9438daa01f3172ad81e3))


### Other

* Add @RealAnna to Maintainers list ([34175bb](https://github.com/keptn/keptn/commit/34175bb5aa9cb03b1ba2600ae1d93e5e8602d13d))
* Add environment variables for setting log levels of Keptn services ([#5373](https://github.com/keptn/keptn/issues/5373)) ([#5911](https://github.com/keptn/keptn/issues/5911)) ([809baea](https://github.com/keptn/keptn/commit/809baea2672fdbcebe236a9b4a6760223cb84870))
* Add flowcharts that describe components of the shipyard controller ([#5919](https://github.com/keptn/keptn/issues/5919)) ([8aa4dd8](https://github.com/keptn/keptn/commit/8aa4dd85fc24d09fa362812518d1314ec4a41c79))
* add missing release notes ([#5781](https://github.com/keptn/keptn/issues/5781)) ([dab9844](https://github.com/keptn/keptn/commit/dab9844c7213638663faf9b37879432e85d4f312))
* Add odubajDT as maintainer ([#60](https://github.com/keptn/keptn/issues/60)) ([#6049](https://github.com/keptn/keptn/issues/6049)) ([65ae6cf](https://github.com/keptn/keptn/commit/65ae6cfa505dc07a5f7207b7625b36dfac8549cc))
* Add TannerGilbert as project member ([#5899](https://github.com/keptn/keptn/issues/5899)) ([65148be](https://github.com/keptn/keptn/commit/65148be195779baeed0a7462cfc3d753ec88c861))
* Add the correct label for bug reports ([#5908](https://github.com/keptn/keptn/issues/5908)) ([dc296a5](https://github.com/keptn/keptn/commit/dc296a5a504a453925c5a761690554a9b2fd896b))
* Added go-sdk and webhook-service to dependencies-and-licenses check ([#5898](https://github.com/keptn/keptn/issues/5898)) ([6481ca3](https://github.com/keptn/keptn/commit/6481ca30fdf8068d3626c32de8b967606023a0e2))
* Cancel integration tests when mismatch between CLI and kube context is detected ([#5743](https://github.com/keptn/keptn/issues/5743)) ([#5824](https://github.com/keptn/keptn/issues/5824)) ([5596611](https://github.com/keptn/keptn/commit/559661149928ed7cfc19ea1e99da60e2fb1322a9)), closes [#5734](https://github.com/keptn/keptn/issues/5734)
* fixing imports according to snyc ([#5936](https://github.com/keptn/keptn/issues/5936)) ([391ace2](https://github.com/keptn/keptn/commit/391ace23fa8df83be6415d61c65ba9b5de66f637))
* **helm-service:** More meaningful error messages ([#6089](https://github.com/keptn/keptn/issues/6089)) ([80d59cb](https://github.com/keptn/keptn/commit/80d59cbb68633659cd26ded62a18c12c92f84155))
* Increase timeout of DeliveryAssistant integration test ([#6067](https://github.com/keptn/keptn/issues/6067)) ([b141ce4](https://github.com/keptn/keptn/commit/b141ce435491b7e01a3c1f499da90a125a1c27af))
* **jmeter-service:** bump version of jmeter binary to 5.4.1 ([#6032](https://github.com/keptn/keptn/issues/6032)) ([3c250d2](https://github.com/keptn/keptn/commit/3c250d206c6694202b3116bdf28207778ecaa67b))
* **jmeter-service:** cleanups ([#6014](https://github.com/keptn/keptn/issues/6014)) ([5e779eb](https://github.com/keptn/keptn/commit/5e779eb91ac0cbdc8d64729bfcbe9c8d739b1b31))
* Mitigating racecondition in unit tests ([#5901](https://github.com/keptn/keptn/issues/5901)) ([5a642a5](https://github.com/keptn/keptn/commit/5a642a538aee5d4a4e162de7dbb722216b5794a4))
* **mongodb-datastore:** Refactoring ([#5917](https://github.com/keptn/keptn/issues/5917)) ([#6002](https://github.com/keptn/keptn/issues/6002)) ([3242094](https://github.com/keptn/keptn/commit/324209413043e79561198075b0da71eb42de063c))
* Polish HTTP(S) headers ([a4f52b4](https://github.com/keptn/keptn/commit/a4f52b409aa77a64e89eba637e71b5cdeefb22e1))
* Remove sequence migration integration test because component has been removed ([#6101](https://github.com/keptn/keptn/issues/6101)) ([afeb7fc](https://github.com/keptn/keptn/commit/afeb7fc718ecad0babf6cdb4d7e8517e5c3cf721))
* removed cluster role binding ([#5955](https://github.com/keptn/keptn/issues/5955)) ([391a3ba](https://github.com/keptn/keptn/commit/391a3bae209254dd9c104c2e4782f89804f14900))
* Removed obsolete files ([#4818](https://github.com/keptn/keptn/issues/4818)) ([#5932](https://github.com/keptn/keptn/issues/5932)) ([588a76d](https://github.com/keptn/keptn/commit/588a76ded965af36f25790e22a5e6a82d11c66a9))
* **shipyard-controller:** Adapted log level ([#5978](https://github.com/keptn/keptn/issues/5978)) ([3cbfcd7](https://github.com/keptn/keptn/commit/3cbfcd7bb5eceea810095964f7892eb0886a18c3))
* **shipyard-controller:** cleaning up package(s) ([#5786](https://github.com/keptn/keptn/issues/5786)) ([a6e51d4](https://github.com/keptn/keptn/commit/a6e51d42f9cfe633437900a23d5d2d7d0f6ae317))
* **shipyard-controller:** cleanups & refactorings 2 ([#5937](https://github.com/keptn/keptn/issues/5937)) ([adf4078](https://github.com/keptn/keptn/commit/adf4078fc277076ea3e057b2f4448cdc2b44243e))
* **shipyard-controller:** Do not interpret absence of configurationChange property as an error ([#5979](https://github.com/keptn/keptn/issues/5979)) ([#5982](https://github.com/keptn/keptn/issues/5982)) ([28a9a92](https://github.com/keptn/keptn/commit/28a9a92dfdda33d4952d4322c14c2363401a38d3))
* **shipyard-controller:** Extract shipyard retrieval into its own component ([#5243](https://github.com/keptn/keptn/issues/5243)) ([#5821](https://github.com/keptn/keptn/issues/5821)) ([a1d18ae](https://github.com/keptn/keptn/commit/a1d18ae4c277bfaebe0c19f408a764033057f1f2))
* **shipyard-controller:** move event operations to event repo ([#5902](https://github.com/keptn/keptn/issues/5902)) ([730864b](https://github.com/keptn/keptn/commit/730864b9e57c8d4a5cfb82567a575eb8e9661cb7))
* Updated dependencies according to ArtifactHub and Snyk ([#5543](https://github.com/keptn/keptn/issues/5543)) ([#5951](https://github.com/keptn/keptn/issues/5951)) ([48fc51c](https://github.com/keptn/keptn/commit/48fc51c1c2b4689b6658e36a8751d0d020c1912c))
* Updated go-utils dependency ([#5968](https://github.com/keptn/keptn/issues/5968)) ([#5969](https://github.com/keptn/keptn/issues/5969)) ([f2c796e](https://github.com/keptn/keptn/commit/f2c796e19b8cac6ae737b2a85e7868b763816e0b))
* use correct link in CLI upgrade message ([961ea2a](https://github.com/keptn/keptn/commit/961ea2a6f9572d3289abb61b2827289e5e2ac224))
* Version 0.10.0 into master ([9eb12ec](https://github.com/keptn/keptn/commit/9eb12ec3f0119e9e0ea30050019e3acf4cca535a))


### Refactoring

* **bridge:** Reduce number of API calls for project endpoint ([#5450](https://github.com/keptn/keptn/issues/5450)) ([25fd876](https://github.com/keptn/keptn/commit/25fd8766b781797b0e043dc929ec92841745982c))
* **bridge:** Refactoring of project settings / create project ([#5100](https://github.com/keptn/keptn/issues/5100)) ([03fc3d2](https://github.com/keptn/keptn/commit/03fc3d298592423cef72a12aaac2a1a8001c2a66))
* **bridge:** Refactoring of service screen ([#4918](https://github.com/keptn/keptn/issues/4918)) ([#5244](https://github.com/keptn/keptn/issues/5244)) ([8f3b810](https://github.com/keptn/keptn/commit/8f3b810e60a6ca7e3cfb02adffb62375cabc9a27))
* **bridge:** Refactoring of services settings ([#5100](https://github.com/keptn/keptn/issues/5100)) ([771ec59](https://github.com/keptn/keptn/commit/771ec595b482d24c9e0d2473d2eabf303e139bc2))
* **cli:** use viper to manage keptn config ([#5694](https://github.com/keptn/keptn/issues/5694)) ([498d893](https://github.com/keptn/keptn/commit/498d89345e42605fe24aea731e66b03dea722be7))
