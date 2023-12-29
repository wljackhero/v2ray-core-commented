# v2ray source code comment

## config loading sequence

```
2023/12/29 14:05:07 config.go:204 init
2023/12/29 14:05:07 config.go:56 RegisterConfigLoader reg configLoaderByName protobuf
2023/12/29 14:05:07 config.go:56 RegisterConfigLoader reg configLoaderByName pb
2023/12/29 14:05:07 config.go:65 RegisterConfigLoader reg configLoaderByExt .pb
2023/12/29 14:05:07 infra/conf/jsonpb/jsonpb.go:46 init
2023/12/29 14:05:07 config.go:56 RegisterConfigLoader reg configLoaderByName jsonpb
2023/12/29 14:05:07 config.go:65 RegisterConfigLoader reg configLoaderByExt .pb.json
2023/12/29 14:05:07 config.go:65 RegisterConfigLoader reg configLoaderByExt .pbjson
2023/12/29 14:05:07 infra/conf/v2jsonpb/v2jsonpb.go:45 init
2023/12/29 14:05:07 config.go:56 RegisterConfigLoader reg configLoaderByName v2jsonpb
2023/12/29 14:05:07 config.go:65 RegisterConfigLoader reg configLoaderByExt .v2pb.json
2023/12/29 14:05:07 config.go:65 RegisterConfigLoader reg configLoaderByExt .v2pbjson
2023/12/29 14:05:07 infra/conf/v5cfg/init.go:18 init
2023/12/29 14:05:07 config.go:56 RegisterConfigLoader reg configLoaderByName jsonv5
2023/12/29 14:05:07 config.go:65 RegisterConfigLoader reg configLoaderByExt .v5.json
2023/12/29 14:05:07 config.go:65 RegisterConfigLoader reg configLoaderByExt .v5.jsonc
2023/12/29 14:05:07 infra/conf/mergers/mergers.go:15 init
2023/12/29 14:05:07 infra/conf/mergers/mergers.go:59 RegisterMerger reg json
2023/12/29 14:05:07 infra/conf/mergers/mergers.go:67 RegisterMerger reg ext .json
2023/12/29 14:05:07 infra/conf/mergers/mergers.go:67 RegisterMerger reg ext .jsonc
2023/12/29 14:05:07 infra/conf/mergers/mergers.go:59 RegisterMerger reg toml
2023/12/29 14:05:07 infra/conf/mergers/mergers.go:67 RegisterMerger reg ext .toml
2023/12/29 14:05:07 infra/conf/mergers/mergers.go:59 RegisterMerger reg yaml
2023/12/29 14:05:07 infra/conf/mergers/mergers.go:67 RegisterMerger reg ext .yml
2023/12/29 14:05:07 infra/conf/mergers/mergers.go:67 RegisterMerger reg ext .yaml
2023/12/29 14:05:07 infra/conf/mergers/mergers.go:59 RegisterMerger reg auto
2023/12/29 14:05:07 main/formats/formats.go:15 init
2023/12/29 14:05:07 main/formats/formats.go:23 formats register auto
2023/12/29 14:05:07 config.go:56 RegisterConfigLoader reg configLoaderByName auto
2023/12/29 14:05:07 main/formats/formats.go:23 formats register json
2023/12/29 14:05:07 config.go:56 RegisterConfigLoader reg configLoaderByName json
2023/12/29 14:05:07 config.go:65 RegisterConfigLoader reg configLoaderByExt .json
2023/12/29 14:05:07 config.go:65 RegisterConfigLoader reg configLoaderByExt .jsonc
2023/12/29 14:05:07 main/formats/formats.go:23 formats register toml
2023/12/29 14:05:07 config.go:56 RegisterConfigLoader reg configLoaderByName toml
2023/12/29 14:05:07 config.go:65 RegisterConfigLoader reg configLoaderByExt .toml
2023/12/29 14:05:07 main/formats/formats.go:23 formats register yaml
2023/12/29 14:05:07 config.go:56 RegisterConfigLoader reg configLoaderByName yaml
2023/12/29 14:05:07 config.go:65 RegisterConfigLoader reg configLoaderByExt .yml
2023/12/29 14:05:07 config.go:65 RegisterConfigLoader reg configLoaderByExt .yaml
/Users/shareit/Desktop/v2ray-core-commented/main/main.go:12: main starts
```

## app/

> this folder holds all apps you can config in your config file.  commonly used apps are:
```
log
policy
router
stats
etc.
```

## main/commands/all/

>    this folder holds all the command that has been integreted in v2ray itself, including:
```
run
api
tls
verify
uuid
etc.
```

>    all the commands can be listed using v2ray help

>    all commands are added with **RegisterCommand**

>   follow **RegisterCommand** to get all supported commands
```
shareit@MacBook-Pro-10-Jack v2ray-core-commented # ./v2ray help
A unified platform for anti-censorship.

Usage:

    v2ray <command> [arguments]

The commands are:

    run           run V2Ray with config
    api           call V2Ray API
    convert       convert config files
    test          test config files
    tls           TLS tools
    uuid          generate new UUID
    verify        verify if a binary is officially signed
    version       print V2Ray version

Use "v2ray help <command>" for more information about a command.

Additional help topics:

    config-merge  config merge logic
    format-loader config formats and loading

Use "v2ray help <topic>" for more information about that topic.
```

## main/commands/run.go

- **executeRun** *key entrance*
    - **getConfigFilePath** *get valid config file*
        > all supported formats are registered with **RegisterConfigLoader**
    - **startV2Ray** *service starts here*
        - **core.LoadConfig** load config with method registered with **RegisterConfigLoader**
        - **core.New** create server with config just loaded
            > call v2ray.go New()
    - **server.Start** *start service*

---

## load config

> loaded config
```
(dlv) p config
("*github.com/v2fly/v2ray-core/v5.Config")(0x14000231ef0)
*github.com/v2fly/v2ray-core/v5.Config {
	state: google.golang.org/protobuf/internal/impl.MessageState {
		NoUnkeyedLiterals: google.golang.org/protobuf/internal/pragma.NoUnkeyedLiterals {},
		DoNotCompare: google.golang.org/protobuf/internal/pragma.DoNotCompare [],
		DoNotCopy: google.golang.org/protobuf/internal/pragma.DoNotCopy [],
		atomicMessageInfo: *google.golang.org/protobuf/internal/impl.MessageInfo nil,},
	sizeCache: 0,
	unknownFields: []uint8 len: 0, cap: 0, nil,
	Inbound: []*github.com/v2fly/v2ray-core/v5.InboundHandlerConfig len: 2, cap: 2, [
		*(*"github.com/v2fly/v2ray-core/v5.InboundHandlerConfig")(0x140005365a0),
		*(*"github.com/v2fly/v2ray-core/v5.InboundHandlerConfig")(0x14000537450),
	],
	Outbound: []*github.com/v2fly/v2ray-core/v5.OutboundHandlerConfig len: 2, cap: 2, [
		*(*"github.com/v2fly/v2ray-core/v5.OutboundHandlerConfig")(0x14004710ea0),
		*(*"github.com/v2fly/v2ray-core/v5.OutboundHandlerConfig")(0x14004710f60),
	],
	App: []*google.golang.org/protobuf/types/known/anypb.Any len: 8, cap: 12, [
		*(*"google.golang.org/protobuf/types/known/anypb.Any")(0x14000116c30),
		*(*"google.golang.org/protobuf/types/known/anypb.Any")(0x14000116500),
		*(*"google.golang.org/protobuf/types/known/anypb.Any")(0x14000116550),
		*(*"google.golang.org/protobuf/types/known/anypb.Any")(0x140001165a0),
		*(*"google.golang.org/protobuf/types/known/anypb.Any")(0x14000116910),
		*(*"google.golang.org/protobuf/types/known/anypb.Any")(0x14000116960),
		*(*"google.golang.org/protobuf/types/known/anypb.Any")(0x14000117360),
		*(*"google.golang.org/protobuf/types/known/anypb.Any")(0x14000117cc0),
	],
	Transport: *github.com/v2fly/v2ray-core/v5/transport.Config nil,
	Extension: []*google.golang.org/protobuf/types/known/anypb.Any len: 0, cap: 0, nil,}
(dlv) p config.Inbound
[]*github.com/v2fly/v2ray-core/v5.InboundHandlerConfig len: 2, cap: 2, [
	*{
		state: (*"google.golang.org/protobuf/internal/impl.MessageState")(0x140005365a0),
		sizeCache: 0,
		unknownFields: []uint8 len: 0, cap: 0, nil,
		Tag: "api",
		ReceiverSettings: *(*"google.golang.org/protobuf/types/known/anypb.Any")(0x140005361e0),
		ProxySettings: *(*"google.golang.org/protobuf/types/known/anypb.Any")(0x14000536550),},
	*{
		state: (*"google.golang.org/protobuf/internal/impl.MessageState")(0x14000537450),
		sizeCache: 0,
		unknownFields: []uint8 len: 0, cap: 0, nil,
		Tag: "dft",
		ReceiverSettings: *(*"google.golang.org/protobuf/types/known/anypb.Any")(0x14000536eb0),
		ProxySettings: *(*"google.golang.org/protobuf/types/known/anypb.Any")(0x14000537400),},
]
(dlv) p config.Inbound[1]
("*github.com/v2fly/v2ray-core/v5.InboundHandlerConfig")(0x14000537450)
*github.com/v2fly/v2ray-core/v5.InboundHandlerConfig {
	state: google.golang.org/protobuf/internal/impl.MessageState {
		NoUnkeyedLiterals: google.golang.org/protobuf/internal/pragma.NoUnkeyedLiterals {},
		DoNotCompare: google.golang.org/protobuf/internal/pragma.DoNotCompare [],
		DoNotCopy: google.golang.org/protobuf/internal/pragma.DoNotCopy [],
		atomicMessageInfo: *google.golang.org/protobuf/internal/impl.MessageInfo nil,},
	sizeCache: 0,
	unknownFields: []uint8 len: 0, cap: 0, nil,
	Tag: "dft",
	ReceiverSettings: *google.golang.org/protobuf/types/known/anypb.Any {
		state: (*"google.golang.org/protobuf/internal/impl.MessageState")(0x14000536eb0),
		sizeCache: 0,
		unknownFields: []uint8 len: 0, cap: 0, nil,
		TypeUrl: "types.v2fly.org/v2ray.core.app.proxyman.ReceiverConfig",
		Value: []uint8 len: 123, cap: 123, [10,6,8,144,78,16,144,78,18,6,10,4,127,0,0,1,34,105,18,92,18,79,10,62,116,121,112,101,115,46,118,50,102,108,121,46,111,114,103,47,118,50,114,97,121,46,99,111,114,101,46,116,114,97,110,115,112,111,114,116,46,105,110,116,...+59 more],},
	ProxySettings: *google.golang.org/protobuf/types/known/anypb.Any {
		state: (*"google.golang.org/protobuf/internal/impl.MessageState")(0x14000537400),
		sizeCache: 0,
		unknownFields: []uint8 len: 0, cap: 0, nil,
		TypeUrl: "types.v2fly.org/v2ray.core.proxy.vmess.inbound.Config",
		Value: []uint8 len: 107, cap: 107, [10,105,18,9,120,120,120,64,103,46,99,111,109,26,92,10,46,116,121,112,101,115,46,118,50,102,108,121,46,111,114,103,47,118,50,114,97,121,46,99,111,114,101,46,112,114,111,120,121,46,118,109,101,115,115,46,65,99,99,111,117,110,116,18,...+43 more],},}
(dlv) p config.Inbound[1].ReceiverSettings
("*google.golang.org/protobuf/types/known/anypb.Any")(0x14000536eb0)
*google.golang.org/protobuf/types/known/anypb.Any {
	state: google.golang.org/protobuf/internal/impl.MessageState {
		NoUnkeyedLiterals: google.golang.org/protobuf/internal/pragma.NoUnkeyedLiterals {},
		DoNotCompare: google.golang.org/protobuf/internal/pragma.DoNotCompare [],
		DoNotCopy: google.golang.org/protobuf/internal/pragma.DoNotCopy [],
		atomicMessageInfo: *google.golang.org/protobuf/internal/impl.MessageInfo nil,},
	sizeCache: 0,
	unknownFields: []uint8 len: 0, cap: 0, nil,
	TypeUrl: "types.v2fly.org/v2ray.core.app.proxyman.ReceiverConfig",
	Value: []uint8 len: 123, cap: 123, [10,6,8,144,78,16,144,78,18,6,10,4,127,0,0,1,34,105,18,92,18,79,10,62,116,121,112,101,115,46,118,50,102,108,121,46,111,114,103,47,118,50,114,97,121,46,99,111,114,101,46,116,114,97,110,115,112,111,114,116,46,105,110,116,...+59 more],}
(dlv) p config.Inbound[1].ReceiverSettings.Value
[]uint8 len: 123, cap: 123, [10,6,8,144,78,16,144,78,18,6,10,4,127,0,0,1,34,105,18,92,18,79,10,62,116,121,112,101,115,46,118,50,102,108,121,46,111,114,103,47,118,50,114,97,121,46,99,111,114,101,46,116,114,97,110,115,112,111,114,116,46,105,110,116,...+59 more]
```

## server

> after **initInstanceWithConfig**, print server

```
(dlv) p server
("*github.com/v2fly/v2ray-core/v5.Instance")(0x14005f95020)
*github.com/v2fly/v2ray-core/v5.Instance {
	access: sync.Mutex {state: 0, sema: 0},
	features: []github.com/v2fly/v2ray-core/v5/features.Feature len: 9, cap: 16, [
		...,
		...,
		...,
		...,
		...,
		...,
		...,
		...,
		...,
	],
	featureResolutions: []github.com/v2fly/v2ray-core/v5.resolution len: 0, cap: 0, nil,
	running: false,
	env: github.com/v2fly/v2ray-core/v5/common/environment.RootEnvironment(*github.com/v2fly/v2ray-core/v5/common/environment.rootEnvImpl) *{
		transientStorage: github.com/v2fly/v2ray-core/v5/features/extension/storage.ScopedTransientStorage(*github.com/v2fly/v2ray-core/v5/common/environment/transientstorageimpl.scopedTransientStorageImpl) ...,
		systemDialer: github.com/v2fly/v2ray-core/v5/transport/internet.SystemDialer(github.com/v2fly/v2ray-core/v5/common/environment/systemnetworkimpl.systemDefaultDialer) *(*"github.com/v2fly/v2ray-core/v5/transport/internet.SystemDialer")(0x140002d9990),
		systemListener: github.com/v2fly/v2ray-core/v5/transport/internet.SystemListener(github.com/v2fly/v2ray-core/v5/common/environment/systemnetworkimpl.systemDefaultDialer) *(*"github.com/v2fly/v2ray-core/v5/transport/internet.SystemListener")(0x140002d99a0),
		ctx: context.Context(context.backgroundCtx) *(*context.Context)(0x140002d99b0),},
	ctx: context.Context(context.backgroundCtx) {
		emptyCtx: context.emptyCtx {},},}
(dlv) p server.features
[]github.com/v2fly/v2ray-core/v5/features.Feature len: 9, cap: 16, [
	*github.com/v2fly/v2ray-core/v5/app/log.Instance {
		RWMutex: (*sync.RWMutex)(0x14005fa40a0),
		config: *(*"github.com/v2fly/v2ray-core/v5/app/log.Config")(0x140002d99c0),
		accessLogger: github.com/v2fly/v2ray-core/v5/common/log.Handler(*github.com/v2fly/v2ray-core/v5/common/log.generalLogger) ...,
		errorLogger: github.com/v2fly/v2ray-core/v5/common/log.Handler(*github.com/v2fly/v2ray-core/v5/common/log.generalLogger) ...,
		followers: map[reflect.Value]func(github.com/v2fly/v2ray-core/v5/common/log.Message) nil,
		active: true,},
	*github.com/v2fly/v2ray-core/v5/app/dispatcher.DefaultDispatcher {
		ohm: github.com/v2fly/v2ray-core/v5/features/outbound.Manager(*github.com/v2fly/v2ray-core/v5/app/proxyman/outbound.Manager) ...,
		router: github.com/v2fly/v2ray-core/v5/features/routing.Router(*github.com/v2fly/v2ray-core/v5/app/router.Router) ...,
		policy: github.com/v2fly/v2ray-core/v5/features/policy.Manager(*github.com/v2fly/v2ray-core/v5/app/policy.Instance) ...,
		stats: github.com/v2fly/v2ray-core/v5/features/stats.Manager(*github.com/v2fly/v2ray-core/v5/app/stats.Manager) ...,},
	*github.com/v2fly/v2ray-core/v5/app/proxyman/inbound.Manager {
		ctx: context.Context(*context.valueCtx) ...,
		access: (*sync.RWMutex)(0x14005fa4100),
		untaggedHandler: []github.com/v2fly/v2ray-core/v5/features/inbound.Handler len: 0, cap: 0, nil,
		taggedHandlers: map[string]github.com/v2fly/v2ray-core/v5/features/inbound.Handler [...],
		running: false,},
	*github.com/v2fly/v2ray-core/v5/app/proxyman/outbound.Manager {
		access: (*sync.RWMutex)(0x14005fa4140),
		defaultHandler: github.com/v2fly/v2ray-core/v5/features/outbound.Handler(*github.com/v2fly/v2ray-core/v5/app/proxyman/outbound.Handler) ...,
		taggedHandler: map[string]github.com/v2fly/v2ray-core/v5/features/outbound.Handler [...],
		untaggedHandlers: []github.com/v2fly/v2ray-core/v5/features/outbound.Handler len: 0, cap: 0, nil,
		running: false,},
	*github.com/v2fly/v2ray-core/v5/app/commander.Commander {
		Mutex: (*sync.Mutex)(0x14005fa4280),
		server: *google.golang.org/grpc.Server nil,
		services: []github.com/v2fly/v2ray-core/v5/app/commander.Service len: 2, cap: 2, [
			...,
			...,
		],
		ohm: github.com/v2fly/v2ray-core/v5/features/outbound.Manager(*github.com/v2fly/v2ray-core/v5/app/proxyman/outbound.Manager) ...,
		tag: "api",},
	*github.com/v2fly/v2ray-core/v5/app/stats.Manager {
		access: (*sync.RWMutex)(0x140003dade0),
		counters: map[string]*github.com/v2fly/v2ray-core/v5/app/stats.Counter [...],
		channels: map[string]*github.com/v2fly/v2ray-core/v5/app/stats.Channel [],
		running: false,},
	*github.com/v2fly/v2ray-core/v5/app/router.Router {
		domainStrategy: DomainStrategy_AsIs (0),
		rules: []*github.com/v2fly/v2ray-core/v5/app/router.Rule len: 2, cap: 2, [
			*(*"github.com/v2fly/v2ray-core/v5/app/router.Rule")(0x140003db5c0),
			*(*"github.com/v2fly/v2ray-core/v5/app/router.Rule")(0x140003db740),
		],
		balancers: map[string]*github.com/v2fly/v2ray-core/v5/app/router.Balancer [],
		dns: github.com/v2fly/v2ray-core/v5/features/dns.Client(*github.com/v2fly/v2ray-core/v5/features/dns/localdns.Client) ...,},
	*github.com/v2fly/v2ray-core/v5/app/policy.Instance {
		levels: map[uint32]*github.com/v2fly/v2ray-core/v5/app/policy.Policy [...],
		system: *(*"github.com/v2fly/v2ray-core/v5/app/policy.SystemPolicy")(0x140002d9e40),},
	*github.com/v2fly/v2ray-core/v5/features/dns/localdns.Client {},
]
```

> features
```
(dlv) p server.features[2]
github.com/v2fly/v2ray-core/v5/features.Feature(*github.com/v2fly/v2ray-core/v5/app/proxyman/inbound.Manager) *{
	ctx: context.Context(*context.valueCtx) *{
		Context: context.Context(*context.valueCtx) ...,
		key: interface {}(string) *(*interface {})(0x140003da8e0),
		val: interface {}(*github.com/v2fly/v2ray-core/v5/common/environment.appEnvImpl) ...,},
	access: sync.RWMutex {
		w: (*sync.Mutex)(0x14005fa4100),
		writerSem: 0,
		readerSem: 0,
		readerCount: (*"sync/atomic.Int32")(0x14005fa4110),
		readerWait: (*"sync/atomic.Int32")(0x14005fa4114),},
	untaggedHandler: []github.com/v2fly/v2ray-core/v5/features/inbound.Handler len: 0, cap: 0, nil,
	taggedHandlers: map[string]github.com/v2fly/v2ray-core/v5/features/inbound.Handler [
		"api": ...,
		"dft": ...,
	],
	running: false,}
(dlv) p server.features[2].taggedHandlers
map[string]github.com/v2fly/v2ray-core/v5/features/inbound.Handler [
	"api": *github.com/v2fly/v2ray-core/v5/app/proxyman/inbound.AlwaysOnInboundHandler {
		proxy: github.com/v2fly/v2ray-core/v5/proxy.Inbound(*github.com/v2fly/v2ray-core/v5/proxy/dokodemo.Door) ...,
		workers: []github.com/v2fly/v2ray-core/v5/app/proxyman/inbound.worker len: 1, cap: 1, [
			...,
		],
		mux: *(*"github.com/v2fly/v2ray-core/v5/common/mux.Server")(0x140005a31d0),
		tag: "api",},
	"dft": *github.com/v2fly/v2ray-core/v5/app/proxyman/inbound.AlwaysOnInboundHandler {
		proxy: github.com/v2fly/v2ray-core/v5/proxy.Inbound(*github.com/v2fly/v2ray-core/v5/proxy/vmess/inbound.Handler) ...,
		workers: []github.com/v2fly/v2ray-core/v5/app/proxyman/inbound.worker len: 1, cap: 1, [
			...,
		],
		mux: *(*"github.com/v2fly/v2ray-core/v5/common/mux.Server")(0x140005a3360),
		tag: "dft",},
]
(dlv) p server.features[2].taggedHandlers["dft"]
github.com/v2fly/v2ray-core/v5/features/inbound.Handler(*github.com/v2fly/v2ray-core/v5/app/proxyman/inbound.AlwaysOnInboundHandler) *{
	proxy: github.com/v2fly/v2ray-core/v5/proxy.Inbound(*github.com/v2fly/v2ray-core/v5/proxy/vmess/inbound.Handler) *{
		policyManager: github.com/v2fly/v2ray-core/v5/features/policy.Manager(*github.com/v2fly/v2ray-core/v5/app/policy.Instance) ...,
		inboundHandlerManager: github.com/v2fly/v2ray-core/v5/features/inbound.Manager(*github.com/v2fly/v2ray-core/v5/app/proxyman/inbound.Manager) ...,
		clients: *(*"github.com/v2fly/v2ray-core/v5/proxy/vmess.TimedUserValidator")(0x1400065aa10),
		usersByEmail: *(*"github.com/v2fly/v2ray-core/v5/proxy/vmess/inbound.userByEmail")(0x1400000f6b0),
		detours: *github.com/v2fly/v2ray-core/v5/proxy/vmess/inbound.DetourConfig nil,
		sessionHistory: *(*"github.com/v2fly/v2ray-core/v5/proxy/vmess/encoding.SessionHistory")(0x140003dbb90),
		secure: false,},
	workers: []github.com/v2fly/v2ray-core/v5/app/proxyman/inbound.worker len: 1, cap: 1, [
		...,
	],
	mux: *github.com/v2fly/v2ray-core/v5/common/mux.Server {
		dispatcher: github.com/v2fly/v2ray-core/v5/features/routing.Dispatcher(*github.com/v2fly/v2ray-core/v5/app/dispatcher.DefaultDispatcher) ...,},
	tag: "dft",}
(dlv) p server.features[2].taggedHandlers["dft"].workers
[]github.com/v2fly/v2ray-core/v5/app/proxyman/inbound.worker len: 1, cap: 1, [
	*github.com/v2fly/v2ray-core/v5/app/proxyman/inbound.tcpWorker {
		address: github.com/v2fly/v2ray-core/v5/common/net.Address(github.com/v2fly/v2ray-core/v5/common/net.ipv4Address) *(*"github.com/v2fly/v2ray-core/v5/common/net.Address")(0x140002f6e60),
		port: 10000,
		proxy: github.com/v2fly/v2ray-core/v5/proxy.Inbound(*github.com/v2fly/v2ray-core/v5/proxy/vmess/inbound.Handler) ...,
		stream: *(*"github.com/v2fly/v2ray-core/v5/transport/internet.MemoryStreamConfig")(0x14005fa46e0),
		recvOrigDest: false,
		tag: "dft",
		dispatcher: github.com/v2fly/v2ray-core/v5/features/routing.Dispatcher(*github.com/v2fly/v2ray-core/v5/common/mux.Server) ...,
		sniffingConfig: *github.com/v2fly/v2ray-core/v5/app/proxyman.SniffingConfig nil,
		uplinkCounter: github.com/v2fly/v2ray-core/v5/features/stats.Counter(*github.com/v2fly/v2ray-core/v5/app/stats.Counter) ...,
		downlinkCounter: github.com/v2fly/v2ray-core/v5/features/stats.Counter(*github.com/v2fly/v2ray-core/v5/app/stats.Counter) ...,
		hub: github.com/v2fly/v2ray-core/v5/transport/internet.Listener nil,
		ctx: context.Context(*context.valueCtx) ...,},
]
(dlv) p server.features[2].taggedHandlers["dft"].proxy
github.com/v2fly/v2ray-core/v5/proxy.Inbound(*github.com/v2fly/v2ray-core/v5/proxy/vmess/inbound.Handler) *{
	policyManager: github.com/v2fly/v2ray-core/v5/features/policy.Manager(*github.com/v2fly/v2ray-core/v5/app/policy.Instance) *{
		levels: map[uint32]*github.com/v2fly/v2ray-core/v5/app/policy.Policy [...],
		system: *(*"github.com/v2fly/v2ray-core/v5/app/policy.SystemPolicy")(0x140002d9e40),},
	inboundHandlerManager: github.com/v2fly/v2ray-core/v5/features/inbound.Manager(*github.com/v2fly/v2ray-core/v5/app/proxyman/inbound.Manager) *{
		ctx: context.Context(*context.valueCtx) ...,
		access: (*sync.RWMutex)(0x14005fa4100),
		untaggedHandler: []github.com/v2fly/v2ray-core/v5/features/inbound.Handler len: 0, cap: 0, nil,
		taggedHandlers: map[string]github.com/v2fly/v2ray-core/v5/features/inbound.Handler [...],
		running: false,},
	clients: *github.com/v2fly/v2ray-core/v5/proxy/vmess.TimedUserValidator {
		RWMutex: (*sync.RWMutex)(0x1400065aa10),
		users: []*github.com/v2fly/v2ray-core/v5/proxy/vmess.user len: 1, cap: 16, [
			*(*"github.com/v2fly/v2ray-core/v5/proxy/vmess.user")(0x140003dbc80),
		],
		userHash: map[[16]uint8]github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair [...],
		hasher: github.com/v2fly/v2ray-core/v5/common/protocol.DefaultIDHash,
		baseTime: 1703826520,
		task: *(*"github.com/v2fly/v2ray-core/v5/common/task.Periodic")(0x140003dbb00),
		behaviorSeed: 2487058318055934096,
		behaviorFused: false,
		aeadDecoderHolder: *(*"github.com/v2fly/v2ray-core/v5/proxy/vmess/aead.AuthIDDecoderHolder")(0x140005a32d0),
		legacyWarnShown: false,},
	usersByEmail: *github.com/v2fly/v2ray-core/v5/proxy/vmess/inbound.userByEmail {
		Mutex: (*sync.Mutex)(0x1400000f6b0),
		cache: map[string]*github.com/v2fly/v2ray-core/v5/common/protocol.MemoryUser [...],
		defaultLevel: 0,
		defaultAlterIDs: 0,},
	detours: *github.com/v2fly/v2ray-core/v5/proxy/vmess/inbound.DetourConfig nil,
	sessionHistory: *github.com/v2fly/v2ray-core/v5/proxy/vmess/encoding.SessionHistory {
		RWMutex: (*sync.RWMutex)(0x140003dbb90),
		cache: map[github.com/v2fly/v2ray-core/v5/proxy/vmess/encoding.sessionID]time.Time [],
		task: *(*"github.com/v2fly/v2ray-core/v5/common/task.Periodic")(0x140003dbbc0),},
	secure: false,}
(dlv) p server.features[2].taggedHandlers["dft"].proxy.clients
("*github.com/v2fly/v2ray-core/v5/proxy/vmess.TimedUserValidator")(0x1400065aa10)
*github.com/v2fly/v2ray-core/v5/proxy/vmess.TimedUserValidator {
	RWMutex: sync.RWMutex {
		w: (*sync.Mutex)(0x1400065aa10),
		writerSem: 0,
		readerSem: 0,
		readerCount: (*"sync/atomic.Int32")(0x1400065aa20),
		readerWait: (*"sync/atomic.Int32")(0x1400065aa24),},
	users: []*github.com/v2fly/v2ray-core/v5/proxy/vmess.user len: 1, cap: 16, [
		*(*"github.com/v2fly/v2ray-core/v5/proxy/vmess.user")(0x140003dbc80),
	],
	userHash: map[[16]uint8]github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair [
		[92,23,91,77,11,242,195,194,162,41,131,209,223,245,93,175]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fae1d8),
		[196,243,253,1,36,133,52,144,26,104,228,83,16,69,54,220]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fae1f0),
		[29,76,208,58,243,90,109,4,220,217,235,101,144,220,177,18]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fae478),
		[165,19,139,200,46,221,50,185,223,113,54,77,122,1,50,92]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fae5c8),
		[181,226,61,61,107,91,142,246,130,207,86,65,120,17,103,31]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fae718),
		[178,233,81,80,186,45,100,178,212,18,133,184,107,231,97,85]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fae730),
		[87,122,248,239,5,92,93,245,58,254,2,99,225,201,239,7]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fae9b8),
		[239,61,180,226,108,220,130,51,162,133,115,199,143,180,121,54]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fae9d0),
		[186,171,28,4,173,242,185,78,149,51,253,52,252,98,107,104]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fae9e8),
		[154,68,190,177,19,32,123,180,228,248,153,235,154,134,94,200]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005faec58),
		[56,11,225,166,149,40,48,236,69,132,250,3,85,81,138,3]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005faec70),
		[122,109,45,36,27,81,121,164,218,104,251,181,120,152,134,148]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005faf048),
		[120,61,70,139,63,49,18,65,127,13,248,16,123,120,255,146]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005faf060),
		[86,243,173,115,218,210,249,155,223,134,202,206,85,194,159,235]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005faf198),
		[62,71,233,24,205,162,214,159,236,29,127,37,25,7,255,77]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005faf588),
		[63,224,218,154,114,63,142,172,99,23,73,117,22,252,29,102]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005faf6d8),
		[25,200,76,14,69,49,100,7,45,225,227,235,117,249,189,140]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005faf6f0),
		[235,111,159,129,152,163,155,45,132,138,144,3,35,168,244,194]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005faf978),
		[90,247,96,16,107,49,167,20,111,35,4,58,186,117,73,24]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005faf990),
		[20,55,254,58,33,192,252,0,203,155,119,198,67,161,184,198]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005faf9a8),
		[71,99,71,65,135,246,129,200,63,53,67,222,33,58,232,21]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fafd68),
		[173,26,248,11,4,20,235,2,252,239,4,56,203,131,225,250]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb0008),
		[146,141,186,75,201,68,35,225,51,183,99,17,162,213,95,144]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb0020),
		[83,84,61,35,3,178,60,107,53,192,172,146,145,54,130,224]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb0038),
		[140,49,37,17,230,189,191,143,190,25,237,252,231,127,249,124]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb0548),
		[197,3,114,67,109,209,205,50,149,117,111,42,226,6,58,151]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb0560),
		[67,53,103,172,129,159,50,9,141,170,175,163,189,21,193,141]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb0578),
		[46,147,234,171,222,131,103,30,160,169,221,211,65,197,227,232]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb0698),
		[67,225,85,218,227,178,223,121,155,59,43,222,131,199,194,6]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb07e8),
		[120,48,3,131,139,33,13,252,155,80,239,87,117,233,158,44]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb0938),
		[40,158,219,199,42,197,137,247,166,101,148,190,135,195,155,234]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb0a88),
		[231,236,170,244,7,225,0,36,29,125,207,63,104,182,50,69]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb0bd8),
		[86,247,144,55,96,253,94,30,2,119,7,65,210,54,1,220]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb0bf0),
		[244,132,191,62,120,155,10,104,51,250,147,44,49,111,201,223]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb0c08),
		[156,2,11,7,51,44,31,157,219,84,27,232,142,49,130,194]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb0c20),
		[252,125,191,121,161,127,220,109,199,175,225,221,75,61,222,85]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb0e78),
		[14,16,20,111,113,2,64,112,72,255,147,125,228,189,198,23]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb1118),
		[31,236,81,191,154,125,178,113,14,200,111,196,249,59,115,49]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb1268),
		[129,143,241,221,66,46,59,56,107,90,96,154,202,193,44,11]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb1280),
		[176,233,50,147,114,237,219,1,175,80,16,4,92,206,179,61]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb13b8),
		[76,211,190,157,55,2,68,135,38,119,7,211,145,137,180,207]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb13d0),
		[25,253,73,177,7,37,164,58,251,251,74,207,89,231,204,71]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb17a8),
		[191,18,198,157,20,211,89,127,53,104,13,217,36,16,174,187]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb18f8),
		[145,192,50,129,247,174,59,70,52,139,161,245,160,74,58,252]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb1a48),
		[120,178,123,53,71,225,189,147,34,87,5,165,171,195,141,171]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb1a60),
		[5,121,148,119,154,170,73,249,236,7,50,212,221,54,51,21]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb1ce8),
		[159,104,244,155,255,134,196,154,219,151,66,68,75,120,2,232]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb1d00),
		[99,107,104,149,174,56,202,139,50,60,148,241,25,206,55,244]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb1e38),
		[192,23,54,152,31,185,31,138,210,60,195,106,148,150,44,111]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb1e50),
		[162,34,98,177,232,170,208,252,73,169,143,134,143,191,146,125]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb1e68),
		[39,104,29,64,92,254,162,224,157,114,85,253,31,236,30,5]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb1e80),
		[153,131,18,85,122,231,254,226,165,224,25,15,37,124,159,222]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb2618),
		[240,237,105,130,190,138,61,8,205,13,171,218,252,21,81,212]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb2768),
		[40,119,223,122,234,51,158,105,131,85,7,204,121,164,72,162]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb2780),
		[79,246,81,85,27,99,128,185,80,69,81,222,36,128,166,16]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb28b8),
		[199,143,196,125,189,161,157,144,128,14,159,247,190,219,162,56]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb2a08),
		[113,73,24,73,24,171,79,162,4,50,197,38,211,130,204,211]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb2b58),
		[87,247,182,113,73,241,104,72,45,93,151,143,248,110,244,179]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb2df8),
		[46,160,17,226,204,245,85,125,132,2,217,197,140,142,14,201]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb2f48),
		[2,171,8,204,114,144,77,136,221,247,219,128,204,37,7,50]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb2f60),
		[9,180,141,106,59,206,161,135,132,231,224,236,244,160,221,186]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb2f78),
		[219,69,21,213,216,170,23,40,47,231,231,146,95,57,90,56]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb31e8),
		[15,95,1,10,178,130,107,67,16,154,81,248,130,120,165,216]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb3200),
		[248,181,37,170,218,238,55,47,85,59,128,92,13,48,1,54]: (*"github.com/v2fly/v2ray-core/v5/proxy/vmess.indexTimePair")(0x14005fb3218),
		...+177 more
	],
	hasher: github.com/v2fly/v2ray-core/v5/common/protocol.DefaultIDHash,
	baseTime: 1703826520,
	task: *github.com/v2fly/v2ray-core/v5/common/task.Periodic {
		Interval: github.com/v2fly/v2ray-core/v5/proxy/vmess.updateInterval (10000000000),
		Execute: github.com/v2fly/v2ray-core/v5/proxy/vmess.NewTimedUserValidator.func1,
		access: (*sync.Mutex)(0x140003dbb10),
		timer: *(*time.Timer)(0x14005fa4640),
		running: true,},
	behaviorSeed: 2487058318055934096,
	behaviorFused: false,
	aeadDecoderHolder: *github.com/v2fly/v2ray-core/v5/proxy/vmess/aead.AuthIDDecoderHolder {
		decoders: map[string]*github.com/v2fly/v2ray-core/v5/proxy/vmess/aead.AuthIDDecoderItem [...],
		filter: *(*"github.com/v2fly/v2ray-core/v5/common/antireplay.ReplayFilter")(0x140003dbad0),},
	legacyWarnShown: false,}
(dlv) p server.features[2].taggedHandlers["dft"].proxy.clients.users
[]*github.com/v2fly/v2ray-core/v5/proxy/vmess.user len: 1, cap: 16, [
	*{
		user: (*"github.com/v2fly/v2ray-core/v5/common/protocol.MemoryUser")(0x140003dbc80),
		lastSec: 1703826880,},
]
(dlv) p server.features[2].taggedHandlers["dft"].proxy.clients.users[0]
("*github.com/v2fly/v2ray-core/v5/proxy/vmess.user")(0x140003dbc80)
*github.com/v2fly/v2ray-core/v5/proxy/vmess.user {
	user: github.com/v2fly/v2ray-core/v5/common/protocol.MemoryUser {
		Account: github.com/v2fly/v2ray-core/v5/common/protocol.Account(*github.com/v2fly/v2ray-core/v5/proxy/vmess.MemoryAccount) ...,
		Email: "xxx@g.com",
		Level: 0,},
	lastSec: 1703826880,}
(dlv) p server.features[2].taggedHandlers["dft"].proxy.clients.users[0].user
github.com/v2fly/v2ray-core/v5/common/protocol.MemoryUser {
	Account: github.com/v2fly/v2ray-core/v5/common/protocol.Account(*github.com/v2fly/v2ray-core/v5/proxy/vmess.MemoryAccount) *{
		ID: *(*"github.com/v2fly/v2ray-core/v5/common/protocol.ID")(0x14000616660),
		AlterIDs: []*github.com/v2fly/v2ray-core/v5/common/protocol.ID len: 0, cap: 0, [],
		Security: SecurityType_CHACHA20_POLY1305 (4),
		AuthenticatedLengthExperiment: false,
		NoTerminationSignal: false,},
	Email: "xxx@g.com",
	Level: 0,}
(dlv) p server.features[2].taggedHandlers["dft"].proxy.clients.users[0].user.Account
github.com/v2fly/v2ray-core/v5/common/protocol.Account(*github.com/v2fly/v2ray-core/v5/proxy/vmess.MemoryAccount) *{
	ID: *github.com/v2fly/v2ray-core/v5/common/protocol.ID {
		uuid: github.com/v2fly/v2ray-core/v5/common/uuid.UUID [222,32,217,55,202,143,175,20,234,7,32,180,84,71,211,113],
		cmdKey: [16]uint8 [254,181,240,101,221,227,114,180,98,22,68,23,39,164,31,18],},
	AlterIDs: []*github.com/v2fly/v2ray-core/v5/common/protocol.ID len: 0, cap: 0, [],
	Security: SecurityType_CHACHA20_POLY1305 (4),
	AuthenticatedLengthExperiment: false,
	NoTerminationSignal: false,}
(dlv) p server.features[2].taggedHandlers["dft"].proxy.clients.users[0].user.Account.ID
("*github.com/v2fly/v2ray-core/v5/common/protocol.ID")(0x14000616660)
*github.com/v2fly/v2ray-core/v5/common/protocol.ID {
	uuid: github.com/v2fly/v2ray-core/v5/common/uuid.UUID [222,32,217,55,202,143,175,20,234,7,32,180,84,71,211,113],
	cmdKey: [16]uint8 [254,181,240,101,221,227,114,180,98,22,68,23,39,164,31,18],}
(dlv) p server.features[2].taggedHandlers["dft"].proxy.clients.users[0].user.Account.ID.uuid
github.com/v2fly/v2ray-core/v5/common/uuid.UUID [222,32,217,55,202,143,175,20,234,7,32,180,84,71,211,113]
```


## server start

```
   352:	func (s *Instance) Start() error {
   353:		s.access.Lock()
   354:		defer s.access.Unlock()
   355:
   356:		s.running = true
   357:		for _, f := range s.features {
=> 358:			if err := f.Start(); err != nil {
   359:				return err
   360:			}
   361:		}
```

### inbound

> github.com/v2fly/v2ray-core/v5/app/proxyman/inbound.Manager

```
github.com/v2fly/v2ray-core/v5/features.Feature(*github.com/v2fly/v2ray-core/v5/app/proxyman/inbound.Manager) *{
	ctx: context.Context(*context.valueCtx) *{
		Context: context.Context(*context.valueCtx) ...,
		key: interface {}(string) *(*interface {})(0x140003da8e0),
		val: interface {}(*github.com/v2fly/v2ray-core/v5/common/environment.appEnvImpl) ...,},
	access: sync.RWMutex {
		w: (*sync.Mutex)(0x14005fa4100),
		writerSem: 0,
		readerSem: 0,
		readerCount: (*"sync/atomic.Int32")(0x14005fa4110),
		readerWait: (*"sync/atomic.Int32")(0x14005fa4114),},
	untaggedHandler: []github.com/v2fly/v2ray-core/v5/features/inbound.Handler len: 0, cap: 0, nil,
	taggedHandlers: map[string]github.com/v2fly/v2ray-core/v5/features/inbound.Handler [
		"api": ...,
		"dft": ...,
	],
	running: false,}
```

### inbound handler start

```
    91:	// Start implements common.Runnable.
    92:	func (m *Manager) Start() error {
    93:		m.access.Lock()
    94:		defer m.access.Unlock()
    95:
    96:		m.running = true
    97:
    98:		for _, handler := range m.taggedHandlers {
=>  99:			if err := handler.Start(); err != nil {
   100:				return err
   101:			}
   102:		}
```

```
(dlv) p handler
github.com/v2fly/v2ray-core/v5/features/inbound.Handler(*github.com/v2fly/v2ray-core/v5/app/proxyman/inbound.AlwaysOnInboundHandler) *{
	proxy: github.com/v2fly/v2ray-core/v5/proxy.Inbound(*github.com/v2fly/v2ray-core/v5/proxy/dokodemo.Door) *{
		policyManager: github.com/v2fly/v2ray-core/v5/features/policy.Manager(*github.com/v2fly/v2ray-core/v5/app/policy.Instance) ...,
		config: *(*"github.com/v2fly/v2ray-core/v5/proxy/dokodemo.Config")(0x1400065a930),
		address: github.com/v2fly/v2ray-core/v5/common/net.Address(github.com/v2fly/v2ray-core/v5/common/net.ipv4Address) *(*"github.com/v2fly/v2ray-core/v5/common/net.Address")(0x14000112f98),
		port: 0,
		sockopt: *github.com/v2fly/v2ray-core/v5/common/session.Sockopt nil,},
	workers: []github.com/v2fly/v2ray-core/v5/app/proxyman/inbound.worker len: 1, cap: 1, [
		...,
	],
	mux: *github.com/v2fly/v2ray-core/v5/common/mux.Server {
		dispatcher: github.com/v2fly/v2ray-core/v5/features/routing.Dispatcher(*github.com/v2fly/v2ray-core/v5/app/dispatcher.DefaultDispatcher) ...,},
	tag: "api",}
```

### outbound

> github.com/v2fly/v2ray-core/v5/app/proxyman/outbound.Manager

```
(dlv) p f
github.com/v2fly/v2ray-core/v5/features.Feature(*github.com/v2fly/v2ray-core/v5/app/proxyman/outbound.Manager) *{
	access: sync.RWMutex {
		w: (*sync.Mutex)(0x14005fa4140),
		writerSem: 0,
		readerSem: 0,
		readerCount: (*"sync/atomic.Int32")(0x14005fa4150),
		readerWait: (*"sync/atomic.Int32")(0x14005fa4154),},
	defaultHandler: github.com/v2fly/v2ray-core/v5/features/outbound.Handler(*github.com/v2fly/v2ray-core/v5/app/proxyman/outbound.Handler) *{
		ctx: context.Context(*context.valueCtx) ...,
		tag: "allowed",
		senderSettings: *(*"github.com/v2fly/v2ray-core/v5/app/proxyman.SenderConfig")(0x14005fa4730),
		streamSettings: *(*"github.com/v2fly/v2ray-core/v5/transport/internet.MemoryStreamConfig")(0x14005fa4780),
		proxy: github.com/v2fly/v2ray-core/v5/proxy.Outbound(*github.com/v2fly/v2ray-core/v5/proxy/freedom.Handler) ...,
		outboundManager: github.com/v2fly/v2ray-core/v5/features/outbound.Manager(*github.com/v2fly/v2ray-core/v5/app/proxyman/outbound.Manager) ...,
		mux: *github.com/v2fly/v2ray-core/v5/common/mux.ClientManager nil,
		uplinkCounter: github.com/v2fly/v2ray-core/v5/features/stats.Counter nil,
		downlinkCounter: github.com/v2fly/v2ray-core/v5/features/stats.Counter nil,
		dns: github.com/v2fly/v2ray-core/v5/features/dns.Client nil,},
	taggedHandler: map[string]github.com/v2fly/v2ray-core/v5/features/outbound.Handler [
		"allowed": ...,
		"blocked": ...,
	],
	untaggedHandlers: []github.com/v2fly/v2ray-core/v5/features/outbound.Handler len: 0, cap: 0, nil,
	running: false,}
(dlv) p f.defaultHandler
github.com/v2fly/v2ray-core/v5/features/outbound.Handler(*github.com/v2fly/v2ray-core/v5/app/proxyman/outbound.Handler) *{
	ctx: context.Context(*context.valueCtx) *{
		Context: context.Context(*context.valueCtx) ...,
		key: interface {}(string) *(*interface {})(0x140003dbdb0),
		val: interface {}(*github.com/v2fly/v2ray-core/v5/common/environment.proxyEnvImpl) ...,},
	tag: "allowed",
	senderSettings: *github.com/v2fly/v2ray-core/v5/app/proxyman.SenderConfig {
		state: (*"google.golang.org/protobuf/internal/impl.MessageState")(0x14005fa4730),
		sizeCache: 0,
		unknownFields: []uint8 len: 0, cap: 0, nil,
		Via: *github.com/v2fly/v2ray-core/v5/common/net.IPOrDomain nil,
		StreamSettings: *github.com/v2fly/v2ray-core/v5/transport/internet.StreamConfig nil,
		ProxySettings: *github.com/v2fly/v2ray-core/v5/transport/internet.ProxyConfig nil,
		MultiplexSettings: *github.com/v2fly/v2ray-core/v5/app/proxyman.MultiplexingConfig nil,
		DomainStrategy: SenderConfig_AS_IS (0),},
	streamSettings: *github.com/v2fly/v2ray-core/v5/transport/internet.MemoryStreamConfig {
		ProtocolName: "tcp",
		ProtocolSettings: interface {}(*github.com/v2fly/v2ray-core/v5/transport/internet/tcp.Config) ...,
		SecurityType: "",
		SecuritySettings: interface {} nil,
		SocketSettings: *github.com/v2fly/v2ray-core/v5/transport/internet.SocketConfig nil,},
	proxy: github.com/v2fly/v2ray-core/v5/proxy.Outbound(*github.com/v2fly/v2ray-core/v5/proxy/freedom.Handler) *{
		policyManager: github.com/v2fly/v2ray-core/v5/features/policy.Manager(*github.com/v2fly/v2ray-core/v5/app/policy.Instance) ...,
		dns: github.com/v2fly/v2ray-core/v5/features/dns.Client(*github.com/v2fly/v2ray-core/v5/features/dns/localdns.Client) ...,
		config: *(*"github.com/v2fly/v2ray-core/v5/proxy/freedom.Config")(0x14000113900),},
	outboundManager: github.com/v2fly/v2ray-core/v5/features/outbound.Manager(*github.com/v2fly/v2ray-core/v5/app/proxyman/outbound.Manager) *{
		access: (*sync.RWMutex)(0x14005fa4140),
		defaultHandler: github.com/v2fly/v2ray-core/v5/features/outbound.Handler(*github.com/v2fly/v2ray-core/v5/app/proxyman/outbound.Handler) ...,
		taggedHandler: map[string]github.com/v2fly/v2ray-core/v5/features/outbound.Handler [...],
		untaggedHandlers: []github.com/v2fly/v2ray-core/v5/features/outbound.Handler len: 0, cap: 0, nil,
		running: false,},
	mux: *github.com/v2fly/v2ray-core/v5/common/mux.ClientManager nil,
	uplinkCounter: github.com/v2fly/v2ray-core/v5/features/stats.Counter nil,
	downlinkCounter: github.com/v2fly/v2ray-core/v5/features/stats.Counter nil,
	dns: github.com/v2fly/v2ray-core/v5/features/dns.Client nil,}
(dlv) p f.defaultHandler.streamSettings
("*github.com/v2fly/v2ray-core/v5/transport/internet.MemoryStreamConfig")(0x14005fa4780)
*github.com/v2fly/v2ray-core/v5/transport/internet.MemoryStreamConfig {
	ProtocolName: "tcp",
	ProtocolSettings: interface {}(*github.com/v2fly/v2ray-core/v5/transport/internet/tcp.Config) *{
		state: (*"google.golang.org/protobuf/internal/impl.MessageState")(0x140001138c0),
		sizeCache: 0,
		unknownFields: []uint8 len: 0, cap: 0, nil,
		HeaderSettings: *google.golang.org/protobuf/types/known/anypb.Any nil,
		AcceptProxyProtocol: false,},
	SecurityType: "",
	SecuritySettings: interface {} nil,
	SocketSettings: *github.com/v2fly/v2ray-core/v5/transport/internet.SocketConfig nil,}
```

### outbound handler start

```
    40:	// Start implements core.Feature
    41:	func (m *Manager) Start() error {
    42:		m.access.Lock()
    43:		defer m.access.Unlock()
    44:
    45:		m.running = true
    46:
    47:		for _, h := range m.taggedHandler {
=>  48:			if err := h.Start(); err != nil {
    49:				return err
    50:			}
    51:		}
```

```
(dlv) p h
github.com/v2fly/v2ray-core/v5/features/outbound.Handler(*github.com/v2fly/v2ray-core/v5/app/proxyman/outbound.Handler) *{
	ctx: context.Context(*context.valueCtx) *{
		Context: context.Context(*context.valueCtx) ...,
		key: interface {}(string) *(*interface {})(0x140003dbdb0),
		val: interface {}(*github.com/v2fly/v2ray-core/v5/common/environment.proxyEnvImpl) ...,},
	tag: "allowed",
	senderSettings: *github.com/v2fly/v2ray-core/v5/app/proxyman.SenderConfig {
		state: (*"google.golang.org/protobuf/internal/impl.MessageState")(0x14005fa4730),
		sizeCache: 0,
		unknownFields: []uint8 len: 0, cap: 0, nil,
		Via: *github.com/v2fly/v2ray-core/v5/common/net.IPOrDomain nil,
		StreamSettings: *github.com/v2fly/v2ray-core/v5/transport/internet.StreamConfig nil,
		ProxySettings: *github.com/v2fly/v2ray-core/v5/transport/internet.ProxyConfig nil,
		MultiplexSettings: *github.com/v2fly/v2ray-core/v5/app/proxyman.MultiplexingConfig nil,
		DomainStrategy: SenderConfig_AS_IS (0),},
	streamSettings: *github.com/v2fly/v2ray-core/v5/transport/internet.MemoryStreamConfig {
		ProtocolName: "tcp",
		ProtocolSettings: interface {}(*github.com/v2fly/v2ray-core/v5/transport/internet/tcp.Config) ...,
		SecurityType: "",
		SecuritySettings: interface {} nil,
		SocketSettings: *github.com/v2fly/v2ray-core/v5/transport/internet.SocketConfig nil,},
	proxy: github.com/v2fly/v2ray-core/v5/proxy.Outbound(*github.com/v2fly/v2ray-core/v5/proxy/freedom.Handler) *{
		policyManager: github.com/v2fly/v2ray-core/v5/features/policy.Manager(*github.com/v2fly/v2ray-core/v5/app/policy.Instance) ...,
		dns: github.com/v2fly/v2ray-core/v5/features/dns.Client(*github.com/v2fly/v2ray-core/v5/features/dns/localdns.Client) ...,
		config: *(*"github.com/v2fly/v2ray-core/v5/proxy/freedom.Config")(0x14000113900),},
	outboundManager: github.com/v2fly/v2ray-core/v5/features/outbound.Manager(*github.com/v2fly/v2ray-core/v5/app/proxyman/outbound.Manager) *{
		access: (*sync.RWMutex)(0x14005fa4140),
		defaultHandler: github.com/v2fly/v2ray-core/v5/features/outbound.Handler(*github.com/v2fly/v2ray-core/v5/app/proxyman/outbound.Handler) ...,
		taggedHandler: map[string]github.com/v2fly/v2ray-core/v5/features/outbound.Handler [...],
		untaggedHandlers: []github.com/v2fly/v2ray-core/v5/features/outbound.Handler len: 0, cap: 0, nil,
		running: true,},
	mux: *github.com/v2fly/v2ray-core/v5/common/mux.ClientManager nil,
	uplinkCounter: github.com/v2fly/v2ray-core/v5/features/stats.Counter nil,
	downlinkCounter: github.com/v2fly/v2ray-core/v5/features/stats.Counter nil,
	dns: github.com/v2fly/v2ray-core/v5/features/dns.Client nil,}
```
