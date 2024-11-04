all: package

# Compile vulnreport client binary
vrclient:
	cd vulnreport && go build -o vrclient ./client

# Compile vulnreport server binary
vrserver:
	cd vulnreport && go build -o vrserver ./server

# Remove application binaries
binclean:
	cd vulnreport && rm vrclient vrserver

# Compile package binaries and move to appropriate directories
package: vrclient vrserver
	cp vulnreport/vrclient package/usr/bin/vulnreport

# Remove binaries from package directory
packclean:
	rm package/usr/bin/*

# Remove all binaries
clean: binclean packclean