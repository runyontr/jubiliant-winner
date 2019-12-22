# Litecoin Docker Image

## Dockerfile

Much improvement could be made to reduce the layer sizes of the Docker image by combining multiple lines into one command. However,
due to my in-experience with the product, I believed a longer/verbose Dockerfile provided better clarity on what was occuring
in the build process.

The configuration file currently contains an RPCUser and Password. This is not ideal, but didn't see a way to inject additional information
into the runtime outside of this configuration file.

## Makefile

The makefile provided here has four simple commands:

- _build_ - build the docker image
- _push_ - pushes the Docker image to the configured registry
- _scan_ - provides a Trivy Scan (output below)

## Trivy Scan

Using [Trivy](https://github.com/aquasecurity/trivy) from Aqua Security, we see the following vulnerabilities:

```bash
$ make scan
trivy -severity HIGH,CRITICAL runyonsolutions/litecoind:0.17.1
2019-12-22T05:55:58.680-0500    INFO    Detecting Ubuntu vulnerabilities...

runyonsolutions/litecoind:0.17.1 (ubuntu 18.04)
===============================================
Total: 16 (UNKNOWN: 0, LOW: 0, MEDIUM: 0, HIGH: 16, CRITICAL: 0)

+---------------+------------------+----------+---------------------+---------------+--------------------------------+
|    LIBRARY    | VULNERABILITY ID | SEVERITY |  INSTALLED VERSION  | FIXED VERSION |             TITLE              |
+---------------+------------------+----------+---------------------+---------------+--------------------------------+
| bsdutils      | CVE-2018-7738    | HIGH     | 2.31.1-0.4ubuntu3.4 |               | util-linux: Shell command      |
|               |                  |          |                     |               | injection in unescaped         |
|               |                  |          |                     |               | bash-completed mount point     |
|               |                  |          |                     |               | names                          |
+---------------+------------------+          +---------------------+---------------+--------------------------------+
| dpkg          | CVE-2017-8283    |          | 1.19.0.5ubuntu2.3   |               | dpkg-source in dpkg 1.3.0      |
|               |                  |          |                     |               | through 1.18.23 is able to use |
|               |                  |          |                     |               | a non-GNU...                   |
+---------------+------------------+          +---------------------+---------------+--------------------------------+
| fdisk         | CVE-2018-7738    |          | 2.31.1-0.4ubuntu3.4 |               | util-linux: Shell command      |
|               |                  |          |                     |               | injection in unescaped         |
|               |                  |          |                     |               | bash-completed mount point     |
|               |                  |          |                     |               | names                          |
+---------------+                  +          +                     +---------------+                                +
| libblkid1     |                  |          |                     |               |                                |
|               |                  |          |                     |               |                                |
|               |                  |          |                     |               |                                |
|               |                  |          |                     |               |                                |
+---------------+------------------+          +---------------------+---------------+--------------------------------+
| libc-bin      | CVE-2018-11236   |          | 2.27-3ubuntu1       |               | glibc: Integer overflow in     |
|               |                  |          |                     |               | stdlib/canonicalize.c on       |
|               |                  |          |                     |               | 32-bit architectures leading   |
|               |                  |          |                     |               | to stack-based buffer...       |
+               +------------------+          +                     +---------------+--------------------------------+
|               | CVE-2019-9169    |          |                     |               | glibc: regular-expression      |
|               |                  |          |                     |               | match via proceed_next_node    |
|               |                  |          |                     |               | in posix/regexec.c leads to    |
|               |                  |          |                     |               | heap-based buffer over-read... |
+---------------+------------------+          +                     +---------------+--------------------------------+
| libc6         | CVE-2018-11236   |          |                     |               | glibc: Integer overflow in     |
|               |                  |          |                     |               | stdlib/canonicalize.c on       |
|               |                  |          |                     |               | 32-bit architectures leading   |
|               |                  |          |                     |               | to stack-based buffer...       |
+               +------------------+          +                     +---------------+--------------------------------+
|               | CVE-2019-9169    |          |                     |               | glibc: regular-expression      |
|               |                  |          |                     |               | match via proceed_next_node    |
|               |                  |          |                     |               | in posix/regexec.c leads to    |
|               |                  |          |                     |               | heap-based buffer over-read... |
+---------------+------------------+          +---------------------+---------------+--------------------------------+
| libfdisk1     | CVE-2018-7738    |          | 2.31.1-0.4ubuntu3.4 |               | util-linux: Shell command      |
|               |                  |          |                     |               | injection in unescaped         |
|               |                  |          |                     |               | bash-completed mount point     |
|               |                  |          |                     |               | names                          |
+---------------+                  +          +                     +---------------+                                +
| libmount1     |                  |          |                     |               |                                |
|               |                  |          |                     |               |                                |
|               |                  |          |                     |               |                                |
|               |                  |          |                     |               |                                |
+---------------+------------------+          +---------------------+---------------+--------------------------------+
| libpcre3      | CVE-2017-11164   |          | 2:8.39-9            |               | pcre: OP_KETRMAX feature       |
|               |                  |          |                     |               | in the match function in       |
|               |                  |          |                     |               | pcre_exec.c                    |
+---------------+------------------+          +---------------------+---------------+--------------------------------+
| libsmartcols1 | CVE-2018-7738    |          | 2.31.1-0.4ubuntu3.4 |               | util-linux: Shell command      |
|               |                  |          |                     |               | injection in unescaped         |
|               |                  |          |                     |               | bash-completed mount point     |
|               |                  |          |                     |               | names                          |
+---------------+------------------+          +---------------------+---------------+--------------------------------+
| libtasn1-6    | CVE-2018-1000654 |          | 4.13-2              |               | libtasn1: Infinite loop in     |
|               |                  |          |                     |               | _asn1_expand_object_id(ptree)  |
|               |                  |          |                     |               | leads to memory exhaustion     |
+---------------+------------------+          +---------------------+---------------+--------------------------------+
| libuuid1      | CVE-2018-7738    |          | 2.31.1-0.4ubuntu3.4 |               | util-linux: Shell command      |
|               |                  |          |                     |               | injection in unescaped         |
|               |                  |          |                     |               | bash-completed mount point     |
|               |                  |          |                     |               | names                          |
+---------------+                  +          +                     +---------------+                                +
| mount         |                  |          |                     |               |                                |
|               |                  |          |                     |               |                                |
|               |                  |          |                     |               |                                |
|               |                  |          |                     |               |                                |
+---------------+                  +          +                     +---------------+                                +
| util-linux    |                  |          |                     |               |                                |
|               |                  |          |                     |               |                                |
|               |                  |          |                     |               |                                |
|               |                  |          |                     |               |                                |
+---------------+------------------+----------+---------------------+---------------+--------------------------------+
```

Other (smaller) images could be looked at to reduce the risk exposure, but without guides on installing on `alpine`, I felt it was out of scope
for this exercise.

## Running

```bash
$ docker run -it --rm runyonsolutions/litecoind:0.17.1
2019-12-22T11:47:09Z Litecoin Core version v0.17.1 (release build)
2019-12-22T11:47:09Z InitParameterInteraction: parameter interaction: -bind set -> setting -listen=1
2019-12-22T11:47:09Z InitParameterInteraction: parameter interaction: -whitelistforcerelay=1 -> setting -whitelistrelay=1
2019-12-22T11:47:09Z Assuming ancestors of block a601455787cb65ffc325dda4751a99cf01d1567799ec4b04f45bb05f9ef0cbde have valid signatures.
2019-12-22T11:47:09Z Setting nMinimumChainWork=0000000000000000000000000000000000000000000001488d32719b8eb30150
2019-12-22T11:47:09Z Using the 'sse4(1way),sse41(4way),avx2(8way)' SHA256 implementation
2019-12-22T11:47:09Z Using RdRand as an additional entropy source
2019-12-22T11:47:09Z Default data directory /home/litecoin/.litecoin
2019-12-22T11:47:09Z Using data directory /data/litecoin
2019-12-22T11:47:09Z Using config file /data/litecoin/litecoin.conf
2019-12-22T11:47:09Z Using at most 125 automatic connections (1048576 file descriptors available)
2019-12-22T11:47:09Z Using 16 MiB out of 32/2 requested for signature cache, able to store 524288 elements
2019-12-22T11:47:09Z Using 16 MiB out of 32/2 requested for script execution cache, able to store 524288 elements
2019-12-22T11:47:09Z Using 6 threads for script verification
2019-12-22T11:47:09Z scheduler thread start
2019-12-22T11:47:09Z WARNING: option -rpcbind was ignored because -rpcallowip was not specified, refusing to allow everyone to connect
2019-12-22T11:47:09Z libevent: getaddrinfo: address family for nodename not supported
2019-12-22T11:47:09Z Binding RPC on address ::1 port 9332 failed.
2019-12-22T11:47:09Z HTTP: creating work queue of depth 16
2019-12-22T11:47:09Z Config options rpcuser and rpcpassword will soon be deprecated. Locally-run instances may remove rpcuser to use cookie-based auth, or may be replaced with rpcauth. Please see share/rpcauth for rpcauth auth generation.
2019-12-22T11:47:09Z HTTP: starting 4 worker threads
2019-12-22T11:47:09Z Using wallet directory /data/litecoin
2019-12-22T11:47:09Z init message: Verifying wallet(s)...
2019-12-22T11:47:09Z Using BerkeleyDB version Berkeley DB 4.8.30: (April  9, 2010)
2019-12-22T11:47:09Z Using wallet wallet.dat
2019-12-22T11:47:09Z BerkeleyEnvironment::Open: LogDir=/data/litecoin/database ErrorFile=/data/litecoin/db.log
2019-12-22T11:47:09Z Cache configuration:
2019-12-22T11:47:09Z * Using 2.0MiB for block index database
2019-12-22T11:47:09Z * Using 56.0MiB for transaction index database
2019-12-22T11:47:09Z * Using 8.0MiB for chain state database
2019-12-22T11:47:09Z * Using 384.0MiB for in-memory UTXO set (plus up to 286.1MiB of unused mempool space)
2019-12-22T11:47:09Z init message: Loading block index...
2019-12-22T11:47:09Z Opening LevelDB in /data/litecoin/blocks/index
2019-12-22T11:47:09Z Opened LevelDB successfully
2019-12-22T11:47:09Z Using obfuscation key for /data/litecoin/blocks/index: 0000000000000000
2019-12-22T11:47:09Z LoadBlockIndexDB: last block file = 0
2019-12-22T11:47:09Z LoadBlockIndexDB: last block file info: CBlockFileInfo(blocks=0, size=0, heights=0...0, time=1970-01-01...1970-01-01)
2019-12-22T11:47:09Z Checking all blk files are present...
2019-12-22T11:47:09Z Initializing databases...
2019-12-22T11:47:09Z Pre-allocating up to position 0x1000000 in blk00000.dat
2019-12-22T11:47:09Z Opening LevelDB in /data/litecoin/chainstate
2019-12-22T11:47:09Z Opened LevelDB successfully
2019-12-22T11:47:09Z Wrote new obfuscate key for /data/litecoin/chainstate: e3cad79a0194fb6f
2019-12-22T11:47:09Z Using obfuscation key for /data/litecoin/chainstate: e3cad79a0194fb6f
2019-12-22T11:47:09Z init message: Rewinding blocks...
2019-12-22T11:47:09Z  block index              14ms
2019-12-22T11:47:09Z Opening LevelDB in /data/litecoin/indexes/txindex
2019-12-22T11:47:09Z Opened LevelDB successfully
2019-12-22T11:47:09Z Using obfuscation key for /data/litecoin/indexes/txindex: 0000000000000000
2019-12-22T11:47:09Z init message: Loading wallet...
2019-12-22T11:47:09Z txindex thread start
2019-12-22T11:47:09Z txindex is enabled
2019-12-22T11:47:09Z txindex thread exit
2019-12-22T11:47:09Z [default wallet] nFileVersion = 170100
2019-12-22T11:47:09Z [default wallet] Keys: 0 plaintext, 0 encrypted, 0 w/ metadata, 0 total. Unknown wallet records: 0
2019-12-22T11:47:09Z [default wallet] Performing wallet upgrade to 169900
2019-12-22T11:47:10Z [default wallet] keypool added 2000 keys (1000 internal), size=2000 (1000 internal)
2019-12-22T11:47:10Z [default wallet] Wallet completed loading in             976ms
2019-12-22T11:47:10Z [default wallet] setKeyPool.size() = 2000
2019-12-22T11:47:10Z [default wallet] mapWallet.size() = 0
2019-12-22T11:47:10Z [default wallet] mapAddressBook.size() = 0
2019-12-22T11:47:10Z UpdateTip: new best=12a765e31ffd4059bada1e25190f6e98c99d9714d334efa41a195a7e7e04bfe2 height=0 version=0x00000001 log2_work=20.000022 tx=1 date='2011-10-07T07:31:05Z' progress=0.000000 cache=0.0MiB(0txo)
2019-12-22T11:47:10Z Failed to open mempool file from disk. Continuing anyway.
2019-12-22T11:47:10Z mapBlockIndex.size() = 1
2019-12-22T11:47:10Z nBestHeight = 0
2019-12-22T11:47:10Z torcontrol thread start
2019-12-22T11:47:10Z Bound to 127.0.0.1:9333
2019-12-22T11:47:10Z init message: Loading P2P addresses...
2019-12-22T11:47:10Z ERROR: DeserializeFileDB: Failed to open file /data/litecoin/peers.dat
2019-12-22T11:47:10Z Invalid or missing peers.dat; recreating
2019-12-22T11:47:10Z init message: Loading banlist...
2019-12-22T11:47:10Z ERROR: DeserializeFileDB: Failed to open file /data/litecoin/banlist.dat
2019-12-22T11:47:10Z Invalid or missing banlist.dat; recreating
2019-12-22T11:47:10Z init message: Starting network threads...
2019-12-22T11:47:10Z net thread start
2019-12-22T11:47:10Z init message: Done loading
2019-12-22T11:47:10Z dnsseed thread start
2019-12-22T11:47:10Z opencon thread start
2019-12-22T11:47:10Z Loading addresses from DNS seeds (could take a while)
2019-12-22T11:47:10Z msghand thread start
2019-12-22T11:47:10Z addcon thread start
2019-12-22T11:47:11Z New outbound peer connected: version: 70015, blocks=1757608, peer=0
2019-12-22T11:47:12Z New outbound peer connected: version: 70015, blocks=1757608, peer=1
2019-12-22T11:47:14Z New outbound peer connected: version: 70015, blocks=1757608, peer=2
...
```
