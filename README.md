1. Proto validate
	# Install protovalidate
		go get buf.build/go/protovalidate
		Docs: https://protovalidate.com
	
	# Install buf
		https://buf.build/docs/cli/installation/ (tab github)
	# add buf.yaml and buf.gen.yaml
	# run:
		buf dep update
		buf generate
