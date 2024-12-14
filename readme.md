## Csim AT command go Library

A little helper library to exchange APDU with a SIM card inside device which supports `AT+CSIM` command (most GSM/3G/LTE modems).

The library has been tested with https://go.bug.st/serial/ package as transport layer, but can be used with anything that implements `io.ReadWriteCloser` interface.

The code is less than 70 lines long so could be considered self-explained. Use `Csim()` function to send APDU command and its response is in returned value. It does feature a primitive AT response parser (func `ExpectATResp`) which is command-agnostic. Other functions are also exposed because I need them sometimes.