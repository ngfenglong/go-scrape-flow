# Define the binary name
BINARY_NAME=goscrape
OUTPUT_FOLDER=output

# ==============================================================================

# Define the default make action
all: build setup

# Build the project
build:
	go build -o ${BINARY_NAME} .
	chmod +x ${BINARY_NAME}

# To setup project folder
setup: 
	mkdir ./${OUTPUT_FOLDER}

# Run the project
run:
	./${BINARY_NAME}

# Clean up
clean:
	go clean
	rm ${BINARY_NAME}

# Clean up
clean-output:
	rm -r ./${OUTPUT_FOLDER}


# ==============================================================================
# dev test command

dev-test:
	./${BINARY_NAME} start https://www.rocktherankings.com/sitemap.xml

	
	

