all: package

# Compile vulnreport client binary
vrclient:
	cd vulnreport && go build -o vrclient ./client

# Compile vulnreport server binary
vrserver:
	cd vulnreport && go build -o vrserver ./server

# Run vulnreport server binary
run: vrserver
	cd vulnreport && ./vrserver

# Remove application binaries
binclean:
	cd vulnreport && rm vrclient vrserver

# Compile package binaries and move to appropriate directories
package: vrclient
	cp vulnreport/vrclient penguinsay_1.0-1/usr/bin/vulnreport
	cp penguinsay/* penguinsay_1.0-1/usr/share/penguinsay
	chmod 755 penguinsay_1.0-1/DEBIAN
	chmod 755 penguinsay_1.0-1/DEBIAN/preinst
	chmod 755 penguinsay_1.0-1/DEBIAN/postinst
	chmod 755 penguinsay_1.0-1/DEBIAN/prerm
	chmod 755 penguinsay_1.0-1/DEBIAN/postrm
	chmod +x penguinsay_1.0-1/usr/bin/penguinsay
	dpkg-deb --build penguinsay_1.0-1

# Remove binaries from package directory
packclean:
	rm penguinsay_1.0-1/usr/bin/*

# Remove all binaries
clean: binclean packclean