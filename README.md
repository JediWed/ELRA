# E.L.R.A.
## **E**asy **L**ND **R**EST **A**PI

**ELRA** is an open source REST extension for [LND](https://github.com/lightningnetwork/lnd) with a lot of additional features compared to LND's default REST or gRPC API built on Go.

## Comparison

|                                   |ELRA|LND REST|LND gRPC|
|-----------------------------------|:--:|:------:|:------:|
|Easy to use                        |<span style="color:green">**yes**</span>|<span style="color:red">**no**  |<span style="color:red">**no**        |
|User Management                    |<span style="color:green">**yes**</span>|<span style="color:red">**no**  |<span style="color:red">**no**        |
|Role Management                    |<span style="color:green">**yes**</span>|<span style="color:red">**no**  |<span style="color:red">**no**        |
|Reveals Macaroons to Frontend      |<span style="color:green">**no**</span> |<span style="color:green">**no**|<span style="color:red">**yes**       |
|Reveals RPC Credentials to Frontend|<span style="color:green">**no**</span> |<span style="color:red">**yes** |<span style="color:green">**no**</span>|

## License

ELRA is released under the terms of the GPL-3.0 license. See [LICENSE](LICENSE) for more information or see https://www.gnu.org/licenses/.