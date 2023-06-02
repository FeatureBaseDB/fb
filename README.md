# fb

The `featurebasedb/fb` repository contains shared types used across various
featurebase services.

The idea is that we can start to share common types across services written in
Go. So for example, a service might import this package and use its types like
this:

```
"github.com/featurebasedb/fb"

type addr fb.Address
```
