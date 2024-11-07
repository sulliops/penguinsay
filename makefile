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
package: vrclient
	cp vulnreport/vrclient penguinsay_1.0-1/usr/bin/vulnreport
	dpkg-deb --build penguinsay_1.0-1

# Remove binaries from package directory
packclean:
	rm penguinsay_1.0-1/usr/bin/*

# Remove all binaries
clean: binclean packclean