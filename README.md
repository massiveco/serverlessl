# serverlessl

Serverless powered TLS PKI.  

## Overview

`serverlessl` leverages cfssl to provide a TLS PKI solution with minimal requirements.  It was initially concieved to provide a PKI for etcd/kubernetes clusters to allow for zero input cluster standup when combined with an instance init tool like [massiveco/headstart](https://github.com/massiveco/headstart) 

## Components

### ca

The `ca` function is responsible for bootstrapping the PKI

### signer

Leverages the Certificate Authority created by the [ca function](#ca) to sign CSRs

