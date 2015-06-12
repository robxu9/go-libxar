VERSION ?= 1.6.1

all: vendor/xar/lib/libxar.a
	go build -tags vendor

vendor/xar/lib/libxar.a: vendor/xar
	cd vendor/xar && ./configure
	${MAKE} -C vendor/xar

vendor/xar: vendor/cache.tar.gz
	tar xf vendor/cache.tar.gz -C vendor
	mv vendor/xar-${VERSION} vendor/xar

vendor/cache.tar.gz:
	mkdir -p vendor
	curl -LRo vendor/cache.tar.gz https://github.com/downloads/mackyle/xar/xar-${VERSION}.tar.gz

clean:
	rm -rf vendor/xar vendor/cache.tar.gz

.PHONY: clean
