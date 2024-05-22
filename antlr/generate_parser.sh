#!/bin/sh

# Set the working directory to the root of the project
cd "$(dirname "$0")/.."

# Directory to place the generated files
OUTPUT_DIR="pkg/parser"
TEMP_DIR="pkg/antlr_temp"

# ANTLR4 grammar file
GRAMMAR_FILE="antlr/Epha.g4"

# Create the output directories if they don't exist
mkdir -p $OUTPUT_DIR
mkdir -p $TEMP_DIR

# Generate the parser files with ANTLR4 in a temporary directory
antlr4 -Dlanguage=Go -package parser -o $TEMP_DIR $GRAMMAR_FILE

# Move the generated files to the desired directory
mv $TEMP_DIR/antlr/* $OUTPUT_DIR/

# Clean up temporary directory and unnecessary files
rm -rf $TEMP_DIR
rm -f $OUTPUT_DIR/*.interp $OUTPUT_DIR/*.tokens

echo "Parser files generated in $OUTPUT_DIR"
