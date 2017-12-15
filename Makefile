

image:
	make -C image/marathonserver
	make -C image/groupmanager

charts:
	make -C charts charts

clean:
	make -C charts clean

.PHONY: charts clean image
