#!/bin/bash

# The directory containing the original script and where the output files will be stored
KOMPOSE_ROOT=$(readlink -f $(dirname "${BASH_SOURCE}")/../../..)
OUTPUT_DIR="$KOMPOSE_ROOT/output"

# Create output directory if it does not exist
mkdir -p "$OUTPUT_DIR"

# Define the input test script file
TEST_SCRIPT="$KOMPOSE_ROOT/script/test/cmd/tests.sh"
echo $TEST_SCRIPT

# Read through tests.sh and process each line with cmd=
grep "^cmd=" $TEST_SCRIPT | while read -r line; do
    # Extract the command
    eval $line
    
    # Use sed to extract the output path after redirection '>'
    output_path=$(echo $cmd | sed -n 's/.*>\s*\(.*\)$/\1/p')

    # Check if no output path was found in the command
    if [ -z "$output_path" ]; then
        # Default paths based on provider
        if [[ "$cmd" == *"--provider=openshift"* ]]; then
            output_path="/tmp/output-os.json"
        else
            output_path="/tmp/output-k8s.json"
        fi
    fi
    
    # Execute the command and redirect output
    eval "$cmd"
    
    # Log the output path used
    echo "Output for command ($cmd) stored in $output_path"
done