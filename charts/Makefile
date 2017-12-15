

charts:
	helm dependency build marathonserver
	helm dependency build marathon

clean:
	find ./ -name "*.tgz" -exec rm {} \;
	find ./ -name "*.lock" -exec rm {} \;
