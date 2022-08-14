# index-price

Index price application listen external exchange services.

All quotes/prices are stored in buffers.

Buffers use segments based on truncated unix timestamp.

App prints average price for ticker every minute.

Prices ticking based on system clock.

### Usage

Build and run:

```shell
make
./index-price
```

Sample of output to console:

```text
1660515900 0.5088131075804161
1660515960 0.4997630491348989
```