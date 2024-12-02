default:
	# Generate packaged executables for Linux & Windows.
	# In the case of Linux, a Makefile is created
	# to facilitate its installation in the operating system.
	# 
	# Run "create-bundles" to generate the bundles

create-bundles:
	wails build -clean -upx
	wails build -skipbindings -s -platform windows/amd64 -upx
	sh scripts/packaging-linux.sh 'Nu-i uita' nuiuita
	sh scripts/packaging-windows.sh nuiuita