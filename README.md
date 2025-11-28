1. Proto validate
	# Install protovalidate
		go get buf.build/go/protovalidate
		Docs: https://protovalidate.com
	
	# Install brew
		curl -sSL https://github.com/bufbuild/buf/releases/download/v1.61.0/buf-Linux-x86_64 -o buf
		(Replace v1.61.0 with the latest version if needed.)
		chmod +x buf
		sudo mv buf /usr/local/bin/
		buf --version
	# add buf.yaml and buf.gen.yaml
	# run:
		buf dep update
		buf generate
